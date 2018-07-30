package tim

import (
	"math/rand"
	"testing"
	"time"
)

func TestCreateGroup(t *testing.T) {
	def := GroupDefine{
		GroupID:         "go1",
		Name:            "goSdkTestGroup",
		MaxMemberCount:  100,
		Type:            TIMGroupPublic,
		FaceURL:         "AAAAA",
		OwnerAccount:    "1",
		ApplyJoinOption: TIMGroupNeedPermission,
	}
	resp, err := testTimAPP.CreateGroup(def)
	if err != nil {
		t.Error(err)
	}
	if resp.ActionStatus == ResponseFail {
		t.Error(resp.ErrorInfo, resp.ErrorCode)
	}
}
func TestAddGroupMember(t *testing.T) {
	groupMember := GroupMemberAccount{}
	groupMember.MemberAccount = "10"
	resp, err := testTimAPP.AddGroupMember("go1", []GroupMemberAccount{groupMember}, 1)
	if err != nil {
		t.Error(err)
	}
	if resp.ActionStatus == ResponseFail {
		t.Error(resp.ErrorInfo)
	}
}

func TestDeleteGroupMember(t *testing.T) {
	resp, err := testTimAPP.DeleteGroupMember("go1", []string{"10"}, "测试", 1)
	if err != nil {
		t.Error(err)
	}
	if resp.ActionStatus == ResponseFail {
		t.Error(resp.ErrorInfo)
	}
}
func TestImportGroupMember(t *testing.T) {
	m1 := ImportMemberAccount{
		JoinTime:     time.Date(2018, 7, 24, 12, 0, 0, 0, utc8).Unix(),
		UnreadMsgNum: 10,
	}
	m1.MemberAccount = "10"
	m1.Role = "Admin"

	m2 := ImportMemberAccount{
		JoinTime:     time.Date(2018, 7, 28, 12, 0, 0, 0, utc8).Unix(),
		UnreadMsgNum: 1,
	}
	m2.MemberAccount = "notexists"
	resp, err := testTimAPP.ImportGroupMember("go1", []ImportMemberAccount{m1, m2})
	if err != nil {
		t.Error(err)
	}
	if resp.ActionStatus == ResponseFail {
		t.Error(resp.ErrorInfo)
	}
	t.Log(resp.MemberList)
}
func TestImportGroup(t *testing.T) {
	def := GroupDefine{
		GroupID:         "go1",
		Name:            "goSdkTestGroup",
		MaxMemberCount:  1000,
		Type:            TIMGroupPublic,
		FaceURL:         "AAAAA",
		OwnerAccount:    "1",
		CreateTime:      1532275200,
		ApplyJoinOption: TIMGroupNeedPermission,
	}
	resp, err := testTimAPP.CreateGroup(def)
	if err != nil {
		t.Error(err)
	}
	if resp.ActionStatus == ResponseFail {
		t.Error(resp.ErrorInfo, resp.ErrorCode)
	}
}

func TestDestroyGroup(t *testing.T) {
	groupID := "go1"
	resp, err := testTimAPP.DestroyGroup(groupID)
	if err != nil {
		t.Error(err)
	}
	if resp.ActionStatus == ResponseFail {
		t.Error(resp.ErrorCode, resp.ErrorInfo)
	}
}

func TestImportGroupMsg(t *testing.T) {
	m1 := Message{
		FromAccount: "1",
		Random:      rand.Int31(),
		SendTime:    time.Date(2018, 7, 30, 16, 12, 3, 0, utc8).Unix(),
		MsgBody: []MsgElement{
			MsgElement{
				MsgType: TIMText,
				MsgContent: MsgText{
					Text: "import message1ERER",
				},
			},
		},
	}
	resp, err := testTimAPP.ImportGroupMsg("go1", []Message{m1})
	if err != nil {
		t.Error(err)
	}
	if resp.ActionStatus == ResponseFail {
		t.Error(resp.ErrorCode, resp.ErrorInfo)
	}
	t.Log(resp)
}

func TestSendGroupMsg(t *testing.T) {
	text := MsgElement{
		MsgType: TIMText,
		MsgContent: MsgText{
			Text: "代发群消息",
		},
	}
	msgGroup := GroupMessage{
		GroupID:     "go1",
		FromAccount: "1",
		Random:      rand.Int31(),
		MsgBody:     []MsgElement{text},
	}
	resp, err := testTimAPP.SendGroupMsg(msgGroup)

	if err != nil {
		t.Error(err)
	}
	if resp.ActionStatus == ResponseFail {
		t.Error(resp.ErrorInfo)
	}
	t.Log(resp.MsgSeq)
}
