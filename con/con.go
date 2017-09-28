/*
author: foolbread
file: con/con.go
date: 2017/9/11
*/
package con

import (
	"github.com/foolbread/fbcommon/golog"

	"github.com/foolbread/fbftp/config"
)

func InitCon(){
	golog.Info("fbftp con initing......")

	g_pasvportManager.init(config.GetConfig().GetPasvMinPort(),config.GetConfig().GetPasvMaxPort())

	go g_pasvportManager.run()
}

func GetFreePassPort()int{
	return g_pasvportManager.GetFreePasvPort()
}

type DataCon interface {
	Read([]byte)(int,error)
	Write([]byte)(int,error)
	Close()
}