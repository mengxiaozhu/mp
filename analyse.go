package mp

import (
	"encoding/json"
	"mp/base"
	"net/url"
	"qiniupkg.com/x/log.v7"
	"time"
)

type UserAnalyseRequest struct {
	BeginDate string `json:"begin_date"`
	EndDate   string `json:"end_date"`
}

type CumulateUserItem struct {
	RefDate      string `json:"ref_date"`
	CumulateUser int64  `json:"cumulate_user"`
}

type CumulateUserResp struct {
	List []*CumulateUserItem `json:"list"`
}

// 获取累计的用户
func (this *Mp) GetUserCumulateStruct(start time.Time, count int64) (result *CumulateUserResp, err error) {
	result = &CumulateUserResp{}
	rawData, err := this.GetUserCumulate(start, count)
	if err != nil {
		return
	}
	log.Println(string(rawData))
	err = json.Unmarshal(rawData, result)
	if err != nil {
		return
	}
	return
}

func (mp *Mp) GetUserCumulate(start time.Time, count int64) (result []byte, err error) {
	userAnalyseRequest := &UserAnalyseRequest{
		EndDate:   start.Format("2006-01-02"),
		BeginDate: start.AddDate(0, 0, -(int)(count)).Format("2006-01-02"),
	}
	params := &url.Values{}
	params.Add("access_token", mp.Token)
	u := "https://api.weixin.qq.com/datacube/getusercumulate?" + params.Encode()
	log.Println(userAnalyseRequest)
	body, err := base.PostJSON(u, userAnalyseRequest)
	if err != nil {
		return nil, err
	}
	return body, nil
}
