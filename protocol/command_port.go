/*
author: foolbread
file: protocol/command_port.go
date: 2018/5/24
*/
package protocol

import (
	"github.com/foolbread/fbftp/session"
	"net"
	"strings"
	"strconv"
	"github.com/foolbread/fbftp/con"
)

const(
	port_arg_len = 6
	ipport_reserved = 1024
)

type commandPort struct {

}

func (p *commandPort)IsExtend()bool{
	return false
}

func (p *commandPort)RequireAuth()bool{
	return true
}

func (p *commandPort)RequireParam()bool{
	return true
}

func (p *commandPort)Execute(sess *session.FTPSession,arg string)error{
	remote := sess.CtrlCon.GetRemoteHost()
	remotehost,_,_ := net.SplitHostPort(remote)

	vals := strings.Split(arg,",")
	if len(vals) != port_arg_len{
		return sess.CtrlCon.WriteMsg(FTP_BADCMD,"Illegal PORT command.")
	}

	porthost := strings.Join(vals[:4],".")
	p1,_ := strconv.Atoi(vals[4])
	p2,_ := strconv.Atoi(vals[5])
	remoteport := p1<<8|p2
	if remotehost != porthost || remoteport < ipport_reserved{
		return sess.CtrlCon.WriteMsg(FTP_BADCMD,"Illegal PORT command.")
	}

	sess.DataCon = con.NewPortCon(porthost,strconv.Itoa(remoteport))

	return sess.CtrlCon.WriteMsg(FTP_PORTOK,"PORT command successful. Consider using PASV.")
}