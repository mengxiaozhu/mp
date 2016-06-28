package mp

import (
	"net/url"
	"qiniupkg.com/x/errors.v7"
)

type Tag struct {
	Id    int64  `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Count int64  `json:"count,omitempty"`
}

type WechatTag struct {
	T *Tag `json:"tag"`
}

func (mp *Mp) NewTag(name string) (*WechatTag, error) {
	newTag := &WechatTag{
		T: &Tag{Name: name},
	}
	err := mp.Cgi(newTag, "/cgi-bin/tags/create", url.Values{}, newTag)
	return newTag, err
}

type WechatTags struct {
	Tags []*Tag `json:"tags"`
}

var ErrorTagNotFound = errors.New("tag not found")

func (this *WechatTags) FindTag(name string) (*Tag, error) {
	for _, t := range this.Tags {
		if t.Name == name {
			return t, nil
		}
	}
	return nil, ErrorTagNotFound

}

func (mp *Mp) GetTags() (*WechatTags, error) {
	tags := &WechatTags{}
	err := mp.CgiGet(tags, "/cgi-bin/tags/get", url.Values{})
	return tags, err
}

func (mp *Mp) EditTag(id int64, name string) (*WechatApiError, error) {
	editTag := &WechatTag{
		T: &Tag{
			Id:   id,
			Name: name,
		},
	}
	wechatApiError := &WechatApiError{}
	err := mp.Cgi(wechatApiError, "/cgi-bin/tags/update", url.Values{}, editTag)
	return wechatApiError, err
}

func (mp *Mp) DeleteTag(id int64) error {
	deleteTag := &WechatTag{
		T: &Tag{
			Id: id,
		},
	}
	wechatApiError := &WechatApiError{}
	err := mp.Cgi(wechatApiError, "/cgi-bin/tags/delete", url.Values{}, deleteTag)
	return err
}

type ListTagedUserRequest struct {
	TagId      int64  `json:"tagid"`
	NextOpenid string `json:"next_openid"`
}

type TagedUserList struct {
	Count int64 `json:"count"`
	Data  struct {
		Openid []string `json:"openid"`
	} `json:"data"`
	NextOpenid string `json:"next_openid"`
}

func (mp *Mp) GetTagedUserList(tagId int64, nextOpenid string) (*TagedUserList, error) {
	listTagedUserRequest := &ListTagedUserRequest{
		TagId:      tagId,
		NextOpenid: nextOpenid,
	}
	tagedUserList := &TagedUserList{}
	err := mp.Cgi(tagedUserList, "/cgi-bin/user/tag/get", url.Values{}, listTagedUserRequest)
	return tagedUserList, err
}

type TagUsersRequest struct {
	OpenidList []string `json:"openid_list"`
	TagId      int64    `json:"tagid"`
}

func (mp *Mp) TagUsers(tagId int64, openidList []string) (*WechatApiError, error) {
	tagUsersRequest := &TagUsersRequest{
		TagId:      tagId,
		OpenidList: openidList,
	}
	wechatApiError := &WechatApiError{}
	err := mp.Cgi(wechatApiError, "/cgi-bin/tags/members/batchtagging", url.Values{}, tagUsersRequest)
	return wechatApiError, err
}

func (mp *Mp) UnTagUsers(tagId int64, openidList []string) (*WechatApiError, error) {
	tagUsersRequest := &TagUsersRequest{
		TagId:      tagId,
		OpenidList: openidList,
	}
	wechatApiError := &WechatApiError{}
	err := mp.Cgi(wechatApiError, "/cgi-bin/tags/members/batchuntagging", url.Values{}, tagUsersRequest)
	return wechatApiError, err
}

type UserTagRequset struct {
	Openid string `json:"openid"`
}

type UserTagList struct {
	TagIdList []int64 `json:"tagid_list"`
}

func (mp *Mp) GetUserTags(openid string) (*UserTagList, error) {
	userTagRequset := &UserTagRequset{
		Openid: openid,
	}
	userTagList := &UserTagList{}
	err := mp.Cgi(userTagList, "/cgi-bin/tags/getidlist", url.Values{}, userTagRequset)
	return userTagList, err
}
