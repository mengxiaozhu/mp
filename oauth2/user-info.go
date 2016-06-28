package oauth2

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// GetUserInfo 获取用户信息.
//  accessToken: 网页授权接口调用凭证
//  openId:      用户的唯一标识
//  lang:        返回国家地区语言版本，zh_CN 简体，zh_TW 繁体，en 英语, 如果留空 "" 则默认为 zh_CN
//  httpClient:  如果不指定则默认为 http.DefaultClient
func GetUserInfo(accessToken, openid, lang string) (info *UserInfo, err error) {
	switch lang {
	case "":
		lang = LanguageZhCN
	case LanguageZhCN, LanguageZhTW, LanguageEN:
	default:
		lang = LanguageZhCN
	}

	url := GetUserInfoURL(accessToken, openid, lang)
	httpResp, err := http.Get(url)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode != http.StatusOK {
		err = fmt.Errorf("http.Status: %s", httpResp.Status)
		return
	}
	body, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return nil, err
	}
	var result struct {
		Error
		UserInfo
	}
	if err = json.Unmarshal(body, &result); err != nil {
		return
	}
	if result.ErrCode != ErrCodeOK {
		err = errors.New(result.ErrMsg)
		return
	}
	info = &result.UserInfo
	return
}
