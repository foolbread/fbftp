/*
author: foolbread
file: user/user_test.go
date: 2017/11/28
*/
package user

import (
	"testing"
	"github.com/foolbread/fbftp/config"
)

func init(){
	config.InitConfig("conf.ini")
}

func Test_updateCommonUser(t *testing.T){
	g_userManager.updateCommonUser()
}

func Test_updateCloudUser(t *testing.T){
	g_userManager.updateCloudUser()
}