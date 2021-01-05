package reports

import (
	"io/ioutil"
	"os"
	"regexp"

	transactions "github.com/lyticaa/lyticaa-app/internal/data/reports/ingest/amazon/reports/custom_transactions"
	sponsored "github.com/lyticaa/lyticaa-app/internal/data/reports/ingest/amazon/reports/sponsored_products"
	"github.com/lyticaa/lyticaa-app/internal/data/reports/ingest/amazon/reports/types"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	customTransaction = regexp.MustCompile(`CustomTransaction`).MatchString
	sponsoredProducts = regexp.MustCompile(`Sponsored\+Products`).MatchString
)

type Reports struct {
	logger zerolog.Logger
	db     *sqlx.DB
}

func NewReports(db *sqlx.DB) *Reports {
	return &Reports{
		logger: log.With().Str("module", os.Getenv("APP_NAME")).Logger(),
		db:     db,
	}
}

func (r *Reports) Run(username, filename string) {
	result := r.s3Object(filename)
	validType := types.ValidMime(*result.ContentType)

	if validType {
		body, err := ioutil.ReadAll(result.Body)
		if err != nil {
			return
		}

		_ = r.processReport(filename, username, *result.ContentType, body)
	} else {
		r.logger.Info().Msgf("invalid content type %v", *result.ContentType)
	}
}

func (r *Reports) ShouldProcess(filename string) bool {
	process := false
	if customTransaction(filename) || sponsoredProducts(filename) {
		process = true
	}

	return process
}

func (r *Reports) processReport(filename, username, contentType string, body []byte) error {
	rows := r.ToMap(contentType, body)
	r.logger.Info().Msgf("total rows to process: %v", len(rows))

	if customTransaction(filename) {
		t := transactions.NewCustomTransactions(r.db)
		_ = t.Process(rows, username)
	}

	var errorList []error
	if sponsoredProducts(filename) {
		s := sponsored.NewSponsoredProducts(r.db)
		errorList = s.Process(rows, username)
	}

	if len(errorList) > 0 {
		for _, err := range errorList {
			r.logger.Error().Err(err).Msg("failed while attempting to save the data")
		}
	}

	return nil
}
