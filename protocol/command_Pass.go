/*
author: foolbread
file: protocol/command_Pass.go
date: 2017/9/12
*/
package protocol

import (
	"github.com/foolbread/fbftp/session"
)

type commandPass struct {

}

func (p *commandPass)IsExtend()bool{
	return false
}

func (p *commandPass)RequireAuth()bool{
	return false
}

func (p *commandPass)RequireParam()bool{
	return true
}

func (p *commandPass)Execute(sess *session.FTPSession, arg string)error{
	return nil
}