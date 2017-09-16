/*
author: foolbread
file: protocol/command_list.go
date: 2017/9/12
*/
package protocol

import (
	"github.com/foolbread/fbftp/session"
)

type commandList struct {

}

func (p *commandList)IsExtend()bool{
	return false
}

func (p *commandList)RequireAuth()bool{
	return true
}

func (p *commandList)RequireParam()bool{
	return false
}

func (p *commandList)Execute(sess *session.FTPSession, arg string)error{
	return nil
}