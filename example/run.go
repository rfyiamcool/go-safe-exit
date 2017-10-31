package main

import (
	"log"
	safe_exit "github.com/rfyiamcool/go-safe-exit"
)

func main(){
	timeout := 3
	EngineWaitGroup := safe_exit.NewControlGroup()

	EngineWaitGroup.MakeSignal()
	EngineWaitGroup.RecvSignal()

	if EngineWaitGroup.WaitTimeout(timeout) {
		log.Println("timeout")
	} else {
		log.Println("正常退出")
	}
}
