package report

import (
	"bytes"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestToMap(t *testing.T) {
	r, _, complete := SetupTests(t)
	defer complete(r)

	content := r.toMap(typeCsv, readFile(transactionReportFile, t))
	if len(content) == 0 {
		t.Error("no rows generated")
	}

	content = r.toMap(typeXlsx, readFile(sponsoredProductReportFile, t))
	if len(content) == 0 {
		t.Error("no rows generated")
	}
}

func TestMapCsv(t *testing.T) {
	r, _, complete := SetupTests(t)
	defer complete(r)

	content := r.mapCsv(bytes.NewBuffer(readFile(transactionReportFile, t)))
	if len(content) == 0 {
		t.Error("no rows generated")
	}

	if content[0]["date/time"] != "Dec 1, 2019 12:07:47 AM PST" {
		t.Error()
	}

	if content[0]["settlement id"] != "12447169531" {
		t.Error()
	}

	if content[0]["type"] != "Order" {
		t.Error()
	}

	if content[0]["order id"] != "113-0688349-7048213" {
		t.Error()
	}

	if content[0]["sku"] != "PF-EV1C-1R5B" {
		t.Error()
	}

	if content[0]["description"] != "Trained Flag Football Set,10 Man Set,Premium Football Gear, Massive 46 Piece Set, Flags, Belts, Cones, More, Bonus: Stylish Carry Bag & Flag Football" {
		t.Error()
	}

	if content[0]["quantity"] != "1" {
		t.Error()
	}

	if content[0]["marketplace"] != "amazon.com" {
		t.Error()
	}

	if content[0]["fulfillment"] != "Amazon" {
		t.Error()
	}

	if content[0]["order city"] != "Milford" {
		t.Error()
	}

	if content[0]["order state"] != "DE" {
		t.Error()
	}

	if content[0]["order postal"] != "19963" {
		t.Error()
	}

	if content[0]["tax collection model"] != "" {
		t.Error()
	}

	if content[0]["product sales"] != "26.5" {
		t.Error()
	}

	if content[0]["product sales tax"] != "0" {
		t.Error()
	}

	if content[0]["shipping credits"] != "0" {
		t.Error()
	}

	if content[0]["shipping credits tax"] != "0" {
		t.Error()
	}

	if content[0]["gift wrap credits"] != "0" {
		t.Error()
	}

	if content[0]["giftwrap credits tax"] != "0" {
		t.Error()
	}

	if content[0]["promotional rebates"] != "-0.27" {
		t.Error()
	}

	if content[0]["promotional rebates tax"] != "0" {
		t.Error()
	}

	if content[0]["marketplace withheld tax"] != "0" {
		t.Error()
	}

	if content[0]["selling fees"] != "-3.93" {
		t.Error()
	}

	if content[0]["fba fees"] != "-5.26" {
		t.Error()
	}

	if content[0]["other transaction fees"] != "0" {
		t.Error()
	}

	if content[0]["other"] != "0" {
		t.Error()
	}

	if content[0]["total"] != "17.04" {
		t.Error()
	}
}

func TestMapXlsx(t *testing.T) {
	r, _, complete := SetupTests(t)
	defer complete(r)

	content := r.mapXlsx(readFile(sponsoredProductReportFile, t))
	if len(content) == 0 {
		t.Error("no rows generated")
	}

	if len(content) > 0 {
		if content[0]["Start Date"] != "Dec 01, 2019" {
			t.Error()
		}

		if content[0]["End Date"] != "Dec 21, 2019" {
			t.Error()
		}

		if content[0]["Portfolio name"] != "Not grouped" {
			t.Error()
		}

		if content[0]["Currency"] != "USD" {
			t.Error()
		}

		if content[0]["Campaign Name"] != "Flag Football Auto" {
			t.Error()
		}

		if content[0]["Ad Group Name"] != "Ad Group 1" {
			t.Error()
		}

		if content[0]["Advertised SKU"] != "PF-EV1C-1R5B" {
			t.Error()
		}

		if content[0]["Advertised ASIN"] != "B01AQKSLMC" {
			t.Error()
		}

		if content[0]["Impressions"] != "50293" {
			t.Error()
		}

		if content[0]["Clicks"] != "47" {
			t.Error()
		}

		if content[0]["Click-Thru Rate (CTR)"] != "0.0935%" {
			t.Error()
		}

		if content[0]["Cost Per Click (CPC)"] != "$ 0.35" {
			t.Error()
		}

		if content[0]["Spend"] != "$ 16.22" {
			t.Error()
		}

		if content[0]["7 Day Total Sales"] != "$ 86.48" {
			t.Error()
		}

		if content[0]["Total Advertising Cost of Sales (ACoS)"] != "18.7558%" {
			t.Error()
		}

		if content[0]["Total Return on Advertising Spend (RoAS)"] != "5.33" {
			t.Error()
		}

		if content[0]["7 Day Total Orders (#)"] != "3" {
			t.Error()
		}

		if content[0]["7 Day Total Units (#)"] != "3" {
			t.Error()
		}

		if content[0]["7 Day Conversion Rate"] != "6.3830%" {
			t.Error()
		}

		if content[0]["7 Day Advertised SKU Units (#)"] != "3" {
			t.Error()
		}

		if content[0]["7 Day Other SKU Units (#)"] != "1" {
			t.Error()
		}

		if content[0]["7 Day Advertised SKU Sales"] != "$ 86.48" {
			t.Error()
		}

		if content[0]["7 Day Other SKU Sales"] != "$ 1.00" {
			t.Error()
		}
	}
}

func readFile(filename string, t *testing.T) []byte {
	absPath, _ := filepath.Abs(filename)
	body, err := ioutil.ReadFile(absPath)
	if err != nil {
		t.Error(err)
	}

	return body
}
