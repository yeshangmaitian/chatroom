package process

import (
	"encoding/json"
	"fmt"
	"gocode/chat/client/utils"
	"gocode/chat/common/message"

	"net"
	"time"
)

type SmsProcess struct {
	Conn *net.Conn
}

func (this *SmsProcess) SendToAllProcess(content string, id int) {
	mes := message.Message{}
	mes.Type = message.SmsToAllMesType
	smsToAllMes := message.SmsToAllMes{}
	smsToAllMes.UserId = id
	smsToAllMes.Content = content
	smsToAllMes.SendTime = time.Now()

	data, err := json.Marshal(smsToAllMes)
	if err != nil {
		fmt.Println("json.Marshal(smsToAllMes) err=", err)
	}
	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal(mes) err=", err)
	}

	tf := &utils.Transfer{
		Conn: *this.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println(" tf.Writepkg(data) err=", err)
	}
}

func (this *SmsProcess) ShowToAllMessage(mesData string, id int) {
	smsToAllMes := message.SmsToAllMes{}
	err := json.Unmarshal([]byte(mesData), &smsToAllMes)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(mesData), &smsToAllMes) err=", err)
	}
	if smsToAllMes.UserId != id {
		fmt.Printf("新消息\n %d在%s发出群发信息:%s", smsToAllMes.UserId, smsToAllMes.SendTime.Format("2006.01.02 15:04:05"), smsToAllMes.Content)
	}
}
