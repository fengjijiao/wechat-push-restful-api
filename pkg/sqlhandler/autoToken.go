package sqlhandler

import (
	"github.com/fengjijiao/wechat-push-restful-api/pkg/commonio"
)

func UpdateTokenAuto() error {
	aInfo, err := commonio.GetTokenAuto()
	if err != nil {
		return err
	}
	return UpdateToken(aInfo.AccessToken)
}