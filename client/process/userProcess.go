package process

import (
	"encoding/json"
	"errors"
	"fmt"
	"gocode/chat/client/utils"
	"gocode/chat/common/message"
	"gocode/chat/common/user"

	"net"
	"time"
)

type UserProcess struct {
	Conn *net.Conn
}

func (this *UserProcess) getConn() net.Conn {
	return *this.Conn
}
func (this *UserProcess) Logout(id int) (err error) {

	// 2、包装数据
	var mes message.Message
	mes.Type = message.LogoutMesType

	var LogoutMes message.LogoutMes
	LogoutMes.UserId = id

	data, err := json.Marshal(LogoutMes)
	if err != nil {
		fmt.Println("json.Marshal(LogoutMes) err=", err)
	}
	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal(mes) err=", err)
	}

	// 3、发送数据
	tf := utils.Transfer{
		Conn: this.getConn(),
	}
	//发送data给服务器端
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("发送注销信息错误 err=", err)
	}
	time.Sleep(time.Second * 3)
	return
}

func (this *UserProcess) LogoutLastPart(mesData string) {
	resLogoutMes := message.ResLogoutMes{}
	//将mes的Data部分反序列化成 ResLogoutMes
	err := json.Unmarshal([]byte(mesData), &resLogoutMes)
	// fmt.Println("注销的后部分~")
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(mes.Data), &resRegisterMes)")
		return
	}
	if resLogoutMes.Code == 200 {
		fmt.Println("注销成功，正在退出微信~")
		return
	} else {
		err = errors.New(resLogoutMes.Error)
		return
	}
}
func (this *UserProcess) Register(userId int, userPwd, userName string) {
	//1. 链接到服务器
	conn, err := net.Dial("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	//延时关闭
	defer conn.Close()

	// 2、包装数据
	var mes message.Message
	mes.Type = message.RegisterMesType

	var registerMes message.RegisterMes
	registerMes.UserId = userId
	registerMes.UserPwd = userPwd
	registerMes.UserName = userName

	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("json.Marshal(registerMes) err=", err)
	}
	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal(mes) err=", err)
	}

	// 3、发送数据
	tf := &utils.Transfer{
		Conn: conn,
	}
	//发送data给服务器端
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("注册发送信息错误 err=", err)
	}

	// 这里还需要处理服务器端返回的消息.
	mes, err = tf.ReadPkg() // mes 就是
	if err != nil {
		fmt.Println("tf.ReadPkg() err=", err)
		return
	}

	//将mes的Data部分反序列化成 ResRegisterMes
	var resRegisterMes message.ResRegisterMes
	err = json.Unmarshal([]byte(mes.Data), &resRegisterMes)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(mes.Data), &resRegisterMes)")
		return
	}

	if resRegisterMes.Code == 200 {
		fmt.Println("已注册成功，请登录")
	} else {
		fmt.Println(resRegisterMes.Error)
	}
}
func (this *UserProcess) Login(userId int, userPwd string) {
	//1. 链接到服务器
	conn, err := net.Dial("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	//延时关闭
	defer conn.Close()

	// 2、包装数据
	var mes message.Message
	mes.Type = message.LoginMesType

	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal(loginMes) err=", err)
	}
	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal(mes) err=", err)
	}

	// 3、发送数据
	tf := &utils.Transfer{
		Conn: conn,
	}
	//发送data给服务器端
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("注册发送信息错误 err=", err)
	}

	// 这里还需要处理服务器端返回的消息.
	mes, err = tf.ReadPkg() // mes 就是
	if err != nil {
		fmt.Println("tf.ReadPkg() err=", err)
		return
	}

	//将mes的Data部分反序列化成 LoginResMes
	var ResloginMes message.ResLoginMes
	err = json.Unmarshal([]byte(mes.Data), &ResloginMes)
	if err != nil {
		fmt.Println("部分反序列json.Unmarshal([]byte(mes.Data), &ResloginMes)")
		return
	}

	if ResloginMes.Code == 200 {
		user := user.User{}
		err = json.Unmarshal([]byte(ResloginMes.Data), &user)
		if err != nil {
			fmt.Println("200json.Unmarshal([]byte(mes.Data), &ResloginMes)")
			return
		}
		fmt.Println("登录成功")
		process := Process{
			Conn: &conn,
			user: user,
		}
		go serverProcessMes(&process)
		process.menu()
	} else {
		fmt.Println(ResloginMes.Error)
	}
}
