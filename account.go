package tim

import (
	"errors"
)

const (
	// TimAccountOnline 客户端登录后和云通信后台有长连接
	TimAccountOnline = "Online"
	// TimAccountOffline 客户端主动退出登录或者客户端自上一次登录起 7 天之内未登录过
	TimAccountOffline = "Offline"
	// TimAccountPushOnline iOS 客户端退到后台或进程被杀或因网络问题掉线，进入 PushOnline 状态，此时仍然可以接收消息离线 APNS推送。
	// 注意，云通信后台只会保存 PushOnline 状态 7 天时间，若从掉线时刻起 7 天之内未登录过，则进入 Offline 状态。
	TimAccountPushOnline = "PushOnline"
)

// AccountImport 独立模式账号导入。100次/秒
func (api *TimApp) AccountImport(id, nick, avatar string) (*CommonResp, error) {
	reqdata := map[string]string{
		"Identifier": id,
		"Nick":       nick,
		"FaceUrl":    avatar,
	}
	resp := new(CommonResp)
	err := api.do(openLoginSvc, "account_import", reqdata, resp)
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
	err := api.do(openLoginSvc, "multicaccount_import", reqdata, resp)
	return resp, err
}

// Kick 账号登录态失效(踢下线)
// 1000次/秒
func (api *TimApp) Kick(identifier string) (*CommonResp, error) {
	reqdata := map[string]string{
		"Identifier": identifier,
	}
	resp := new(CommonResp)
	err := api.do(openLoginSvc, "kick", reqdata, resp)
	return resp, err
}

// AccountState 在线状态
type AccountState struct {
	ToAccount string `json:"To_Account"`
	State     string
}

// QueryAccountStateResp 用户登录状态返回值
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
