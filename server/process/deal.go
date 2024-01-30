package process

import (
	"fmt"
	"gocode/chat/common/message"
	"gocode/chat/server/utils"
	"io"
	"net"
)

type Process struct {
	Conn net.Conn
}

func (this *Process) Deal() (err error) {
	for {
		tf := &utils.Transfer{
			Conn: this.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务器端也退出..")
				return err
			} else {
				fmt.Println("tf.ReadPkg() err=", err.Error())
				return err
			}
		}
		err = this.dealProcess(&mes)
		if err != nil {
			return err
		}
	}
}
func (this *Process) dealProcess(mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		userProcess := &UserProcess{
			Conn: this.Conn,
		}
		err = userProcess.LoginProcess(mes.Data)
	case message.RegisterMesType:
		userProcess := &UserProcess{
			Conn: this.Conn,
		}
		err = userProcess.RegisterProcess(mes.Data)
	case message.SmsToAllMesType:
		smsProcess := &SmsProcess{
			Conn: &this.Conn,
		}
		smsProcess.ToAllUsersMes(mes)
	case message.LogoutMesType:
		userProcess := &UserProcess{
			Conn: this.Conn,
		}
		err = userProcess.LogoutProcess(mes.Data)
	default:
		fmt.Println("服务端无法处理的信息类型")
	}
	return
}
