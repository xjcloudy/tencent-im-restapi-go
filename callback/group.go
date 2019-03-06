package callback

import tim "github.com/xjcloudy/tencent-im-restapi-go"

type GroupCallbackBase struct {
	CallbackCommand string
	GroupID         string        `json:"GroupId"` // 群组 ID
	Type            tim.GroupType // 群组类型
}
type GroupCallbackAfterSendMsg struct {
	GroupCallbackBase
	FromAccount string           `json:"From_Account"` // 发送者
	MsgSeq      int64            // 消息的序列号
	MsgTime     int64            // 消息的时间
	MsgBody     []tim.MsgElement // 消息体
}
type GroupCallbackBeforeSendMsg struct {
	GroupCallbackBase
	FromAccount     string           `json:"From_Account"`     // 发送者
	OperatorAccount string           `json:"Operator_Account"` // 请求的发起者
	Random          int32            // 随机数
	MsgBody         []tim.MsgElement // 消息体
}
