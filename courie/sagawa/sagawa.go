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

	reader := strings.NewReader("<dl class=\"okurijo_info\">\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t<dt>\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t再配達希望日時：\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t</dt>\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t<dd>\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t05月25日　午前中")
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return err
	}

	op.tableWriterOP.Write(&tablewriter.TableWriterParameter{
		Header: op.headers(),
		Data:   op.parseData(doc),
	})

	return nil
}

// header文字列の配列を生成して返す
func (op *sagawaOperator) headers() []string {
	return []string{
		"お問い合せ送り状NO",
		"お届け予定日時",
		"出荷日",
		"集荷営業所",
		"配達営業所",
		"お荷物個数",
		"荷物状況",
		"日時",
		"担当営業所",
	}
}

func (op *sagawaOperator) parseData(doc *goquery.Document) [][]string {
	var data [][]string

	return data
}
