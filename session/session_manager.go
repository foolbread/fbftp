/*
author: foolbread
file: session/session_manager.go
date: 2017/8/18
*/
package session

import (
	"github.com/foolbread/fbcommon/golog"
)

const(
	sess_msg_create = 1
	sess_msg_readypasv = 2
	sess_msg_pasv = 3
	sess_msg_timeout = 4
)

func InitSession(){
	golog.Info("fbftp session initing......")
}

type SessionMsg struct {
	MsgType int
	Sess *FBFTPSession
	ExternData interface{}
}

func CommitSessionMsg(m *SessionMsg){
	switch m.MsgType {
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

func RecvSessionMsg()*SessionMsg{
	return <-g_sessManager.msgCH
}

var g_sessManager *sessionManager = newsessionManager()

type sessionManager struct {
	msgCH chan *SessionMsg
}

func newsessionManager()*sessionManager{
	r := new(sessionManager)
	r.msgCH = make(chan *SessionMsg,100)

	return r
}

func (s *sessionManager)handlerCreate(m *SessionMsg){

}

func (s *sessionManager)handlerReadyPasv(m *SessionMsg){

}

func (s *sessionManager)handlerPasv(m *SessionMsg){

}

func (s *sessionManager)handlerTimeOut(m *SessionMsg){

}
