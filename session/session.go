/*
author: foolbread
file: session/session.go
date: 2017/9/11
*/
package session

import (
	"github.com/foolbread/fbcommon/golog"

	"github.com/foolbread/fbftp/con"
	"github.com/foolbread/fbftp/user"
)

func InitSession(){
	golog.Info("fbftp session initing......")
}

type FTPSession struct {
	UserName string
	UserInfo user.FBFTPUser
	CurPath string
	ctrlCon  *con.CmdCon
	dataCon  *con.PasvCon
}

func NewFTPSession()*FTPSession{
	r := new(FTPSession)

	return r
}

func (s *FTPSession)SetCMDCon(con *con.CmdCon){
	s.ctrlCon = con
}

func (s *FTPSession)SetDataCon(con *con.PasvCon){
	s.dataCon = con
}

func (s *FTPSession)Close(){
	if s.ctrlCon != nil{
		s.ctrlCon.Close()
	}

	if s.dataCon != nil{
		s.dataCon.Close()
	}
}

func (s *FTPSession)RecvCMD()(*con.FTPCmdReq,error){
	return s.ctrlCon.ReadCMDReq()
}

func (s *FTPSession)WriteMsg(code int,msg string)error{
	return s.ctrlCon.WriteMsg(code,msg)
}

func (s *FTPSession)IsLogin()bool{
	return s.UserInfo != nil
}