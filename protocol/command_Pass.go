/*
author: foolbread
file: protocol/command_pass.go
date: 2017/9/12
*/
package protocol

import (
	"github.com/foolbread/fbftp/session"
	"github.com/foolbread/fbftp/user"
	"github.com/foolbread/fbftp/acl"
)

type commandPass struct {

}

func (p *commandPass)IsExtend()bool{
	return false
}

func (p *commandPass)RequireAuth()bool{
	return false
}

func (p *commandPass)RequireParam()bool{
	return true
}

func (p *commandPass)Execute(sess *session.FTPSession, arg string)error{
	var login bool = false
	userobj := user.UserLogin(sess.UserName,arg)
	if userobj != nil{
		sess.UserInfo = userobj
		useracl := acl.GetUserAcl(sess.UserName)
		if useracl != nil{
			sess.UserAcl = useracl
			sess.CurPath = "/"
			login = true
		}
	}

	if login{
		sess.CtrlCon.WriteMsg(FTP_LOGINOK,"Login successful.")
	}else{
		sess.CtrlCon.WriteMsg(FTP_BADPBSZ,"Login incorrect.")
	}

	return nil
}