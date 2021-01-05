package reports

import (
	"bytes"
	"encoding/csv"
	"io"
	"strings"

	"github.com/lyticaa/lyticaa-app/internal/data/reports/ingest/amazon/reports/types"

	"github.com/tealeg/xlsx"
)

func (r *Reports) ToMap(contentType string, body []byte) []map[string]string {
	var rows []map[string]string

	if types.IsCSV(contentType) {
		rows = r.MapCSV(bytes.NewBuffer(body))
	} else if types.IsXLSX(contentType) {
		rows = r.MapXLSX(body)
	}

	return rows
}

func (r *Reports) MapCSV(reader io.Reader) []map[string]string {
	rr := csv.NewReader(reader)
	var rows []map[string]string
	var header []string

	for {
		record, err := rr.Read()
		if err == io.EOF {
			break
		}

		if len(record) > 0 {
			shouldSkip := types.ShouldIgnore(record[0])
			if shouldSkip {
				r.logger.Info().Msgf("skipping record: %v", record[0])
				continue
			}
		}

		if header == nil {
			header = record
		} else {
			dict := map[string]string{}
			for i := range header {
				dict[r.translateHeader(header[i])] = record[i]
			}

			rows = append(rows, dict)
		}
	}

	return rows
}

func (r *Reports) MapXLSX(body []byte) []map[string]string {
	var rows []map[string]string
	var header []string

	xlFile, err := xlsx.OpenBinary(body)
	if err != nil {
		return rows
	}

	for _, sheet := range xlFile.Sheets {
		for idx, row := range sheet.Rows {
			dict := map[string]string{}
			if header == nil {
				for _, cell := range sheet.Rows[0].Cells {
					header = append(header, strings.TrimSpace(cell.String()))
				}
			}

			for i, cell := range row.Cells {
				if idx > 0 && len(header) >= i {
					dict[r.translateHeader(header[i])] = strings.ReplaceAll(cell.String(), "\\", "")
				}
			}

			if idx > 0 {
				rows = append(rows, dict)
			}
		}
	}

	return rows
}
