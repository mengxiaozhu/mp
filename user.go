package mp

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

func (mp *Mp) GetUserList(nextOpenid string) (result []byte, err error) {
	params := &url.Values{}
	params.Add("access_token", mp.Token)
	params.Add("next_openid", nextOpenid)
	u := "https://api.weixin.qq.com/cgi-bin/user/get?" + params.Encode()
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
