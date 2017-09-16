/*
author: foolbread
file: protocol/command_pass.go
date: 2017/9/12
*/
package protocol

import (
	"github.com/foolbread/fbftp/session"
	"github.com/foolbread/fbftp/user"
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
	usrobj := user.UserLogin(sess.UserName,arg)
	if usrobj != nil{
		sess.UserInfo = usrobj
		sess.WriteMsg(FTP_LOGINOK,"Login successful.")
	}else{
		sess.WriteMsg(FTP_BADPBSZ,"Login incorrect.")
	}

	return nil
}