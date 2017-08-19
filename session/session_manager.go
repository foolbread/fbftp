/*
author: foolbread
file: session/session_manager.go
date: 2017/8/18
*/
package session

import (
	"github.com/foolbread/fbcommon/golog"
)



func InitSession(){
	golog.Info("fbftp session initing......")
}

func CommitSessionMsg(m *sessionMsg){
	switch m.msgType {
	case sess_msg_create:
		g_sessManager.handlerCreate(m)
	case sess_msg_readypasv:
		g_sessManager.handlerReadyPasv(m)
	case sess_msg_pasv:
		g_sessManager.handlerPasv(m)
	case sess_msg_timeout:
		g_sessManager.handlerTimeOut(m)
	}
}

func RecvPasvReq()*FBFTPSession {
	return <-g_sessManager.msgCH
}

var g_sessManager *sessionManager = newsessionManager()

type sessionManager struct {
	msgCH chan *FBFTPSession
}

func newsessionManager()*sessionManager{
	r := new(sessionManager)
	r.msgCH = make(chan *FBFTPSession,100)

	return r
}

func (s *sessionManager)handlerCreate(m *sessionMsg){

}

func (s *sessionManager)handlerReadyPasv(m *sessionMsg){

}

func (s *sessionManager)handlerPasv(m *sessionMsg){

}

func (s *sessionManager)handlerTimeOut(m *sessionMsg){

}
