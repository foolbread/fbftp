/*
author: foolbread
file: fbftp/fbftp.go
date: 2017/8/4
*/
package main

import (
	"github.com/foolbread/fbftp/config"
	"github.com/foolbread/fbftp/log"
	"github.com/foolbread/fbftp/server"
	"github.com/foolbread/fbftp/con"
	"github.com/foolbread/fbftp/util"
	"github.com/foolbread/fbftp/session"
	"github.com/foolbread/fbftp/protocol"
	"github.com/foolbread/fbftp/statistics"
	"github.com/foolbread/fbftp/acl"
	"github.com/foolbread/fbftp/storage"
	"github.com/foolbread/fbftp/user"

	"flag"
	"runtime"
)

func init(){
	flag.StringVar(&config_file,"f","conf.ini","config file path!")
	flag.Parse()

	config.InitConfig(config_file)
	log.InitLog(config.GetConfig().GetLogFile())
	statistics.InitStatistics()
	protocol.InitProtocol()
	session.InitSession()
	util.InitUtil()
	user.InitUser()
	acl.InitAcl()
	con.InitCon()
	storage.InitStorage()
	server.InitServer()
}

var config_file string

func main(){
	runtime.GOMAXPROCS(runtime.NumCPU())

	select {}
}