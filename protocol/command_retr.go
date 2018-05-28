/*
author: foolbread
file: protocol/command_retr.go
date: 2017/9/29
*/
package protocol

import (
	"github.com/foolbread/fbftp/session"
)

type commandRetr struct {

}

func (p *commandRetr)IsExtend()bool{
	return false
}

func (p *commandRetr)RequireAuth()bool{
	return true
}

func (p *commandRetr)RequireParam()bool{
	return true
}

func (p *commandRetr)Execute(sess *session.FTPSession, arg string)error{
	if !sess.UserAcl.AllowRead(){
		return sess.CtrlCon.WriteMsg(FTP_NOPERM,"Permission denied.")
	}

	defer sess.DataCon.Close()

	if sess.DataCon == nil{
		return sess.CtrlCon.WriteMsg(FTP_BADSENDCONN,"Use PORT or PASV first.")
	}

	localpath := sess.BuildPath(arg)

	sess.CtrlCon.WriteMsg(FTP_DATACONN,"send to data.")

	_,err := sess.Storage.GetFile(localpath,sess.DataCon)
	if err != nil{
		return sess.CtrlCon.WriteMsg(FTP_BADSENDFILE,err.Error())
	}

	return sess.CtrlCon.WriteMsg(FTP_TRANSFEROK,"Transfer complete.")
}
