/*
author: foolbread
file: acl/acl_user.go
date: 2017/9/16
*/
package acl

import (
	"strings"
)

type UserACL struct {
	username  string
	workpath  string
	authority int
}

func newUserACL(usr string)*UserACL {
	r := new(UserACL)

	return r
}

func (a *UserACL)GetWorkPath()string{
	return a.workpath
}

func (a *UserACL)AllowRead(pa string)bool{
	if a.authority&ONLY_READ != 0{
		return 	strings.HasPrefix(pa,a.workpath)
	}

	return false
}

func (a *UserACL)AllowWrite(pa string)bool{
	if a.authority&ONLY_WRITE != 0{
		return strings.HasPrefix(pa,a.workpath)
	}

	return false
}