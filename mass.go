package mp

import (
	"bytes"
	"encoding/json"
	"github.com/mengxiaozhu/mp/base"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"qiniupkg.com/x/log.v7"
)

// http://mp.weixin.qq.com/wiki/15/40b6865b893947b764e2de8e4a1fb55f.html#.E4.B8.8A.E4.BC.A0.E5.9B.BE.E6.96.87.E6.B6.88.E6.81.AF.E5.86.85.E7.9A.84.E5.9B.BE.E7.89.87.E8.8E.B7.E5.8F.96URL.E3.80.90.E8.AE.A2.E9.98.85.E5.8F.B7.E4.B8.8E.E6.9C.8D.E5.8A.A1.E5.8F.B7.E8.AE.A4.E8.AF.81.E5.90.8E.E5.9D.87.E5.8F.AF.E7.94.A8.E3.80.91
func (mp *Mp) UploadImg(filename string, img []byte) (result []byte, err error) {
	params := &url.Values{}
	params.Add("access_token", mp.Token)
	u := "https://api.weixin.qq.com/cgi-bin/media/uploadimg?" + params.Encode()

	postBody := bytes.NewBuffer(make([]byte, 0, 10<<20))
	multipartWriter := multipart.NewWriter(postBody)
	partWriter, err := multipartWriter.CreateFormFile("media", filename)
	if err != nil {
		return nil, err
	}
	if _, err = io.Copy(partWriter, bytes.NewReader(img)); err != nil {
		return nil, err
	}
	if err = multipartWriter.Close(); err != nil {
		return nil, err
	}
	requestBodyBytes := postBody.Bytes()
	requestBodyType := multipartWriter.FormDataContentType()

	resp, err := http.Post(u, requestBodyType, bytes.NewReader(requestBodyBytes))
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// http://mp.weixin.qq.com/wiki/10/10ea5a44870f53d79449290dfd43d006.html
// 媒体文件类型，分别有图片（image）、语音（voice）、视频（video）和缩略图（thumb）
func (mp *Mp) Upload(typ string, filename string, file []byte) (result []byte, err error) {
	params := &url.Values{}
	params.Add("access_token", mp.Token)
	params.Add("type", typ)
	u := "https://api.weixin.qq.com/cgi-bin/material/add_material?" + params.Encode()

	postBody := bytes.NewBuffer(make([]byte, 0, 10<<20))
	multipartWriter := multipart.NewWriter(postBody)
	partWriter, err := multipartWriter.CreateFormFile("media", filename)
	if err != nil {
		return nil, err
	}
	if _, err = io.Copy(partWriter, bytes.NewReader(file)); err != nil {
		return nil, err
	}
	if err = multipartWriter.Close(); err != nil {
		return nil, err
	}
	requestBodyBytes := postBody.Bytes()
	requestBodyType := multipartWriter.FormDataContentType()

	resp, err := http.Post(u, requestBodyType, bytes.NewReader(requestBodyBytes))
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// http://mp.weixin.qq.com/wiki/9/677a85e3f3849af35de54bb5516c2521.html

//永久素材下载
type GetMediaMessage struct {
	MediaId string `json:"media_id"`
}

func (mp *Mp) GetMedia(mediaId string) (result []byte, err error) {
	params := &url.Values{}
	params.Add("access_token", mp.Token)
	msg := &GetMediaMessage{
		MediaId: mediaId,
	}
	u := "https://api.weixin.qq.com/cgi-bin/material/get_material?" + params.Encode()
	result, err = base.PostJSON(u, msg)
	return result, err
}

func (mp *Mp) GetTempMedia(mediaId string) (result []byte, err error) {
	params := &url.Values{}
	params.Add("access_token", mp.Token)
	params.Add("media_id", mediaId)
	u := "https://api.weixin.qq.com/cgi-bin/media/get?" + params.Encode()
	log.Println(u)
	resp, err := http.Get(u)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

//永久素材数量获取
func (mp *Mp) GetCount() (result []byte, err error) {
	params := &url.Values{}
	params.Add("access_token", mp.Token)
	u := "https://api.weixin.qq.com/cgi-bin/material/get_materialcount?" + params.Encode()
	resp, err := http.Get(u)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

type ListRequest struct {
	Typ    string `json:"type"`
	Offset int    `json:"offset"`
	Count  int    `json:"count"`
}

func (mp *Mp) GetList(listRequest *ListRequest) (result []byte, err error) {
	params := &url.Values{}
	params.Add("access_token", mp.Token)
	u := "https://api.weixin.qq.com/cgi-bin/material/batchget_material?" + params.Encode()
	result, err = base.PostJSON(u, listRequest)
	return
}

// http://mp.weixin.qq.com/wiki/15/40b6865b893947b764e2de8e4a1fb55f.html#.E4.B8.8A.E4.BC.A0.E5.9B.BE.E6.96.87.E6.B6.88.E6.81.AF.E7.B4.A0.E6.9D.90.E3.80.90.E8.AE.A2.E9.98.85.E5.8F.B7.E4.B8.8E.E6.9C.8D.E5.8A.A1.E5.8F.B7.E8.AE.A4.E8.AF.81.E5.90.8E.E5.9D.87.E5.8F.AF.E7.94.A8.E3.80.91
type Article struct {
	ThumbMediaId     string `json:"thumb_media_id"`     //图文消息缩略图的media_id，可以在基础支持-上传多媒体文件接口中获得
	Author           string `json:"author"`             //图文消息的作者
	Title            string `json:"title"`              //图文消息的标题
	ContentSourceUrl string `json:"content_source_url"` //在图文消息页面点击“阅读原文”后的页面
	Content          string `json:"content"`            //图文消息页面的内容，支持HTML标签。具备微信支付权限的公众号，可以使用a标签，其他公众号不能使用
	Digest           string `json:"digest"`             //图文消息的描述
	ShowCoverPic     int    `json:"show_cover_pic"`     //是否显示封面，1为显示，0为不显示
}

type News struct {
	Articles []*Article `json:"articles"` //图文消息，一个图文消息支持1到8条图文
}

func (news *News) Insert(insertion []*Article, index int) {
	result := make([]*Article, len(news.Articles)+len(insertion))
	at := copy(result, news.Articles[:index])
	at += copy(result[at:], insertion)
	copy(result[at:], news.Articles[index:])
	news.Articles = result
}

type UploadNewsResult struct {
	WechatApiError
	MediaId string `json:"media_id"`
	Url     string `json:"url"`
}

func (this *News) String() string {
	bytes, err := json.Marshal(this)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func (this *Mp) UploadNewsResult(news *News) (result *UploadNewsResult, err error) {
	result = &UploadNewsResult{}
	uploadResultBytes, err := this.UploadNews(news)
	if err != nil {
		return
	}
	err = json.Unmarshal(uploadResultBytes, result)
	if err != nil {
		return
	}
	return
}

func (mp *Mp) UploadNews(news *News) (result []byte, err error) {
	params := &url.Values{}
	params.Add("access_token", mp.Token)
	u := "https://api.weixin.qq.com/cgi-bin/material/add_news?" + params.Encode()
	result, err = base.PostJSON(u, news)
	return
}

func (mp *Mp) DeleteMedia(mediaId string) (result []byte, err error) {
	params := &url.Values{}
	params.Add("access_token", mp.Token)
	u := "https://api.weixin.qq.com/cgi-bin/material/del_material?" + params.Encode()
	result, err = base.PostJSON(u, &GetMediaMessage{
		MediaId: mediaId,
	})
	return
}

// 所有数据结构最后仅发送Message, Message可以在多个接口中使用,分组群发时,不要对Touser赋值, 根据OpenId分发时不要对MessageFilter赋值..
// MpNews MpText MpVoice MpImage MpVideo MpWxCard只能选择一个赋值,同时请修改Message中MsgType为相应类型.
// 比如发送MpNews类型的群发,请修改Message中MsgType为mpnews,类型对照信息请参照Message中的json序列化值
// 注意: 发送video时需要提前置换video的media_id 请调用     接口
type MessageFilter struct {
	IsToALL bool `json:"is_to_all"`
	GroupId int  `json:"group_id"`
}
type MpNews struct {
	MediaId string `json:"media_id"`
}
type MpText struct {
	Content string `json:"content"`
}
type MpVoice struct {
	MediaId string `json:"media_id"`
}
type MpImage struct {
	MediaId string `json:"media_id"`
}
type MpVideo struct {
	MediaId string `json:"media_id"`
}
type MpWxCard struct {
	CardId string `json:"card_id"`
}
type Message struct {
	Filter   *MessageFilter `json:"filter,omitempty"`
	Touser   interface{}    `json:"touser,omitempty"`
	Towxname string         `json:"towxname,emitempty"`
	MsgType  string         `json:"msgtype"`
	MpNews   *MpNews        `json:"mpnews,omitempty"`
	Text     *MpText        `json:"text,omitempty"`
	Voice    *MpVoice       `json:"voice,omitempty"`
	Image    *MpImage       `json:"image,omitempty"`
	Video    *MpVideo       `json:"mpvideo,omitempty"`
	WxCard   *MpWxCard      `json:"wxcard,omitempty"`
}

type VideoExchange struct {
	MediaId     string `json:"media_id"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}

func (mp *Mp) ExchangeVideoMediaId(msg *VideoExchange) (result []byte, err error) {
	params := &url.Values{}
	params.Add("access_token", mp.Token)
	u := "https://api.weixin.qq.com/cgi-bin/media/uploadvideo?" + params.Encode()

	result, err = base.PostJSON(u, msg)
	return
}

// http://mp.weixin.qq.com/wiki/15/40b6865b893947b764e2de8e4a1fb55f.html#.E6.A0.B9.E6.8D.AE.E5.88.86.E7.BB.84.E8.BF.9B.E8.A1.8C.E7.BE.A4.E5.8F.91.E3.80.90.E8.AE.A2.E9.98.85.E5.8F.B7.E4.B8.8E.E6.9C.8D.E5.8A.A1.E5.8F.B7.E8.AE.A4.E8.AF.81.E5.90.8E.E5.9D.87.E5.8F.AF.E7.94.A8.E3.80.91
// 群发接口
func (mp *Mp) Send(msg *Message) (result []byte, err error) {
	params := &url.Values{}
	params.Add("access_token", mp.Token)
	var u string
	if msg.Filter != nil {
		u = "https://api.weixin.qq.com/cgi-bin/message/mass/sendall?" + params.Encode()
	} else {
		u = "https://api.weixin.qq.com/cgi-bin/message/mass/send?" + params.Encode()
	}
	result, err = base.PostJSON(u, msg)
	return
}

// http://mp.weixin.qq.com/wiki/15/40b6865b893947b764e2de8e4a1fb55f.html#.E5.88.A0.E9.99.A4.E7.BE.A4.E5.8F.91.E3.80.90.E8.AE.A2.E9.98.85.E5.8F.B7.E4.B8.8E.E6.9C.8D.E5.8A.A1.E5.8F.B7.E8.AE.A4.E8.AF.81.E5.90.8E.E5.9D.87.E5.8F.AF.E7.94.A8.E3.80.91
type DeleteMessage struct {
	MsgId int64 `json:"msg_id"`
}

func (mp *Mp) Delete(msgId int64) (result []byte, err error) {
	params := &url.Values{}
	params.Add("access_token", mp.Token)
	u := "https://api.weixin.qq.com/cgi-bin/message/mass/delete?" + params.Encode()
	result, err = base.PostJSON(u, &DeleteMessage{
		MsgId: msgId,
	})
	return
}

// http://mp.weixin.qq.com/wiki/15/40b6865b893947b764e2de8e4a1fb55f.html#.E6.9F.A5.E8.AF.A2.E7.BE.A4.E5.8F.91.E6.B6.88.E6.81.AF.E5.8F.91.E9.80.81.E7.8A.B6.E6.80.81.E3.80.90.E8.AE.A2.E9.98.85.E5.8F.B7.E4.B8.8E.E6.9C.8D.E5.8A.A1.E5.8F.B7.E8.AE.A4.E8.AF.81.E5.90.8E.E5.9D.87.E5.8F.AF.E7.94.A8.E3.80.91
// 查询消息发送状态
type SearchMessage struct {
	MsgId string `json:"msg_id"`
}

func (mp *Mp) GetMessageStatus(msgId string) (result []byte, err error) {
	params := &url.Values{}
	params.Add("access_token", mp.Token)
	u := "https://api.weixin.qq.com/cgi-bin/message/mass/get?" + params.Encode()
	result, err = base.PostJSON(u, &SearchMessage{
		MsgId: msgId,
	})
	return
}

func (mp *Mp) Preview(msg *Message) (result []byte, err error) {
	params := &url.Values{}
	params.Add("access_token", mp.Token)
	log.Println(msg)
	u := "https://api.weixin.qq.com/cgi-bin/message/mass/preview?" + params.Encode()
	result, err = base.PostJSON(u, msg)
	return
}
