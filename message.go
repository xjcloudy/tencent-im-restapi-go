package tim

import (
	"time"
)

type TimMsgType string
type SyncType int

const (
	// TIMText 文本消息
	TIMText TimMsgType = "TIMTextElem"
	// TIMLocation 位置消息
	TIMLocation TimMsgType = "TIMLocationElem"
	// TIMFace 表情消息
	TIMFace TimMsgType = "TIMFaceElem"
	// TIMCustom 自定义消息
	TIMCustom TimMsgType = "TIMCustomElem"
	// TIMSound 音频消息
	TIMSound TimMsgType = "TIMSoundElem"
	// TIMImage 图片消息
	TIMImage TimMsgType = "TIMImageElem"
	// TIMFile 文件消息
	TIMFile TimMsgType = "TIMFileElem"

	maxOfflineLife int = 604800

	SyncToFrom    SyncType = 1
	NotSyncToFrom SyncType = 2
)

type TimMessage struct {
	From_Account string
	SendTime     int
	Random       int
	MsgBody      []TimMsgElement
}
type TimMsgElement struct {
	MsgType    TimMsgType
	MsgContent interface{}
}
type TimMsgText struct {
	Text string
}
type TimMsgLocation struct {
	Desc      string
	Latitude  float64
	Longitude float64
}
type TimMsgFace struct {
	Index int
	Data  string
}
type TimMsgCustom struct {
	Data  string
	Desc  string
	Ext   string
	Sound string
}
type TimMsgSound struct {
	UUID   string
	Size   int
	Second int
}
type TimMsgImage struct {
	UUID           string
	ImageFormat    int
	ImageInfoArray []TimMsgImageInfo
}
type TimMsgImageInfo struct {
	Type   int
	Size   int
	Width  int
	Height int
	URL    string
}
type TimMsgFile struct {
	UUID     string
	FileSize int
	FileName string
}

type SendMsgData struct {
	SyncOtherMachine SyncType
	FromAccount      string `json:"From_Account"`
	// 单发的时候是string,批量发的时候是[]string
	ToAccount    interface{} `json:"To_Account"`
	MsgLifeTime  int
	MsgRandom    int
	MsgTimeStamp int64
	MsgBody      []TimMsgElement
}
type SendMsgResp struct {
	CommonResp
	MsgTime int
}

// SendMsg 发送单聊消息 100次/秒
func (api *TimApp) SendMsg(sendMsgData SendMsgData) (*SendMsgResp, error) {
	resp := new(SendMsgResp)
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
	To_Account string
	ErrorCode  int
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
		MsgBody:          []TimMsgElement{},
		SyncOtherMachine: sync,
	}
	reqdata.ToAccount = toAccount

	msgEle := TimMsgElement{}
	msgEle.MsgType = TIMText

	txtMsg := new(TimMsgText)
	txtMsg.Text = content

	msgEle.MsgContent = txtMsg

	reqdata.MsgBody = append(reqdata.MsgBody, msgEle)

	return api.BatchSendMsg(reqdata)

}
