package mp

import "errors"
import (
	respErr "h5/response/error"
	"net/url"
)

type Template struct {
	TemplateId      string `json:"template_id"`
	Title           string `json:"title"`
	PrimaryIndustry string `json:"primary_industry"`
	DeputyIndustry  string `json:"deputy_industry"`
	Content         string `json:"content"`
	Example         string `json:"example"`
}

type AllPrivateTemplate struct {
	TemplateList []Template `json:"template_list"`
}

var CantFindTemplateErr respErr.Response = respErr.NewErrorRespones(errors.New("cant find template"))

func (a *AllPrivateTemplate) FindTemplate(title string) (t *Template, err respErr.Response) {
	for _, template := range a.TemplateList {
		if template.Title == title {
			return &template, nil
		}
	}
	return nil, CantFindTemplateErr

}

func (a *AllPrivateTemplate) FindTemplateById(templateId string) (t *Template, err respErr.Response) {
	for _, template := range a.TemplateList {
		if template.TemplateId == templateId {
			return &template, nil
		}
	}
	return nil, CantFindTemplateErr

}

func (m *Mp) GetAllPrivateTemplate() (templates *AllPrivateTemplate, err respErr.Response) {
	templates = &AllPrivateTemplate{}
	err = m.CgiGet(templates, "/cgi-bin/template/get_all_private_template", url.Values{})
	return
}

type IndustryClass struct {
	FirstClass  string `json:"first_class"`
	SecondClass string `json:"second_class"`
}

type Industry struct {
	PrimaryIndustry   *IndustryClass `json:"primary_industry"`
	SecondaryIndustry *IndustryClass `json:"secondary_industry"`
}

func (i *Industry) IsDustry(secondIndustry string) bool {
	return i.PrimaryIndustry.SecondClass == secondIndustry || i.SecondaryIndustry.SecondClass == secondIndustry
}

func (m *Mp) GetIndustry() (industry *Industry, err respErr.Response) {
	industry = &Industry{}
	err = m.CgiGet(industry, "/cgi-bin/template/get_industry", url.Values{})
	return
}

type ApiAddTemplateRequest struct {
	TemplateIdShort string `json:"template_id_short"`
}
type ApiAddTemplateResponse struct {
	WechatApiError
	TemplateId string `json:"template_id"`
}

func (m *Mp) ApiAddTemplate(templateIdShort string) (resp *ApiAddTemplateResponse, err respErr.Response) {
	resp = &ApiAddTemplateResponse{}
	req := &ApiAddTemplateRequest{templateIdShort}
	err = m.Cgi(resp, "/cgi-bin/template/api_add_template", url.Values{}, req)
	return
}

func (m *Mp) FindOrCreateTemplateIdByTitle(title string, shortId string) (templateId string, err respErr.Response) {

	// 检查是否已经添加了对应的模板
	templates, err := m.GetAllPrivateTemplate()
	if err != nil {
		return
	}
	template, err := templates.FindTemplate(title)

	if err == nil {
		templateId = template.TemplateId
		return
	} else {
		respAddTemplate, err := m.ApiAddTemplate(shortId)
		if err != nil {
			return "", err
		}
		templateId = respAddTemplate.TemplateId
		return templateId, nil
	}
}

// 查找模板,当err为空时,一定返回一个Template结构体,否则返回一个可以Response的错误信息
func (m *Mp) FindTemplate(templateId string) (template *Template, err respErr.Response) {
	// 检查是否已经添加了对应的模板
	templates, err := m.GetAllPrivateTemplate()
	if err != nil {
		return
	}
	return templates.FindTemplateById(templateId)

}

type TemplateMessage struct {
	ToUser     string      `json:"touser"`
	TemplateId string      `json:"template_id"`
	Url        string      `json:"url"`
	Data       interface{} `json:"data"`
}

func (m *Mp) SendTemplateMessage(openId string, templateId string, messageUrl string, data interface{}) (err respErr.Response) {
	templateMessage := &TemplateMessage{
		ToUser:     openId,
		TemplateId: templateId,
		Url:        messageUrl,
		Data:       data,
	}
	resp := &WechatApiError{}
	err = m.Cgi(resp, "/cgi-bin/message/template/send", url.Values{}, templateMessage)
	return
}
