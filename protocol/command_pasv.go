/*
author: foolbread
file: protocol/command_pasv.go
date: 2017/9/16
*/
package protocol

import (
	"github.com/foolbread/fbftp/session"
)

type commandPasv struct {

}

func (p *commandPasv)IsExtend()bool{
	return false
}

func (p *commandPasv)RequireAuth()bool{
	return true
}

func (p *commandPasv)RequireParam()bool{
	return false
}

func (p *commandPasv)Execute(sess *session.FTPSession,arg string)error{

	return nil
}