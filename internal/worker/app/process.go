package app

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"reflect"
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

const (
	customTransaction = "CustomTransaction"
	unknown           = int64(1)
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

		var inputType string
		if strings.Contains(file, "CustomTransaction") {
			inputType = "CustomTransaction"
		}

		a.processInput(inputType, username, body)
	} else {
		a.Logger.Info().Str("user", username).Msgf("invalid content type: %v", *result.ContentType)
	}
}

func (a *App) processInput(file, username string, body []byte) {
	rows := a.mapCsv(file, bytes.NewBuffer(body))
	a.Logger.Info().Str("user", username).Msgf("total rows to process: %v", len(rows))

	switch file {
	case customTransaction:
		a.saveTransactions(rows, username)
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

func (a *App) getTransactionTypes() []models.TransactionType {
	return models.GetTransactionTypes(a.Db)
}

func (a *App) getTransactionTypeIdByName(name string, txnTypes []models.TransactionType) (int64, bool) {
	for _, txnType := range txnTypes {
		if txnType.Name == name {
			return txnType.Id, true
		}
	}

	return unknown, false
}

func (a *App) getMarketplaces() []models.Marketplace {
	return models.GetMarketplaces(a.Db)
}

func (a *App) getMarketplaceIdByName(name string, marketplaces []models.Marketplace) (int64, bool) {
	for _, marketplace := range marketplaces {
		if marketplace.Name == name {
			return marketplace.Id, true
		}
	}

	return unknown, false
}

func (a *App) getFulfillments() []models.Fulfillment {
	return models.GetFulfillments(a.Db)
}

func (a *App) getFulfillmentIdByName(name string, fulfillments []models.Fulfillment) (int64, bool) {
	for _, fulfillment := range fulfillments {
		if fulfillment.Name == name {
			return fulfillment.Id, true
		}
	}

	return unknown, false
}

func (a *App) getTaxCollectionModels() []models.TaxCollectionModel {
	return models.GetTaxCollectionModels(a.Db)
}

func (a *App) getTaxCollectionModelIdByName(name string, taxCollectionModels []models.TaxCollectionModel) (int64, bool) {
	for _, taxCollectionModel := range taxCollectionModels {
		if taxCollectionModel.Name == name {
			return taxCollectionModel.Id, true
		}
	}

	return unknown, false
}

func (a *App) saveTransactions(rows []map[string]string, username string) {
	user := models.FindUserByUserId(username, a.Db)
	txnTypes := a.getTransactionTypes()
	marketplaces := a.getMarketplaces()
	fulfillments := a.getFulfillments()
	taxCollectionModels := a.getTaxCollectionModels()

	var txns []models.Transaction
	for idx, row := range rows {
		txnType, ok := a.getTransactionTypeIdByName(row["type"], txnTypes)
		if !ok && row["type"] != "" {
			a.Logger.Error().Msgf("Transaction Type %v not found", row["type"])
		}

		marketplace, ok := a.getMarketplaceIdByName(strings.ToLower(row["marketplace"]), marketplaces)
		if !ok && row["marketplace"] != "" {
			a.Logger.Error().Msgf("Marketplace %v not found", row["marketplace"])
		}

		fulfillment, ok := a.getFulfillmentIdByName(row["fulfillment"], fulfillments)
		if !ok && row["fulfillment"] != "" {
			a.Logger.Error().Msgf("Fulfillment %v not found", row["fulfillment"])
		}

		taxCollectionModel, ok := a.getTaxCollectionModelIdByName(row["tax collection model"], taxCollectionModels)
		if !ok && row["tax collection model"] != "" {
			a.Logger.Error().Msgf("Tax Collection Model %v not found", row["tax collection model"])
		}

		dateTime, _ := time.Parse("Jan 2, 2006 3:04:05 PM MST", row["date/time"])
		settlementId, _ := strconv.ParseInt(row["settlement id"], 10, 64)
		quantity, _ := strconv.ParseInt(row["quantity"], 10, 64)
		productSales, _ := strconv.ParseFloat(row["product sales"], 64)
		productSalesTax, _ := strconv.ParseFloat(row["product sales tax"], 64)
		shippingCredits, _ := strconv.ParseFloat(row["shipping credits"], 64)
		shippingCreditsTax, _ := strconv.ParseFloat(row["shipping credits tax"], 64)
		giftwrapCredits, _ := strconv.ParseFloat(row["giftwrap credits"], 64)
		giftwrapCreditsTax, _ := strconv.ParseFloat(row["giftwrap credits tax"], 64)
		promotionalRebates, _ := strconv.ParseFloat(row["promotional rebates"], 64)
		promotionalRebatesTax, _ := strconv.ParseFloat(row["promotional rebates tax"], 64)
		marketplaceWithheldTax, _ := strconv.ParseFloat(row["marketplace withheld tax"], 64)
		sellingFees, _ := strconv.ParseFloat(row["selling fees"], 64)
		fbaFees, _ := strconv.ParseFloat(row["fba fees"], 64)
		otherTransactionFees, _ := strconv.ParseFloat(row["other transaction fees"], 64)
		other, _ := strconv.ParseFloat(row["other"], 64)
		total, _ := strconv.ParseFloat(row["total"], 64)

		txn := models.Transaction{
			Idx:                    idx,
			DateTime:               dateTime,
			User:                   user,
			SettlementId:           settlementId,
			TransactionType:        models.TransactionType{Id: txnType},
			OrderId:                row["order id"],
			Sku:                    row["sku"],
			Quantity:               quantity,
			Marketplace:            models.Marketplace{Id: marketplace},
			Fulfillment:            models.Fulfillment{Id: fulfillment},
			TaxCollectionModel:     models.TaxCollectionModel{Id: taxCollectionModel},
			ProductSales:           productSales,
			ProductSalesTax:        productSalesTax,
			ShippingCredits:        shippingCredits,
			ShippingCreditsTax:     shippingCreditsTax,
			GiftwrapCredits:        giftwrapCredits,
			GiftwrapCreditsTax:     giftwrapCreditsTax,
			PromotionalRebates:     promotionalRebates,
			PromotionalRebatesTax:  promotionalRebatesTax,
			MarketplaceWithheldTax: marketplaceWithheldTax,
			SellingFees:            sellingFees,
			FBAFees:                fbaFees,
			OtherTransactionFees:   otherTransactionFees,
			Other:                  other,
			Total:                  total,
		}

		txns = append(txns, txn)
	}

	for _, txn := range txns {
		models.SaveTransaction(txn, a.Db)
	}
}

func (a *App) userFromKey(key string) string {
	parts := strings.Split(key, "/")
	return parts[0]
}
