package jp

import (
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
	//url := fmt.Sprintf("%s%s", trackingURL, id)
	//resp, err := http.Get(url)
	//if err != nil {
	//	return err
	//}
	//defer resp.Body.Close()

	r := strings.NewReader("<div id=\"content\">\n    \n    \n    \n\n    <form name=\"srv_searchActionForm\" method=\"post\" action=\"/services/srv/search/\">\n        <h1 class=\"ttl_line m_b20\" style=\"text-align: center;\">検索結果 詳細<br>\n            [国際]\n        </h1>\n\n        <div class=\"indent\">\n            \n            \n            \n            \n            \n            \n            \n                \n            \n            \n            \n                \n                \n                \n                \n                \n                    <div class=\"beige_box m_b15\">\n                        <h2 class=\"beige_box_inner txt_type06 bold\">配達状況詳細</h2>\n                    </div>\n                \n            \n            <table class=\"tableType01 txt_c m_b5\" summary=\"配達状況詳細\">\n              <tbody><tr>\n                \n                    \n                    \n                    \n                        <th scope=\"row\" class=\"bg01 w_180\">お問い合わせ番号</th>\n                        <th scope=\"row\" class=\"bg01 w_380\">商品種別</th>\n                        <th scope=\"row\" class=\"bg01 w_100\">付加サービス</th>\n                    \n                \n              </tr>\n              <tr>\n                \n                    \n                    \n                    \n                        <td class=\"w_180\">ET 700 332 656 VN</td>\n                        <td class=\"w_380\">EMS</td>\n                        <td class=\"w_100\"></td>\n                    \n                \n              </tr>\n              \n            </tbody></table>\n            \n            <br>\n            \n            <div class=\"beige_box m_b15\">\n                <h2 class=\"beige_box_inner txt_type06 bold\">履歴情報</h2>\n            </div>\n            <table class=\"tableType01 txt_c m_b5\" summary=\"履歴情報\">\n              <tbody><tr>\n                \n                    \n                        <th scope=\"row\" rowspan=\"2\" class=\"bg01 w_120\">状態発生日<br>（海外で発生した場合は現地時間）</th>\n                    \n                    \n                \n                <th scope=\"row\" rowspan=\"2\" class=\"bg01 w_150\">配送履歴</th>\n                <th scope=\"row\" rowspan=\"2\" class=\"bg01 w_180\">詳細</th>\n                <th scope=\"row\" class=\"bg01 w_105\">取扱局</th>\n                \n                    \n                        <th scope=\"row\" rowspan=\"2\" class=\"bg01 w_105\">県名・国名</th>\n                    \n                    \n                \n              </tr>\n              <tr>\n                <th scope=\"row\" class=\"bg01 w_105\">郵便番号</th>\n              </tr>\n              \n                  <tr>\n                    <td rowspan=\"2\" class=\"w_120\">2021/06/14 08:14</td>\n                    <td rowspan=\"2\" class=\"w_150\">引受</td>\n                    <td rowspan=\"2\" class=\"w_180\"></td>\n                    <td class=\"w_105\">&nbsp;</td>\n                    <td rowspan=\"2\" class=\"w_105\">VIET NAM        </td>\n                  </tr>\n                  <tr>\n                    <td class=\"w_105\">&nbsp;</td>\n                  </tr>\n              \n                  <tr>\n                    <td rowspan=\"2\" class=\"w_120\">2021/06/14 08:54</td>\n                    <td rowspan=\"2\" class=\"w_150\">国際交換局に到着</td>\n                    <td rowspan=\"2\" class=\"w_180\"></td>\n                    <td class=\"w_105\">HOCHIMINH EM</td>\n                    <td rowspan=\"2\" class=\"w_105\">VIET NAM        </td>\n                  </tr>\n                  <tr>\n                    <td class=\"w_105\">&nbsp;</td>\n                  </tr>\n              \n                  <tr>\n                    <td rowspan=\"2\" class=\"w_120\">2021/06/14 08:56</td>\n                    <td rowspan=\"2\" class=\"w_150\">税関検査のため税関へ提示</td>\n                    <td rowspan=\"2\" class=\"w_180\"></td>\n                    <td class=\"w_105\">&nbsp;</td>\n                    <td rowspan=\"2\" class=\"w_105\">VIET NAM        </td>\n                  </tr>\n                  <tr>\n                    <td class=\"w_105\">&nbsp;</td>\n                  </tr>\n              \n                  <tr>\n                    <td rowspan=\"2\" class=\"w_120\">2021/06/15 11:49</td>\n                    <td rowspan=\"2\" class=\"w_150\">税関から受領</td>\n                    <td rowspan=\"2\" class=\"w_180\"></td>\n                    <td class=\"w_105\">&nbsp;</td>\n                    <td rowspan=\"2\" class=\"w_105\">VIET NAM        </td>\n                  </tr>\n                  <tr>\n                    <td class=\"w_105\">&nbsp;</td>\n                  </tr>\n              \n                  <tr>\n                    <td rowspan=\"2\" class=\"w_120\">2021/06/15 11:51</td>\n                    <td rowspan=\"2\" class=\"w_150\">国際交換局から発送</td>\n                    <td rowspan=\"2\" class=\"w_180\"></td>\n                    <td class=\"w_105\">HOCHIMINH EM</td>\n                    <td rowspan=\"2\" class=\"w_105\">VIET NAM        </td>\n                  </tr>\n                  <tr>\n                    <td class=\"w_105\">&nbsp;</td>\n                  </tr>\n              \n                  <tr>\n                    <td rowspan=\"2\" class=\"w_120\">2021/06/16 18:24</td>\n                    <td rowspan=\"2\" class=\"w_150\">国際交換局に到着</td>\n                    <td rowspan=\"2\" class=\"w_180\"></td>\n                    <td class=\"w_105\">東京国際郵便局</td>\n                    <td rowspan=\"2\" class=\"w_105\">東京都</td>\n                  </tr>\n                  <tr>\n                    <td class=\"w_105\">138-8799</td>\n                  </tr>\n              \n                  <tr>\n                    <td rowspan=\"2\" class=\"w_120\">2021/06/18 05:00</td>\n                    <td rowspan=\"2\" class=\"w_150\">保税運送中</td>\n                    <td rowspan=\"2\" class=\"w_180\"></td>\n                    <td class=\"w_105\">東京国際郵便局</td>\n                    <td rowspan=\"2\" class=\"w_105\">東京都</td>\n                  </tr>\n                  <tr>\n                    <td class=\"w_105\">138-8799</td>\n                  </tr>\n              \n                  <tr>\n                    <td rowspan=\"2\" class=\"w_120\">2021/06/18 13:32</td>\n                    <td rowspan=\"2\" class=\"w_150\">保税運送到着</td>\n                    <td rowspan=\"2\" class=\"w_180\"></td>\n                    <td class=\"w_105\">新福岡郵便局</td>\n                    <td rowspan=\"2\" class=\"w_105\">福岡県</td>\n                  </tr>\n                  <tr>\n                    <td class=\"w_105\">811-8799</td>\n                  </tr>\n              \n                  <tr>\n                    <td rowspan=\"2\" class=\"w_120\">2021/06/18 13:40</td>\n                    <td rowspan=\"2\" class=\"w_150\">通関手続中</td>\n                    <td rowspan=\"2\" class=\"w_180\"></td>\n                    <td class=\"w_105\">新福岡郵便局</td>\n                    <td rowspan=\"2\" class=\"w_105\">福岡県</td>\n                  </tr>\n                  <tr>\n                    <td class=\"w_105\">811-8799</td>\n                  </tr>\n              \n                  <tr>\n                    <td rowspan=\"2\" class=\"w_120\">2021/06/18 14:12</td>\n                    <td rowspan=\"2\" class=\"w_150\">通関手続中</td>\n                    <td rowspan=\"2\" class=\"w_180\"></td>\n                    <td class=\"w_105\">新福岡郵便局</td>\n                    <td rowspan=\"2\" class=\"w_105\">福岡県</td>\n                  </tr>\n                  <tr>\n                    <td class=\"w_105\">811-8799</td>\n                  </tr>\n              \n                  <tr>\n                    <td rowspan=\"2\" class=\"w_120\">2021/06/19 10:26</td>\n                    <td rowspan=\"2\" class=\"w_150\">通関手続中</td>\n                    <td rowspan=\"2\" class=\"w_180\"></td>\n                    <td class=\"w_105\">新福岡郵便局</td>\n                    <td rowspan=\"2\" class=\"w_105\">福岡県</td>\n                  </tr>\n                  <tr>\n                    <td class=\"w_105\">811-8799</td>\n                  </tr>\n              \n                  <tr>\n                    <td rowspan=\"2\" class=\"w_120\">2021/06/19 16:46</td>\n                    <td rowspan=\"2\" class=\"w_150\">植物検疫中</td>\n                    <td rowspan=\"2\" class=\"w_180\"></td>\n                    <td class=\"w_105\">新福岡郵便局</td>\n                    <td rowspan=\"2\" class=\"w_105\">福岡県</td>\n                  </tr>\n                  <tr>\n                    <td class=\"w_105\">811-8799</td>\n                  </tr>\n              \n            </tbody></table>\n            \n            <br>\n            \n            \n            \n            \n            \n              <p class=\"txt_type02\">\n<font color=\"red\">\n※配達を担当する郵便局に到着等した際に履歴を更新いたします。<br>\n　 輸送中などにより履歴情報の更新までお時間をいただく場合がございます。<br>\n　 その他、よくあるご質問・お問い合わせについては\n<a href=\"https://www.post.japanpost.jp/question/haisou/category01.html\">こちら。</a>\n</font>\n\n<br><br></p>\n\n<div class=\"beige_box m_b15\">\n    <h2 class=\"beige_box_inner txt_type06 bold\">お客様サービス相談センター</h2>\n</div>\n\n<table class=\"tableType01 txt_c m_b5\" summary=\"お客様サービス相談センター\">\n    <tbody>\n    <tr>\n      <th class=\"bg01 w_300\">固定電話から<br>(フリーダイヤル)</th>\n      <th class=\"bg01 w_300\">携帯電話から<br>(通話料有料)</th>\n      <th class=\"bg01 w_300\">英語受付(For English)<br>(通話料有料 Chargeable call)</th>\n    </tr>\n    <tr>\n      <td class=\"w_300\">0120-23-28-86</td>\n      <td class=\"w_300\">0570-046-666</td>\n      <td class=\"w_300\">0570-046-111</td>\n    </tr>\n  </tbody>\n</table>\n\n<p class=\"txt_type02\">　受付時間 平日 8:00～21:00 土・日・休日 9:00～21:00</p><br>\n            \n            \n            \n            \n            \n            \n            \n            \n            \n        </div>\n    \n        \n        \n        \n    </form>\n\n    \n    <div class=\"indent\">\n<ul class=\"w_380 m_auto m_b30 clearfix\">\n<li class=\"float_l\"><a href=\"/services/srv/search/\"><img src=\"/services/common/displayImage/searchDisplayImage?registeredIdentifyId=480404BI1P30ja02202101251410&amp;imageFileId=btn_service_top.gif\" alt=\"追跡サービスTOP\"></a></li>\n<li class=\"float_r\"><a href=\"http://www.post.japanpost.jp/office_search/index.html\"><img src=\"/services/common/displayImage/searchDisplayImage?registeredIdentifyId=480404BI1P30ja02202101251410&amp;imageFileId=btn_search_shop.gif\" alt=\"取扱店を調べる\"></a></li></ul>\n<ul class=\"w_380 m_auto m_b30 clearfix\">\n<li class=\"float_l\">\n  <a href=\"/services/srv/search/input\"><img src=\"/services/common/displayImage/searchDisplayImage?registeredIdentifyId=480404BI1P30ja02202101251410&amp;imageFileId=btn_num_search01.gif\" alt=\"個別番号検索\"></a>\n</li>\n<li class=\"float_r\">\n  <a href=\"/services/srv/sequenceNoSearch/input\"><img src=\"/services/common/displayImage/searchDisplayImage?registeredIdentifyId=480404BI1P30ja02202101251410&amp;imageFileId=btn_num_search02.gif\" alt=\"連続番号検索\"></a>\n</li>\n</ul>\n<div class=\"beige_box m_b15\">\n<h2 class=\"beige_box_inner txt_type06 bold\">郵便物</h2>\n</div>\n<div class=\"clearfix dot_center m_b20\">\n<div class=\"float_l w_328\">\n  <div class=\"clearfix p_l20 p_r20 m_b15\">\n   <ul class=\"li_mb5 float_l\">\n    <li class=\"icon_radius_r\">一般書留</li>\n    <li class=\"icon_radius_r\">現金書留</li>\n    <li class=\"icon_radius_r\">簡易書留</li>\n    <li class=\"icon_radius_r\">特定記録郵便</li>\n   </ul>\n   <p class=\"float_r\"><img src=\"/services/common/displayImage/searchDisplayImage?registeredIdentifyId=480404BI1P30ja02202101251410&amp;imageFileId=img_service01.gif\" alt=\"郵便物のイメージ\"></p>\n  </div>\n</div>\n<div class=\"float_r w_328\">\n  <div class=\"clearfix p_l20 p_r20 m_b15\">\n   <ul class=\"li_mb5 float_l\">\n    <li class=\"icon_radius_r\">レターパック</li>\n    <li class=\"icon_radius_r\">レタックス</li>\n    <li class=\"icon_radius_r\">配達時間帯指定郵便</li>\n    <li class=\"icon_radius_r\">新特急郵便</li>\n   </ul>\n   <p class=\"float_r\"><img src=\"/services/common/displayImage/searchDisplayImage?registeredIdentifyId=480404BI1P30ja02202101251410&amp;imageFileId=img_service02.gif\" alt=\"郵便物のイメージ\"></p>\n  </div>\n</div>\n</div>\n<div class=\"clearfix m_b20\">\n<div class=\"float_l w_328\">\n<div class=\"beige_box_h m_b15\">\n   <h3 class=\"beige_box_h_inner txt_type06 bold\">荷物</h3>\n  </div>\n  <div class=\"clearfix p_l20 p_r20 m_b15\">\n   <ul class=\"li_mb5 float_l\">\n    <li class=\"icon_radius_r\">ゆうパック</li>\n    <li class=\"icon_radius_r\">ゆうパケット</li>\n    <li class=\"icon_radius_r\">クリックポスト</li>\n   </ul>\n   <p class=\"float_r\"><img src=\"/services/common/displayImage/searchDisplayImage?registeredIdentifyId=480404BI1P30ja02202101251410&amp;imageFileId=img_service03.gif\" alt=\"荷物のイメージ\"></p>\n  </div>\n</div>\n<div class=\"float_r w_328\">\n  <div class=\"beige_box_h m_b15\">\n   <h3 class=\"beige_box_h_inner txt_type06 bold\">国際郵便物</h3>\n  </div>\n  <div class=\"clearfix p_l20 p_r20 m_b15\">\n   <ul class=\"li_mb5 float_l\">\n    <li class=\"icon_radius_r\">EMS</li>\n    <li class=\"icon_radius_r\">国際小包</li>\n    <li class=\"icon_radius_r\">国際書留・保険付</li>\n    <li><br></li>\n   </ul>\n   <p class=\"float_r\"><img src=\"/services/common/displayImage/searchDisplayImage?registeredIdentifyId=480404BI1P30ja02202101251410&amp;imageFileId=img_service04.gif\" alt=\"国際郵便物のイメージ\"></p>\n  </div>\n</div>\n</div>\n<ul class=\"w_380 m_auto m_b30 clearfix\">\n<li class=\"float_l\">\n  <a href=\"/services/srv/search/input\"><img src=\"/services/common/displayImage/searchDisplayImage?registeredIdentifyId=480404BI1P30ja02202101251410&amp;imageFileId=btn_num_search01.gif\" alt=\"個別番号検索\"></a>\n</li>\n<li class=\"float_r\">\n  <a href=\"/services/srv/sequenceNoSearch/input\"><img src=\"/services/common/displayImage/searchDisplayImage?registeredIdentifyId=480404BI1P30ja02202101251410&amp;imageFileId=btn_num_search02.gif\" alt=\"連続番号検索\"></a>\n</li>\n</ul>\n<div class=\"beige_box m_b20\"><div class=\"beige_box_inner\"><ul class=\"disc li_mb5\"><li>EMS、国際小包、国際書留・保険付について、一部の国・地域は試験接続中です。\n<br>配達状況確認が可能な国・地域は<a href=\"http://www.post.japanpost.jp/int/ems/delivery/ems_search.html\">こちら<span class=\"hidden\">\n：配達状況確認が可能な国・地域へのリンク</span></a>をご覧ください。</li><li>追跡を行える期間は、郵便物をお取扱いしてから約100日間、国際郵便物は約12ヵ月間です。</li>\n<li class=\"m_b0\">普通郵便の「はがき・封書」についてはお取り扱いしておりません。</li></ul></div></div><ul class=\"clearfix m_b10\"><li class=\"w_320 float_l clearfix\">\n<p class=\"float_l\"><img src=\"/services/common/displayImage/searchDisplayImage?registeredIdentifyId=480404BI1P30ja02202101251410&amp;imageFileId=img_service05.gif\" alt=\"連続番号検索\"></p><div class=\"w_170 float_r\">\n\n\n<p class=\"m_b20\">配達完了電子メール通知サービスのお知らせ</p><p class=\"txt_r\">\n<a href=\"http://www.post.japanpost.jp/oshirase/index.html\">詳しくはこちら\n</a></p></div></li><li class=\"w_320 float_r clearfix\"><p class=\"float_l\">\n<img src=\"/services/common/displayImage/searchDisplayImage?registeredIdentifyId=480404BI1P30ja02202101251410&amp;imageFileId=img_service06.gif\" alt=\"連続番号検索\"></p>\n<div class=\"w_170 float_r\"><p class=\"m_b20\">スマートフォンでも配達状況を調べることできます。\n</p><p class=\"txt_r\"><a href=\"https://trackings.post.japanpost.jp/services/sp/srv/search/input\">詳しくはこちら\n</a></p></div></li></ul><ul class=\"clearfix m_b20\"><li class=\"w_320 float_l clearfix\">\n<p class=\"float_l\"><img src=\"/services/common/displayImage/searchDisplayImage?registeredIdentifyId=480404BI1P30ja02202101251410&amp;imageFileId=img_service07.gif\" alt=\"連続番号検索\"></p><div class=\"w_170 float_r\">\n<p class=\"m_b10\">検索表示結果のデータをCSV形式のファイルによりダウンロードができます。</p>\n<p class=\"txt_r\"><a href=\"http://www.post.japanpost.jp/tsuiseki/download.html\">詳しくはこちら\n</a></p></div></li><li class=\"w_320 float_r clearfix\"><p class=\"float_l\"><img src=\"/services/common/displayImage/searchDisplayImage?registeredIdentifyId=480404BI1P30ja02202101251410&amp;imageFileId=img_service08.gif\" alt=\"連続番号検索\"></p>\n<div class=\"w_170 float_r\"><p>郵便追跡システムには、追跡データをCSV形式等のファイルで 提供する各種サービスがあります。\n</p><p class=\"txt_r\"><a href=\"http://www.post.japanpost.jp/tsuiseki/tuiseki_teikyo.html\">詳しくはこちら\n</a></p></div></li></ul><p class=\"pagetop\"><a href=\"#wrap\">▲ ページトップ</a></p></div>\n\n    \n</div>")

	doc, err := goquery.NewDocumentFromReader(r)
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
		"詳細",
		"取扱局",
		"郵便番号",
		"県名・国名",
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
	//var dates, statuses, details []string
	content.Find(".tableType01.txt_c.m_b5 > tbody > tr").Each(func(i int, s *goquery.Selection) {
		// 状態発生日
		date := s.Find("td.w_120").Text()
		if len(date) > 0 && date != "" {
			deliveryStatus = append(deliveryStatus, date)
		}

		// 配送履歴
		status := s.Find("td.w_150").Text()
		if len(status) > 0 && status != "" {
			deliveryStatus = append(deliveryStatus, status)
		}

		// 詳細
		//detail := s.Find("td.w_105").Text()
		//if len(detail) > 0 && detail != "" {
		//	//items = append(items, detail)
		//	details = append(details, detail)
		//}
	})

	var data [][]string
	var items = make([]string, 0, 5)
	for _, v := range deliveryStatus {
		if len(items) < 2 {
			items = append(items, v)
		}

		if len(items) == 2 {
			items = append([]string{id, itemType, additionalService}, items...)
			data = append(data, items)

			items = nil
		}
	}

	data = append(data, items)

	return data
}