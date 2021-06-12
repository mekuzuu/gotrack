package sagawa

import (
	"net/http"
	"strings"

	"gotrack/tablewriter"

	"github.com/PuerkitoBio/goquery"
)

const truckingURL = "http://k2k.sagawa-exp.co.jp/p/web/okurijosearch.do?okurijoNo"

type sagawaOperator struct {
	tableWriterOP tablewriter.ITableWriterOperator
}

func NewSagawaOperator(
	tableWriterOP tablewriter.ITableWriterOperator,
) *sagawaOperator {
	return &sagawaOperator{
		tableWriterOP: tableWriterOP,
	}
}

func (op *sagawaOperator) TrackShipment(id string) error {
	resp, err := http.Get(truckingURL + id)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	//reader := transform.NewReader(resp.Body, japanese.ShiftJIS.NewDecoder())

	reader := strings.NewReader(
		"<table class=\"table_basic ttl01\">\n\t\t\t\t\t\t\t\t\t\t\n\t\t\t\t\t\t\t\t\t\t<tbody><tr>\n\t\t\t\t\t\t\t\t\t\t\t<th class=\"detail\">詳細</th>\n\t\t\t\t\t\t\t\t\t\t\t<th class=\"number\">お問い合せ送り状NO</th>\n\t\t\t\t\t\t\t\t\t\t\t<th>最新荷物状況</th>\n\t\t\t\t\t\t\t\t\t\t</tr>\n\t\t\t\t\t\t\t\t\t</tbody></table>",
	)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return err
	}

	op.tableWriterOP.Write(&tablewriter.TableWriterParameter{
		Header: op.parseHeader(doc),
		Data:   op.parseData(doc),
	})

	return nil
}

func (op *sagawaOperator) parseHeader(doc *goquery.Document) []string {
	var header []string
	no := doc.Find("table.table_basic.ttl01").Text()
	header = append(header, no)

	return header
}

func (op *sagawaOperator) parseData(doc *goquery.Document) [][]string {
	var data [][]string

	return data
}
