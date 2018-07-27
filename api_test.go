package tim

import (
	"encoding/json"
	"fmt"
	"testing"
)

var testTimAPP *TimApp

func init() {
	testTimAPP = new(TimApp)
	testTimAPP.AppID = "yourAppId"
	testTimAPP.Identifiner = "yourIdentifiner"
	testTimAPP.Sig = "yourUserSig"
	testTimAPP.Debug = true
}

func TestGetSig(t *testing.T) {
	resp := new(MultiAccountImportResp)
	GetResp(resp)
	fmt.Println(resp.FailAccounts)
}
func GetResp(v interface{}) {
	jsonStr := `{"ActionStatus":"OK","ErrorInfo":"以前ok",
		"ErrorCode":0,"FailAccounts":[]
	 }`
	if err := json.Unmarshal([]byte(jsonStr), v); err != nil {
		fmt.Println(err)
	}
}
