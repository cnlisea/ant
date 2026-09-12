package mp

import (
	"errors"
	jsoniter "github.com/json-iterator/go"
	"net/http"
	"strings"
	"time"
)

type AccessTokenInfo struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"` // 有效时间
	ErrCode     int    `json:"errcode"`    // 错误码
	ErrMsg      string `json:"errmsg"`     // 错误信息
}

func (w *WeChatMp) AccessToken() (string, int64, error) {
	var s strings.Builder
	s.WriteString("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=")
	s.WriteString(w.AppId)
	s.WriteString("&secret=")
	s.WriteString(w.AppSecret)

	// settings 3 second request timeout
	client := &http.Client{Timeout: time.Second * 3}
	res, err := client.Get(s.String())
	if err != nil {
		return "", 0, err
	}
	defer res.Body.Close()

	var accessTokenInfo AccessTokenInfo
	if err = jsoniter.NewDecoder(res.Body).Decode(&accessTokenInfo); err != nil {
		return "", 0, err
	}

	if accessTokenInfo.ErrCode != 0 {
		return "", 0, errors.New(accessTokenInfo.ErrMsg)
	}

	return accessTokenInfo.AccessToken, accessTokenInfo.ExpiresIn, nil
}
