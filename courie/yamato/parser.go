package yamato

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type parser struct {
	doc     *goquery.Document
	headers []string
	body    [][]string
}

type parseHeaderFuncs func(p *parser)
type parseBodyFuncs func(p *parser)

func parseHeaders(doc *goquery.Document, funcs ...parseHeaderFuncs) []string {
	p := &parser{doc: doc}
	for _, f := range funcs {
		f(p)
	}
	return p.headers
}

func parseMeisaiBaseHeaders() parseHeaderFuncs {
	return func(p *parser) {
		meisaiBase := p.doc.Find("table.meisaibase").Find("tr").Text()
		replaced := strings.Replace(meisaiBase[1:], "\n", ",", -1)
		split := strings.Split(replaced, ",")
		p.headers = append(p.headers, split[0], split[1])
	}
}

func parseMeisaiBaseData() parseBodyFuncs {
	return func(p *parser) {
		p.doc.Find("table.meisaibase").Each(func(i int, s *goquery.Selection) {
			meisaiBase := s.Find("tr").Next().Text()
			replaced := strings.Replace(meisaiBase[1:], "\n", ",", -1)
			split := strings.Split(replaced, ",")
			p.body = append(p.body, []string{split[0], split[1]})
		})
	}
}

func parseBody(doc *goquery.Document, funcs ...parseBodyFuncs) [][]string {
	p := &parser{doc: doc}
	for _, f := range funcs {
		f(p)
	}
	return p.body
}
