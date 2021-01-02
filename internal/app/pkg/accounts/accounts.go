package accounts

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/lyticaa/lyticaa-app/internal/app/models"
	"github.com/lyticaa/lyticaa-app/internal/app/pkg/accounts/payments"
	"github.com/lyticaa/lyticaa-app/internal/app/types"

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

	monthlyPlan = "monthly"
	annualPlan  = "annual"
)

func Subscription(ctx context.Context, userID int64, db *sqlx.DB) models.AccountSubscriptionModel {
	accountSubscriptionModel := models.AccountSubscriptionModel{
		UserID: userID,
	}

	return accountSubscriptionModel.FetchOne(ctx, db).(models.AccountSubscriptionModel)
}

func UpdateSubscription(ctx context.Context, userID int64, planID string, db *sqlx.DB) error {
	accountSubscriptionModel := models.AccountSubscriptionModel{
		UserID: userID,
	}

	accountSubscription := accountSubscriptionModel.FetchOne(ctx, db).(models.AccountSubscriptionModel)
	if accountSubscription.StripePlanID.String == planID {
		return fmt.Errorf("user %v already on plan %v", userID, accountSubscription.StripePlanID)
	}

	if err := payments.NewStripePayments().UpdateSubscription(accountSubscription.StripeSubscriptionID.String, planID); err != nil {
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

	accountSubscription := accountSubscriptionModel.FetchOne(ctx, db).(models.AccountSubscriptionModel)
	if err := payments.NewStripePayments().CancelSubscription(accountSubscription.StripeSubscriptionID.String); err != nil {
		return err
	}

	if err := accountSubscription.Delete(ctx, db); err != nil {
		return err
	}

	return nil
}

func ReactivateSubscription(ctx context.Context, user *models.UserModel, planID string, db *sqlx.DB) error {
	sub, err := payments.NewStripePayments().CreateSubscription(user.StripeCustomerID.String, planID)
	if err != nil {
		return err
	}

	var accountSubscription models.AccountSubscriptionModel
	accountSubscription.UserID = user.ID

	var subscription sql.NullString
	if err := subscription.Scan(sub.ID); err != nil {
		return err
	}
	accountSubscription.StripeSubscriptionID = subscription

	var plan sql.NullString
	if err := plan.Scan(planID); err != nil {
		return err
	}
	accountSubscription.StripePlanID = plan

	if err = accountSubscription.Create(ctx, db); err != nil {
		return err
	}

	return nil
}

func Invoices(ctx context.Context, accountInvoices *types.Invoices, userID string, db *sqlx.DB) {
	userModel := models.UserModel{
		UserID: userID,
	}
	user := userModel.FetchOne(ctx, db).(models.UserModel)

	invoices := payments.NewStripePayments().InvoicesByUser(user.StripeCustomerID.String)
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

	userModel := models.UserModel{
		UserID: stripePayments.CustomerRefID(&session),
	}
	user := userModel.FetchOne(ctx, db).(models.UserModel)
	if user.ID == 0 {
		return fmt.Errorf("unable to find the user %v", userModel.UserID)
	}

	accountSubscription.UserID = user.ID

	var subscriptionID sql.NullString
	if err := subscriptionID.Scan(*stripePayments.SubscriptionID(&session)); err != nil {
		return err
	}
	accountSubscription.StripeSubscriptionID = subscriptionID

	plan, err := stripePayments.PlanID(&session)
	if err != nil {
		return err
	}

	var planID sql.NullString
	if err := planID.Scan(*plan); err != nil {
		return err
	}
	accountSubscription.StripePlanID = planID

	if err := accountSubscription.Create(ctx, db); err != nil {
		return err
	}

	var stripeCustomerID sql.NullString
	if err := stripeCustomerID.Scan(*stripePayments.CustomerID(&session)); err != nil {
		return err
	}
	user.StripeCustomerID = stripeCustomerID

	if err := user.Update(ctx, db); err != nil {
		return err
	}

	return nil
}

func NewStripeSessions(userID, email, sessionID string) (*string, *string, error) {
	monthly, err := payments.NewStripePayments().CheckoutSession(userID, email, monthlyPlan, sessionID)
	if err != nil {
		return nil, nil, err
	}

	annual, err := payments.NewStripePayments().CheckoutSession(userID, email, annualPlan, sessionID)
	if err != nil {
		return nil, nil, err
	}

	return &monthly.ID, &annual.ID, nil
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

	data := make(map[string]interface{})
	data["UserID"] = userID

	accountNotifications.RecordsTotal = accountNotificationModel.Count(ctx, data, db)
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
