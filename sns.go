package tim

type AddType string
type ForceAddFlag int
type DeleteType string

const (
	// 单项
	ADD_TYPE_SINGLE AddType = "Add_Type_Single"
	// 双向
	ADD_TYPE_BOTH AddType = "Add_Type_Both"

	// 强制添加
	FORCE_ADD ForceAddFlag = 1
	// 常规添加
	NORMAL_ADD ForceAddFlag = 0

	//单项删好友
	DELETE_TYPE_SINGLE DeleteType = "Delete_Type_Single"
	// 双向删好友
	DELETE_TYPE_BOTH DeleteType = "Delete_Type_Both"
)

type AddFriendItem struct {
	To_Account string //
	AddSource  string // 来源 前缀 AddSource_Type_ 关键字 8字节 英文
	Remark     string // 备注 96字节
	GroupName  string // 分组
	AddWording string // 附言 256字节
}
type ImportFriendItem struct {
	To_Account string //
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

// 添加好友
func (api *TimApp) FriendAdd(from_account string, addFriendItems []AddFriendItem, addType AddType, forceAddFlags ForceAddFlag) (error, *FriendAddResp) {
	req_data := map[string]interface{}{
		"From_Account":  from_account,
		"AddFriendItem": addFriendItems,
		"AddType":       addType,
		"ForceAddFlags": forceAddFlags,
	}
	resp := new(FriendAddResp)
	err := api.do(snsSvc, "friend_add", req_data, resp)
	return err, resp
}

// 导入好友
func (api *TimApp) FriendImport(from_account string, importFriendItems []ImportFriendItem) (error, *FriendAddResp) {
	req_data := map[string]interface{}{
		"From_Account":  from_account,
		"AddFriendItem": importFriendItems,
	}
	resp := new(FriendAddResp)
	err := api.do(snsSvc, "friend_import", req_data, resp)
	return err, resp
}

// 删除好友
func (api *TimApp) FriendDelete(from_account string, to_account []string, deleteType DeleteType) (error, *FriendAddResp) {
	req_data := map[string]interface{}{
		"From_Account": from_account,
		"To_Account":   to_account,
		"DeleteType":   deleteType,
	}
	resp := new(FriendAddResp)
	err := api.do(snsSvc, "friend_delete", req_data, resp)
	return err, resp
}

// 清空好友
func (api *TimApp) FriendDeleteAll(from_account string) (error, *CommonResp) {
	req_data := map[string]string{
		"From_Account": from_account,
	}
	resp := new(CommonResp)
	err := api.do(snsSvc, "friend_delete_all", req_data, resp)
	return err, resp
}
