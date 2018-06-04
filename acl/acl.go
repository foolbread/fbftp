/*
author: foolbread
file: acl/acl.go
date: 2017/8/10
*/
package acl

import (
	"github.com/foolbread/fbcommon/golog"
	fbconfig "github.com/foolbread/fbcommon/config"
	"github.com/foolbread/fbftp/config"
	"path"
	"strings"
)

func InitAcl(){
	golog.Info("fbftp acl initing......")
}

const (
	ONLY_READ = 1
	ONLY_WRITE = 2
	READ_WRITE = ONLY_READ|ONLY_WRITE
)

func GetUserAcl(username string)*UserACL{
	dir := config.GetConfig().GetUserConfDir()

	conf,err := fbconfig.LoadConfigByFile(path.Join(dir,username))
	if err != nil{
		golog.Error(err)
		return nil
	}

	var retacl *UserACL = newUserACL(username)

	str := conf.MustString("acl","authority","")
	switch strings.ToLower(str) {
	case "r":
		retacl.authority = ONLY_READ
	case "w":
		retacl.authority = ONLY_WRITE
	case "rw":
		retacl.authority = READ_WRITE
	}

	str = conf.MustString("acl","allow_list","")
	ips := strings.Split(str,"|")
	if len(ips) > 0{
		retacl.allowMap = make(map[string]struct{})
		for _, v:= range ips{
			retacl.allowMap[v] = struct{}{}
		}
	}

	str = conf.MustString("data","data_path","")
	retacl.workpath = str


	return retacl
}