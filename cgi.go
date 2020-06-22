package mp

import (
	"encoding/json"
	"fmt"
	mpBase "github.com/mengxiaozhu/mp/base"
	"net/url"
	"errors"
)

// 微信的错误信息
type WechatApiError struct {
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// 微信API提供出的错误信息
type ErrorResponse struct {
	mpBase.ErrResp
	Data WechatApiError `json:"data"`
}

func (m *Mp) Cgi(dest interface{}, method string, params url.Values, body interface{}) (err *mpBase.ErrResp) {

	params.Add("access_token", m.Token)
	u := "https://api.weixin.qq.com" + method + "?" + params.Encode()
	var bytesTransfer []byte
	var errTransfer error
	if "" != body {
		bodyBytes, bodyErr := json.Marshal(body)

		if bodyErr != nil {
			bytesTransfer, errTransfer = AuthPostTransfer([]byte(u), []byte(""))
		} else {
			bytesTransfer, errTransfer = AuthPostTransfer([]byte(u), mpBase.ClearUnicode(bodyBytes))
		}

	} else {
		bytesTransfer, errTransfer = AuthGetTransfer([]byte(u))
	}

	fmt.Println(string(bytesTransfer), errTransfer)
	if errTransfer != nil {
		return mpBase.NewErrResp(-1, errTransfer)
	}

	wechatApiErr := WechatApiError{}
	// 解析错误信息
	errUnmarshalApiErr := json.Unmarshal(bytesTransfer, &wechatApiErr)

	if errUnmarshalApiErr != nil {
		return mpBase.NewErrResp(-1, errUnmarshalApiErr)
	}

	if wechatApiErr.ErrCode != 0 {
		return mpBase.NewErrResp(wechatApiErr.ErrCode, errors.New(wechatApiErr.ErrMsg))
	}

	errUnmarshalResponse := json.Unmarshal(bytesTransfer, dest)

	if errUnmarshalResponse != nil {
		return mpBase.NewErrResp(-1, errUnmarshalResponse)
	}
	return nil
}

func (m *Mp) CgiGet(dest interface{}, method string, params url.Values) (err *mpBase.ErrResp) {
	return m.Cgi(dest, method, params, "")
}
