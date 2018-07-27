package tim

import (
	"testing"
)

func TestPortraitSet(t *testing.T) {
	pi := ProfileKV{
		Tag:   "Tag_Profile_IM_Nick",
		Value: "改个名",
	}

	resp, err := testTimAPP.PortraitSet("goTest", []ProfileKV{pi})
	if err != nil {
		t.Error(err)
	}
	if resp.ActionStatus == ResponseFail {
		t.Error(resp.ErrorInfo)
	}

}
func TestPortraitGet(t *testing.T) {
	resp, err := testTimAPP.PortraitGet([]string{"1", "goTest"}, []string{"Tag_Profile_IM_Nick"})
	if err != nil {
		t.Error(err)
	}
	if resp.ActionStatus == ResponseFail {
		t.Error(resp.ErrorInfo)
	} else {
		t.Log(resp.FailAccount, resp.InvalidAccount, resp.UserProfileItem)
	}

}
