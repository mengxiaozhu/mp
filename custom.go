package mp

import (
	respErr "github.com/mengxiaozhu/mp/base"
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

// 图文消息
type ArticleMessage struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	PicUrl      string `json:"picurl"`
}
type NewsMessage struct {
	Articles []*ArticleMessage `json:"articles"`
}

type ArticleCustomMessageRequest struct {
	CustomMessageHeader
	News *NewsMessage `json:"news"`
}

func NewArticleCustomMessageRequest() (articleCustomMessageRequest *ArticleCustomMessageRequest) {
	articleCustomMessageRequest = &ArticleCustomMessageRequest{
		CustomMessageHeader: CustomMessageHeader{MsgType: `news`},
		News:                &NewsMessage{},
	}
	return articleCustomMessageRequest
}

func (t *ArticleCustomMessageRequest) Send(m *Mp) (*WechatApiError, respErr.Response) {
	return m.SendCustomMessage(t)
}

func (t *ArticleCustomMessageRequest) To(openid string) *ArticleCustomMessageRequest {
	t.ToUser = openid
	return t
}

func (t *ArticleCustomMessageRequest) Article() *ArticleCustomMessageRequest {
	t.News.Articles = append(t.News.Articles, &ArticleMessage{})
	return t
}

func (t *ArticleCustomMessageRequest) Title(title string) *ArticleCustomMessageRequest {
	t.News.Articles[len(t.News.Articles)-1].Title = title
	return t
}
func (t *ArticleCustomMessageRequest) Description(description string) *ArticleCustomMessageRequest {
	t.News.Articles[len(t.News.Articles)-1].Description = description
	return t
}
func (t *ArticleCustomMessageRequest) Url(url string) *ArticleCustomMessageRequest {
	t.News.Articles[len(t.News.Articles)-1].Url = url
	return t
}
func (t *ArticleCustomMessageRequest) PicUrl(picUrl string) *ArticleCustomMessageRequest {
	t.News.Articles[len(t.News.Articles)-1].PicUrl = picUrl
	return t
}

// 发送客服消息
func (m *Mp) SendCustomMessage(req CustomMessageRequest) (resp *WechatApiError, err respErr.Response) {
	resp = &WechatApiError{}
	err = m.Cgi(resp, "/cgi-bin/message/custom/send", url.Values{}, req)
	return
}
