/*
author: foolbread
file: user/common_user.go
date: 2017/9/12
*/
package user

type CommonUser struct {
	user string
	pass string
}

func NewCommonUser(usr string,pass string)*CommonUser{
	r := new(CommonUser)
	r.user = usr
	r.pass = pass

	return r
}