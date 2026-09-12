package mp

import (
	"errors"
	"github.com/cnlisea/ant/typex"
	"net/http"
	"strings"
)

func (w *WeChatMp) ManualLogin(code string) (*UserInfo, error) {
	var s strings.Builder
	s.WriteString("https://api.weixin.qq.com/sns/oauth2/access_token?appid=")
	s.WriteString(w.AppId)
	s.WriteString("&secret=")
	s.WriteString(w.AppSecret)
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

	// 从token获取用户信息
	s.WriteString("https://api.weixin.qq.com/sns/userinfo?access_token=")
	s.WriteString(code2token.AccessToken)
	s.WriteString("&openid=")
	s.WriteString(code2token.OpenId)
	s.WriteString("&lang=zh_CN")
	res, err = http.Get(s.String())
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	userInfo := new(UserInfo)
	if err = typex.JsonUnmarshalReader(res.Body, userInfo); err != nil {
		return nil, err
	}

	if userInfo.ErrCode != 0 {
		return nil, errors.New(userInfo.ErrMsg)
	}

	return userInfo, nil
}

type Code2Token struct {
	AccessToken string `json:"access_token"`
	OpenId      string `json:"openid"`
	ErrCode     int64  `json:"errcode"` // 错误码
	ErrMsg      string `json:"errmsg"`  // 错误信息
}
