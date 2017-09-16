/*
author: foolbread
file: user/user.go
date: 2017/8/10
*/
package user

import (
	"github.com/foolbread/fbcommon/golog"
)

func InitUser(){
	golog.Info("fbftp user initing......")

	g_userManager.initUserManager()
}

type UserType int

const (
	COMMON_USER = 1
)

func UserLogin(usr string,pwd string)FBFTPUser{
	var ret FBFTPUser

	if !g_userManager.checkUser(usr,pwd){
		return ret
	}

	switch g_userManager.getUserType(usr) {
	case COMMON_USER:
		ret = NewCommonUser(usr,pwd)
	}

	return ret
}

type FBFTPUser interface {

}

var g_userManager *fbFTPUserManager = newfbFTPUserManager()

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

func (u *fbFTPUserManager)initUserManager(){

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

