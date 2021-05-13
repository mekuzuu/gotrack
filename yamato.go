package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// https://qiita.com/the_red/items/39eea9ea20f5a81d66e7#web-api%E7%9A%84%E3%81%AA%E4%BB%95%E6%A7%98
const (
	truckingURL      = "https://toi.kuronekoyamato.co.jp/cgi-bin/tneko"
	detailKey        = "number00"
	detailValNotNeed = "0"
	detailValNeed    = "1"
	numberFrom       = 1
)

type Yamato struct{}

func (y *Yamato) FindShipments(ids []string) ([]Shipment, error) {
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

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(body))
	return nil, nil
}
