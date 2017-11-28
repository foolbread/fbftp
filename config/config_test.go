/*
author: foolbread
file: test/config_test.go
date: 2017/9/19
*/
package config

import (
	"testing"

)

func Test_Config(t *testing.T){
	InitConfig("conf.ini")
}
