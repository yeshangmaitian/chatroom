package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"gocode/chat/common/message"

	"net"
)

// 收发信息的结构体
type Transfer struct {
	Conn net.Conn
	Buf  [8096]byte
}

/*
读信息：确保信息收取完整
1、获取发送端先传递的长度
2、获取收到的数据及数据长度
3、检验长度完整后，返回信息
*/
func (this Transfer) ReadPkg() (mes message.Message, err error) {
	fmt.Println("读取客户端", this.Conn.RemoteAddr(), "发来的数据...")
	//conn.Read 在conn没有被关闭的情况下，才会阻塞。
	// 获取长度的字节到Buf
	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil {
		return
	}
	// 将收到的Buf四个字节通过大端模式转化为数字
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[:4])

	// 根据pkgLen读取数据到对应的Buf中。检查数据传输失败或者数据丢失
	n, err := this.Conn.Read(this.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		return
	}

	fmt.Println(this.Conn.RemoteAddr(), "发来", n, "位信息：", string(this.Buf[:n]))

	// 反序列化 -> message.Message
	err = json.Unmarshal(this.Buf[:n], &mes)
	if err != nil {
		return
	}
	return
}

/*
发信息：确保信息传递完整
第一步，先发数据的长度
第二步，再发数据，
第三步，第一步中的长度是否与第二步中发送数据后收到的长度相等.相等则成功，否则失败
*/
func (this Transfer) WritePkg(data []byte) (err error) {
	//先发送一个长度给对方
	var pkgLen uint32
	pkgLen = uint32(len(data))
	binary.BigEndian.PutUint32(this.Buf[0:4], pkgLen)
	n, err := this.Conn.Write(this.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("this.Conn.Write(this.Buf[:4]) fail=", err)
		return
	}

	//发送data本身，检查数据传输失败或者数据丢失
	n, err = this.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("this.Conn.Write(data) fail=", err)
		return
	}
	return
}
