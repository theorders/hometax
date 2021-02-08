package hometax

import (
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/imroc/req"
	"github.com/theorders/aefire"
	"github.com/theorders/popbill"
	"net/url"
	"regexp"
	"strings"
)

type User struct {
	UserId string `json:"userId"`
	Name   string `json:"name"`
	RegNum string `json:"regNum"`

	BizList []*CorpInfo `json:"bizList"`
}

type CorpInfo struct {
	popbill.CorpInfo
	*popbill.CloseDown `json:"closeDown"`
}

type EncPasswordRes struct {
	XMLName xml.Name `xml:"map"`
	Text    string   `xml:",chardata"`
	ID      string   `xml:"id,attr"`
	Pswd    string   `xml:"pswd"`
	Map     struct {
		Text      string `xml:",chardata"`
		ID        string `xml:"id,attr"`
		DetailMsg string `xml:"detailMsg"`
		Msg       string `xml:"msg"`
		Code      string `xml:"code"`
		Result    string `xml:"result"`
	} `xml:"map"`
}

type LoginRes struct {
	Code         string      `json:"code"`
	ErrCode      interface{} `json:"errCode"`
	ErrMsg       string      `json:"errMsg"`
	LgnRsltCd    string      `json:"lgnRsltCd"`
	PswdErrNbcnt string      `json:"pswdErrNbcnt"`
	Tin          string      `json:"tin"`
	SecCardID    *string     `json:"secCardId"`
}

type UserInfoRes struct {
	XMLName xml.Name `xml:"map"`
	Text    string   `xml:",chardata"`
	ID      string   `xml:"id,attr"`
	Map     struct {
		Text      string `xml:",chardata"`
		ID        string `xml:"id,attr"`
		DetailMsg string `xml:"detailMsg"`
		Msg       string `xml:"msg"`
		Map       struct {
			Text                   string `xml:",chardata"`
			ID                     string `xml:"id,attr"`
			TxprDscmNo             string `xml:"txprDscmNo"`
			BmanOfbDt              string `xml:"bmanOfbDt"`
			EtxivPkcYn             string `xml:"etxivPkcYn"`
			TxofOgzCd              string `xml:"txofOgzCd"`
			LgnUserClCd            string `xml:"lgnUserClCd"`
			TxprDscmNoCnfrYn       string `xml:"txprDscmNoCnfrYn"`
			HaboCl                 string `xml:"haboCl"`
			NtplBmanAthYn          string `xml:"ntplBmanAthYn"`
			CharId                 string `xml:"charId"`
			SmprYn                 string `xml:"smprYn"`
			PubcUserNo             string `xml:"pubcUserNo"`
			ThofOgzCd              string `xml:"thofOgzCd"`
			Tin                    string `xml:"tin"`
			RprsTin                string `xml:"rprsTin"`
			BmanUnitEngeTrerJnngYn string `xml:"bmanUnitEngeTrerJnngYn"`
			RprsResno              string `xml:"rprsResno"`
			TxprClsfCd             string `xml:"txprClsfCd"`
			RprsFnm                string `xml:"rprsFnm"`
			UserClsfCd             string `xml:"userClsfCd"`
			DprtUserYn             string `xml:"dprtUserYn"`
			TxpAgnYn               string `xml:"txpAgnYn"`
			CnvrTin                string `xml:"cnvrTin"`
			AfaTxprYn              string `xml:"afaTxprYn"`
			SsnAltPsbYn            string `xml:"ssnAltPsbYn"`
			MpbNo                  string `xml:"mpbNo"`
			LgnClientIp            string `xml:"lgnClientIp"`
			UserId                 string `xml:"userId"`
			ChrgDutsCd             string `xml:"chrgDutsCd"`
			DataMaagClCd           string `xml:"dataMaagClCd"`
			CrpBmanAthYn           string `xml:"crpBmanAthYn"`
			LgnCertCd              string `xml:"lgnCertCd"`
			TxaaYn                 string `xml:"txaaYn"`
			UserCertClCd           string `xml:"userCertClCd"`
			WhlPmtMnpt             string `xml:"whlPmtMnpt"`
			NtplAthYn              string `xml:"ntplAthYn"`
			UserNm                 string `xml:"userNm"`
		} `xml:"map"`
		Code   string `xml:"code"`
		Result string `xml:"result"`
	} `xml:"map"`
	ScrnId    string `xml:"scrnId"`
	OriginTin string `xml:"originTin"`
}

