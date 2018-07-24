package tim

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type TimApp struct {
	AppId       string
	Identifiner string
	Sig         string
}
type CommonResp struct {
	ActionStatus string
	ErrorInfo    string
	ErrorCode    int
	ErrorDisplay string
}

const (
	// service name

	iMopenLoginSvc = "im_open_login_svc"
	openImSvc      = "openim"
	snsSvc         = "sns"
	groupSvc       = "group_open_http_svc"

	timAPIHost  = "https://console.tim.qq.com/v4"
	contentType = "json"
	apn         = '0'
)

func (api *TimApp) httpReq(service, cmd string, params []byte) {
	param := new(url.Values)
	param.Add("identifier", api.Identifiner)
	if resp, err := http.Post(fmt.Sprintf("%s/%s/%s?%s", timAPIHost, service, cmd, params), "text/plain", bytes.NewReader(params)); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp)
	}
}
func (api *TimApp) do(service, cmd string, params interface{}, resp interface{}) error {
	jsonData, encodeError := json.Marshal(params)
	if encodeError != nil {
		return encodeError
	}
	fmt.Println(service, cmd, string(jsonData))
	return nil
}
