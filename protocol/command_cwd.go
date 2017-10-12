/*
author: foolbread
file: protocol/command_cwd.go
date: 2017/9/29
*/
package protocol

import (
	"github.com/foolbread/fbftp/session"
	"strings"
	"github.com/foolbread/fbcommon/golog"
)

type commandCwd struct {

}

func (p *commandCwd)IsExtend()bool{
	return false
}

func (p *commandCwd)RequireAuth()bool{
	return true
}

func (p *commandCwd)RequireParam()bool{
	return true
}

func (p *commandCwd)Execute(sess *session.FTPSession, arg string)error{
	if !sess.UserAcl.AllowRead(){
		return sess.CtrlCon.WriteMsg(FTP_NOPERM,"Permission denied.")
	}

	pa := sess.BuildPath(arg)
	ok,err := sess.Storage.ChangeDir(pa)
	if err != nil || !ok{
		sess.CtrlCon.WriteMsg(FTP_FILEFAIL,"Failed to change directory.")
		return err
	}

	if !ok{
		sess.CtrlCon.WriteMsg(FTP_FILEFAIL,"Failed to change directory.")
		return nil
	}

	cp := strings.TrimPrefix(pa,sess.UserAcl.GetWorkPath())
	if len(cp) == 0{
		sess.CurPath = "/"
	}else{
		sess.CurPath = cp
	}

	golog.Info("curpath:",sess.CurPath)

	return sess.CtrlCon.WriteMsg(FTP_CWDOK,"Directory successfully changed.")
}