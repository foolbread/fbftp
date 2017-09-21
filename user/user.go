/*
author: foolbread
file: user/user.go
date: 2017/8/10
*/
package user

import (
	"github.com/foolbread/fbcommon/golog"
	"github.com/foolbread/fbftp/config"
)

func InitUser(){
	golog.Info("fbftp user initing......")

	initUserManager()
}

type UserType int

const (
	COMMON_USER = 1
	CLOUD_USER = 2
)

func UserLogin(usr string,pwd string)FBFTPUser{
	var ret FBFTPUser

	if !g_userManager.checkUser(usr,pwd){
		return ret
	}

	switch g_userManager.getUserType(usr) {
	case COMMON_USER:
		ret = newCommonUser(usr,pwd)
	case CLOUD_USER:
		ret = newCloudUser(usr,pwd)
	}

	return ret
}

var g_userManager *fbFTPUserManager = newfbFTPUserManager()

func initUserManager(){
	g_userManager.updateCommonUser()

	g_userManager.updateCloudUser()
}

type fbFTPUserManager struct {
	userMap map[string]string
	userTypeMap map[string]UserType
}

func newfbFTPUserManager()*fbFTPUserManager{
	r := new(fbFTPUserManager)
	r.userMap = make(map[string]string)
	r.userTypeMap = make(map[string]UserType)

	return r
}

func (u *fbFTPUserManager)updateCommonUser(){
	commonUsers := config.GetConfig().GetAllCommonUsers()
	for _,v := range commonUsers.Users{
		u.userMap[v.UserName] = v.PassWord
		u.userTypeMap[v.UserName] = COMMON_USER
	}
}

func (u *fbFTPUserManager)updateCloudUser(){
	cloudUsers := config.GetConfig().GetAllCloudUsers()
	for _,v := range cloudUsers.Users{
		u.userMap[v.UserName] = v.PassWord
		u.userTypeMap[v.UserName] = CLOUD_USER
	}
}

func (u *fbFTPUserManager)getUserType(usr string)UserType{
	return u.userTypeMap[usr]
}

func (u *fbFTPUserManager)checkUser(usr string,password string)bool{
	ps,ok := u.userMap[usr]
	if !ok{
		return false
	}

	return ps == password
}

type FBFTPUser interface {
	GetUserType()int
	GetUserName()string
}