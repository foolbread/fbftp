/*
author: foolbread
file: protocol/command_pwd.go
date: 2017/9/25
*/
package protocol

import (
	"github.com/foolbread/fbftp/session"
	"fmt"
)

type commandPwd struct {

}

func (p *commandPwd)IsExtend()bool{
	return false
}

func (p *commandPwd)RequireAuth()bool{
	return true
}

func (p *commandPwd)RequireParam()bool{
	return false
}

func (p *commandPwd)Execute(sess *session.FTPSession, arg string)error{
	return 	sess.CtrlCon.WriteMsg(FTP_PWDOK,fmt.Sprintf("\"%s\" is the current directory",sess.CurPath))
}