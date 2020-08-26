package hometax

import (
	"fmt"
	"github.com/imroc/req"
	"github.com/pkg/errors"
	"github.com/theorders/aefire"
	"github.com/theorders/popbill"
	"regexp"
	"strings"
	"time"
)

type CloseDownConfig struct {
	State struct {
		Date         string `json:"date" firestore:"date"`
		Close        string `json:"close" firestore:"close"`
		Down         string `json:"down" firestore:"down"`
		Unregistered string `json:"unregistered" firestore:"unregistered"`
	} `json:"state" firestore:"state"`

	TaxType struct {
		Date      string `json:"date" firestore:"date"`
		Free      string `json:"free" firestore:"free"`
		NonProfit string `json:"nonProfit" firestore:"nonProfit"`
		Normal    string `json:"normal" firestore:"normal"`
		Simple    string `json:"simple" firestore:"simple"`
	} `json:"taxType" firestore:"taxType"`
}

var homeTaxConfig = CloseDownConfig{
	State: struct {
		Date         string `json:"date" firestore:"date"`
		Close        string `json:"close" firestore:"close"`
		Down         string `json:"down" firestore:"down"`
		Unregistered string `json:"unregistered" firestore:"unregistered"`
	}{
		Date:         "([0-9]{4}-[0-9]{2}-[0-9]{2})",
		Close:        "폐업",
		Down:         "휴업",
		Unregistered: "사업을 하지 않고 있습니다",
	},
	TaxType: struct {
		Date      string `json:"date" firestore:"date"`
		Free      string `json:"free" firestore:"free"`
		NonProfit string `json:"nonProfit" firestore:"nonProfit"`
		Normal    string `json:"normal" firestore:"normal"`
		Simple    string `json:"simple" firestore:"simple"`
	}{
		Date:      "과세유형 전환된 날짜는 ([0-9]{4}년.[0-9]{2}월.[0-9]{2}일)",
		Free:      "면세",
		NonProfit: "단체",
		Normal:    "일반",
		Simple:    "간이",
	},
}

func ParseHomeTaxCloseDown(resBody string, conf *CloseDownConfig) popbill.CloseDown {
	cd := popbill.CloseDown{
		CheckDate: time.Now().In(aefire.KRTimezone()).Format("20060102"),
	}

	regex := regexp.MustCompile(conf.State.Date)
	cd.StateDate = strings.Replace(regex.FindString(resBody), "-", "", -1)

	if strings.Contains(resBody, conf.State.Close) {
		cd.State = popbill.CloseDownStateClose
	} else if strings.Contains(resBody, conf.State.Down) {
		cd.State = popbill.CloseDownStateDown
	} else if strings.Contains(resBody, conf.State.Unregistered) {
		cd.State = popbill.CloseDownStateUnregistered
	} else {
		cd.State = popbill.CloseDownStateInBusiness
	}

	regex = regexp.MustCompile(conf.TaxType.Date)
	m := regex.FindStringSubmatch(resBody)
	if len(m) > 1 {
		cd.TypeDate = m[1]
		cd.TypeDate = strings.Replace(cd.TypeDate, "년", "", -1)
		cd.TypeDate = strings.Replace(cd.TypeDate, "월", "", -1)
		cd.TypeDate = strings.Replace(cd.TypeDate, "일", "", -1)
		cd.TypeDate = strings.Replace(cd.TypeDate, " ", "", -1)
	}

	if strings.Contains(resBody, conf.TaxType.Simple) {
		cd.TaxType = popbill.TaxTypeSimple
	} else if strings.Contains(resBody, conf.TaxType.NonProfit) {
		cd.TaxType = popbill.TaxTypeNonProfit
	} else if strings.Contains(resBody, conf.TaxType.Free) {
		cd.TaxType = popbill.TaxTypeFree
	} else if strings.Contains(resBody, conf.TaxType.Normal) {
		cd.TaxType = popbill.TaxTypeNormal
	}

	return cd
}

func (c *Client) GetCloseDown(corpNum string) (*popbill.CloseDown, error) {
	if !aefire.ValidateCorpNum(corpNum) {
		return nil, errors.New("GetCloseDown: 사업자등록번호 형식이 맞지 않습니다")
	}

	res, err := req.Post(
		"https://teht.hometax.go.kr/wqAction.do?actionId=ATTABZAA001R08&screenId=UTEABAAA13&popupYn=false&realScreenId=",
		c.httpCli,
		req.BodyXML(fmt.Sprintf(`<map id='ATTABZAA001R08'><pubcUserNo/><mobYn>N</mobYn><inqrTrgtClCd>1</inqrTrgtClCd><txprDscmNo>%s</txprDscmNo><dongCode>__MIDDLE__</dongCode><psbSearch>Y</psbSearch><map id='userReqInfoVO'/></map>`, corpNum)))

	if aefire.LogIfError(err) {
		return nil, errors.Wrap(err, "홈택스 연결이 실패했습니다")
	}

	closeDown := ParseHomeTaxCloseDown(res.String(), &homeTaxConfig)
	closeDown.CorpNum = corpNum

	return &closeDown, nil
}
