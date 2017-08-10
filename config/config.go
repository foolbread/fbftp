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
	var val int
	section = "server"
	val = conf.MustInt(section,"port",21)
	g_conf.port = val

	val = conf.MustInt(section,"pasv_min_port",0)
	g_conf.pasv_min_port = val

	val = conf.MustInt(section,"pasv_max_port",0)
	g_conf.pasv_max_port = val

	//print config result
	golog.Info("server","port",g_conf.port)
	golog.Info("server","pasv_min_port",g_conf.pasv_min_port)
	golog.Info("server","pasv_max_port",g_conf.pasv_max_port)
}

func GetConfig()*fbFTPConfig{
	return g_conf
}


var g_conf *fbFTPConfig = new(fbFTPConfig)

type fbFTPConfig struct {
	port int
	pasv_min_port int
	pasv_max_port int
}

func (c *fbFTPConfig)GetPort()int{
	return c.port
}

func (c *fbFTPConfig)GetPasvMinPort()int{
	return c.pasv_min_port
}

func (c *fbFTPConfig)GetPasvMaxPort()int{
	return c.pasv_max_port
}

