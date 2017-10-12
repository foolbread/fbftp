/*
author: foolbread
file: protocol/command_cdup.go
date: 2017/10/12
*/
package protocol

import (
	"github.com/foolbread/fbftp/session"
	"strings"
)

type commandCdup struct {

}

func (p *commandCdup)IsExtend() bool{
	return false
}

func (p *commandCdup)RequireAuth() bool{
	return true
}

func (p *commandCdup)RequireParam() bool{
	return false
}

func (p *commandCdup)Execute(sess *session.FTPSession, arg string)error{
	if !sess.UserAcl.AllowRead(){
		return sess.CtrlCon.WriteMsg(FTP_NOPERM,"Permission denied.")
	}

	localpath := sess.BuildPath("..")

	cp := strings.TrimPrefix(localpath,sess.UserAcl.GetWorkPath())
	if len(cp) == 0{
		sess.CurPath = "/"
	}else{
		sess.CurPath = cp
	}

	return sess.CtrlCon.WriteMsg(FTP_CWDOK,"Directory successfully changed.")
}
