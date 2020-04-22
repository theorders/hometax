package hometax

import (
	"errors"
	"github.com/imroc/req"
	"github.com/theorders/aefire"
	"regexp"
)

func (c *Client) SsoToken() (ssoToken, userClcd string, err error) {
	res, err := req.Post(
		"https://hometax.go.kr/token.do?query=_dShuaODDu5ZY0dDeAHOa&postfix=2020_02_25",
		c.httpCli,
		req.BodyXML(`<map id='postParam'><popupYn></popupYn></map>`))

	if aefire.LogIfError(err) {
		return "", "", errors.New("홈택스 권한 획득이 실패했습니다")
	}

	matches := regexp.MustCompile(`<ssoToken>(.*)<\/ssoToken><userClCd>(.*)<\/userClCd>`).
		FindStringSubmatch(res.String())

	if len(matches) < 3 {
		return "", "", errors.New("홈택스 권한 획득이 실패했습니다 : 홈택스 응답 해석 실패")
	}

	return matches[1], matches[2], nil
}
