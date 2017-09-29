/*
author: foolbread
file: protocol/command_size.go
date: 2017/9/29
*/
package protocol

import (
	"github.com/foolbread/fbftp/session"
	"fmt"
)

type commandSize struct {

}

func (p *commandSize)IsExtend()bool{
	return false
}

func (p *commandSize)RequireAuth()bool{
	return true
}

func (p *commandSize)RequireParam()bool{
	return true
}

func (p *commandSize)Execute(sess *session.FTPSession, arg string)error{
	if !sess.UserAcl.AllowRead(){
		return sess.CtrlCon.WriteMsg(FTP_NOPERM,"Permission denied.")
	}

	infos,err := sess.Storage.ListFile(sess.BuildPath(arg))
	if err != nil{
		fmt.Println(err)
	}

	if infos == nil || infos[0].IsDir{
		return sess.CtrlCon.WriteMsg(FTP_FILEFAIL,"Could not get file size.")
	}

	return sess.CtrlCon.WriteMsg(FTP_SIZEOK,fmt.Sprintf("%d",infos[0].Size))
}
