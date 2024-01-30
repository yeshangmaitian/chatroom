package process

import (
	"fmt"
	"gocode/chat/common/user"
)

var Online map[int]user.OnlineUser

func ShowAllOnlineUser() {
	fmt.Println("显示在线用户：")
	if len(Online) == 0 {
		fmt.Println("你是第一个在线用户哦，暂无其他在线用户~")
		return
	} else {
		for i, v := range Online {
			fmt.Printf("账号 %s(%d) 在线,IP为：%s\n", v.Name, i, v.IpAddress)
		}
	}
}

func AddAOnlineUser(user *user.OnlineUser) {
	Online[user.UserId] = *user
	fmt.Printf("账号 %s(%d) 上线了,IP为：%s\n", user.Name, user.UserId, user.IpAddress)
}
func AddManyOnlineUser(users []user.OnlineUser) {
	Online = make(map[int]user.OnlineUser)
	for _, v := range users {
		Online[v.UserId] = v
	}
}
