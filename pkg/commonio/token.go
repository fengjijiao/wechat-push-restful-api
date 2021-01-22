package commonio

import (
	"fmt"
	"errors"
	"github.com/imroc/req"
)

type AccessTokenInfo struct {
	AccessToken string `json:"access_token"`
	ExpiresIn int `json:"expires_in", default:0`
	ErrCode int `json:"errcode", default:0`
	ErrMsg string `json:"errmsg"`
}

func GetToken(appId, appSecret string) (*AccessTokenInfo, error) {
	var accessTokenInfo AccessTokenInfo
	r, err := req.Get(fmt.Sprintf(`https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s`, appId, appSecret))
	if err != nil {
		return nil, err
	}
	r.ToJSON(&accessTokenInfo)
	if accessTokenInfo.ErrCode != 0 {
		return nil, errors.New(accessTokenInfo.ErrMsg)
	}
	return &accessTokenInfo, nil
}