/*
author: foolbread
file: util/tick.go
date: 2017/8/18
*/
package util

import (
	"time"
	fbtime "github.com/foolbread/fbcommon/time"
)

func initTick(){
	go g_tick.Start()
}

func AddSecondTickEvent(d time.Duration, f func()){
	g_tick.NewTimer(d,f)
}

var g_tick *fbtime.Timer = fbtime.New(1*time.Second)