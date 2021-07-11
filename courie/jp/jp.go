package jp

import (
	"fmt"
	"net/http"
	"strings"

	"gotrack/tablewriter"

	"github.com/PuerkitoBio/goquery"
)

const trackingURL = "http://tracking.post.japanpost.jp/service/singleSearch.do?org.apache.struts.taglib.html.TOKEN=&searchKind=S002&locale=ja&SVID=&reqCodeNo1="

type jpOperator struct {
	tableWriterOP tablewriter.ITableWriterOperator
}

func NewJPOperator(
	tableWriterOP tablewriter.ITableWriterOperator,
) *jpOperator {
	return &jpOperator{
		tableWriterOP: tableWriterOP,
	}
}

func (op *jpOperator) TrackShipment(id string) error {
	url := fmt.Sprintf("%s%s", trackingURL, id)
	resp, err := http.Get(url)
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
		SetHeaderAlignmentCenter: true,
		SetAlignmentCenter:       true,
	})

	return nil
}

// headerはhtmlからパースした結果をしようしていないので、実際のheaderであることは保証されていない
func (op *jpOperator) headers() []string {
	return []string{
		"お問い合せ番号",
		"商品種別",
		"付加サービス",
		"状態発生日\n（海外で発生した場合は現地時間）",
		"配送情報",
		"取扱局",
	}
}

func (op *jpOperator) parseData(doc *goquery.Document) [][]string {
	var id, itemType, additionalService string
	content := doc.Find("#content")
	// お問い合わせ番号
	id = content.Find(".tableType01.txt_c.m_b5 > tbody > tr > td.w_180").Text()
	// 商品種別
	itemType = content.Find(".tableType01.txt_c.m_b5 > tbody > tr > td.w_380").Text()
	// 付加サービス
	additionalService = content.Find(".tableType01.txt_c.m_b5 > tbody > tr > td.w_100").Text()

	var deliveryStatus []string
	var hasDate bool
	var hasStatus bool
	//var dates, statuses, details []string
	content.Find(".tableType01.txt_c.m_b5 > tbody > tr").Each(func(i int, s *goquery.Selection) {
		// 状態発生日
		date := s.Find("td.w_120").Text()
		if len(date) > 0 && date != "" {
			deliveryStatus = append(deliveryStatus, date)
			hasDate = true
		}

		// 配送履歴
		status := s.Find("td.w_150").Text()
		if len(status) > 0 && status != "" {
			deliveryStatus = append(deliveryStatus, status)
			hasStatus = true
		}

		// 詳細
		if hasDate && hasStatus {
			var defaultVal, postStr = "海外", "郵便局"
			detail := s.Find("td.w_105").Text()

			if strings.Contains(detail, postStr) {
				to := strings.LastIndex(detail, postStr)
				defaultVal = fmt.Sprintf("%s%s", detail[0:to], postStr)
			}

			deliveryStatus = append(deliveryStatus, defaultVal)
		}

		// reset
		hasDate, hasStatus = false, false
	})

	var data [][]string
	var items = make([]string, 0, 6)
	for _, v := range deliveryStatus {
		if len(items) < 3 {
			items = append(items, v)
		}

		if len(items) == 3 {
			items = append([]string{id, itemType, additionalService}, items...)
			data = append(data, items)

			items = nil
		}
	}

	data = append(data, items)

	return data
}
