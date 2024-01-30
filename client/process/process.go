package process

import (
	"encoding/json"
	"errors"
	"fmt"
	"gocode/chat/client/utils"
	"gocode/chat/common/message"
	"gocode/chat/common/user"

	"net"
)

type Process struct {
	Conn *net.Conn
	user user.User
	key  int
}

func (this *Process) menu() {
	fmt.Println("欢迎", this.user.Name, "登录微信")

	for {
		this.key = 1
		fmt.Println("操作选项如下：")
		fmt.Println("————1、查看在线列表")
		fmt.Println("————2、单线发送信息")
		fmt.Println("————3、发送全局信息")
		fmt.Println("————4、修改个人信息")
		fmt.Println("————5、注销微信账号")
		fmt.Println("————0、退出微信登录")
		fmt.Println("请输入选项表号0-5")
		fmt.Scanln(&this.key)
		switch this.key {
		case 1:
			this.showAllOnlineUser()
		case 2:
			this.sendToOneMessage()
		case 3:
			this.sendToAllMessage()
		case 5:
			err := this.userLogout()
			if err == nil {
				return
			}
			if err.Error() == "取消" {
				break
			}
			fmt.Println("this.userLogout() = ", err)
			fmt.Println("注销失败", err)
		case 0:
			fmt.Println("————0、退出微信登录")
			return
		default:
			fmt.Println("您输入的选项有误，请重新输入")
		}
	}
}

// 和服务器保持通讯
func serverProcessMes(process *Process) {
	//创建一个transfer实例, 不停的读取服务器发送的消息
	tf := &utils.Transfer{
		Conn: *process.Conn,
	}
	for {
		mes, err := tf.ReadPkg()
		if err != nil {
			// fmt.Println("tf.ReadPkg err=", err)
			return
		}
		// fmt.Println("mes", mes)
		//如果读取到消息，又是下一步处理逻辑
		switch mes.Type {
		case message.NoticeUserOnlineMesType: // 有人上线了
			process.addAOnlineUser(mes.Data)
		case message.ResLogoutMesType:
			process.logoutLastPart(mes.Data)
		case message.NoticeMeOnlineMesType:
			process.addManyOnlineUser(mes.Data)
		case message.SmsToAllMesType:
			process.showToAllMessage(mes.Data)
		default:
			fmt.Println("服务器端返回了未知的消息类型")
		}
	}
}

func (this *Process) showToAllMessage(mesData string) {
	sp := &SmsProcess{
		Conn: this.Conn,
	}
	sp.ShowToAllMessage(mesData, this.user.Id)
}
func (this *Process) addManyOnlineUser(mesData string) {
	var onlineUsers []user.OnlineUser
	json.Unmarshal([]byte(mesData), &onlineUsers)
	AddManyOnlineUser(onlineUsers)
}
func (this *Process) logoutLastPart(mesData string) {
	userProcess := &UserProcess{
		Conn: this.Conn,
	}
	userProcess.LogoutLastPart(mesData)
}
func (this *Process) addAOnlineUser(mesData string) {
	var onlineUser user.OnlineUser
	json.Unmarshal([]byte(mesData), &onlineUser)
	AddAOnlineUser(&onlineUser)
}
func (this *Process) showAllOnlineUser() {
	ShowAllOnlineUser()
}

func (this *Process) sendToOneMessage() {
}
func (this *Process) sendToAllMessage() {
	fmt.Println("发送全局信息")
	fmt.Println("请输入要发送的内容")
	content := ""
	fmt.Scanln(&content)
	sp := &SmsProcess{
		Conn: this.Conn,
	}
	sp.SendToAllProcess(content, this.user.Id)

}

func (this *Process) userLogout() (err error) {
	fmt.Println("注销微信账号")
	fmt.Println("确认是否注销微信账号：yes/no")
	var userSelect string
	for {
		userSelect = ""
		fmt.Scanln(&userSelect)
		switch userSelect {
		case "yes":
			up := &UserProcess{
				Conn: this.Conn,
			}
			err = up.Logout(this.user.Id)
			return
		case "no":
			fmt.Println("已取消注销微信账号")
			err := errors.New("取消")
			return err
		default:
			fmt.Println("您的输入有误，请重新输入~")
			fmt.Println("确认是否注销微信账号:yes/no")
		}
	}
}
