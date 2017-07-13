package mp

type QQSendTemplateMessageReq struct {
	Tousername string             `json:"tousername"`
	Templateid string             `json:"templateid"`
	Data       map[string]QQValue `json:"data"`
	Button     map[string]QQValue `json:"button"`
}

type QQValue struct {
	Value string `json:"value"`
}

type QQSendTemplateMessageResp struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	Msgid   string `json:"msgid"`
}