type BusinessListRes struct {
	XMLName xml.Name `xml:"map"`
	Text    string   `xml:",chardata"`
	ID      string   `xml:"id,attr"`
	Map     struct {
		Text      string `xml:",chardata"`
		ID        string `xml:"id,attr"`
		DetailMsg string `xml:"detailMsg"`
		Msg       string `xml:"msg"`
		Map       struct {
			Text                   string `xml:",chardata"`
			ID                     string `xml:"id,attr"`
			TxprDscmNo             string `xml:"txprDscmNo"`
			BmanOfbDt              string `xml:"bmanOfbDt"`
			EtxivPkcYn             string `xml:"etxivPkcYn"`
			TxofOgzCd              string `xml:"txofOgzCd"`
			LgnUserClCd            string `xml:"lgnUserClCd"`
			TxprDscmNoCnfrYn       string `xml:"txprDscmNoCnfrYn"`
			HaboCl                 string `xml:"haboCl"`
			NtplBmanAthYn          string `xml:"ntplBmanAthYn"`
			CharId                 string `xml:"charId"`
			SmprYn                 string `xml:"smprYn"`
			PubcUserNo             string `xml:"pubcUserNo"`
			ThofOgzCd              string `xml:"thofOgzCd"`
			Tin                    string `xml:"tin"`
			BmanUnitEngeTrerJnngYn string `xml:"bmanUnitEngeTrerJnngYn"`
			TxprClsfCd             string `xml:"txprClsfCd"`
			UserClsfCd             string `xml:"userClsfCd"`
			DprtUserYn             string `xml:"dprtUserYn"`
			TxpAgnYn               string `xml:"txpAgnYn"`
			CnvrTin                string `xml:"cnvrTin"`
			AfaTxprYn              string `xml:"afaTxprYn"`
			SsnAltPsbYn            string `xml:"ssnAltPsbYn"`
			MpbNo                  string `xml:"mpbNo"`
			LgnClientIp            string `xml:"lgnClientIp"`
			UserId                 string `xml:"userId"`
			ChrgDutsCd             string `xml:"chrgDutsCd"`
			DataMaagClCd           string `xml:"dataMaagClCd"`
			CrpBmanAthYn           string `xml:"crpBmanAthYn"`
			LgnCertCd              string `xml:"lgnCertCd"`
			TxaaYn                 string `xml:"txaaYn"`
			UserCertClCd           string `xml:"userCertClCd"`
			WhlPmtMnpt             string `xml:"whlPmtMnpt"`
			NtplAthYn              string `xml:"ntplAthYn"`
			UserNm                 string `xml:"userNm"`
		} `xml:"map"`
		Code   string `xml:"code"`
		Result string `xml:"result"`
	} `xml:"map"`
	List struct {
		Text string        `xml:",chardata"`
		ID   string        `xml:"id,attr"`
		Map  []CorpInfoSum `xml:"map"`
	} `xml:"list"`
}

type CorpInfoSum struct {
	Text              string `xml:",chardata"`
	ID                string `xml:"id,attr"`
	TxprStatNm        string `xml:"txprStatNm"`
	Tin               string `xml:"tin"`
	StatusValue       string `xml:"statusValue"`
	TxprDscmDt        string `xml:"txprDscmDt"`
	TxprNm            string `xml:"txprNm"`
	TxprDscmNoEncCntn string `xml:"txprDscmNoEncCntn"`
}

