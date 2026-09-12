package app

import (
	"errors"
	"github.com/cnlisea/ant/typex"
	"net/http"
	"strings"
)

type Code2Token struct {
	AccessToken string `json:"access_token"`
	OpenId      string `json:"openid"`
	ErrCode     int64  `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
}

type Token2UserInfo struct {
	Openid     string `json:"openid"`
	UnionId    string `json:"unionid"`
	Nickname   string `json:"nickname"`
	HeadImgUrl string `json:"headimgurl"`
	Sex        int    `json:"sex"`
	ErrCode    int64  `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

func Login(appId string, appSecret string, code string) (*Token2UserInfo, error) {
	var s strings.Builder
	s.WriteString(BaseUrl)
	s.WriteString("/oauth2/access_token?appid=")
	s.WriteString(appId)
	s.WriteString("&secret=")
	s.WriteString(appSecret)
	s.WriteString("&code=")
	s.WriteString(code)
	s.WriteString("&grant_type=authorization_code")
	res, err := http.Get(s.String())
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var code2token Code2Token
	if err = typex.JsonUnmarshalReader(res.Body, &code2token); err != nil {
		return nil, err
	}

	if code2token.ErrCode != 0 {
		return nil, errors.New(code2token.ErrMsg)
	}

	s.Reset()

	s.WriteString(BaseUrl)
	s.WriteString("/userinfo?access_token=")
	s.WriteString(code2token.AccessToken)
	s.WriteString("&openid=")
	s.WriteString(code2token.OpenId)
	res, err = http.Get(s.String())
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var token2UserInfo Token2UserInfo
	if err = typex.JsonUnmarshalReader(res.Body, &token2UserInfo); err != nil {
		return nil, err
	}

	if token2UserInfo.ErrCode != 0 {
		return nil, errors.New(token2UserInfo.ErrMsg)
	}

	return &token2UserInfo, nil
}
