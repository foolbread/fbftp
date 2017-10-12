/*
author: foolbread
file: protocol/command_dele.go
date: 2017/10/12
*/
package protocol

import (
	"github.com/foolbread/fbftp/session"
	"github.com/foolbread/fbcommon/golog"
)

type commandDele struct {

}

func (p *commandDele)IsExtend() bool{
	return false
}

func (p *commandDele)RequireAuth() bool{
	return true
}

func (p *commandDele)RequireParam() bool{
	return true
}

func (p *commandDele)Execute(sess *session.FTPSession, arg string)error{
	if !sess.UserAcl.AllowWrite(){
		return sess.CtrlCon.WriteMsg(FTP_NOPERM,"Permission denied.")
	}

	localpath := sess.BuildPath(arg)
	golog.Info("delete file:",localpath)
	err := sess.Storage.DeleteFile(localpath)
	if err != nil{
		return sess.CtrlCon.WriteMsg(FTP_FILEFAIL,"Delete operation failed.")
	}

	return sess.CtrlCon.WriteMsg(FTP_DELEOK,"Delete operation successful.")
}