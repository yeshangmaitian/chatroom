# chatroom

#### 一、概要

本地电脑建立服务端和客户端，使用go实现简易聊天室

#### 二、具体功能

###### 1、注册、登录、注销

###### 2、当有新成员上线时，提醒上线。也可以单独查看在线列表

###### 3、给全部在线用户发送信息，即群发信息

#### 三、运行步骤

###### 1、下载redis，开启redis服务端

​    redis下载网址：https://redis.io/download/

###### 2、开启服务端(端口为8889)

在路径 chat/server/main 路径下开启服务端，可以通过 go run main.go redis.go 运行

或者通过 go build 指令编译后执行 exe 文件

###### 3、开启客户端

在路径 chat/client/main 路径下开启服务端，可以通过 go run main.go 运行

或者通过 go build 指令编译后执行 exe 文件

