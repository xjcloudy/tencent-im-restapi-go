package tim

type AddType string
type ForceAddFlag int
type DeleteType string

const (
	// TIMAddTypeSingle 单项
	TIMAddTypeSingle AddType = "Add_Type_Single"
	// TIMAddTypeBoth 双向
	TIMAddTypeBoth AddType = "Add_Type_Both"

	// TIMForceAdd 强制添加
	TIMForceAdd ForceAddFlag = 1
	// TIMNormalAdd 常规添加
	TIMNormalAdd ForceAddFlag = 0

	// TIMDeleteTypeSingle 单项删好友
	TIMDeleteTypeSingle DeleteType = "Delete_Type_Single"
	// TIMDeleteTypeBoth 双向删好友
	TIMDeleteTypeBoth DeleteType = "Delete_Type_Both"
)

type AddFriendItem struct {
	ToAccount  string `json:"To_Account"` //
	AddSource  string // 来源 前缀 AddSource_Type_ 关键字 8字节 英文
	Remark     string // 备注 96字节
	GroupName  string // 分组
	AddWording string // 附言 256字节
}
type ImportFriendItem struct {
	ToAccount  string `json:"To_Account"` //
	AddSource  string // 来源 前缀 AddSource_Type_ 关键字 8字节 英文
	Remark     string // 备注 96字节
	RemarkTime int
	GroupName  []string // 分组
	AddWording string   // 附言 256字节
	AddTime    int
}

type ResultItem struct {
	To_Account string
	ResultCode int
	ResultInfo string
}
type FriendAddResp struct {
	CommonResp
	ResultItem      []ResultItem
	Fail_Account    []string
	Invalid_Account []string
}

// FriendAdd 添加好友
func (api *TimApp) FriendAdd(fromAccount string, addFriendItems []AddFriendItem, addType AddType, forceAddFlags ForceAddFlag) (*FriendAddResp, error) {
	reqdata := map[string]interface{}{
		"From_Account":  fromAccount,
		"AddFriendItem": addFriendItems,
		"AddType":       addType,
		"ForceAddFlags": forceAddFlags,
	}
	resp := new(FriendAddResp)
	err := api.do(snsSvc, "friend_add", reqdata, resp)
	return resp, err
}

// FriendImport 导入好友
func (api *TimApp) FriendImport(fromAccount string, importFriendItems []ImportFriendItem) (*FriendAddResp, error) {
	reqdata := map[string]interface{}{
		"From_Account":  fromAccount,
		"AddFriendItem": importFriendItems,
	}
	resp := new(FriendAddResp)
	err := api.do(snsSvc, "friend_import", reqdata, resp)
	return resp, err
}

// FriendDelete 删除好友
func (api *TimApp) FriendDelete(fromAccount string, toAccount []string, deleteType DeleteType) (*FriendAddResp, error) {
	reqdata := map[string]interface{}{
		"From_Account": fromAccount,
		"To_Account":   toAccount,
		"DeleteType":   deleteType,
	}
	resp := new(FriendAddResp)
	err := api.do(snsSvc, "friend_delete", reqdata, resp)
	return resp, err
}

// FriendDeleteAll 清空好友
func (api *TimApp) FriendDeleteAll(fromAccount string) (*CommonResp, error) {
	reqdata := map[string]string{
		"From_Account": fromAccount,
	}
	resp := new(CommonResp)
	err := api.do(snsSvc, "friend_delete_all", reqdata, resp)
	return resp, err
}
