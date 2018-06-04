/*
author: foolbread
file: acl/acl_user.go
date: 2017/9/16
*/
package acl

type UserACL struct {
	username  string
	workpath  string
	authority int
	allowMap map[string]struct{}
}

func newUserACL(usr string)*UserACL {
	r := new(UserACL)

	return r
}

func (a *UserACL)GetWorkPath()string{
	return a.workpath
}

func (a *UserACL)AllowRead()bool{
	return a.authority&ONLY_READ != 0
}

func (a *UserACL)AllowWrite()bool{
	return a.authority&ONLY_WRITE != 0
}

func (a *UserACL)IsAllowList(client string)bool{
	if len(a.allowMap) == 0{
		return true
	}

	_,allow := a.allowMap[client]

	return allow
}