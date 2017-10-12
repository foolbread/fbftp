/*
author: foolbread
file: protocol/command_mkd.go
date: 2017/10/12
*/
package protocol

import (
	"github.com/foolbread/fbftp/session"
	"fmt"
)

type commandMkd struct {

}

func (p *commandMkd)IsExtend() bool{
	return false
}
func (p *commandMkd)RequireAuth() bool{
	return true
}

func (p *commandMkd)RequireParam() bool{
	return true
}

func (p *commandMkd)Execute(sess *session.FTPSession, arg string)error{
	if !sess.UserAcl.AllowWrite(){
		return sess.CtrlCon.WriteMsg(FTP_NOPERM,"Permission denied.")
	}

	localpath := sess.BuildPath(arg)
	err := sess.Storage.MKDir(localpath)
	if err != nil{
		return sess.CtrlCon.WriteMsg(FTP_FILEFAIL,"Create directory operation failed.")
	}

	return sess.CtrlCon.WriteMsg(FTP_MKDIROK,fmt.Sprintf(`"%s" created`,arg))
}
