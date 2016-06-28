package mp

import (
	"encoding/json"
	//	"errors"
)

type DownloadedNews struct {
	NewsItem []*Article `json:"news_item"`
}

//func (this *Mp) GetTargetNews(mediaId string, index *int) (news *DownloadedNews, err error) {
//	if mediaId != "" {
//		news, err := this.GetNews(mediaId)
//		if err != nil {
//			return
//		}
//		if index > len(news.NewsItem) {
//			err = errors.New("index超过限制")
//			return
//		}
//	} else {
//		index = 0
//	}
//}

func (this *Mp) GetNews(mediaId string) (news *DownloadedNews, err error) {

	news = &DownloadedNews{}
	newsBody, err := this.GetMedia(mediaId)
	err = json.Unmarshal(newsBody, &news)
	if err != nil {
		return
	}
	return
}

type MsgStatus struct {
	MsgId     int64  `json:"msg_id"`
	MsgStatus string `json:"msg_status"`
}

func (this *Mp) GetMsgStatus(msgId string) (status *MsgStatus, err error) {

	status = &MsgStatus{}
	statusBody, err := this.GetMessageStatus(msgId)
	err = json.Unmarshal(statusBody, &status)
	if err != nil {
		return
	}
	return
}
