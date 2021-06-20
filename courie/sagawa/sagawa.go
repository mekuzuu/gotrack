package sagawa

import (
	"fmt"
	"net/http"
	"strings"

	"gotrack/tablewriter"

	"github.com/PuerkitoBio/goquery"
)

const trackingURL = "http://k2k.sagawa-exp.co.jp/p/web/okurijosearch.do?okurijoNo="

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
	resp, err := http.Get(trackingURL + id)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return err
	}

	op.tableWriterOP.Write(&tablewriter.TableWriterParameter{
		Header:                   op.headers(),
		Data:                     op.parseData(doc),
		MergeCellsByColumnIndex:  []int{0, 1, 2, 5},
		SetHeaderAlignmentCenter: true,
		SetAlignmentCenter:       true,
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
	var id, etaOr string
	doc.Find("#list1").Each(func(i int, s *goquery.Selection) {
		// お問い合わせ送り状NO
		id = s.Find(".number.nowrap").Text()

		oi := s.Find(".okurijo_info")

		// 配達予定 or 完了日時
		etaOr = oi.Find("dd").Text()
		etaOr = strings.ReplaceAll(etaOr, "\t", "")
		etaOr = strings.ReplaceAll(etaOr, "\n", "")
	})

	var std []string
	var shipDate, pickUp, delivery, itemNum string
	doc.Find("#detail1").Each(func(i int, s *goquery.Selection) {
		od := s.Find(".table_basic.table_okurijo_detail2")

		td := od.Find("td").Text()
		td = strings.ReplaceAll(td, "\t", "")
		std = strings.Split(td, "\n")

		// 0オリジンで配列のn番目にあることを前提としているので、順番が保証されていなかったら壊れる

		// 出荷日
		shipDate = std[2]
		// 集荷営業所
		pickUp = fmt.Sprintf("%s %s %s", std[7], std[12], std[16])
		// 配達営業所
		delivery = fmt.Sprintf("%s %s %s", std[22], std[28], std[32])
		// お荷物個数
		itemNum = std[36]
	})

	var data [][]string
	var items = make([]string, 0, 9)
	deliveryStatus := std[38:len(std)]

	for _, v := range deliveryStatus {
		if isEmpty(v) || isWhiteSpace(v) {
			continue
		}

		if len(items) < 3 {
			items = append(items, v)
		}

		if len(items) == 3 {
			items = append([]string{id, etaOr, shipDate, pickUp, delivery, itemNum}, items...)
			data = append(data, items)

			// sliceをリセットする
			items = nil
		}
	}

	return data
}

func isEmpty(s string) bool {
	return s == ""
}

func isWhiteSpace(s string) bool {
	return s == " "
}
