/*
author: foolbread
file: protocol/command_quit.go
date: 2017/9/26
*/
package protocol

import (
	"github.com/foolbread/fbftp/session"
)

type commandQuit struct {

}

func (p *commandQuit)IsExtend()bool{
	return false
}

func (p *commandQuit)RequireAuth()bool{
	return false
}

func (p *commandQuit)RequireParam()bool{
	return false
}

func (p *commandQuit)Execute(sess *session.FTPSession, arg string)error{
	sess.CtrlCon.WriteMsg(FTP_GOODBYE,"Goodbye.")

	sess.Close()
	return nil
}
