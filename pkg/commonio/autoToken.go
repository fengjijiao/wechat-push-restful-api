package commonio

import (
	"../../../wechat-push-restful-api/pkg/conf"
)

func GetTokenAuto() (*AccessTokenInfo, error) {
	return GetToken(conf.Config.WechatAppId, conf.Config.WechatAppSecret)
}