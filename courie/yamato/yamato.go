package yamato

import (
	"fmt"
	"net/url"
	"strings"

	"gotrack/tablewriter"

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

func (op *yamatoOperator) TrackShipments(ids []string) error {
	queryParams := url.Values{}
	queryParams.Add("number00", "1")

	for i, id := range ids {
		// 伝票番号のパラメータの形式がnumber01～number10なので、1はじまり
		queryParams.Add(fmt.Sprintf("number%02d", 1+i), id)
	}

	//resp, err := http.PostForm(truckingURL, queryParams)
	//if err != nil {
	//	return nil, err
	//}
	//defer resp.Body.Close()
	//
	//reader := transform.NewReader(resp.Body, japanese.ShiftJIS.NewDecoder())
	//
	reader := strings.NewReader(
		"<table class=\"saisin\">\n<tbody><tr>\n <td class=\"image\"><img src=\"/images/arrow_001.gif\" alt=\"最新\"></td>\n <td class=\"number\">　１件目</td>\n <td class=\"bold\">伝票番号 3970-0685-0170</td>\n</tr>\n<tr>\n <td rowspan=\"3\"><br></td>\n <td rowspan=\"3\"><br></td>\n <td class=\"font14\">配達完了</td>\n</tr>\n<tr>\n <td class=\"bold\">このお品物はお届けが済んでおります。</td>\n</tr>\n<tr>\n <td>お問い合わせはサービスセンターまでお願いいたします。</td>\n</tr>\n</tbody></table>" +
			"<table class=\"meisaibase\">\n<tbody><tr>\n <th width=\"50%\">商品名</th>\n <th width=\"50%\">お届け予定日時</th>\n</tr>\n<tr>\n <td nowrap=\"\">宅急便<br></td>\n <td nowrap=\"\">05/13　08:00-12:00<br></td>\n</tr>\n</tbody></table>" +
			"<table class=\"meisai\">\n<tbody><tr>\n <th width=\"55\"><br></th>\n <th>荷物状況</th>\n <th>日 付</th>\n <th>時 刻</th>\n <th>担当店名</th>\n <th>担当店コード</th>\n</tr>\n<tr class=\"odd\">\n <td class=\"image\"><img src=\"/images/ya_02.gif\" alt=\"経過\"></td>\n <td>荷物受付</td>\n <td>05/12</td>\n <td>15:31</td>\n <td>芝浦３丁目センター</td>\n <td>136251</td>\n</tr>\n<tr class=\"even\">\n <td class=\"image\"><img src=\"/images/ya_02.gif\" alt=\"経過\"></td>\n <td>発送</td>\n <td>05/12</td>\n <td>15:31</td>\n <td><a href=\"http://sneko2.kuronekoyamato.co.jp/sneko2/Sngp?ID=NET_C&amp;JC=136251&amp;DN=&amp;MD=&amp;F=2\" target=\"_blank\">芝浦３丁目センター</a></td>\n <td>136251</td>\n</tr>\n<tr class=\"odd\">\n <td class=\"image\"><img src=\"/images/ya_02.gif\" alt=\"経過\"></td>\n <td>作業店通過</td>\n <td>05/13</td>\n <td>04:13</td>\n <td>中部ゲートウェイベース</td>\n <td>053990</td>\n</tr>\n<tr class=\"even\">\n <td class=\"image\"><img src=\"/images/nimotsu_01.gif\" alt=\"最新\"></td>\n <td>配達完了</td>\n <td>05/13</td>\n <td>10:58</td>\n <td><a href=\"http://sneko2.kuronekoyamato.co.jp/sneko2/Sngp?ID=NET_C&amp;JC=053212&amp;DN=&amp;MD=&amp;F=2\" target=\"_blank\">岡崎六名センター</a></td>\n <td>053212</td>\n</tr>\n</tbody></table>",
	)

	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return err
	}

	op.tableWriter.Write(op.parseHeader(doc), op.parseData(doc))

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
	var id, item, eta string
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
				ms = append([]string{id, item, eta}, ms...)
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
