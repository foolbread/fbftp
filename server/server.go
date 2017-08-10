/*
author: foolbread
file: server/server.go
date: 2017/8/4
*/
package server

import (
	"github.com/foolbread/fbcommon/golog"
	"github.com/foolbread/fbftp/config"

	"net"
	"fmt"
)

func InitServer(){
	golog.Info("fbftp server initing......")
}

type fbFTPSever struct {
	server_port int
	pasv_min_port int
	pasv_max_port int
}

func newFBFTPServer()*fbFTPSever{
	r := new(fbFTPSever)
	r.server_port = config.GetConfig().GetPort()
	r.pasv_min_port = config.GetConfig().GetPasvMinPort()
	r.pasv_max_port = config.GetConfig().GetPasvMaxPort()

	return r
}

func (s *fbFTPSever)run(){

}

func (s *fbFTPSever)startServerListen(){
	serverAddr,err := net.ResolveTCPAddr("tcp",fmt.Sprintf(":%d",s.server_port))
	if err != nil{
		golog.Critical(err)
	}

	li,err := net.ListenTCP("tcp",serverAddr)
	if err != nil{
		golog.Critical(err)
	}

	for{
		li.Accept()
	}
}

func (s *fbFTPSever)startPasvListen(){
	for i:=s.pasv_min_port; i<=s.pasv_max_port; i++{
		pasvAddr,err := net.ResolveTCPAddr("tcp",fmt.Sprintf(":%d",i))
		if err != nil{
			golog.Critical(err)
		}

		li,err := net.ListenTCP("tcp",pasvAddr)

		go func(){
			li.Accept()
		}()
	}
}