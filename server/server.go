/*
author: foolbread
file: server/server.go
date: 2017/8/18
*/
package server

import (
	"github.com/foolbread/fbcommon/golog"

	"github.com/foolbread/fbftp/config"
	"github.com/foolbread/fbftp/session"
	"github.com/foolbread/fbftp/con"
	"net"
	"fmt"
	"github.com/foolbread/fbftp/protocol"
)

func InitServer(){
	golog.Info("fbftp server initing......")

	go g_server.serve()
}

var g_server *fbFTPServer = newfbFTPServer()

type fbFTPServer struct {

}

func newfbFTPServer()*fbFTPServer{
	r := new(fbFTPServer)

	return r
}

func (s *fbFTPServer)serve(){
	addr,err := net.ResolveTCPAddr("tcp",fmt.Sprintf(":%d",config.GetConfig().GetPort()))
	if err != nil{
		golog.Critical(err)
	}

	li,err := net.ListenTCP("tcp",addr)
	if err != nil{
		golog.Critical(err)
	}

	for{
		c,err := li.AcceptTCP()
		if err != nil{
			golog.Error(err)
			c.Close()
			continue
		}

		go s.work(c)
	}
}

func (s *fbFTPServer)work(c *net.TCPConn){
	sess := session.NewFTPSession()
	sess.SetCMDCon(con.NewCmdCon(c))

	for{
		req,err := sess.RecvCMD()
		if err != nil{
			golog.Error(err)
			sess.Close()
			break
		}

		protocol.DisPatchCmd(sess,req)
	}
}


