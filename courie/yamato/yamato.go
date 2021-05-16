package yamato

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"gotrack/shipment"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

// cf. https://qiita.com/the_red/items/39eea9ea20f5a81d66e7#web-api%E7%9A%84%E3%81%AA%E4%BB%95%E6%A7%98
const (
	courierName      = "Yamato"
	truckingURL      = "https://toi.kuronekoyamato.co.jp/cgi-bin/tneko"
	detailKey        = "number00"
	detailValNotNeed = "0"
	detailValNeed    = "1"
	numberFrom       = 1
)

type yamato struct{}

func NewYamato() yamato {
	return yamato{}
}

func (y *yamato) FindShipments(ids []string) ([]shipment.Shipment, error) {
	//var shipments = make([]shipment.Shipment, len(ids))

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

	var itemNames, etas []string
	doc.Find("table.meisaibase").Each(func(i int, s *goquery.Selection) {
		meisaiBase := s.Find("tr").Next().Text()
		replaced := strings.Replace(meisaiBase, "\n", ",", -1)
		splited := strings.Split(replaced, ",")
		itemNames = append(itemNames, splited[1])
		etas = append(etas, splited[2])
	})

	fmt.Println(itemNames)
	fmt.Println(etas)

	return nil, nil
}