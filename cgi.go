package mp

import (
	"encoding/json"
	"fmt"
	mpBase "mp/base"
	"net/url"
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

func (m *Mp) Cgi(dest interface{}, method string, params url.Values, body interface{}) (err mpBase.Response) {

	params.Add("access_token", m.Token)
	u := "http://api.weixin.qq.com" + method + "?" + params.Encode()
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
		return &mpBase.ErrResp{mpBase.Resp{-1, "api error:" + string(bytesTransfer)}}
	}

	if wechatApiErr.ErrCode != 0 {
		return &ErrorResponse{
			ErrResp: mpBase.ErrResp{mpBase.Resp{-1, "api error:" + string(bytesTransfer)}},
			Data:          wechatApiErr,
		}
	}

	errUnmarshalResponse := json.Unmarshal(bytesTransfer, dest)

	if errUnmarshalResponse != nil {
		return &mpBase.ErrResp{mpBase.Resp{-1, "api json:" + string(bytesTransfer)}}
	}
	return nil
}

func (m *Mp) CgiGet(dest interface{}, method string, params url.Values) (err mpBase.Response) {
	return m.Cgi(dest, method, params, "")
}
