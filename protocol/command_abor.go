/*
author: foolbread
file: protocol/command_abor.go
date: 2018/5/8
*/
package protocol

import "github.com/foolbread/fbftp/session"

type commandAbor struct {

}

func (p *commandAbor)IsExtend()bool{
	return false
}

func (p *commandAbor)RequireAuth()bool{
	return  true
}

func (p *commandAbor)RequireParam()bool{
	return false
}

func (p *commandAbor)Execute(sess *session.FTPSession, arg string)error{
	if sess.DataCon != nil{
		sess.DataCon.Close()
	}

	return sess.CtrlCon.WriteMsg(FTP_ABOR_NOCONN,"No transfer to ABOR.")
}
