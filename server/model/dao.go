package model

import (
	"encoding/json"
	"fmt"
	"gocode/chat/common/user"

	"github.com/garyburd/redigo/redis"
)

type UserDao struct {
	pool *redis.Pool
}

// 定义一个全局变量
var (
	MyUserDao *UserDao
)

func NewUserDao(pool *redis.Pool) *UserDao {
	return &UserDao{
		pool: pool,
	}
}
func (this *UserDao) getUserById(conn redis.Conn, userId int) (user user.User, err error) {
	res, err := redis.String(conn.Do("hget", "users", userId))
	if err != nil {
		if err == redis.ErrNil {
			err = ERROR_USER_NOTEXISTS
		}
		return
	}
	err = json.Unmarshal([]byte(res), &user)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(res),user) err", err)
		return
	}
	return
}
func (this *UserDao) Logout(id int) (err error) {
	conn := this.pool.Get()
	_, err = this.getUserById(conn, id)
	if err == ERROR_USER_NOTEXISTS {
		return
	} else if err == nil {
		_, err = conn.Do("del", "users", id)
		if err != nil {
			fmt.Println("del  err=", err)
			return
		}
		fmt.Println("服务器已注销")
		return
	}
	return
}

func (this *UserDao) Register(id int, pwd, name string) (err error) {
	conn := this.pool.Get()
	_, err = this.getUserById(conn, id)
	if err == ERROR_USER_NOTEXISTS {
		user := user.User{
			Id:   id,
			Name: name,
			Pwd:  pwd,
		}
		var data []byte
		data, err = json.Marshal(user)
		if err != nil {
			fmt.Println("data, err = json.Marshal(user) err=", err)
			return
		}

		_, err = conn.Do("HSet", "users", id, string(data))

		r, a := redis.Strings(conn.Do("HMGet", "users", id))
		if a != nil {
			fmt.Println("hget  err=", a)

		}
		for i, v := range r {
			fmt.Printf("r[%d]=%s\n", i, v)
		}

		fmt.Println("服务器已注册")
		if err != nil {
			fmt.Println("conn.Do(, string(data)) err=")
			return err
		}
		return
	} else if err == nil {
		return ERROR_USER_EXISTS
	}
	return
}
func (this *UserDao) Login(id int, pwd string) (user user.User, err error) {
	conn := this.pool.Get()
	user, err = this.getUserById(conn, id)
	if err != nil {
		return
	}
	if user.Pwd != pwd {
		err = ERROR_USER_PWD
		return
	}
	return
}
