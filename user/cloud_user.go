/*
author: foolbread
file: user/cloud_user.go
date: 2017/9/21
*/
package user

type cloudUser struct {
	user string
	pass string
}

func newCloudUser(usr string,pass string)*cloudUser{
	r := new(cloudUser)
	r.user = usr
	r.pass = pass

	return r
}

func (u *cloudUser)GetUserType()int{
	return CLOUD_USER
}

func (u *cloudUser)GetUserName()string{
	return u.user
}