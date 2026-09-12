package mp

import (
	"net/url"
	"strings"
)

func (w *WeChatMp) AutomaticAuthorizeUrl(u string) string {
	var s strings.Builder

	s.WriteString("https://open.weixin.qq.com/connect/oauth2/authorize?appid=")
	s.WriteString(w.AppId)
	s.WriteString("&redirect_uri=")
	s.WriteString(url.QueryEscape(u))
	s.WriteString("&response_type=code&scope=snsapi_base#wechat_redirect")

	return s.String()
}
