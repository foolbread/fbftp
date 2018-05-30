/*
author: foolbread
file: protocol/command_type.go
date: 2017/9/29
*/
package protocol

import (
	"github.com/foolbread/fbftp/session"
)

type commandType struct {

}

func (p *commandType)IsExtend()bool{
	return false
}

func (p *commandType)RequireAuth()bool{
	return true
}

func (p *commandType)RequireParam()bool{
	return true
}

func (p *commandType)Execute(sess *session.FTPSession, arg string)error{
	if arg == "I" || arg == "L8" || arg == "L 8"{
		return sess.CtrlCon.WriteMsg(FTP_TYPEOK,"Switching to Binary mode.")
	}

	return sess.CtrlCon.WriteMsg(FTP_BADCMD,"Unrecognised TYPE command.")
}
