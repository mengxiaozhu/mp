package oauth2

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// GetAccessToken 获取用户access_token.
//  appId:               公众号的唯一标识
//  code:                填写第一步获取的code参数
//  componentAppId:      服务方的appid，在申请创建公众号服务成功后，可在公众号服务详情页找到
//  componentAccessToken:服务开发方的access_token
func GetAccessToken(appId, code, componentAppId, componentAccessToken string) (accessToken *AccessToken, err error) {

	url := AuthAccessTokenURL(appId, code, componentAppId, componentAccessToken)
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
		AccessToken
	}
	if err = json.Unmarshal(body, &result); err != nil {
		return
	}
	if result.ErrCode != ErrCodeOK {
		err = errors.New(result.ErrMsg)
		return
	}
	accessToken = &result.AccessToken
	return
}
