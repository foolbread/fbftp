/*
author: foolbread
file: protocol/command_rnfr.go
date: 2017/10/12
*/
package protocol

import (
	"github.com/foolbread/fbftp/session"
)

type commandRnfr struct {

}

func (p *commandRnfr)IsExtend() bool{
	return false
}

func (p *commandRnfr)RequireAuth() bool{
	return true
}

func (p *commandRnfr)RequireParam() bool{
	return true
}

func (p *commandRnfr)Execute(sess *session.FTPSession, arg string)error{
	if !sess.UserAcl.AllowWrite(){
		return sess.CtrlCon.WriteMsg(FTP_NOPERM,"Permission denied.")
	}

	localpath := sess.BuildPath(arg)
	_,err := sess.Storage.Stat(localpath)
	if err != nil{
		return sess.CtrlCon.WriteMsg(FTP_FILEFAIL,"RNFR command failed.")
	}

	sess.RnfrStr = arg

	return sess.CtrlCon.WriteMsg(FTP_RNFROK,"Ready for RNTO.")
}
