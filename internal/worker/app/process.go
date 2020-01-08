package app

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"

	"gitlab.com/getlytica/lytica/internal/models"
	"gitlab.com/getlytica/lytica/internal/worker/app/types"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func (a *App) parseMessage(msg *sqs.Message) error {
	body := []byte(*msg.Body)
	var r types.Response

	err := json.Unmarshal(body, &r)
	if err != nil {
		a.Logger.Error().Err(err)
		return nil
	}

	for _, record := range r.Records {
		a.Logger.Info().Msgf("processing %v....", record.S3.Object.Key)
		go a.processFile(record.S3.Object.Key)
	}

	return nil
}

func (a *App) getS3Object(file string) *s3.GetObjectOutput {
	cfg := &aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
	}
	sess, err := session.NewSession(cfg)
	if err != nil {
		a.Logger.Fatal().Err(err)
	}

	svc := s3.New(sess)
	input := &s3.GetObjectInput{
		Bucket: aws.String(os.Getenv("AWS_S3_UPLOAD_BUCKET")),
		Key:    aws.String(file),
	}

	result, err := svc.GetObject(input)
	if err != nil {
		a.Logger.Error().Err(err)
	}

	return result
}

func (a *App) processFile(file string) {
	result := a.getS3Object(file)
	validType := types.ValidMime(*result.ContentType)
	username := a.userFromKey(file)

	if validType {
		body, err := ioutil.ReadAll(result.Body)
		if err != nil {
			a.Logger.Error().Err(err)
			return
		}

		rows := a.mapCsv(file, bytes.NewBuffer(body))
		a.Logger.Info().Str("user", username).Msgf("total rows to process: %v", len(rows))

		if strings.Contains(file, "CustomTransaction") {
			a.processCustomTransaction(rows)
		}
	} else {
		a.Logger.Info().Str("user", username).Msgf("invalid content type: %v", *result.ContentType)
	}
}

// TODO : support for files in other languages. Translate headers into English accordingly.
func (a *App) mapCsv(file string, reader io.Reader) []map[string]string {
	r := csv.NewReader(reader)
	var rows []map[string]string
	var header []string
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		if len(record) > 0 {
			shouldSkip := types.ShouldIgnore(record[0])
			if shouldSkip {
				a.Logger.Info().Str("user", a.userFromKey(file)).Msgf("skipping record: %v", record[0])
				continue
			}
		}

		if err != nil {
			a.Logger.Fatal().Err(err)
		}

		if header == nil {
			header = record
		} else {
			dict := map[string]string{}
			for i := range header {
				dict[header[i]] = record[i]
			}
			rows = append(rows, dict)
		}
	}

	return rows
}

func (a *App) getOrderTypes() []models.OrderType {
	return models.GetOrderTypes(a.Db)
}

func (a *App) getOrderTypeIdByName(name string, orderTypes []models.OrderType) (*models.OrderType, bool) {
	for _, orderType := range orderTypes {
		if orderType.Name == name {
			return &orderType, true
		}
	}

	return &models.OrderType{}, false
}

func (a *App) getMarketplaces() []models.Marketplace {
	return models.GetMarketplaces(a.Db)
}

func (a *App) getMarketplaceIdByName(name string, marketplaces []models.Marketplace) (*models.Marketplace, bool) {
	for _, marketplace := range marketplaces {
		if marketplace.Name == name {
			return &marketplace, true
		}
	}

	return &models.Marketplace{}, false
}

func (a *App) getFulfillments() []models.Fulfillment {
	return models.GetFulfillments(a.Db)
}

func (a *App) getFulfillmentIdByName(name string, fulfillments []models.Fulfillment) (*models.Fulfillment, bool) {
	for _, fulfillment := range fulfillments {
		if fulfillment.Name == name {
			return &fulfillment, true
		}
	}

	return &models.Fulfillment{}, false
}

func (a *App) getTaxCollectionModels() []models.TaxCollectionModel {
	return models.GetTaxCollectionModels(a.Db)
}

