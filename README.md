# chatroom
go实现简易聊天室，本地电脑开服务端和客户端

运行步骤如下：
###### 1、下载redis，开启redis服务端

​    redis下载网址：https://redis.io/download/

###### 2、开启服务端

在路径 chat/server/main 路径下开启服务端，可以通过 go run main.go redis.go 运行

或者通过 go build 指令编译后执行 exe 文件

###### 3、开启客户端

在路径 chat/client/main 路径下开启服务端，可以通过 go run main.go 运行

或者通过 go build 指令编译后执行 exe 文件
