/*
author: foolbread
file: protocol/command_pasv.go
date: 2017/9/16
*/
package protocol

import (
	"github.com/foolbread/fbftp/session"
	"github.com/foolbread/fbftp/con"
	"strings"
	"fmt"
)

type commandPasv struct {

}

func (p *commandPasv)IsExtend()bool{
	return false
}

func (p *commandPasv)RequireAuth()bool{
	return true
}

func (p *commandPasv)RequireParam()bool{
	return false
}

func (p *commandPasv)Execute(sess *session.FTPSession,arg string)error{
	port := con.GetFreePassPort()
	pasv,err := con.NewPasvCon(port)
	if err != nil{
		return err
	}

	sess.DataCon = pasv
	local := sess.CtrlCon.GetLocalHost()
	idx := strings.Index(local,":")
	ip := strings.Split(local[:idx],".")
	p1 := port/256
	p2 := port&255

	return sess.CtrlCon.WriteMsg(FTP_PASVOK,fmt.Sprintf("Entering Passive Mode (%s,%s,%s,%s,%d,%d)",
		ip[0],ip[1],ip[2],ip[3],p1,p2))
}