package mp

type WeChatMp struct {
	AppId       string // 凭证
	AppSecret   string // 凭证密钥
	accessToken string // 调用凭据
	expiresTime int64  // 过期时间戳
}

func NewWeChatMp(appId string, appSecret string) *WeChatMp {
	return &WeChatMp{
		AppId:     appId,
		AppSecret: appSecret,
	}
}
