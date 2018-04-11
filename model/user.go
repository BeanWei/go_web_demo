package model

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/BeanWei/go-web-demo/config"
	"github.com/garyburd/redigo/redis"
)

// User 用户
type User struct {
	ID         int64      `form:"userId" xorm:"pk not null autoincr"`
	Nickname   string     `form:"nickname" xorm:"varchar(20) not null"`
	Account    string     `form:"account" xorm:"varchar(60) not null"`
	Password   string     `form:"password" xorm:"varchar(60) not null"`
	Status     int        `form:"status" xorm:"int(1) default 0"`
	CreateTime *time.Time `xorm:"datetime created"`
}

// CheckPassword 密码验证
func (u *User) CheckPassword(password string) bool {
	if password == "" || u.Password == "" {
		return false
	}
	return u.EncryptPassword(password, u.Salt()) == u.Password
}

//EncryptPassword 加密
func (u *User) EncryptPassword(password, salt string) (hash string) {
	password = fmt.Sprintf("%x", md5.Sum([]byte(password)))
	hash = salt + password + config.ServerConfig.PassSalt
	hash = salt + fmt.Sprintf("%x", md5.Sum([]byte(hash)))
	return
}

//Salt 加盐
func (u *User) Salt() string {
	var userSalt string
	if u.Password == "" {
		userSalt = strconv.Itoa(int(time.Now().Unix()))
	} else {
		userSalt = u.Password[0:10]
	}
	return userSalt
}

//UserFromRedis 从redis中取出用户信息
func UserFromRedis(userID int) (User, error) {
	loginUser := fmt.Sprintf("%s%d", LoginUser, userID)
	RedisConn := RedisPool.Get()
	defer RedisConn.Close()
	userBytes, err := redis.Bytes(RedisConn.Do("GET", loginUser))
	if err != nil {
		fmt.Println(err)
		return User{}, errors.New("未登录")
	}
	var user User
	bytesErr := json.Unmarshal(userBytes, &user)
	if bytesErr != nil {
		fmt.Println(bytesErr)
		return user, errors.New("未登录")
	}
	return user, nil
}

//UserToRedis 将用户信息存到redis
func UserToRedis(user User) error {
	userBytes, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return errors.New("error")
	}
	loginUserKey := fmt.Sprintf("%s%d", LoginUser, user.ID)
	RedisConn := RedisPool.Get()
	if _, redisErr := RedisConn.Do("SET", loginUserKey, userBytes, "EX", config.ServerConfig.TokenMaxAge); redisErr != nil {
		fmt.Println("redis set failed: ", redisErr.Error())
		return errors.New("error")
	}
	return nil
}

//用户等级
const (
	UserLevelAdmin  = 0
	UserLevelNormal = 1
)
