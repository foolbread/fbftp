/*
author: foolbread
file: server/server.go
date: 2017/8/18
*/
package server

import (
	"github.com/foolbread/fbcommon/golog"

	"github.com/foolbread/fbftp/session"
	"github.com/foolbread/fbftp/config"
)

func InitServer(){
	golog.Info("fbftp server initing......")
}

const (
	svr_msg_create = 1
	svr_msg_pasv = 2
	svr_msg_timeout = 3
)

type serverMsg struct {
	msgType int
	session *session.FBFTPSession
	externData interface{}
}

var g_server *fbFTPServer = newfbFTPServer()

type fbFTPServer struct {
	cmdServer *fbFTPCmdSever
	pasvServers []*fbFTPPasvServer
	msgCH chan *serverMsg
}

func newfbFTPServer()*fbFTPServer{
	r := new(fbFTPServer)
	r.msgCH = make(chan *serverMsg,100)
	r.cmdServer = newfbFTPCmdServer(config.GetConfig().GetPort())
	for i := config.GetConfig().GetPasvMinPort(); i <= config.GetConfig().GetPasvMaxPort(); i++{
		r.pasvServers = append(r.pasvServers, newfbFTPPasvServer(i))
	}

	return r
}

func (s *fbFTPServer)run(){
	go s.cmdServer.run()

	for _,v := range s.pasvServers{
		go v.run()
	}
	
	for {
		select {
		case msg := s.recvServerMsg():
				go s.handlerServerMsg(msg)
		case sess := session.RecvPasvReq():
				go s.handlerPasvReq(sess)
		}
	}
}

func (s *fbFTPServer)recvServerMsg()*serverMsg{
	return <-s.msgCH
}

func (s *fbFTPServer)handlerServerMsg(m *serverMsg){
	switch m.msgType {
	case svr_msg_create:
		session.CommitSessionMsg(session.NewSessionCreateMsg(m.session,m.externData))
	case svr_msg_pasv:
		session.CommitSessionMsg(session.NewSessionPasvMsg(m.session,m.externData))
	case svr_msg_timeout:
		session.CommitSessionMsg(session.NewSessionTimeoutMsg(m.session,m.externData))
	}
}

func (s *fbFTPServer)handlerPasvReq(sess *session.FBFTPSession){

}

///////////////////////////////////////////////////////////////////////////////////////
func sendOwerServer(m *serverMsg){
	g_server.msgCH <- m
}
