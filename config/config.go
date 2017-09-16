/*
author: foolbread
file: config/config.go
date: 2017/8/4
*/
package config

import (
	"github.com/foolbread/fbcommon/golog"
	fbconfig "github.com/foolbread/fbcommon/config"
)

func InitConfig(configfile string){
	golog.Info("fbftp config initing......")
	conf,err := fbconfig.LoadConfigByFile(configfile)
	if err != nil{
		golog.Critical(err)
	}

	var section string
	var val string
	var vai int
	section = "server"
	vai = conf.MustInt(section,"port",21)
	g_conf.port = vai

	vai = conf.MustInt(section,"pasv_min_port",0)
	g_conf.pasvMinPort = vai

	vai = conf.MustInt(section,"pasv_max_port",0)
	g_conf.pasvMaxPort = vai

	section = "ftp"
	val = conf.MustString(section,"welcome_msg","")
	g_conf.welcomeMsg = val

	section = "log"
	val = conf.MustString(section,"log_file","")
	g_conf.logfile = val

	//print config result
	golog.Info("[server]","port:",g_conf.port)
	golog.Info("[server]","pasv_min_port:",g_conf.pasvMinPort)
	golog.Info("[server]","pasv_max_port:",g_conf.pasvMaxPort)
	golog.Info("[ftp]","welcome_msg:",g_conf.welcomeMsg)
	golog.Info("[log]","log_file:",g_conf.logfile)
}

func GetConfig()*fbFTPConfig{
	return g_conf
}

var g_conf *fbFTPConfig = new(fbFTPConfig)

type fbFTPConfig struct {
	port        int
	pasvMinPort int
	pasvMaxPort int

	welcomeMsg string
	logfile string
}

func (c *fbFTPConfig)GetPort()int{
	return c.port
}

func (c *fbFTPConfig)GetPasvMinPort()int{
	return c.pasvMinPort
}

func (c *fbFTPConfig)GetPasvMaxPort()int{
	return c.pasvMaxPort
}

func (c *fbFTPConfig)GetWelcomeMsg()string{
	return c.welcomeMsg
}

func (c *fbFTPConfig)GetLogFile()string{
	return c.logfile
}

