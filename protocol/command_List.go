/*
author: foolbread
file: protocol/command_list.go
date: 2017/9/12
*/
package protocol

import (
	"github.com/foolbread/fbftp/session"
	"path"
	"fmt"
)

const(
	list_fmt = "%s 1 %s %s %d %s %s\r\n"//mod owner group size time name
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
	if !sess.UserAcl.AllowRead(){
		return sess.CtrlCon.WriteMsg(FTP_NOPERM,"Permission denied.")
	}

	arg = path.Clean(arg)

	infos,err := sess.Storage.ListFile(path.Join(sess.UserAcl.GetWorkPath(),arg))
	if err !=nil{
		return err
	}

	sess.CtrlCon.WriteMsg(FTP_DATACONN,"Here comes the directory listing.")
	if sess.DataCon == nil{
		return sess.CtrlCon.WriteMsg(FTP_BADSENDNET,"Failure writing network stream.")
	}

	for _,v := range infos{
		sess.DataCon.Write([]byte(fmt.Sprintf(list_fmt,v.Mode.String(),
			v.Owner,v.Group,v.Size,v.ModTime.Format("Jan 2 15:04"),v.Name)))
	}
	sess.DataCon.Close()
	sess.DataCon = nil

	return 	sess.CtrlCon.WriteMsg(FTP_TRANSFEROK,"Directory send OK.")
}