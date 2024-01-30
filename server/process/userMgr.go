package process

import (
	"encoding/json"
	"fmt"
	"gocode/chat/common/message"
	"gocode/chat/common/user"
	"gocode/chat/server/utils"

	"net"
)

var online []OnlineInfo

type OnlineInfo struct {
	Conn      net.Conn
	Id        int
	Name      string
	IpAddress string
}

func noticeOnline(conn *net.Conn, userOld user.User) {
	var mes message.Message
	mes.Type = message.NoticeUserOnlineMesType

	//给其他用户发送登录信息
	newUser := user.OnlineUser{
		UserId:    userOld.Id,
		IpAddress: (*conn).RemoteAddr().String(),
		Name:      userOld.Name,
	}
	data, err := json.Marshal(newUser)
	if err != nil {
		fmt.Println(err)
	}
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range online {
		tf := &utils.Transfer{
			Conn: v.Conn,
		}
		err := tf.WritePkg(data)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func noticeToMe(conn *net.Conn, userOld user.User) {
	// 给自己发送其他人的在线信息
	onlineSlice := make([]user.OnlineUser, 0)
	for _, v := range online {
		onlineSlice = append(onlineSlice, user.OnlineUser{
			UserId:    v.Id,
			IpAddress: v.IpAddress,
			Name:      v.Name,
		})
	}
	data, _ := json.Marshal(onlineSlice)

	var mes message.Message
	mes.Type = message.NoticeMeOnlineMesType
	mes.Data = string(data)
	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println(err)
	}
	tf := &utils.Transfer{
		Conn: *conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println(err)
	}
	online = append(online, OnlineInfo{
		Conn:      *conn,
		Id:        userOld.Id,
		Name:      userOld.Name,
		IpAddress: (*conn).RemoteAddr().String(),
	})
}
