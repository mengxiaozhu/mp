package mp

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var AuthUrlPost string = "http://auth.mengxiaozhu.cn/Interface/base64_transfer_post_clear"
var AuthUrlGet string = "http://auth.mengxiaozhu.cn/Interface/base64_transfer_get_clear"
var Transfer string = "http://component-trans.mengxiaozhu.cn"
var TransMethod string = os.Getenv("TRANS")

// 转发Get请求
// 同PostTransfer
func GetDirect(u []byte) (resp []byte, err error) {
	return PostDirect(u, nil)
}

// 转发Post请求
// 返回的错误值不为空的时候,一定会有一个不为空的byte数组
func PostDirect(u []byte, data []byte) (resp []byte, err error) {
	var respPost *http.Response
	var errRespPost error
	// 转发
	if data != nil {
		respPost, errRespPost = http.Post(string(u), "application/json", bytes.NewReader(data))
	} else {
		respPost, errRespPost = http.Get(string(u))
	}

	if errRespPost != nil {
		if respPost.Body != nil {
			respPost.Body.Close()
		}
		err = errRespPost
		return
	}
	// 读取内容
	defer respPost.Body.Close()
	return ioutil.ReadAll(respPost.Body)
}

// 转发Get请求
// 同PostTransfer
func GetTransfer(u []byte) (resp []byte, err error) {
	return PostTransfer(u, nil)
}

// 转发Post请求
// 返回的错误值不为空的时候,一定会有一个不为空的byte数组
func PostTransfer(u []byte, data []byte) (resp []byte, err error) {
	queryUrl := url.Values{}
	// 发送不同的转发请求
	if data != nil {
		queryUrl.Add("method", "POST")
	} else {
		queryUrl.Add("method", "GET")
	}
	// 添加URL地址
	queryUrl.Add("url", string(u))
	// 转发
	respPost, errRespPost := http.Post(Transfer+"/trans?"+queryUrl.Encode(), "application/json", bytes.NewReader(data))
	if errRespPost != nil {
		if respPost.Body != nil {
			respPost.Body.Close()
		}
		err = errRespPost
		return
	}
	// 读取内容
	defer respPost.Body.Close()
	return ioutil.ReadAll(respPost.Body)
}

// 转发Get请求
func AuthGetTransfer(u []byte) (resp []byte, err error) {
	// 新的转发方式
	if strings.ToUpper(TransMethod) == "TRANS" {
		return GetTransfer(u)
	}

	if strings.ToUpper(TransMethod) == "DIRECT" {
		return GetDirect(u)
	}
	queryBody := url.Values{}
	queryBody.Add("url", base64.StdEncoding.EncodeToString(u))
	respPost, errRespPost := http.PostForm(AuthUrlGet, queryBody)
	if errRespPost != nil {
		if respPost.Body != nil {
			respPost.Body.Close()
		}
		err = errRespPost
		return
	}

	b, e := ioutil.ReadAll(respPost.Body)
	fmt.Println(respPost.Body, errRespPost, b, e)
	defer respPost.Body.Close()
	return b, e

}

// 转发Post请求
func AuthPostTransfer(u []byte, data []byte) (resp []byte, err error) {
	// 新的转发方式
	if strings.ToUpper(TransMethod) == "TRANS" {
		return PostTransfer(u, data)
	}
	if strings.ToUpper(TransMethod) == "DIRECT" {
		return PostDirect(u, data)
	}

	queryBody := url.Values{}

	queryBody.Add("url", base64.StdEncoding.EncodeToString(u))
	queryBody.Add("data", base64.StdEncoding.EncodeToString(data))

	respPost, errRespPost := http.PostForm(AuthUrlPost, queryBody)
	if errRespPost != nil {
		if respPost.Body != nil {
			respPost.Body.Close()
		}
		err = errRespPost
		return
	}

	defer respPost.Body.Close()
	return ioutil.ReadAll(respPost.Body)
}
