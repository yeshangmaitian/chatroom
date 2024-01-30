package process

import (
	"encoding/json"
	"fmt"
	"gocode/chat/common/message"
	"gocode/chat/server/utils"

	"net"
)

type SmsProcess struct {
	Conn *net.Conn
}

func (this *SmsProcess) ToAllUsersMes(mes *message.Message) {
	data, _ := json.Marshal(mes)
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
