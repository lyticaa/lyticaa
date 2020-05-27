package report

import (
	"bytes"
	"encoding/csv"
	"io"
	"strings"

	"gitlab.com/getlytica/lytica-app/internal/worker/app/report/types"

	"github.com/tealeg/xlsx"
)

func (r *Report) toMap(contentType string, body []byte) []map[string]string {
	var rows []map[string]string

	if types.IsCsv(contentType) {
		rows = r.mapCsv(bytes.NewBuffer(body))
	} else if types.IsXlsx(contentType) {
		rows = r.mapXlsx(body)
	}

	return rows
}

func (r *Report) mapCsv(reader io.Reader) []map[string]string {
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
				r.Logger.Info().Msgf("skipping record: %v", record[0])
				continue
			}
		}

		if err != nil {
			r.Logger.Fatal().Err(err)
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

func (r *Report) mapXlsx(body []byte) []map[string]string {
	var rows []map[string]string
	var header []string

	xlFile, err := xlsx.OpenBinary(body)
	if err != nil {
		r.Logger.Error().Err(err)
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
