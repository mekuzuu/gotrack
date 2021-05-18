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
	truckingURL      = "https://toi.kuronekoyamato.co.jp/cgi-bin/tneko"
	detailKey        = "number00"
	detailValNotNeed = "0"
	detailValNeed    = "1"
	numberFrom       = 1
)

const (
	tableHeaderDescription = "商品名"
	tableHeaderETA         = "お届け予定日時"
	tableHeaderStatus      = "荷物状況"
	tableHeaderDate        = "日付"
	tableHeaderTime        = "時刻"
	tableHeaderStoreName   = "担当店名"
	tableHeaderStoreCode   = "担当店名コード"
)

type yamato struct{}

func NewYamato() yamato {
	return yamato{}
}

type shipment struct {
	itemName  string
	eta       string
	status    string
	date      string
	time      string
	storeName string
	storeCode string
}

func (y *yamato) FindShipmentsTable(ids []string) (*tablewriter.TableWriter, error) {
	queryParams := url.Values{}
	queryParams.Add(detailKey, detailValNeed)

	for i, id := range ids {
		queryParams.Add(fmt.Sprintf("number%02d", numberFrom+i), id)
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
		"<table class=\"meisaibase\">\n<tbody><tr>\n <th width=\"50%\">商品名</th>\n <th width=\"50%\">お届け予定日時</th>\n</tr>\n<tr>\n <td nowrap=\"\">宅急便<br></td>\n <td nowrap=\"\">05/13　08:00-12:00<br></td>\n</tr>\n</tbody></table>" +
			"<table class=\"meisaibase\">\n<tbody><tr>\n <th width=\"50%\">商品名</th>\n <th width=\"50%\">お届け予定日時</th>\n</tr>\n<tr>\n <td nowrap=\"\">宅急便<br></td>\n <td nowrap=\"\">05/13　08:00-12:00<br></td>\n</tr>\n</tbody></table>",
	)

	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return nil, err
	}

	headers := parseHeaders(doc, parseMeisaiBaseHeaders())
	data := parseBody(doc, parseMeisaiBaseData())

	return &tablewriter.TableWriter{
		Header: headers,
		Data:   data,
	}, nil
}
