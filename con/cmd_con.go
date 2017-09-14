/*
author: foolbread
file: con/cmd_con.go
date: 2017/9/11
*/
package con

import (
	"bufio"
	"net"
	"strings"
	"errors"
	"fmt"
)

const (
	msg_fmt = "%d %s\r\n"
)

type FTPCmdReq struct {
	Cmd string
	Arg string
}

func (c *FTPCmdReq)IsArgNULL()bool{
	return len(c.Arg) == 0
}

type CmdCon struct {
	rawcon *net.TCPConn
	reader *bufio.Reader
	writer *bufio.Writer
}

func NewCmdCon(con *net.TCPConn)*CmdCon{
	r := new(CmdCon)
	r.rawcon = con
	r.reader = bufio.NewReader(con)
	r.writer = bufio.NewWriter(con)

	return r
}

func (c *CmdCon)Close(){
	if c.rawcon != nil{
		c.rawcon.Close()
	}
}

func (c *CmdCon)ReadCMDReq()(*FTPCmdReq,error){
	str,err := c.reader.ReadString('\n')
	if err != nil{
		return nil,err
	}

	fields := strings.SplitN(str," ",2)
	if len(fields) == 0{
		return nil,errors.New("ftp decode error!")
	}

	r := new(FTPCmdReq)
	r.Cmd = fields[0]
	if len(fields) > 1{
		r.Arg = strings.TrimSpace(fields[1])
	}

	return r, nil
}

func (c *CmdCon)WriteMsg(code int,msg string)error{
	_,err := c.writer.Write([]byte(fmt.Sprintf(msg_fmt,code,msg)))

	return err
}