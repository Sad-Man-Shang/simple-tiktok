package service

import (
	"fmt"
	"github.com/kuanyuh/simple-tiktok/dao"
	"testing"
)

func TestIsExist(t *testing.T) {
	dao.Init()
	IsExist("13388338208@qq.com")
}

func TestGetUserinfoById(t *testing.T) {
	dao.Init()
	GetUserinfoById("1")
}

func TestGetUser(t *testing.T) {
	dao.Init()
	GetUser("1119526092@qq.com","123456")
}

func TestGetRandomString(t *testing.T) {
	dao.Init()
	t.Logf(GetRandomString(1))
}

func TestIsFollow(t *testing.T) {
	dao.Init()
	user := GetUserinfoById("1")
	IsFollow(1, &user)
	fmt.Printf("%v",user)
}