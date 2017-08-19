/*
author: foolbread
file: session/session.go
date: 2017/8/10
*/
package session

import (
	"net"
)

type FBFTPSession struct {
	cmdCon *net.TCPConn
	dataCon *net.TCPConn
}

func NewFBFTPSession()*FBFTPSession {
	r := new(FBFTPSession)

	return r
}

func (s *FBFTPSession)setCmdConnect(con *net.TCPConn){
	s.cmdCon = con
}

func (s *FBFTPSession)setDataConnect(con *net.TCPConn){
	s.dataCon = con
}

func (s *FBFTPSession)GetClientAddr()string{
	var str string
	if s.cmdCon != nil{
		str = s.cmdCon.RemoteAddr().String()
	}

	return str
}

func (s *FBFTPSession)Notify(){

}