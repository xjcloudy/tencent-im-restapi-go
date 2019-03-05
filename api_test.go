package tim

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

var testTimAPP *TimApp

var utc8 *time.Location

func init() {
	testTimAPP = new(TimApp)
	testTimAPP.AppID = "xxx"
	testTimAPP.Identifiner = "xxx"
	testTimAPP.Sig = "xxx"
	testTimAPP.Debug = true

	utc8, _ = time.LoadLocation("Asia/Shanghai")

	rand.Seed(time.Now().Unix())

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
