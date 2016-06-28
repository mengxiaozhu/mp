package mp

import (
	"net/url"
)

type CreateQRCodeRequest struct {
	ExpireSeconds int64  `json:"expire_seconds"`
	ActionName    string `json:"action_name"`
	ActionInfo    struct {
		Scene struct {
			SceneStr string `json:"scene_str,omitempty"`
			SceneId  int32  `json:"scene_id"`
		} `json:"scene"`
	} `json:"action_info"`
}

type CreateQRCodeResp struct {
	Ticket        string `json:"ticket"`
	ExpireSeconds int64  `json:"expire_seconds"`
	Url           string `json:"url"`
}

func (this *Mp) CreateQRCode(req *CreateQRCodeRequest) (resp *CreateQRCodeResp, err error) {
	resp = &CreateQRCodeResp{}
	err = this.Cgi(resp, "/cgi-bin/qrcode/create", url.Values{}, req)
	return
}
