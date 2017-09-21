/*
author: foolbread
file: user/common_user.go
date: 2017/9/12
*/
package user

type commonUser struct {
	user string
	pass string
}

func newCommonUser(usr string,pass string)*commonUser{
	r := new(commonUser)
	r.user = usr
	r.pass = pass

	return r
}

func (u *commonUser)GetUserType()int{
	return COMMON_USER
}

func (u *commonUser)GetUserName()string{
	return u.user
}