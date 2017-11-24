/*
author: foolbread
file: config/config.go
date: 2017/8/4
*/
package config

import (
	"github.com/foolbread/fbcommon/golog"
	fbconfig "github.com/foolbread/fbcommon/config"
	"io/ioutil"
	"encoding/xml"
)

func InitConfig(configfile string){
	golog.Info("fbftp config initing......")
	conf,err := fbconfig.LoadConfigByFile(configfile)
	if err != nil{
		golog.Critical(err)
	}

	golog.Info("==============load config file==============")

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

	section = "user"
	val = conf.MustString(section,"user_conf_dir","")
	g_conf.userConfDir = val

	val = conf.MustString(section,"user_file","")
	g_conf.userFile = val

	//print config result
	golog.Info("[server]","port:",g_conf.port)
	golog.Info("[server]","pasv_min_port:",g_conf.pasvMinPort)
	golog.Info("[server]","pasv_max_port:",g_conf.pasvMaxPort)
	golog.Info("[ftp]","welcome_msg:",g_conf.welcomeMsg)
	golog.Info("[user]","user_conf_dir:",g_conf.userConfDir)
	golog.Info("[user]","user_file:",g_conf.userFile)
	golog.Info("[log]","log_file:",g_conf.logfile)

	golog.Info("==============load ftpuser file==============")

	data,err := ioutil.ReadFile(g_conf.userFile)
	if err != nil{
		golog.Critical(err)
	}

	err = xml.Unmarshal(data,&g_conf.user)
	if err != nil{
		golog.Critical(err)
	}

	for _,v := range g_conf.user.CommonUsers.Users{
		golog.Info("common user:",v.UserName)
	}

	for _,v := range g_conf.user.CloudUsers.Users{
		golog.Info("cloud user:",v.UserName)
	}
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

	userConfDir string
	userFile string

	user fbUserConfig
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

func (c *fbFTPConfig)GetUserConfDir()string{
	return c.userConfDir
}

func (c *fbFTPConfig)GetAllCommonUsers()*commonUsers{
	return &c.user.CommonUsers
}

func (c *fbFTPConfig)GetAllCloudUsers()*cloudUsers{
	return &c.user.CloudUsers
}

////////////////////////////////////////////////////////
type fbUserConfig struct {
	XMLName xml.Name `xml:"ftpuser"`
	CommonUsers commonUsers `xml:"common_user"`
	CloudUsers cloudUsers `xml:"cloud_user"`
}

type commonUsers struct {
	Users []*fbCommonUserConfig `xml:"user"`
}

type cloudUsers struct {
	Users []*fbCloudUserConfig	`xml:"user"`
}

type fbCommonUserConfig struct {
	XMLName xml.Name `xml:"user"`
	UserName string `xml:"username"`
	PassWord string `xml:"password"`
}

type fbCloudUserConfig struct {
	XMLName xml.Name `xml:"user"`
	UserName string	`xml:"username"`
	PassWord string `xml:"password"`
	Bucket   string `xml:"bucket"`
	EndPoint string `xml:"endpoint"`
	Token  string   `xml:"token"`
	AccKey string	`xml:"acckey"`
	SecKey string	`xml:"seckey"`
}
////////////////////////////////////////////////////////
