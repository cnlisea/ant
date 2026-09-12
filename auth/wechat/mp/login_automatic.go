package mp

import (
	"errors"
	"github.com/cnlisea/ant/typex"
	"net/http"
	"strings"
	"time"
)

func (w *WeChatMp) AutomaticLogin(code string) (*UserInfo, error) {
	if w.accessToken == "" || w.expiresTime < time.Now().Unix() {
		accessToken, expiresIn, err := w.AccessToken()
		if err != nil {
			return nil, err
		}

		w.accessToken = accessToken
		w.expiresTime = time.Now().Unix() + expiresIn - 30
	}

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
	s.WriteString("https://api.weixin.qq.com/cgi-bin/user/info?access_token=")
	s.WriteString(w.accessToken)
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
