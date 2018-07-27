package tim

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type TimApp struct {
	AppID       string
	Identifiner string
	Sig         string
	Debug       bool
}
type CommonResp struct {
	ActionStatus string
	ErrorInfo    string
	ErrorCode    int
	ErrorDisplay string
}

const (
	// service name

	openLoginSvc = "im_open_login_svc"
	openImSvc    = "openim"
	snsSvc       = "sns"
	groupSvc     = "group_open_http_svc"
	profileSvc   = "profile"
	version      = "v4"
	timAPIHost   = "https://console.tim.qq.com"
	contentType  = "json"
	apn          = '0'

	// ResponseOK 成功
	ResponseOK = "OK"
	// ResponseFail 错误
	ResponseFail = "FAIL"
)

func (api *TimApp) httpReq(service, cmd string, params []byte) ([]byte, error) {
	param := make(url.Values)
	param.Add("usersig", api.Sig)
	param.Add("identifier", api.Identifiner)
	param.Add("sdkappid", api.AppID)
	param.Add("contenttype", "json")

	url := fmt.Sprintf("%s/%s/%s/%s?%s", timAPIHost, version, service, cmd, param.Encode())
	resp, err := http.Post(url, "text/plain", bytes.NewReader(params))
	if err != nil {
		return nil, err
	}
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return responseBody, nil
}
func (api *TimApp) do(service, cmd string, params interface{}, resp interface{}) error {
	jsonData, encodeError := json.Marshal(params)
	if encodeError != nil {
		return encodeError
	}

	if api.Debug {
		fmt.Println("post data", string(jsonData))
	}
	responsBody, err := api.httpReq(service, cmd, jsonData)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(responsBody, resp); err != nil {
		return err
	}
	return nil
}
