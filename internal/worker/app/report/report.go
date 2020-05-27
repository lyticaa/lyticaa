package report

import (
	"io/ioutil"
	"os"
	"strings"

	"gitlab.com/getlytica/lytica-app/internal/worker/app/report/types"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	customTransaction = "CustomTransaction"
	sponsoredProducts = "Sponsored Products"
	unknown           = int64(1)
)

type Report struct {
	Logger zerolog.Logger
	Db     *sqlx.DB
}

func NewReport(db *sqlx.DB) *Report {
	return &Report{
		Logger: log.With().Str("module", os.Getenv("APP_NAME")).Logger(),
		Db:     db,
	}
}

func (r *Report) Run(filename string) {
	result := r.getS3Object(filename)
	validType := types.ValidMime(*result.ContentType)
	username := r.userFromFilename(filename)

	if validType {
		body, err := ioutil.ReadAll(result.Body)
		if err != nil {
			r.Logger.Error().Err(err)
			return
		}

		_ = r.processReport(filename, username, *result.ContentType, body)
	} else {
		r.Logger.Info().Str("user", username).Msgf("invalid content type: %v", *result.ContentType)
	}
}

func (r *Report) processReport(filename, username, contentType string, body []byte) error {
	rows := r.toMap(contentType, body)
	r.Logger.Info().Str("user", username).Msgf("total rows to process: %v", len(rows))

	if strings.Contains(filename, customTransaction) {
		err := r.processTransactions(rows, username)
		if err != nil {
			return err
		}
	} else if strings.Contains(filename, sponsoredProducts) {
		err := r.processSponsoredProducts(rows, username)
		if err != nil {
			return err
		}
	} else {
		r.Logger.Info().Msgf("filename %v is not recognised", filename)
	}

	return nil
}

func (r *Report) processTransactions(rows []map[string]string, username string) error {
	transactions := r.formatTransactions(rows, username)
	for _, transaction := range transactions {
		err := r.saveTransaction(transaction)
		if err != nil {
			r.Logger.Error().Err(err)
			return err
		}
	}

	return nil
}

func (r *Report) processSponsoredProducts(rows []map[string]string, username string) error {
	sponsoredProducts := r.formatSponsoredProducts(rows, username)
	for _, sponsoredProduct := range sponsoredProducts {
		err := r.saveSponsoredProduct(sponsoredProduct)
		if err != nil {
			r.Logger.Error().Err(err)
			return err
		}
	}

	return nil
}

func (r *Report) userFromFilename(filename string) string {
	parts := strings.Split(filename, "/")
	return parts[0]
}
