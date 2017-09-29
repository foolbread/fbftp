/*
author: foolbread
file: protocol/command_stat.go
date: 2017/9/29
*/
package protocol

import (
	"github.com/foolbread/fbftp/session"
	"fmt"
)

type commandStat struct {

}

func (p *commandStat)IsExtend()bool{
	return false
}

func (p *commandStat)RequireAuth()bool{
	return true
}

func (p *commandStat)RequireParam()bool{
	return true
}

func (p *commandStat)Execute(sess *session.FTPSession, arg string)error{
	if !sess.UserAcl.AllowRead(){
		return sess.CtrlCon.WriteMsg(FTP_NOPERM,"Permission denied.")
	}

	infos,err := sess.Storage.ListFile(sess.BuildPath(arg))
	if err !=nil{
		return err
	}

	sess.CtrlCon.WriteHyphen(FTP_STATFILE_OK,"Status follows:")

	for _,v := range infos{
		sess.CtrlCon.WriteRaw(fmt.Sprintf(list_fmt,v.Mode.String(),
			v.Owner,v.Group,v.Size,v.ModTime.Format("Jan 2 15:04"),v.Name))
	}

	return 	sess.CtrlCon.WriteMsg(FTP_STATFILE_OK,"End of status")
}
