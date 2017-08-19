/*
author: foolbread
file: server/pasv_server.go
date: 2017/8/14
*/
package server

import (
	"sync"
	"fmt"
	"net"

	"github.com/foolbread/fbcommon/golog"
	"github.com/foolbread/fbftp/session"
	"github.com/foolbread/fbftp/util"
	"strings"
	"time"
)

type waitItem struct {
	sess *session.FBFTPSession
	addTime time.Time
}

type fbFTPPasvServer struct {
	port int
	waitMap map[string]*waitItem

	lo *sync.RWMutex
}

func newfbFTPPasvServer(port int)*fbFTPPasvServer {
	r := new(fbFTPPasvServer)
	r.waitMap = make(map[string]*waitItem)
	r.lo = new(sync.RWMutex)
	r.port = port

	return r
}

func (s *fbFTPPasvServer)run(){
	util.AddSecondTickEvent(5*time.Second,s.tickCheckWait)

	s.startServerListen()
}

func (s *fbFTPPasvServer)startServerListen(){
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

		addr := strings.Split(con.RemoteAddr().String(),":")
		if s.isWait(addr[0]){
			sess := s.getFromWait(addr[0])
			sess.SetDataConnect(con)

			s.deleteFromWait(addr[0])

			msg := &serverMsg{svr_msg_pasv,sess,nil}
			sendOwerServer(msg)
		}else{
			con.Close()
		}
	}
}

func (s *fbFTPPasvServer)tickCheckWait(){
	s.lo.Lock()
	for k,v := range s.waitMap{
		if v.addTime.Add(10*time.Second).Before(time.Now()){
			delete(s.waitMap,k)
			msg := &serverMsg{svr_msg_timeout,v.sess,nil}
			sendOwerServer(msg)
		}
	}
	s.lo.Unlock()

	util.AddSecondTickEvent(5*time.Second,s.tickCheckWait)
}

func (s *fbFTPPasvServer)isWait(ip string)bool{
	s.lo.RLock()
	defer s.lo.RUnlock()

	_,ok := s.waitMap[ip]

	return ok
}

func (s *fbFTPPasvServer)getFromWait(ip string)*session.FBFTPSession{
	s.lo.RLock()
	defer s.lo.RUnlock()

	return s.waitMap[ip].sess
}

func (s *fbFTPPasvServer)commitToWait(ip string,sess *session.FBFTPSession){
	s.lo.Lock()
	defer s.lo.Unlock()

	s.waitMap[ip] = &waitItem{sess,time.Now()}
}

func (s *fbFTPPasvServer)deleteFromWait(ip string){
	s.lo.Lock()
	defer s.lo.Unlock()

	delete(s.waitMap,ip)
}

