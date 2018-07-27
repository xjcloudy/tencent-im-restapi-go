package tim

import "testing"

func TestAccountImport(t *testing.T) {
	resp, err := testTimAPP.AccountImport("goTest", "gopher", "")
	if err != nil {
		t.Error(err)
	}
	if resp.ActionStatus == ResponseFail {
		t.Error(resp.ErrorInfo)
	}

}

func TestQueryAccountStatus(t *testing.T) {
	acccount := "goTest"
	resp, err := testTimAPP.QueryState([]string{acccount})
	if err != nil {
		t.Error(err)
	}
	if resp.ActionStatus == ResponseFail {
		t.Error(resp.ErrorInfo)
	} else {
		for _, acs := range resp.QueryResult {
			t.Logf("user `%s` is %s ", acs.ToAccount, acs.State)
		}
	}

}
