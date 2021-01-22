package httphandler

import (
	"fmt"
    "net/http"
	"github.com/fengjijiao/wechat-push-restful-api/pkg/conf"
	"github.com/fengjijiao/wechat-push-restful-api/pkg/logio"
	"go.uber.org/zap"
	"path"
	"github.com/fengjijiao/wechat-push-restful-api/pkg/sqlhandler"
	"encoding/json"
	"github.com/imroc/req"
)

type SendInfo struct {
	Touser     string `json:"touser"`
	TemplateID string `json:"template_id"`
	URL        string `json:"url"`
	Topcolor   string `json:"topcolor"`
	Data       struct {
		Text struct {
			Value string `json:"value"`
			Color string `json:"color"`
		} `json:"text"`
	} `json:"data"`
}

type ErrorInfo struct {
	ErrCode int `json:"errcode"`
	ErrMsg string `json:"errmsg"`
}

func sendHttpHandler(w http.ResponseWriter, hr *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var res ErrorInfo
	var sendInfo SendInfo
	sendInfo.Touser = conf.Config.WechatOpenId
	sendInfo.TemplateID = conf.Config.WechatTemplateId
	sendInfo.URL = path.Join("https://", hr.Header.Get("Host"), "/")
	sendInfo.Topcolor = "#FF0000"
	sendInfo.Data.Text.Value = hr.PostFormValue("context")
	sendInfo.Data.Text.Color = "#173177"
	param, err := json.Marshal(&sendInfo)
	if err != nil {
		json.NewEncoder(w).Encode(&ErrorInfo{-1, err.Error()})
		logio.Logger.Error("sendHttpHandler: ", zap.Error(err))
		return
	}
	token, err := sqlhandler.GetToken()
	if err != nil {
		json.NewEncoder(w).Encode(&ErrorInfo{-1, err.Error()})
		logio.Logger.Error("sendHttpHandler: ", zap.Error(err))
		return
	}
	r, err := req.Post(fmt.Sprintf(`https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=%s`, token), param)
	if err != nil {
		json.NewEncoder(w).Encode(&ErrorInfo{-1, err.Error()})
		logio.Logger.Error("sendHttpHandler: ", zap.Error(err))
		return
	}
	r.ToJSON(&res)
	if res.ErrCode == 0 {
		json.NewEncoder(w).Encode(&ErrorInfo{0, "send message success!"})
	}else {
		json.NewEncoder(w).Encode(&ErrorInfo{-1, "send message failed!"})
	}
}