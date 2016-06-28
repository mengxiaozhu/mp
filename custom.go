package mp

import (
	respErr "h5/response/error"
	"net/url"
)

// 可以发送的客服消息
type CustomMessageRequest interface {
	Send(*Mp) (*WechatApiError, respErr.Response)
}

// 客服消息
type CustomMessageHeader struct {
	ToUser  string `json:"touser"`
	MsgType string `json:"msgtype"`
}

// 文本类型的客服消息内容
type TextMessage struct {
	Content string `json:"content"`
}

// 文本类型的客服消息请求
type TextCustomMessageRequest struct {
	CustomMessageHeader
	Text TextMessage `json:"text"`
}

func (t *TextCustomMessageRequest) Send(m *Mp) (*WechatApiError, respErr.Response) {
	return m.SendCustomMessage(t)
}

// 发送客服消息
func (m *Mp) SendCustomMessage(req *TextCustomMessageRequest) (resp *WechatApiError, err respErr.Response) {
	resp = &WechatApiError{}
	err = m.Cgi(resp, "/cgi-bin/message/custom/send", url.Values{}, req)
	return
}
