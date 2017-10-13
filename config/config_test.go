/*
author: foolbread
file: test/config_test.go
date: 2017/9/19
*/
package test

import (
	"testing"
	"github.com/foolbread/fbftp/config"
)

func Test_Config(t *testing.T){
	config.InitConfig("conf.ini")
}