//홈택스의 패스워드 암호화
func (c *Client) EncPassword(pw string) (encPw *string, err error) {
	body := fmt.Sprintf(
		`<map id="ATXPPABA001R07"><pswd>%s</pswd></map><nts<nts>nts>%s`,
		base64.StdEncoding.EncodeToString([]byte(pw)),
		aefire.RandString(46))

	r, err := req.Post(
		"https://www.hometax.go.kr/wqAction.do?actionId=ATXPPABA001R07&screenId=UTXPPABA01&popupYn=false&realScreenId=",
		c.httpCli,
		req.BodyXML(body))

	if err != nil {
		return nil, err
	}

	res := EncPasswordRes{}

	err = r.ToXML(&res)

	if err != nil {
		return nil, err
	}

	return &res.Pswd, nil
}

func (c *Client) login(id string, pw string) error {
	res, err := req.Post(
		"https://www.hometax.go.kr/pubcLogin.do?domain=hometax.go.kr&mainSys=Y",
		c.httpCli,
		req.Param{
			"ssoLoginYn":        "N",
			"secCardLoginYn":    "",
			"secCardId":         "",
			"cncClCd":           "01",
			"id":                base64.StdEncoding.EncodeToString([]byte(id)),
			"pswd":              base64.StdEncoding.EncodeToString([]byte(pw)),
			"ssoStatus":         "",
			"portalStatus":      "",
			"scrnId":            "UTXPPABA01",
			"userScrnRslnXcCnt": "2560",
			"userScrnRslnYcCnt": "1440",
		},
	)

	if err != nil {
		return err
	}

	resString := res.String()

	//println(resString)

	errMsg := regexp.MustCompile(`decodeURIComponent\('([^\)]*)'\)`).FindStringSubmatch(resString)[1]

	if errMsg != "" {
		decodedMsg, err := url.QueryUnescape(errMsg)
		if err != nil {
			return errors.New("로그인 오류가 발생했습니다 : 홈택스 로그인 오류 메시지 해석 실패")
		}
		return errors.New(decodedMsg)
	}

	resString = strings.ReplaceAll(resString, `'errMsg' : decodeURIComponent('').replace(/\+/g,' ').replace(/\\n/g,'\n'),`, "")
	resString = strings.ReplaceAll(resString, `'`, `"`)

	r, _ := regexp.Compile(`(\{.*\})`)
	j := r.FindString(resString)

	var loginRes LoginRes

	err = json.Unmarshal([]byte(j), &loginRes)

	if aefire.LogIfError(err) {
		return errors.New("로그인 오류가 발생했습니다 : 홈택스 로그인 결과 해석 실패")
	}

	if loginRes.Code != "S" {
		return errors.New("로그인이 실패했습니다")
	}

	return nil
}

func (c *Client) UserInfo() (userInfo *UserInfoRes, err error) {
	res, err := req.Post(
		"https://www.hometax.go.kr/wqAction.do?actionId=ATXPPAAA003A01&screenId=index&popupYn=false&realScreenId=",
		c.httpCli,
		req.BodyXML(`<map id="ATXPPAAA003A01"><scrnId>index</scrnId></map><nts<nts>nts>`+aefire.RandString(42)))

	if aefire.LogIfError(err) {
		return nil, errors.New("사용자 정보 조회가 실패했습니다 : 홈택스 연결 오류 발생")
	}

	userInfo = &UserInfoRes{}
	if aefire.LogIfError(res.ToXML(userInfo)) {
		return nil, errors.New("사용자 정보 조회가 실패했습니다 : 홈택스 응답 해석 실패")
	}

	if userInfo.Map.Result != "S" {
		if userInfo.Map.Msg != "" {
			return nil, errors.New("사용자 정보 조회가 실패했습니다 : " + userInfo.Map.Msg)
		} else {
			//println(res.String())
			return nil, errors.New("사용자 정보 조회가 실패했습니다 : 홈택스 요청 오류 발생")
		}
	}

	return
}

