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

func ChangePlan(ctx context.Context, user *models.UserModel, planID string, db *sqlx.DB) error {
	if user.StripePlanID.String == planID {
		return fmt.Errorf("user %v already on plan %v", user.UserID, user.StripePlanID)
	}

	if err := payments.NewStripePayments().ChangePlan(user.StripeSubscriptionID.String, planID); err != nil {
		return err
	}

	var plan sql.NullString
	if err := plan.Scan(planID); err != nil {
		return err
	}

	user.StripePlanID = plan
	if err := user.Update(ctx, db); err != nil {
		return err
	}

	return nil
}

func CancelSubscription(ctx context.Context, user *models.UserModel, db *sqlx.DB) error {
	if err := payments.NewStripePayments().CancelSubscription(user.StripeSubscriptionID.String); err != nil {
		return err
	}

	user.StripeSubscriptionID = sql.NullString{}
	user.StripePlanID = sql.NullString{}
	if err := user.Update(ctx, db); err != nil {
		return err
	}

	return nil
}

func Subscribe(ctx context.Context, user *models.UserModel, planID string, db *sqlx.DB) error {
	if user.StripePlanID.String == planID {
		return fmt.Errorf("user %v already on plan %v", user.UserID, user.StripePlanID)
	}

	sub, err := payments.NewStripePayments().CreateSubscription(user.StripeUserID.String, planID)
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

	user.StripeSubscriptionID = subscription
	user.StripePlanID = plan
	if err = user.Update(ctx, db); err != nil {
		return err
	}

	return nil
}

func Invoices(accountInvoices *types.Invoices, stripeUserID string) {
	invoices := payments.NewStripePayments().InvoicesByUser(stripeUserID)
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
	stripePayments := payments.NewStripePayments()

	session, err := stripePayments.EventSession(event)
	if err != nil {
		return err
	}

	var user models.UserModel
	user.UserID = stripePayments.CustomerRefID(&session)
	user.FetchOne(ctx, db)

	var stripeUserID sql.NullString
	if err := stripeUserID.Scan(*stripePayments.CustomerID(&session)); err != nil {
		return err
	}
	user.StripePlanID = stripeUserID

	var subscriptionID sql.NullString
	if err := subscriptionID.Scan(*stripePayments.SubscriptionID(&session)); err != nil {
		return err
	}
	user.StripeSubscriptionID = subscriptionID

	var planID sql.NullString
	if err := planID.Scan(*stripePayments.PlanID(&session)); err != nil {
		return err
	}
	user.StripePlanID = planID

	if err := user.Update(ctx, db); err != nil {
		return err
	}

	return nil
}

func Notifications(ctx context.Context, accountNotifications *types.Notifications, filter *models.Filter, userID int64, db *sqlx.DB) {
	var notificationModel models.NotificationModel
	notificationModel.UserID = userID

	notifications := notificationModel.FetchAll(ctx, nil, filter, db).([]models.NotificationModel)
	for _, notification := range notifications {
		table := types.NotificationTable{
			Notification: notification.Notification,
			Date:         notification.CreatedAt,
		}

		accountNotifications.Data = append(accountNotifications.Data, table)
	}

	accountNotifications.RecordsTotal = notificationModel.Count(ctx, nil, db)
}

func invoiceClass(status string) string {
	return invoiceStatusMap[status]
}
