/*
author: foolbread
file: protocol/command_rnto.go
date: 2017/10/12
*/
package protocol

import (
	"github.com/foolbread/fbftp/session"
)

type commandRnto struct {

}

func (p *commandRnto)IsExtend() bool{
	return false
}

func (p *commandRnto)RequireAuth() bool{
	return true
}

func (p *commandRnto)RequireParam() bool{
	return true
}

func (p *commandRnto)Execute(sess *session.FTPSession, arg string)error{
	if !sess.UserAcl.AllowWrite(){
		return sess.CtrlCon.WriteMsg(FTP_NOPERM,"Permission denied.")
	}

	if len(sess.RnfrStr) == 0{
		return sess.CtrlCon.WriteMsg(FTP_NEEDRNFR,"RNFR required first.")
	}

	rnfrpath := sess.BuildPath(sess.RnfrStr)
	rntopath := sess.BuildPath(arg)

	err := sess.Storage.ReName(rnfrpath,rntopath)
	if err != nil{
		return sess.CtrlCon.WriteMsg(FTP_FILEFAIL,"Rename failed.")
	}

	return sess.CtrlCon.WriteMsg(FTP_RENAMEOK,"Rename successful.")
}
