/*
author: foolbread
file: protocol/protocol.go
date: 2017/9/11
*/
package protocol

import (
	"github.com/foolbread/fbcommon/golog"
	"github.com/foolbread/fbftp/session"
	"github.com/foolbread/fbftp/con"
	"errors"
)

func InitProtocol(){
	golog.Info("fbftp protocol initing......")
}

var commandMap map[string]Command = map[string]Command{
	"USER":&commandUser{},
	"PASS":&commandPass{},
	"LIST":&commandList{},
}

var(
	errUnknowCmd = errors.New("unknow cmd error")
	errArgMiss = errors.New("arg is nil")
	errAuth = errors.New("auth error")
)


type Command interface {
	IsExtend() bool
	RequireAuth() bool
	RequireParam() bool
	Execute(*session.FTPSession, string)error
}

func DisPatchCmd(sess *session.FTPSession,req *con.FTPCmdReq)error{
	e := commandMap[req.Cmd]
	if e == nil{
		return errUnknowCmd
	}

	if e.RequireParam()&& req.IsArgNULL(){
		return errArgMiss
	}

	if e.RequireAuth() && !sess.IsLogin(){
		return errAuth
	}

	return 	e.Execute(sess,req.Arg)
}
