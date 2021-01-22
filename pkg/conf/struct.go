package conf

import (
	//"os"
	"github.com/fengjijiao/wechat-push-restful-api/pkg/logio"
)

type ConfInfo struct {
	WorkDir string `yaml:"work-dir"`
	WechatAppId string `yaml:"app-id"`
	WechatAppSecret string `yaml:"app-secret"`
	HttpServerListen string `yaml:"http-server-listen"`
	BaseUrlPath string `yaml:"base-url-path"`
	SecurityPrefix string `yaml:"security-prefix"`
	WechatToken string `yaml:"token"`
	WechatTemplateId string `yaml:"template-id"`
	WechatOpenId string `yaml:"open-id"`
}

func (ci *ConfInfo) setDefaults() {
	if ci.WorkDir == "" {
		ci.WorkDir = "./"
	}
	if ci.WechatAppId == "" {
		logio.Logger.Fatal("[setDefaults]: WechatAppId can not be empty!")
	}
	if ci.WechatAppSecret == "" {
		logio.Logger.Fatal("[setDefaults]: WechatAppSecret can not be empty!")
	}
	if ci.HttpServerListen == "" {
		ci.HttpServerListen = ":9465"
	}
	if ci.WechatToken == "" {
		logio.Logger.Fatal("[setDefaults]: WechatToken can not be empty!")
	}
	if ci.WechatTemplateId == "" {
		logio.Logger.Fatal("[setDefaults]: WechatTemplateId can not be empty!")
	}
	if ci.WechatOpenId == "" {
		logio.Logger.Fatal("[setDefaults]: WechatOpenId can not be empty!")
	}
}