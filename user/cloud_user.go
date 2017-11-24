/*
author: foolbread
file: user/cloud_user.go
date: 2017/9/21
*/
package user

type CloudUserExternInfo struct {
	Bucket string
	AccKey string
	SecKey string
	EndPoint string
	Token  string
}

type cloudUser struct {
	user string
	pass string
	extern CloudUserExternInfo
}

func newCloudUser(usr string,pass string,bucket string,acckey string,seckey string,epoint string,token string)*cloudUser{
	r := new(cloudUser)
	r.user = usr
	r.pass = pass
	r.extern.Bucket = bucket
	r.extern.AccKey = acckey
	r.extern.SecKey = seckey
	r.extern.EndPoint = epoint
	r.extern.Token = token

	return r
}

func (u *cloudUser)getPassWord()string{
	return u.pass
}

func (u *cloudUser)GetUserType()UserType{
	return CLOUD_USER
}

func (u *cloudUser)GetUserName()string{
	return u.user
}

func (u *cloudUser)GetUserExternInfo()FBFTPUserExternInfo{
	return &u.extern
}