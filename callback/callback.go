package callback

import (
	"github.com/xjcloudy/tencent-im-restapi-go"
)

type CallbackResp struct {
	ActionStatus string
	ErrorInfo    string
	ErrorCode    int
}

func CallbackSuccess() CallbackResp {
	return CallbackResp{
		ActionStatus: tim.ResponseOK,
		ErrorInfo:    "",
		ErrorCode:    0,
	}
}
func CallbackFail(errorcode int, errorinfo string) CallbackResp {
	if errorcode == 0{
		errorcode =1
	}
	return CallbackResp{
		ActionStatus: tim.ResponseFail,
		ErrorInfo:    errorinfo,
		ErrorCode:    errorcode,
	}
}
