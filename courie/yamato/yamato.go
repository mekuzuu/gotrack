package yamato

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"gotrack/tablewriter"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"

	"github.com/thoas/go-funk"

	"github.com/PuerkitoBio/goquery"
)

// cf. https://qiita.com/the_red/items/39eea9ea20f5a81d66e7#web-api%E7%9A%84%E3%81%AA%E4%BB%95%E6%A7%98
const (
	truckingURL = "https://toi.kuronekoyamato.co.jp/cgi-bin/tneko"
)

type yamatoOperator struct {
	tableWriter tablewriter.ITableWriterOperator
}

func NewYamatoOperator(
	tableWriter tablewriter.ITableWriterOperator,
) *yamatoOperator {
	return &yamatoOperator{
		tableWriter: tableWriter,
	}
}

// TODO: 複数の伝票番号に対応する
func (op *yamatoOperator) TrackShipments(ids []string) error {
	queryParams := url.Values{}
	queryParams.Add("number00", "1")

	for i, id := range ids {
		// 伝票番号のパラメータの形式がnumber01～number10なので、1はじまり
		queryParams.Add(fmt.Sprintf("number%02d", 1+i), id)
	}

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

	op.tableWriter.Write(&tablewriter.TableWriterParameter{
		Header:                  op.parseHeader(doc),
		Data:                    op.parseData(doc),
		MergeCellsByColumnIndex: []int{0, 1, 2},
	})

	return nil
}

func (y *yamatoOperator) parseHeader(doc *goquery.Document) []string {
	var header []string
	ss := doc.Find("table.saisin").Find("td.bold").Text()
	header = append(header, ss[:strings.Index(ss, " ")])

	mb := doc.Find("table.meisaibase").Find("tr").Text()
	smb := strings.Split(mb, "\n")

	// 配列の1-3番目までの値を使いたい
	for i := 1; i < 3; i++ {
		header = append(header, smb[i])
	}

	header = append(header, "#")

	m := doc.Find("table.meisai").Find("tr").Text()
	sm := strings.Split(m, "\n")

	// 配列の2-7番目までの値を使いたい
	for i := 2; i < 7; i++ {
		header = append(header, sm[i])
	}

	return header
}

func (y *yamatoOperator) parseData(doc *goquery.Document) [][]string {
	var data [][]string
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
	doc.Find("table.meisai").Each(func(i int, s *goquery.Selection) {
		m := s.Find("tr").Next().Text()
		sm := strings.Split(m, "\n")

		var ms = make([]string, 0, 7)
		for _, v := range sm {
			if isEmpty(v) || isSpace(v) {
				ms = nil
				continue
			}
			if len(ms) < 5 {
				ms = append(ms, v)
			}

			if len(ms) == 5 {
				if funk.Contains(ms, " 配達完了") {
					stMark = "📦"
				} else {
					stMark = "↓"
				}

				ms = append([]string{id, item, eta, stMark}, ms...)
				data = append(data, ms)
			}
		}
	})

	return data
}

func isEmpty(s string) bool {
	return s == ""
}

func isSpace(s string) bool {
	return s == " "
}
