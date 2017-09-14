/*
author: foolbread
file: con/pasv_con.go
date: 2017/9/11
*/
package con

import (
	"net"
	"time"
	"sync"
	"strconv"

	"github.com/foolbread/fbcommon/golog"
)

var g_pasvportManager *pasvPortManager = newPasvPortManager()

type pasvPortItem struct {
	li *net.TCPListener
	t time.Time
}

type pasvPortManager struct {
	portMap map[int]*pasvPortItem
	lo *sync.RWMutex
}

func newPasvPortManager()*pasvPortManager{
	r := new(pasvPortManager)
	r.portMap = make(map[int]*pasvPortItem)
	r.lo = new(sync.RWMutex)

	return r
}

func (c *pasvPortManager)init(minport int, maxport int){
	for i:=minport; i<=maxport; i++{
		c.portMap[i] = nil
	}
}

func (c *pasvPortManager)run(){
	ti := time.Tick(30*time.Second)

	for{
		<-ti

		c.lo.Lock()
		for k,v := range c.portMap{
			if v != nil{
				if v.t.Add(30*time.Second).Before(time.Now()){
					golog.Info("pasv port:",k,"use too long,forced free it.")
					if v.li != nil{
						v.li.Close()
					}

					c.portMap[k] = nil
				}
			}
		}
		c.lo.Unlock()
	}
}

func (c *pasvPortManager)GetFreePasvPort()int{
	var ret int = -1
	for{
		c.lo.Lock()
		for k,v := range c.portMap{
			if v == nil{
				ret = k
				c.portMap[k] = &pasvPortItem{nil,time.Now()}
				break
			}
		}
		c.lo.Unlock()

		if ret != -1{
			return ret
		}

		time.Sleep(2*time.Second)
	}
}

func (c *pasvPortManager)freePasvPort(port int){
	c.lo.Lock()
	defer c.lo.Unlock()

	c.portMap[port] = nil
}


type PasvCon struct {
	 rawcon *net.TCPConn
	 port int

	 wg sync.WaitGroup
	 err error
}

func NewPasvCon(port int)(*PasvCon,error){
	r := new(PasvCon)
	r.port = port

	return r,r.listernAndServe()
}

func (c *PasvCon)listernAndServe()error{
	addr,err := net.ResolveTCPAddr("tcp",net.JoinHostPort("",strconv.Itoa(c.port)))
	if err != nil{
		return err
	}

	li,err := net.ListenTCP("tcp",addr)
	if err != nil{
		return err
	}

	c.wg.Add(1)

	go func(){
		defer li.Close()
		defer c.wg.Done()
		defer g_pasvportManager.freePasvPort(c.port)

		con,err := li.AcceptTCP()
		if err != nil{
			golog.Error(err)
			c.err = err
			return
		}

		c.rawcon = con
	}()

	return nil
}

func (c *PasvCon)waitForOpenCon()error{
	if c.rawcon != nil {
		return nil
	}

	c.wg.Wait()
	return c.err
}

func (c *PasvCon)Read(p []byte) (n int, err error){
	if err := c.waitForOpenCon(); err != nil {
		return 0, err
	}

	return c.rawcon.Read(p)
}

func (c *PasvCon)Write(p []byte) (n int, err error) {
	if err := c.waitForOpenCon(); err != nil {
		return 0, err
	}

	return c.rawcon.Write(p)
}


func (c *PasvCon)Close(){
	if c.rawcon != nil{
		c.rawcon.Close()
	}
}