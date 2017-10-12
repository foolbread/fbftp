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
	"github.com/foolbread/fbftp/storage"

	"path/filepath"
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
	RnfrStr   string
	Storage	  storage.FTPStorage
	CtrlCon   *con.CmdCon
	DataCon   con.DataCon
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

func (s *FTPSession)IsLogin()bool{
	return s.UserInfo != nil
}

func (s *FTPSession)BuildPath(p string)(string){
	var ret string
	if filepath.IsAbs(p){
		ret = filepath.Clean(filepath.Join(s.UserAcl.GetWorkPath(),p))
	}else{
		ret = filepath.Clean(filepath.Join(s.UserAcl.GetWorkPath(),s.CurPath,p))
	}

	if len(ret) < len(s.UserAcl.GetWorkPath()){
		return s.UserAcl.GetWorkPath()
	}

	return ret
}