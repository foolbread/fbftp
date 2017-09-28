/*
author: foolbread
file: protocol/command_help.go
date: 2017/9/28
*/
package protocol

import (
	"github.com/foolbread/fbftp/session"
)

type commandHelp struct {

}

func (p *commandHelp)IsExtend()bool{
	return false
}

func (p *commandHelp)RequireAuth()bool{
	return false
}

func (p *commandHelp)RequireParam()bool{
	return false
}

func (p *commandHelp)Execute(sess *session.FTPSession, arg string)error{
	sess.CtrlCon.WriteHyphen(FTP_HELP,"The following commands are recognized.")
	sess.CtrlCon.WriteRaw(" FEAT HELP LIST PASS PASV PWD QUIT SYST USER\r\n")

	return sess.CtrlCon.WriteMsg(FTP_HELP,"help ok.")
}
