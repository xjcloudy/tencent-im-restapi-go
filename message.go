package tim

import (
	"time"
)

type TimMsgType string
type SyncType int
type CountType int

const (
	// TIMText 文本消息
	TIMText TimMsgType = "TIMTextElem"
	// TIMLocation 位置消息
	TIMLocation TimMsgType = "TIMLocationElem"
	// TIMFace 表情消息
	TIMFace TimMsgType = "TIMFaceElem"
	// TIMCustom 自定义消息
	TIMCustom TimMsgType = "TIMCustomElem"

	maxOfflineLife int = 604800

	// SyncToFrom 同步到发送方
	SyncToFrom SyncType = 1
	// NotSyncToFrom 不同步到发送方
	NotSyncToFrom SyncType = 2

	// SyncAndCount 实时消息导入 计入未读
	SyncAndCount CountType = 1

	// SyncNotCount 历史消息导入 不计入未读
	SyncNotCount CountType = 2
)

// Message 消息结构体
type Message struct {
	FromAccount string `json:"From_Account"`
	SendTime    int
	Random      int
	MsgBody     []MsgElement
}

// MsgElement 消息体中`MsgBody`中的值的结构体
type MsgElement struct {
	MsgType    TimMsgType
	MsgContent interface{} // 具体类型的消息结构体数据
}

// MsgText 文本消息 结构体
type MsgText struct {
	Text string
}

// MsgLocation 位置消息结构体
type MsgLocation struct {
	Desc      string
	Latitude  float64
	Longitude float64
}

// MsgFace 表情消息结构体
type MsgFace struct {
	Index int
	Data  string
}

// MsgCustom 自定义消息结构体
type MsgCustom struct {
	Data  interface{}
	Desc  string
	Ext   string
	Sound string
}

// SendMsgData 发消息接口参数结构体
type SendMsgData struct {
	SyncFromOldSystem CountType
	SyncOtherMachine  SyncType
	FromAccount       string `json:"From_Account"`
	// 单发的时候是string,批量发的时候是[]string
	ToAccount    interface{} `json:"To_Account"`
	MsgLifeTime  int
	MsgRandom    int
	MsgTimeStamp int64
	MsgBody      []MsgElement
}

// SendMsgResp 发消息接口返回数据结构体
type SendMsgResp struct {
	CommonResp
	MsgTime int
}

// SendMsg 发送单聊消息 100次/秒
func (api *TimApp) SendMsg(sendMsgData SendMsgData) (*SendMsgResp, error) {
	resp := new(SendMsgResp)
	if sendMsgData.MsgLifeTime > maxOfflineLife {
		sendMsgData.MsgLifeTime = maxOfflineLife
	}
	if sendMsgData.MsgTimeStamp <= 0 {
		sendMsgData.MsgTimeStamp = time.Now().Unix()
	}
	err := api.do(openImSvc, "sendmsg", sendMsgData, resp)
	return resp, err
}

// ImportMsg 导入单聊消息
func (api *TimApp) ImportMsg(msgData SendMsgData) (*CommonResp, error) {
	resp := new(CommonResp)
	err := api.do(openImSvc, "importmsg", msgData, resp)
	return resp, err
}

type ErrorAccount struct {
	ToAccount string `json:"To_Account"`
	ErrorCode int
}
type BatchSendMsgResp struct {
	CommonResp
	ErrorList []ErrorAccount
}

// BatchSendMsg 批量发 通用方法
func (api *TimApp) BatchSendMsg(batchMsgData SendMsgData) (*BatchSendMsgResp, error) {
	resp := new(BatchSendMsgResp)
	err := api.do(openImSvc, "batchsendmsg", batchMsgData, resp)
	return resp, err
}

// BatchSendTextMsg 批量发文本消息快捷方法
func (api *TimApp) BatchSendTextMsg(fromAccount string, toAccount []string, content string, sync SyncType) (*BatchSendMsgResp, error) {
	reqdata := SendMsgData{
		MsgLifeTime:      maxOfflineLife,
		MsgTimeStamp:     time.Now().Unix(),
		FromAccount:      fromAccount,
		ToAccount:        toAccount,
		MsgBody:          []MsgElement{},
		SyncOtherMachine: sync,
	}
	reqdata.ToAccount = toAccount

	msgEle := MsgElement{}
	msgEle.MsgType = TIMText

	txtMsg := new(MsgText)
	txtMsg.Text = content

	msgEle.MsgContent = txtMsg

	reqdata.MsgBody = append(reqdata.MsgBody, msgEle)

	return api.BatchSendMsg(reqdata)

}
