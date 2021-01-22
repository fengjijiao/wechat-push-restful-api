package commonio

import (
	"github.com/fengjijiao/wechat-push-restful-api/pkg/conf"
)

func GetTokenAuto() (*AccessTokenInfo, error) {
	return GetToken(conf.Config.WechatAppId, conf.Config.WechatAppSecret)
}