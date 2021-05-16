package yamato

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"gotrack/tablewriter"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
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
)

type yamato struct {
}

func NewYamato() yamato {
	return yamato{}
}

type shipment struct {
	itemName string
	eta      string
}

func (y *yamato) FindShipmentsTable(ids []string) (*tablewriter.TableWriterModel, error) {
	queryParams := url.Values{}
	queryParams.Add(detailKey, detailValNeed)

	for i, id := range ids {
		queryParams.Add(fmt.Sprintf("number%02d", numberFrom+i), id)
	}

	resp, err := http.PostForm(truckingURL, queryParams)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	reader := transform.NewReader(resp.Body, japanese.ShiftJIS.NewDecoder())
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return nil, err
	}

	shipments := make([]shipment, len(ids))
	doc.Find("table.meisaibase").Each(func(i int, s *goquery.Selection) {
		meisaiBase := s.Find("tr").Next().Text()
		replaced := strings.Replace(meisaiBase, "\n", ",", -1)
		split := strings.Split(replaced, ",")

		shipments[i].itemName = split[1]
		shipments[i].eta = split[2]
	})

	return &tablewriter.TableWriterModel{
		Header: y.genTableHeaderStrings(),
		Data:   y.shipmentsToData(shipments),
	}, nil
}

func (y *yamato) genTableHeaderStrings() []string {
	return []string{
		tableHeaderDescription,
		tableHeaderETA,
	}
}

func (y *yamato) shipmentsToData(shipments []shipment) [][]string {
	var data [][]string
	for _, shipment := range shipments {
		data = append(data, []string{
			shipment.itemName,
			shipment.eta,
		})
	}

	return data
}
