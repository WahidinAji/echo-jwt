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

type AuthResponse struct {
	Code        int    `json:"code"`
	Status      bool   `json:"status"`
	Message     string `json:"message"`
	JWTResponse JWTResponse `json:"data"`
}
type JWTResponse struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}
