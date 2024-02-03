package controllers

type RespCodeType int64

const (
	CodeSuccess RespCodeType = 1000 + iota
	CodeValidateParams
	CodeValidPassword
	CodeUserExists
	CodeUserNotExists
	CodeServerBusy
)

var responseCodeMessage = map[RespCodeType]string{
	CodeSuccess:        "success",
	CodeValidateParams: "Validate parameters",
	CodeValidPassword:  "Validate password",
	CodeUserExists:     "User exists",
	CodeUserNotExists:  "User not exists or password error",
	CodeServerBusy:     "Service busy",
}

func (code RespCodeType) getMsg() string {
	msg, ok := responseCodeMessage[code]
	if !ok {
		msg = "Service busy"
	}
	return msg

}
