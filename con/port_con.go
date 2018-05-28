/*
author: foolbread
file: con/port_con.go
date: 2018/5/24
*/
package con

import (
	"net"
)

type PortCon struct {
	rawCon *net.TCPConn
	remoteAddr *net.TCPAddr
}

func NewPortCon(addr string,port string)*PortCon{
	r := new(PortCon)
	r.remoteAddr,_ = net.ResolveTCPAddr("tcp",net.JoinHostPort(addr,port))

	return r
}

func (c *PortCon)startConnectClient()error{
	if c.rawCon != nil{
		return nil
	}

	//lr,_ := net.ResolveTCPAddr("tcp",":12020")
	var err error
	c.rawCon,err = net.DialTCP("tcp",nil,c.remoteAddr)

	return err
}

func (c *PortCon)Read(p []byte)(n int, err error){
	if err := c.startConnectClient();err != nil{
		return 0,err
	}

	return c.rawCon.Read(p)
}

func (c *PortCon)Write(p []byte) (n int, err error){
	if err := c.startConnectClient();err != nil{
		return 0,err
	}

	return c.rawCon.Write(p)
}

func (c *PortCon)Close(){
	if c.rawCon != nil{
		c.rawCon.Close()
		c.rawCon = nil
	}
}