package main

import (
	"fmt"
	"gocode/chat/server/model"
	"gocode/chat/server/process"

	"net"
	"time"
)

func keep(conn net.Conn) {
	defer conn.Close()
	process := &process.Process{
		Conn: conn,
	}
	err := process.Deal()
	if err != nil {
		fmt.Println("process.Deal() err=", err.Error())
		return
	}
}

func init() {
	initRedisPool("localhost:6379", 16, 0, 300*time.Second)
	initUserDao()
}

func initUserDao() {
	//这里的pool 本身就是一个全局的变量
	//这里需要注意一个初始化顺序问题
	//initPool, 在 initUserDao
	model.MyUserDao = model.NewUserDao(redisPool)
}
func main() {
	fmt.Println("服务端开启开启")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("net.Listen err=", err.Error())
	}
	defer func() {
		listen.Close()
		fmt.Println("监听关闭")
	}()

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err=", err.Error())
		}
		go keep(conn)
	}

}
