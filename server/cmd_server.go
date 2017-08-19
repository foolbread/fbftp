/*
author: foolbread
file: server/cmd_server.go
date: 2017/8/4
*/
package server

import (
	"github.com/foolbread/fbcommon/golog"

	"net"
	"fmt"
	"github.com/foolbread/fbftp/session"
)

type fbFTPCmdSever struct {
	port int
}

func newfbFTPCmdServer(port int)*fbFTPCmdSever {
	r := new(fbFTPCmdSever)
	r.port = port

	return r
}

func (s *fbFTPCmdSever)run(){
	s.startServerListen()
}

func (s *fbFTPCmdSever)startServerListen(){
	serverAddr,err := net.ResolveTCPAddr("tcp",fmt.Sprintf(":%d",s.port))
	if err != nil{
		golog.Critical(err)
	}

	li,err := net.ListenTCP("tcp",serverAddr)
	if err != nil{
		golog.Critical(err)
	}

	for{
		con,err := li.AcceptTCP()
		if err != nil{
			golog.Critical(err)
		}

		sess := session.NewFBFTPSession()

		msg := &serverMsg{svr_msg_create,sess,con}
		sendOwerServer(msg)
	}
}