func (a *App) getTaxCollectionModelIdByName(name string, taxCollectionModels []models.TaxCollectionModel) (*models.TaxCollectionModel, bool) {
	for _, taxCollectionModel := range taxCollectionModels {
		if taxCollectionModel.Name == name {
			return &taxCollectionModel, true
		}
	}

	return &models.TaxCollectionModel{}, false
}

func (a *App) processCustomTransaction(content []map[string]string) {
	orderTypes := a.getOrderTypes()
	marketplaces := a.getMarketplaces()
	fulfillments := a.getFulfillments()
	taxCollectionModels := a.getTaxCollectionModels()

	var txns []models.CustomTransaction
	for _, row := range content {
		orderType, ok := a.getOrderTypeIdByName(row["type"], orderTypes)
		if !ok {
			a.Logger.Error().Msgf("OrderType %v not found", row["type"])
		}

		marketplace, ok := a.getMarketplaceIdByName(row["marketplace"], marketplaces)
		if !ok {
			a.Logger.Error().Msgf("Marketplace %v not found", row["marketplace"])
		}

		fulfillment, ok := a.getFulfillmentIdByName(row["fulfillment"], fulfillments)
		if !ok {
			a.Logger.Error().Msgf("Fulfillment %v not found", row["fulfillment"])
		}

		taxCollectionModel, ok := a.getTaxCollectionModelIdByName(row["tax collection model"], taxCollectionModels)
		if !ok {
			a.Logger.Error().Msgf("Tax Collection Model %v not found", row["tax collection model"])
		}

		dateTime, _ := time.Parse(time.RFC3339, row["date/time"])
		settlementId, _ := strconv.ParseInt(row["settlement id"], 64, 10)
		quantity, _ := strconv.ParseInt(row["quantity"], 64, 10)
		productSales, _ := strconv.ParseFloat(row["product sales"], 64)
		productSalesTax, _ := strconv.ParseFloat(row["product sales tax"], 64)
		shippingCredits, _ := strconv.ParseFloat(row["shipping credits"], 64)
		shippingCreditsTax, _ := strconv.ParseFloat(row["shipping credits tax"], 64)
		giftWrapCredits, _ := strconv.ParseFloat(row["gift wrap credits"], 64)
		giftWrapCreditsTax, _ := strconv.ParseFloat(row["gift wrap credits tax"], 64)
		promotionalRebates, _ := strconv.ParseFloat(row["promotional rebates"], 64)
		promotionalRebatesTax, _ := strconv.ParseFloat(row["promotional rebates tax"], 64)
		marketplaceWithheldTax, _ := strconv.ParseFloat(row["marketplace withheld tax"], 64)
		sellingFees, _ := strconv.ParseFloat(row["selling fees"], 64)
		fbaFees, _ := strconv.ParseFloat(row["fba fees"], 64)
		otherTransactionFees, _ := strconv.ParseFloat(row["other transaction fees"], 64)
		order, _ := strconv.ParseFloat(row["order"], 64)
		total, _ := strconv.ParseFloat(row["total"], 64)

		txn := models.CustomTransaction{
			DateTime:               dateTime,
			SettlementId:           settlementId,
			OrderType:              *orderType,
			OrderId:                row["order id"],
			Sku:                    row["sku"],
			Quantity:               quantity,
			Marketplace:            *marketplace,
			Fulfillment:            *fulfillment,
			TaxCollectionModel:     *taxCollectionModel,
			ProductSales:           productSales,
			ProductSalesTax:        productSalesTax,
			ShippingCredits:        shippingCredits,
			ShippingCreditsTax:     shippingCreditsTax,
			GiftWrapCredits:        giftWrapCredits,
			GiftWrapCreditsTax:     giftWrapCreditsTax,
			PromotionalRebates:     promotionalRebates,
			PromotionalRebatesTax:  promotionalRebatesTax,
			MarketplaceWithheldTax: marketplaceWithheldTax,
			SellingFees:            sellingFees,
			FBAFees:                fbaFees,
			OtherTransactionFees:   otherTransactionFees,
			Order:                  order,
			Total:                  total,
		}

		txns = append(txns, txn)
	}

	for _, txn := range txns {
		models.SaveCustomTransaction(txn, a.Db)
	}
}

func (a *App) userFromKey(key string) string {
	parts := strings.Split(key, "/")
	return parts[0]
}
