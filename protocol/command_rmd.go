/*
author: foolbread
file: protocol/command_rmd.go
date: 2017/10/12
*/
package protocol

import (
	"github.com/foolbread/fbftp/session"
)

type commandRmd struct {

}

func (p *commandRmd)IsExtend() bool{
	return false
}

func (p *commandRmd)RequireAuth() bool{
	return true
}

func (p *commandRmd)RequireParam() bool{
	return true
}

func (p *commandRmd)Execute(sess *session.FTPSession, arg string)error{
	if !sess.UserAcl.AllowWrite(){
		return sess.CtrlCon.WriteMsg(FTP_NOPERM,"Permission denied.")
	}

	localpath := sess.BuildPath(arg)
	err := sess.Storage.DeleteDir(localpath)
	if err != nil{
		return sess.CtrlCon.WriteMsg(FTP_FILEFAIL,"Remove directory operation failed.")
	}

	return sess.CtrlCon.WriteMsg(FTP_RMDIROK,"Remove directory operation successful.")
}
