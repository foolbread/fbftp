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
	"github.com/foolbread/fbftp/acl"
)

func InitSession(){
	golog.Info("fbftp session initing......")
}

type FTPSession struct {
	UserName  string
	UserInfo  user.FBFTPUser
	UserAcl   *acl.UserACL
	CurPath   string
	LocalHost string
	CtrlCon   *con.CmdCon
	DataCon   *con.PasvCon
}

func NewFTPSession()*FTPSession{
	r := new(FTPSession)

	return r
}

func (s *FTPSession)Close(){
	if s.CtrlCon != nil{
		s.CtrlCon.Close()
	}

	if s.DataCon != nil{
		s.DataCon.Close()
	}
}

func (s *FTPSession)RecvCMD()(*con.FTPCmdReq,error){
	return s.CtrlCon.ReadCMDReq()
}

func (s *FTPSession)WriteMsg(code int,msg string)error{
	return s.CtrlCon.WriteMsg(code,msg)
}

func (s *FTPSession)IsLogin()bool{
	return s.UserInfo != nil
}