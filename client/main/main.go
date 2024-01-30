package main

import (
	"fmt"
	"gocode/chat/client/process"
)

// 定义两个变量，一个表示用户id, 一个表示用户密码
var userId int
var userPwd string
var userName string

func main() {
	//接收用户的选择
	var key int
	//判断是否还继续显示菜单
	//var loop = true

	for {
		key = 0
		fmt.Println("----------------欢迎来到微信------------")
		fmt.Println("\t\t\t 1 登陆微信")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 0 退出系统")
		fmt.Println("\t\t\t 请选择(0-2):")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登陆微信")
			fmt.Println("请输入用户的id")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户的密码")
			fmt.Scanf("%s\n", &userPwd)
			// 完成登录
			//1. 创建一个UserProcess的实例
			up := &process.UserProcess{}
			up.Login(userId, userPwd)

		case 2:
			fmt.Println("注册用户")
			fmt.Println("请输入用户id:")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户密码:")
			fmt.Scanf("%s\n", &userPwd)
			fmt.Println("请输入用户名字(nickname):")
			fmt.Scanf("%s\n", &userName)
			//2. 调用UserProcess，完成注册的请求、
			up := &process.UserProcess{}
			up.Register(userId, userPwd, userName)
		case 0:
			fmt.Println("退出系统")
			return
		default:
			fmt.Println("你的输入有误，请重新输入")
		}
	}
}
