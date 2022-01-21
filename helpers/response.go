package helpers

type MsgDel struct {
	Code   int  `json:"code"`
	Status bool `json:"status"`
}

type UserMsg struct {
	MsgDel   MsgDel `json:"response"`
	Username string `json:"username"`
	Message  string `json:"message"`
}
type JWTResponse struct {
	MsgDel   MsgDel `json:"response"`
	Username string `json:"username"`
	Token    string `json:"token"`
}
