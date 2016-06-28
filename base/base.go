package base

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

func PostJSON(url string, body interface{}) (result []byte, err error) {
	postBody, err := json.Marshal(body)
	postBody = ClearUnicode(postBody)
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(url, "application/json", bytes.NewReader(postBody))
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	result, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// 清除JSON转义中的Unicode编码,方便微信服务器接收
func ClearUnicode(in []byte) (out []byte) {
	s := strings.Replace(string(in), `\u003c`, `<`, -1)
	s = strings.Replace(s, `\u003e`, `>`, -1)
	return []byte(s)

}
