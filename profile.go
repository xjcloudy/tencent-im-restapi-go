package tim

type PortraitGetResp struct {
	CommonResp
	UserProfileItem []UserProfile
	FailAccount     []string `json:"Fail_Account"`
	InvalidAccount  []string `json:"Invalid_Account"`
}
type UserProfile struct {
	ToAccount   string `json:"To_Account"`
	ProfileItem []ProfileKV
}
type ProfileKV struct {
	Tag   string
	Value interface{}
}

// PortraitGet 拉取账号资料
func (api *TimApp) PortraitGet(account []string, tagList []string) (*PortraitGetResp, error) {
	req := map[string][]string{
		"To_Account": account,
		"TagList":    tagList,
	}
	resp := new(PortraitGetResp)
	err := api.do(profileSvc, "portrait_get", req, resp)
	return resp, err
}

// PortraitSet 设置账户资料
func (api *TimApp) PortraitSet(account string, profileItem []ProfileKV) (*CommonResp, error) {
	req := map[string]interface{}{
		"From_Account": account,
		"ProfileItem":  profileItem,
	}
	resp := new(CommonResp)
	err := api.do(profileSvc, "portrait_set", req, resp)
	return resp, err
}
