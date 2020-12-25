package accounts

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/lyticaa/lyticaa-app/internal/models"
	"github.com/lyticaa/lyticaa-app/internal/web/pkg/accounts/payments"
	"github.com/lyticaa/lyticaa-app/internal/web/types"

	"github.com/jmoiron/sqlx"
	"github.com/stripe/stripe-go/v72"
	"golang.org/x/text/currency"
)

var (
	invoiceStatusMap = map[string]string{
		"draft":         "badge-info",
		"open":          "badge-warning",
		"paid":          "badge-success",
		"void":          "badge-info",
		"uncollectible": "badge-danger",
	}
)

func UpdatePlan(ctx context.Context, userID int64, planID string, db *sqlx.DB) error {
	accountSubscriptionModel := models.AccountSubscriptionModel{
		UserID: userID,
	}

	accountSubscription := accountSubscriptionModel.FetchAll(ctx, nil, nil, db).(models.AccountSubscriptionModel)
	if accountSubscription.StripePlanID.String == planID {
		return fmt.Errorf("user %v already on plan %v", userID, accountSubscription.StripePlanID)
	}

	if err := payments.NewStripePayments().UpdatePlan(accountSubscription.StripeSubscriptionID.String, planID); err != nil {
		return err
	}

	var plan sql.NullString
	if err := plan.Scan(planID); err != nil {
		return err
	}

	accountSubscription.StripePlanID = plan
	if err := accountSubscription.Update(ctx, db); err != nil {
		return err
	}

	return nil
}

func CancelSubscription(ctx context.Context, userID int64, db *sqlx.DB) error {
	accountSubscriptionModel := models.AccountSubscriptionModel{
		UserID: userID,
	}

	accountSubscription := accountSubscriptionModel.FetchAll(ctx, nil, nil, db).(models.AccountSubscriptionModel)
	if err := payments.NewStripePayments().CancelSubscription(accountSubscription.StripeSubscriptionID.String); err != nil {
		return err
	}

	accountSubscription.StripeSubscriptionID = sql.NullString{}
	accountSubscription.StripePlanID = sql.NullString{}
	if err := accountSubscription.Update(ctx, db); err != nil {
		return err
	}

	return nil
}

func Subscribe(ctx context.Context, userID int64, planID string, db *sqlx.DB) error {
	accountSubscriptionModel := models.AccountSubscriptionModel{
		UserID: userID,
	}

	accountSubscription := accountSubscriptionModel.FetchAll(ctx, nil, nil, db).(models.AccountSubscriptionModel)
	sub, err := payments.NewStripePayments().CreateSubscription(accountSubscription.StripeUserID.String, planID)
	if err != nil {
		return err
	}

	var subscription sql.NullString
	if err := subscription.Scan(sub.ID); err != nil {
		return err
	}

	var plan sql.NullString
	if err := plan.Scan(planID); err != nil {
		return err
	}

	accountSubscription.StripeSubscriptionID = subscription
	accountSubscription.StripePlanID = plan
	if err = accountSubscription.Update(ctx, db); err != nil {
		return err
	}

	return nil
}

func Invoices(ctx context.Context, accountInvoices *types.Invoices, userID int64, db *sqlx.DB) {
	accountSubscriptionModel := models.AccountSubscriptionModel{
		UserID: userID,
	}

	accountSubscription := accountSubscriptionModel.FetchAll(ctx, nil, nil, db).(models.AccountSubscriptionModel)

	invoices := payments.NewStripePayments().InvoicesByUser(accountSubscription.StripeUserID.String)
	for _, invoice := range *invoices {
		unit, _ := currency.ParseISO(string(invoice.Currency))
		table := types.InvoiceTable{
			Number:      invoice.Number,
			Date:        invoice.Date.Format("2006-01-02"),
			Amount:      fmt.Sprintf("%v %v", currency.Symbol(unit), invoice.Amount),
			Status:      strings.ToUpper(string(invoice.Status)),
			StatusClass: invoiceClass(string(invoice.Status)),
			PDF:         invoice.PDF,
		}

		accountInvoices.Data = append(accountInvoices.Data, table)
	}
}

func Checkout(ctx context.Context, event stripe.Event, db *sqlx.DB) error {
	var accountSubscription models.AccountSubscriptionModel
	stripePayments := payments.NewStripePayments()

	session, err := stripePayments.EventSession(event)
	if err != nil {
		return err
	}

	user := models.UserModel{
		UserID: stripePayments.CustomerRefID(&session),
	}
	user.FetchOne(ctx, db)

	accountSubscription.UserID = user.ID

	var stripeUserID sql.NullString
	if err := stripeUserID.Scan(*stripePayments.CustomerID(&session)); err != nil {
		return err
	}
	accountSubscription.StripePlanID = stripeUserID

	var subscriptionID sql.NullString
	if err := subscriptionID.Scan(*stripePayments.SubscriptionID(&session)); err != nil {
		return err
	}
	accountSubscription.StripeSubscriptionID = subscriptionID

	var planID sql.NullString
	if err := planID.Scan(*stripePayments.PlanID(&session)); err != nil {
		return err
	}
	accountSubscription.StripePlanID = planID

	if err := accountSubscription.Create(ctx, db); err != nil {
		return err
	}

	return nil
}

func Notifications(ctx context.Context, accountNotifications *types.Notifications, filter *models.Filter, userID int64, db *sqlx.DB) {
	var accountNotificationModel models.AccountNotificationModel
	accountNotificationModel.UserID = userID

	notifications := accountNotificationModel.FetchAll(ctx, nil, filter, db).([]models.AccountNotificationModel)
	for _, notification := range notifications {
		table := types.NotificationTable{
			Notification: notification.Notification,
			Date:         notification.CreatedAt,
		}

		accountNotifications.Data = append(accountNotifications.Data, table)
	}

	accountNotifications.RecordsTotal = accountNotificationModel.Count(ctx, nil, db)
}

func Preferences(ctx context.Context, userID int64, db *sqlx.DB) models.AccountPreferenceModel {
	accountPreferenceModel := models.AccountPreferenceModel{
		UserID: userID,
	}

	accountPreferences := accountPreferenceModel.FetchAll(ctx, nil, nil, db).(models.AccountPreferenceModel)

	return accountPreferences
}

func CreatePreferences(ctx context.Context, userID int64, db *sqlx.DB) error {
	accountPreferenceModel := models.AccountPreferenceModel{
		UserID: userID,
	}

	if err := accountPreferenceModel.Create(ctx, db); err != nil {
		return err
	}

	return nil
}

func UpdatePreferences(ctx context.Context, userID int64, preferences map[string]interface{}, db *sqlx.DB) error {
	accountPreferenceModel := models.AccountPreferenceModel{
		UserID: userID,
	}

	accountPreferences := accountPreferenceModel.FetchAll(ctx, nil, nil, db).(models.AccountPreferenceModel)

	if preferences["mailing_list"] == true {
		var subscribe sql.NullBool
		if err := subscribe.Scan(preferences["mailing_list"]); err != nil {
			return err
		}

		accountPreferences.MailingList = subscribe
	}

	if preferences["setup_completed"] == true {
		accountPreferences.SetupCompleted = true
	}

	if err := accountPreferences.Update(ctx, db); err != nil {
		return err
	}

	return nil
}

func invoiceClass(status string) string {
	return invoiceStatusMap[status]
}
