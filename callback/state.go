package callback

const State_StateChange = "State.StateChange"

type AccountStateInfo struct {
	Action    string
	ToAccount string `json:"To_Account"`
	Reason    string
}
type StateChange struct {
	CallbackCommand string
	Info            AccountStateInfo
}
