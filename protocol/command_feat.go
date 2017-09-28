/*
author: foolbread
file: protocol/command_feat.go
date: 2017/9/26
*/
package protocol

import (
	"github.com/foolbread/fbftp/session"
)

type commandFeat struct {

}

func (p *commandFeat)IsExtend()bool{
	return false
}

func (p *commandFeat)RequireAuth()bool{
	return false
}

func (p *commandFeat)RequireParam()bool{
	return false
}

func (p *commandFeat)Execute(sess *session.FTPSession, arg string)error{
	 sess.CtrlCon.WriteHyphen(FTP_FEAT,"Features:")
	 sess.CtrlCon.WriteRaw(" PASV\r\n")

	 return sess.CtrlCon.WriteMsg(FTP_FEAT,"End")
}