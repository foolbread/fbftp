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
	UNKNOWTYPE_USER = 0
	COMMON_USER = 1
	CLOUD_USER = 2
)

func UserLogin(usr string,pwd string)FBFTPUser{
	var ret FBFTPUser

	ret = g_userManager.getUser(usr)

	if ret == nil{
		return ret
	}

	if ret.getPassWord() != pwd{
		return nil
	}

	return ret
}

var g_userManager *fbFTPUserManager = newfbFTPUserManager()

func initUserManager(){
	g_userManager.updateCommonUser()

	g_userManager.updateCloudUser()
}

type fbFTPUserManager struct {
	userMap map[string]FBFTPUser
}

func newfbFTPUserManager()*fbFTPUserManager{
	r := new(fbFTPUserManager)
	r.userMap = make(map[string]FBFTPUser)

	return r
}

func (u *fbFTPUserManager)updateCommonUser(){
	commonUsers := config.GetConfig().GetAllCommonUsers()
	for _,v := range commonUsers.Users{
		golog.Info("[init common user]:","<username>:",v.UserName,"<password>:",v.PassWord)
		u.userMap[v.UserName] = newCommonUser(v.UserName,v.PassWord)
	}
}

func (u *fbFTPUserManager)updateCloudUser(){
	cloudUsers := config.GetConfig().GetAllCloudUsers()
	for _,v := range cloudUsers.Users{
		golog.Info("[init cloud user]:","<username>:",v.UserName,"<password>:",v.PassWord,"<bucket>:",v.Bucket)
		u.userMap[v.UserName] = newCloudUser(v.UserName,v.PassWord,v.Bucket,v.AccKey,v.SecKey,v.EndPoint,v.Token)
	}
}

func (u *fbFTPUserManager)getUser(usr string)FBFTPUser{
	return u.userMap[usr]
}

type FBFTPUser interface {
	GetUserType()UserType
	GetUserName()string
	getPassWord()string
	GetUserExternInfo()FBFTPUserExternInfo
}

type FBFTPUserExternInfo interface {
}