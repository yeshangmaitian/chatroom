package process

import (
	"encoding/json"
	"fmt"
	"gocode/chat/common/message"
	"gocode/chat/server/model"
	"gocode/chat/server/utils"

	"net"
)

type UserProcess struct {
	Conn net.Conn
}

func (this *UserProcess) LogoutProcess(mesData string) (err error) {
	var LogoutMes message.LogoutMes
	err = json.Unmarshal([]byte(mesData), &LogoutMes)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(mesData), &registerMes) err=", err)
		return
	}
	err = model.MyUserDao.Logout(LogoutMes.UserId)

	var data []byte
	var resLogoutMes message.ResLogoutMes
	if err != nil {
		if err == model.ERROR_USER_NOTEXISTS {
			resLogoutMes.Code = 500
			resLogoutMes.Error = err.Error()
		} else {
			resLogoutMes.Code = 505
			resLogoutMes.Error = "服务器内部错误..."
		}
	} else {
		resLogoutMes.Code = 200
	}
	data, err = json.Marshal(resLogoutMes)
	if err != nil {
		fmt.Println("json.Marshal(resRegisterMes) err=", err)
		return
	}

	var mes message.Message
	mes.Type = message.ResLogoutMesType
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal(resLoginMes) err=", err)
		return
	}

	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg([]byte(data))
	if err != nil {
		fmt.Println("json.Marshal(resLoginMes) err=", err)
		return
	}
	return
}
func (this *UserProcess) RegisterProcess(mesData string) (err error) {
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mesData), &registerMes)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(mesData), &registerMes) err=", err)
		return
	}
	err = model.MyUserDao.Register(registerMes.UserId, registerMes.UserPwd, registerMes.UserName)

	var resRegisterMes message.ResRegisterMes
	var data []byte
	if err != nil {
		if err == model.ERROR_USER_PWD {
			resRegisterMes.Code = 403
			resRegisterMes.Error = err.Error()
		} else if err == model.ERROR_USER_NOTEXISTS {
			resRegisterMes.Code = 500
			resRegisterMes.Error = err.Error()
		} else {
			resRegisterMes.Code = 505
			resRegisterMes.Error = "服务器内部错误..."
		}
	} else {
		resRegisterMes.Code = 200
	}
	data, err = json.Marshal(resRegisterMes)
	if err != nil {
		fmt.Println("json.Marshal(resRegisterMes) err=", err)
		return
	}

	var mes message.Message
	mes.Type = message.ResRegisterMesType
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal(resLoginMes) err=", err)
		return
	}

	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg([]byte(data))
	if err != nil {
		fmt.Println("json.Marshal(resLoginMes) err=", err)
		return
	}
	return
}
func (this *UserProcess) LoginProcess(mesData string) (err error) {
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mesData), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(mesData), &loginMes) err=", err)
		return
	}
	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)

	var data []byte
	var resLoginMes message.ResLoginMes
	if err != nil {
		if err == model.ERROR_USER_PWD {
			resLoginMes.Code = 403
			resLoginMes.Error = err.Error()
		} else if err == model.ERROR_USER_NOTEXISTS {
			resLoginMes.Code = 500
			resLoginMes.Error = err.Error()
		} else {
			resLoginMes.Code = 505
			resLoginMes.Error = "服务器内部错误..."
		}
	} else {
		resLoginMes.Code = 200
		data, err = json.Marshal(user)
		if err != nil {
			fmt.Println("json.Marshal(resLoginMes) err=", err)
			return
		}
		resLoginMes.Data = string(data)

	}
	data, err = json.Marshal(resLoginMes)
	if err != nil {
		fmt.Println("json.Marshal(resLoginMes) err=", err)
		return
	}

	var mes message.Message
	mes.Type = message.ResLoginMesType
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal(resLoginMes) err=", err)
		return
	}

	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg([]byte(data))
	if err != nil {
		fmt.Println("json.Marshal(resLoginMes) err=", err)
		return
	}
	if resLoginMes.Code == 200 {
		noticeOnline(&this.Conn, user)
		noticeToMe(&this.Conn, user)
	}
	return
}
