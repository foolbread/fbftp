/*
author: foolbread
file: log/log.go
date: 2017/8/4
*/
package log

import (
	"github.com/foolbread/fbcommon/golog"
	"github.com/cihub/seelog"
)

func InitLog(logfile string){
	golog.Info("fbftp log initing......")

	var err error
	g_log,err = seelog.LoggerFromConfigAsFile(logfile)
	if err != nil{
		golog.Critical(err)
	}
}

func GetLog()seelog.LoggerInterface{
	return g_log
}

var g_log seelog.LoggerInterface