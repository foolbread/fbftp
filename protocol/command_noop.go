/*
author: foolbread
file: protocol/command_noop.go
date: 2018/4/24
*/
package protocol

import (
	"github.com/foolbread/fbftp/session"
)

type commandNoop struct {
}

func (c *commandNoop)IsExtend()bool{
	return false
}

func (c *commandNoop)RequireAuth()bool{
	return true
}

func (c *commandNoop)RequireParam()bool{
	return false
}

func (c *commandNoop)Execute(sess *session.FTPSession, arg string)error{
	return sess.CtrlCon.WriteMsg(FTP_NOOPOK,"NOOP ok.")
}