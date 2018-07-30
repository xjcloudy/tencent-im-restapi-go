package tim

import (
	"errors"
)

type GroupType string
type ApplyJoinOption string

const (
	// 群类型

	// TIMGroupPrivate 私人群
	TIMGroupPrivate GroupType = "Private"
	// TIMGroupPublic 公开群
	TIMGroupPublic GroupType = "Public"
	// TIMGroupChatRoom 聊天室
	TIMGroupChatRoom GroupType = "ChatRoom"
	// TIMGroupAvChatRoom 在线直播群
	TIMGroupAvChatRoom GroupType = "AVChatRoom"
	// TIMGroupBChatRoom 在线广播大群
	TIMGroupBChatRoom GroupType = "BChatRoom"
	// 加群处理方式
	TIMGroupFreeJoin       ApplyJoinOption = "FreeAccess"
	TIMGroupNeedPermission ApplyJoinOption = "NeedPermission"
	TIMGroupDisableApply   ApplyJoinOption = "DisableApply"
)

type KV struct {
	Key   string
	Value string
}
type GroupMember struct {
	MemberAccount string `json:"Member_Account"` // 成员（必填）
	Role          string `json:",omitempty"`     // 赋予该成员的身份，目前备选项只有Admin（选填）
}
type GroupMemberAccount struct {
	GroupMember
	AppMemberDefinedData []KV `json:",omitempty"` //群成员维度自定义字段（选填）
}

// GroupDefine 群信息
type GroupDefine struct {
	OwnerAccount    string               `json:"Owner_Account,omitempty"` // 群主的UserId（选填）
	Type            GroupType            // 群组类型：Private/Public/ChatRoom(不支持AVChatRoom和BChatRoom)（必填）
	GroupID         string               `json:"GroupId"` //用户自定义群组ID（选填）
	Name            string               // 群名称（必填）
	Introduction    string               `json:",omitempty"`        // 群简介（选填）
	Notification    string               `json:",omitempty"`        // 群公告（选填）
	FaceURL         string               `json:"FaceUrl,omitempty"` // 群头像URL（选填）
	MaxMemberCount  int                  `json:",omitempty"`        // 最大群成员数量（选填）
	ApplyJoinOption ApplyJoinOption      `json:",omitempty"`        // 申请加群处理方式（选填）
	AppDefinedData  []KV                 `json:",omitempty"`        // 群组维度的自定义字段（选填）
	CreateTime      int64                `json:",omitempty"`        //建群时间
	MemberList      []GroupMemberAccount `json:",omitempty"`        // 初始群成员列表，最多500个（选填）

}
type CreateGroupResp struct {
	CommonResp
	GroupID string
}

// CreateGroup 创建群
func (api *TimApp) CreateGroup(groupDefine GroupDefine) (*CreateGroupResp, error) {
	resp := new(CreateGroupResp)
	err := api.do(groupSvc, "create_group", groupDefine, resp)
	return resp, err
}

// DestroyGroup 删除群组
func (api *TimApp) DestroyGroup(groupID string) (*CommonResp, error) {
	req := map[string]string{
		"GroupId": groupID,
	}
	resp := new(CommonResp)
	err := api.do(groupSvc, "destroy_group", req, resp)
	return resp, err

}

type MemberAccountResult struct {
	GroupMember
	Result int // 加群返回值使用
}
type AddGroupMemberResp struct {
	CommonResp
	MemberList []MemberAccountResult
}
type ImportGroupMsgResp struct {
	CommonResp
	ImportMsgResults []ImportMsgResult
}
type ImportMsgResult struct {
	Result  int
	MsgSeq  int
	MsgTime int64
}
type ImportMemberAccount struct {
	GroupMember
	JoinTime     int64
	UnreadMsgNum int
}

// AddGroupMember 加群
func (api *TimApp) AddGroupMember(groupID string, memberList []GroupMemberAccount, slience int) (*AddGroupMemberResp, error) {
	reqdata := map[string]interface{}{
		"GroupId":    groupID,
		"Slicence":   slience,
		"MemberList": memberList,
	}
	resp := new(AddGroupMemberResp)
	err := api.do(groupSvc, "add_group_member", reqdata, resp)
	return resp, err
}

// DeleteGroupMember 删除成员
func (api *TimApp) DeleteGroupMember(groupID string, memberToDelAccount []string, reason string, silence int) (*CommonResp, error) {
	reqdata := map[string]interface{}{
		"GroupId":             groupID,
		"MemberToDel_Account": memberToDelAccount,
		"Reason":              reason,
		"Silence":             silence,
	}
	resp := new(CommonResp)
	err := api.do(groupSvc, "delete_group_member", reqdata, resp)
	return resp, err

}

// GroupMessage 消息结构体
type GroupMessage struct {
	GroupID               string   `json:"GroupId"`
	FromAccount           string   `json:"From_Account,omitempty"`
	MsgPriority           string   `json:",omitempty"` // 消息优先级
	ForbidCallbackControl []string `json:",omitempty"` // 消息回调禁止开关，只对单条消息有效，ForbidBeforeSendMsgCallback 表示禁止发消息前回调，ForbidAfterSendMsgCallback 表示禁止发消息后回调。
	OnlineOnlyFlag        int      `json:",omitempty"` // 1: 表示只在线下发(只有在线群成员才能收到)，不存离线及漫游
	Random                int32    `json:",omitempty"`
	MsgBody               []MsgElement
}

// ImportGroupMsg 导入群聊消息
func (api *TimApp) ImportGroupMsg(groupID string, msgList []Message) (*ImportGroupMsgResp, error) {
	if len(msgList) > 20 {
		return nil, errors.New("导入群聊消息条数不能超过20条")
	}
	reqdata := map[string]interface{}{
		"GroupId": groupID,
		"MsgList": msgList,
	}
	resp := new(ImportGroupMsgResp)
	err := api.do(groupSvc, "import_group_msg", reqdata, resp)
	return resp, err
}

// ImportGroupMember 导入群成员
func (api *TimApp) ImportGroupMember(groupID string, memberList []ImportMemberAccount) (*AddGroupMemberResp, error) {
	if len(memberList) > 500 {
		return nil, errors.New("导入群成员数量不能超过500")
	}
	reqdata := map[string]interface{}{
		"GroupId":    groupID,
		"MemberList": memberList,
	}
	resp := new(AddGroupMemberResp)
	err := api.do(groupSvc, "import_group_member", reqdata, resp)
	return resp, err
}

// ImportGroup 导入群 用于同步群数据
func (api *TimApp) ImportGroup(groupDefine GroupDefine) (*CreateGroupResp, error) {
	resp := new(CreateGroupResp)
	err := api.do(groupSvc, "import_group", groupDefine, resp)
	return resp, err
}

// SendGroupMsgResp 发送群消息的返回值结构体
type SendGroupMsgResp struct {
	SendMsgResp
	MsgSeq int
}

// SendGroupMsg 在群里发消息
func (api *TimApp) SendGroupMsg(groupMsg GroupMessage) (*SendGroupMsgResp, error) {
	resp := new(SendGroupMsgResp)
	err := api.do(groupSvc, "send_group_msg", groupMsg, resp)
	return resp, err
}
