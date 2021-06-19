package yamato

import (
	"net/http"
	"net/url"
	"strings"

	"gotrack/tablewriter"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"

	"github.com/thoas/go-funk"

	"github.com/PuerkitoBio/goquery"
)

const truckingURL = "https://toi.kuronekoyamato.co.jp/cgi-bin/tneko"

type yamatoOperator struct {
	tableWriterOP tablewriter.ITableWriterOperator
}

func NewYamatoOperator(
	tableWriterOP tablewriter.ITableWriterOperator,
) *yamatoOperator {
	return &yamatoOperator{
		tableWriterOP: tableWriterOP,
	}
}

func (op *yamatoOperator) TrackShipments(id string) error {
	queryParams := url.Values{}
	queryParams.Add("number00", "1")
	queryParams.Add("number01", id)

	resp, err := http.PostForm(truckingURL, queryParams)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	reader := transform.NewReader(resp.Body, japanese.ShiftJIS.NewDecoder())

	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return err
	}

	op.tableWriterOP.Write(&tablewriter.TableWriterParameter{
		Header:                   op.headers(),
		Data:                     op.parseData(doc),
		MergeCellsByColumnIndex:  []int{0, 1, 2},
		SetHeaderAlignmentCenter: true,
		SetAlignmentCenter:       true,
	})

	return nil
}

// header文字列の配列を生成して返す
func (op *yamatoOperator) headers() []string {
	return []string{
		"伝票番号",
		"商品名",
		"お届け予定日時",
		"#",
		"荷物状況",
		"日付",
		"時刻",
		"担当店名",
		"担当店コード",
	}
}

func (op *yamatoOperator) parseData(doc *goquery.Document) [][]string {
	var id, item, eta, stMark string
	doc.Find("table.saisin").Each(func(i int, s *goquery.Selection) {
		ss := s.Find("td.bold").Text()
		// 伝票番号の開始位置と終了位置(13:27)
		id = ss[13:27]
	})

	doc.Find("table.meisaibase").Each(func(i int, s *goquery.Selection) {
		mb := s.Find("tr").Next().Text()
		smb := strings.Split(mb, "\n")

		item = smb[1]
		eta = smb[2]
	})

	// HACK:
	var data [][]string
	doc.Find("table.meisai").Each(func(i int, s *goquery.Selection) {
		m := s.Find("tr").Next().Text()
		deleiveryStatus := strings.Split(m, "\n")

		var items = make([]string, 0, 7)
		for _, v := range deleiveryStatus {
			if isEmpty(v) || isWhiteSpace(v) {
				continue
			}
			if len(items) < 5 {
				items = append(items, v)
			}

			if len(items) == 5 {
				if funk.Contains(items, " 配達完了") {
					stMark = "📦"
				} else {
					stMark = "↓"
				}

				items = append([]string{id, item, eta, stMark}, items...)
				data = append(data, items)

				// sliceをリセットする
				items = nil
			}
		}
	})

	return data
}

func isEmpty(s string) bool {
	return s == ""
}

func isWhiteSpace(s string) bool {
	return s == " "
}
