package oauth2

import (
	"net/url"
)

// AuthCodeURL 生成网页授权地址.
//  appId:          公众号的唯一标识
//  redirectURI:    授权后重定向的回调链接地址
//  scope:          应用授权作用域 拥有多个作用域用逗号（,）分隔
//  state:          重定向后会带上 state 参数, 开发者可以填写 a-zA-Z0-9 的参数值, 最多128字节
//  componentAppId: 服务方的appid，在申请创建公众号服务成功后，可在公众号服务详情页找到
func AuthCodeURL(appId, redirectURI, scope, state, componentAppId string) string {
	return "https://open.weixin.qq.com/connect/oauth2/authorize?appid=" + url.QueryEscape(appId) +
		"&redirect_uri=" + url.QueryEscape(redirectURI) +
		"&response_type=code" +
		"&scope=" + url.QueryEscape(scope) +
		"&state=" + url.QueryEscape(state) +
		"&component_appid=" + url.QueryEscape(componentAppId) +
		"#wechat_redirect"
}

// AuthAccessTokenURL 生成通过code换取access_token地址.
//  appId:               公众号的唯一标识
//  code:                填写第一步获取的code参数
//  componentAppId:      服务方的appid，在申请创建公众号服务成功后，可在公众号服务详情页找到
//  componentAccessToken:服务开发方的access_token
func AuthAccessTokenURL(appId, code, componentAppId, componentAccessToken string) string {
	return "https://api.weixin.qq.com/sns/oauth2/component/access_token?appid=" + url.QueryEscape(appId) +
		"&code=" + url.QueryEscape(code) +
		"&grant_type=authorization_code" +
		"&component_appid=" + url.QueryEscape(componentAppId) +
		"&component_access_token=" + url.QueryEscape(componentAccessToken)
}

// RefreshAccessTokenURL 生成刷新access_token地址.
//  appId:               公众号的唯一标识
//  code:                填写第一步获取的code参数
//  componentAppId:      服务方的appid，在申请创建公众号服务成功后，可在公众号服务详情页找到
//  componentAccessToken:服务开发方的access_token
func RefreshAccessTokenURL(appId, refreshToken, componentAppId, componentAccessToken string) string {
	return "https://api.weixin.qq.com/sns/oauth2/component/refresh_token?appid=" + url.QueryEscape(appId) +
		"&grant_type=refresh_token" +
		"&refresh_token=" + url.QueryEscape(refreshToken) +
		"&component_appid=" + url.QueryEscape(componentAppId) +
		"&component_access_token=" + url.QueryEscape(componentAccessToken)
}

// GetUserInfoURL 生成获取用户信息地址.
//  accessToken:         网页授权接口调用凭证,注意：此access_token与基础支持的access_token不同
//  openId:              用户的唯一标识
//  language:            返回国家地区语言版本，zh_CN 简体，zh_TW 繁体，en 英语
func GetUserInfoURL(accessToken, openid, language string) string {
	return "https://api.weixin.qq.com/sns/userinfo?access_token=" + url.QueryEscape(accessToken) +
		"&openid=refresh_token" + url.QueryEscape(openid) +
		"&lang=" + url.QueryEscape(language)
}
