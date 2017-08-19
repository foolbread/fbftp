/*
author: foolbread
file: session/session_messge.go
date: 2017/8/19
*/
package session

import "os/exec"

const(
	sess_msg_create = 1
	sess_msg_readypasv = 2
	sess_msg_pasv = 3
	sess_msg_timeout = 4
)

type sessionMsg struct {
	msgType int
	sess *FBFTPSession
	externData interface{}
}

func NewSessionCreateMsg(sess *FBFTPSession, extern interface{})*sessionMsg{
	r := new(sessionMsg)
	r.msgType = sess_msg_create
	r.sess = sess
	r.externData = extern

	return r
}

func NewSessionReadyPasvMsg(sess *FBFTPSession, extern interface{})*sessionMsg{
	r := new(sessionMsg)
	r.msgType = sess_msg_readypasv
	r.sess = sess
	r.externData = extern

	return r
}

func NewSessionPasvMsg(sess *FBFTPSession, extern interface{})*sessionMsg{
	r := new(sessionMsg)
	r.msgType = sess_msg_pasv
	r.sess = sess
	r.externData = extern

	return r
}

func NewSessionTimeoutMsg(sess *FBFTPSession, extern interface{})*sessionMsg{
	r := new(sessionMsg)
	r.msgType = sess_msg_timeout
	r.sess = sess
	r.externData = extern

	return r
}