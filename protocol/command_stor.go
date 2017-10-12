/*
author: foolbread
file: protocol/command_stor.go
date: 2017/9/29
*/
package protocol

import (
	"github.com/foolbread/fbftp/session"
)

type commandStor struct {

}

func (p *commandStor)IsExtend()bool{
	return false
}

func (p *commandStor)RequireAuth()bool{
	return true
}

func (p *commandStor)RequireParam()bool{
	return true
}

func (p *commandStor)Execute(sess *session.FTPSession, arg string)error{
	if !sess.UserAcl.AllowWrite(){
		return sess.CtrlCon.WriteMsg(FTP_NOPERM,"Permission denied.")
	}
	defer sess.DataCon.Close()

	localpath := sess.BuildPath(arg)

	sess.CtrlCon.WriteMsg(FTP_DATACONN,"Ok to send data.")

	_,err := sess.Storage.StoreFile(localpath,sess.DataCon)
	if err != nil{
		return sess.CtrlCon.WriteMsg(FTP_UPLOADFAIL,err.Error())
	}

	return sess.CtrlCon.WriteMsg(FTP_TRANSFEROK,"Transfer complete.")
}