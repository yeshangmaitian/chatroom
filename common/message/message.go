package message

import (
	"gocode/chat/common/user"
	"time"
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

const (
	LoginMesType            = "LoginMes"
	ResLoginMesType         = "ResLoginMes"
	RegisterMesType         = "RegisterMes"
	ResRegisterMesType      = "ResRegisterMes"
	LogoutMesType           = "LogoutMes"
	ResLogoutMesType        = "ResLogoutMes"
	NoticeUserOnlineMesType = "NoticeUserOnlineMes"
	NoticeMeOnlineMesType   = "NoticeMeOnlineMes"
	SmsToAllMesType         = "SmsToAllMes"
)

type LoginMes struct {
	UserId  int    `json:"userId"`
	UserPwd string `json:"userPwd"`
}

type ResLoginMes struct {
	Code  int    `json:"code"`
	Data  string `json:"data"`
	Error string `json:"error"`
}

type RegisterMes struct {
	UserId   int    `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}

type ResRegisterMes struct {
	Code  int    `json:"code"`
	Data  string `json:"data"`
	Error string `json:"error"`
}

type SmsToAllMes struct {
	UserId   int
	Content  string
	SendTime time.Time
}
type LogoutMes struct {
	UserId int `json:"userId"`
}

type ResLogoutMes struct {
	Code  int    `json:"code"`
	Data  string `json:"data"`
	Error string `json:"error"`
}
type NoticeMeOnlineMes struct {
	uSlice []user.OnlineUser
}
