package report

import (
	"bytes"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

const (
	settlementId = 12447169531
	sku          = "PF-EV1C-1R5B"
	quantity     = 1

	transactionReport = `"date/time","settlement id","type","order id","sku","quantity","marketplace","fulfillment","total"
"Dec 1, 2019 12:07:47 AM PST","12447169531","Order","113-0688349-7048213","PF-EV1C-1R5B","1","amazon.com","Amazon","17.04"`
	sponsoredProductsReport = `"Start Date","End Date","Portfolio name","Currency","Impressions","Clicks","Click-Thru Rate (CTR)","Cost Per Click (CPC)","Spend"
"Dec 01, 2019","Dec 21, 2019","Not grouped","USD","50293","47","0.0935%","$ 0.35","$ 16.22"`

	userId                     = "5de89aea5a61280de1f1bf2b"
	transactionReportFile      = userId + "/" + customTransaction + ".csv"
	SponsoredProductReportFile = userId + "/" + sponsoredProducts + ".csv"
)

func SetupTests(t *testing.T) (*Report, sqlmock.Sqlmock, func(*Report)) {
	dbM, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	db := sqlx.NewDb(dbM, "sqlmock")
	r := NewReport(db)

	return r, mock, func(r *Report) {
		r.Db.Close()
	}
}

func TestMapCsv(t *testing.T) {
	r, _, complete := SetupTests(t)
	defer complete(r)

	body := []byte(transactionReport)
	content := r.mapCsv(transactionReportFile, bytes.NewBuffer(body))

	if len(content) == 0 {
		t.Error("no rows generated")
	}

	if content[0]["date/time"] != "Dec 1, 2019 12:07:47 AM PST" {
		t.Error("date/time does not match")
	}

	if content[0]["settlement id"] != "12447169531" {
		t.Error("settlement Id does not match")
	}

	if content[0]["type"] != transactionType {
		t.Error("order does not match")
	}

	if content[0]["order id"] != "113-0688349-7048213" {
		t.Error("order Id does not match")
	}

	if content[0]["sku"] != sku {
		t.Error("sku does not match")
	}

	if content[0]["quantity"] != "1" {
		t.Error("quantity does not match")
	}

	if content[0]["marketplace"] != marketplace {
		t.Error("marketplace does not match")
	}

	if content[0]["fulfillment"] != fulfillment {
		t.Error("fulfillment does not match")
	}

	if content[0]["total"] != "17.04" {
		t.Error("total does not match")
	}
}

func TestUserFromKey(t *testing.T) {
	r, _, complete := SetupTests(t)
	defer complete(r)

	actual := r.userFromKey(transactionReportFile)

	if actual != userId {
		t.Errorf("expected a user Id of %v but got %v instead!", userId, actual)
	}
}

func TestFileType(t *testing.T) {
	r, _, complete := SetupTests(t)
	defer complete(r)

	fileType := r.fileType(transactionReportFile)
	if fileType != customTransaction {
		t.Error("unable to determine file type")
	}
}
