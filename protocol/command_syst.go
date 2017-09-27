/*
author: foolbread
file: protocol/command_syst.go
date: 2017/9/27
*/
package protocol

import (
	"github.com/foolbread/fbftp/session"
)

type commandSyst struct {

}

func (p *commandSyst)IsExtend()bool{
	return false
}

func (p *commandSyst)RequireAuth()bool{
	return true
}

func (p *commandSyst)RequireParam()bool{
	return false
}

func (p *commandSyst)Execute(sess *session.FTPSession, arg string)error{
	return sess.WriteMsg(FTP_SYSTOK,"UNIX Type: L8")
}