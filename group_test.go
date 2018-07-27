package tim

import (
	"testing"
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
