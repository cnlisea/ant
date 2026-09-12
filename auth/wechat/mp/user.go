package mp

import (
	"errors"
	"github.com/cnlisea/ant/typex"
	"net/http"
	"strings"
	"time"
)

type UserInfo struct {
	Subscribe  int    `json:"subscribe"`    // 户是否订阅该公众号标识，值为0时，代表此用户没有关注该公众号，拉取不到其余信息。
	OpenId     string `json:"openid"`       // 用户的标识，对当前公众号唯一
	Nickname   string `json:"nickname"`     // 用户的昵称
	Sex        int    `json:"sex"`          // 性别
	City       string `json:"city"`         // 用户所在城市
	Country    string `json:"country"`      // 用户所在国家
	Province   string `json:"province"`     // 用户所在省份
	Language   string `json:"language"`     // 用户的语言
	HeadImgUrl string `json:"headimgurl"`   // 用户头像
	UnionId    string `json:"unionid"`      // 用户开放平台唯一标识
	QrScene    int64  `json:"qr_scene"`     // 二维码扫码场景
	QrSceneStr string `json:"qr_scene_str"` // 二维码扫码场景描述
	ErrCode    int    `json:"errcode"`      // 错误码
	ErrMsg     string `json:"errmsg"`       // 错误信息
}

func (w *WeChatMp) UserInfo(openId string) (*UserInfo, error) {
	if w.accessToken == "" || w.expiresTime < time.Now().Unix() {
		accessToken, expiresIn, err := w.AccessToken()
		if err != nil {
			return nil, err
		}

		w.accessToken = accessToken
		w.expiresTime = time.Now().Unix() + expiresIn - 30
	}

	var s strings.Builder
	s.WriteString("https://api.weixin.qq.com/cgi-bin/user/info?access_token=")
	s.WriteString(w.accessToken)
	s.WriteString("&openid=")
	s.WriteString(openId)
	s.WriteString("&lang=zh_CN")

	client := &http.Client{Timeout: 3 * time.Second}
	res, err := client.Get(s.String())
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var userInfoData = new(UserInfo)
	if err = typex.JsonUnmarshalReader(res.Body, userInfoData); err != nil {
		return nil, err
	}

	if userInfoData.ErrCode != 0 {
		return nil, errors.New(userInfoData.ErrMsg)
	}

	return userInfoData, err
}
