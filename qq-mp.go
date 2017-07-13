package mp

import (
	"github.com/cocotyty/httpclient"
	"log"
	"encoding/json"
)

// QQ公众号

type qqMP struct {
	AccessToken string
}

func NewQQMP(token string) *qqMP {
	return &qqMP{AccessToken: token}
}

func (handler *qqMP) SendTemplateMessage(req *QQSendTemplateMessageReq) (resp *QQSendTemplateMessageResp, err error) {
	resp = &QQSendTemplateMessageResp{}
	bs, err := json.Marshal(req)
	log.Println(string(bs), err)
	err = httpclient.Post("https://api.uni.qq.com/cgi-bin/message/template/send").
		Query("access_token", handler.AccessToken).
		JSON(req).
		Send().
		JSON(resp)
	return
}
