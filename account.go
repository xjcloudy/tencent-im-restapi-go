package tim

import (
	"errors"
)

// AccountImport 独立模式账号导入。100次/秒
func (api *TimApp) AccountImport(id, nick, avatar string) (*CommonResp, error) {
	reqdata := map[string]string{
		"Identifier": id,
		"Nick":       nick,
		"FaceUrl":    avatar,
	}
	resp := new(CommonResp)
	err := api.do(iMopenLoginSvc, "account_import", reqdata, resp)
	return resp, err
}

type MultiAccountImportResp struct {
	CommonResp
	FailAccounts []string
}

// MultiaccountImport 独立模式账号批量导入
// 10次/秒 100个/次
func (api *TimApp) MultiaccountImport(accounts []string) (*MultiAccountImportResp, error) {
	if len(accounts) > 100 {
		return nil, errors.New("批量导入账号每次不能超过100个")
	}
	reqdata := map[string][]string{"Accounts": accounts}
	resp := new(MultiAccountImportResp)
	err := api.do(iMopenLoginSvc, "multicaccount_import", reqdata, resp)
	return resp, err
}

// Kick 账号登录态失效(踢下线)
// 1000次/秒
func (api *TimApp) Kick(identifier string) (*CommonResp, error) {
	reqdata := map[string]string{
		"Identifier": identifier,
	}
	resp := new(CommonResp)
	err := api.do(iMopenLoginSvc, "kick", reqdata, resp)
	return resp, err
}

type AccountState struct {
	To_Account string
	State      string
}
type QueryAccountStateResp struct {
	CommonResp
	QueryResult []AccountState
}

// QueryState 获取用户当前的登录状态。
func (api *TimApp) QueryState(toAccount []string) (*QueryAccountStateResp, error) {
	if len(toAccount) > 500 {
		return nil, errors.New("一次最多查询500个账号")
	}
	reqdata := map[string][]string{
		"To_Account": toAccount,
	}
	resp := new(QueryAccountStateResp)
	err := api.do(openImSvc, "querystate", reqdata, resp)
	return resp, err
}
