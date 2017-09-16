/*
author: foolbread
file: acl/acl_user.go
date: 2017/9/16
*/
package acl

type AclUser struct {
	username string
	workDir string
	authority int
}

func newAclUser(usr string)*AclUser{
	r := new(AclUser)

	return r
}

func (a *AclUser)AllowRead(pa string)bool{
	return true
}

func (a *AclUser)AllowWrite(pa string)bool{
	return true
}