func (c *Client) BusinessList() (businessList *BusinessListRes, err error) {
	res, err := req.Post(
		"https://www.hometax.go.kr/wqAction.do?actionId=ATXPPAAA003R01&screenId=UTXPPAAA24&popupYn=false&realScreenId=",
		c.httpCli,
		req.BodyXML(`<map id="ATXPPAAA003R01"/><nts<nts>nts>21dSQX2EuiDFkIeuufo3EljtcimQB2biyhNinGnLLthE10`+aefire.RandString(46)))

	if aefire.LogIfError(err) {
		return nil, errors.New("사업자 정보 조회가 실패했습니다 : 홈택스 연결 오류 발생")
	}

	businessList = &BusinessListRes{}
	if aefire.LogIfError(res.ToXML(businessList)) {
		return nil, errors.New("사업자 정보 조회가 실패했습니다 : 홈택스 응답 해석 실패")
	}

	if businessList.Map.Result != "S" {
		if businessList.Map.Msg != "" {
			return nil, errors.New("사업자 정보 조회가 실패했습니다 : " + businessList.Map.Msg)
		} else {
			//println(res.String())
			return nil, errors.New("사업자 정보 조회가 실패했습니다 : 홈택스 요청 오류 발생")
		}
	}

	return
}

func (c *Client) Login(id string, pw string) (*User, error) {
	encPw, bizListErr := c.EncPassword(pw)

	if bizListErr != nil {
		return nil, errors.New("로그인 오류가 발생했습니다 : 비밀번호 암호화 실패")
	}

	bizListErr = c.login(id, *encPw)

	if bizListErr != nil {
		return nil, bizListErr
	}

	userInfo, bizListErr := c.UserInfo()
	if bizListErr != nil {
		return nil, errors.New("로그인 오류가 발생했습니다 : " + bizListErr.Error())
	}

	businessList, bizListErr := c.BusinessList()

	//println(aefire.ToJson(businessList, "    "))

	domainUser := User{
		BizList: []*CorpInfo{},
	}

	//사업자
	if bizListErr == nil {
		domainUser.Name = businessList.Map.Map.UserNm
		domainUser.RegNum = businessList.Map.Map.TxprDscmNo
		domainUser.UserId = businessList.Map.Map.UserId

		for _, biz := range businessList.List.Map {
			if strings.Contains(biz.TxprStatNm, "폐업") {
				continue
			}

			cd, _ := c.GetCloseDown(biz.TxprDscmNoEncCntn)

			domainUser.BizList = append(domainUser.BizList, &CorpInfo{
				CorpInfo: popbill.CorpInfo{
					CorpNum:  biz.TxprDscmNoEncCntn,
					CorpName: biz.TxprNm,
					CEOName:  businessList.Map.Map.UserNm,
				},
				CloseDown: cd,
			})
		}
	} else if len(userInfo.Map.Map.TxprDscmNo) == 10 { //법인?
		domainUser.Name = userInfo.Map.Map.RprsFnm
		domainUser.RegNum = userInfo.Map.Map.RprsResno
		domainUser.UserId = userInfo.Map.Map.UserId

		cd, _ := c.GetCloseDown(userInfo.Map.Map.TxprDscmNo)

		domainUser.BizList = append(domainUser.BizList, &CorpInfo{
			CorpInfo: popbill.CorpInfo{
				CorpNum:  userInfo.Map.Map.TxprDscmNo,
				CorpName: userInfo.Map.Map.UserNm,
				CEOName:  userInfo.Map.Map.RprsFnm,
			},
			CloseDown: cd,
		})

	} else { //개인
		domainUser.Name = userInfo.Map.Map.UserNm
		domainUser.RegNum = userInfo.Map.Map.TxprDscmNo
		domainUser.UserId = userInfo.Map.Map.UserId
	}

	return &domainUser, nil
}

func TextToCloseDownState(s string) popbill.CloseDownState {
	if strings.Contains(s, "폐업") {
		return popbill.CloseDownStateClose
	} else if strings.Contains(s, "휴업") {
		return popbill.CloseDownStateDown
	} else {
		return popbill.CloseDownStateInBusiness
	}
}
