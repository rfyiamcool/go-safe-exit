# go-safe-exit

通过信号、标记位、Channel的方法来控制所有逻辑任务流的安全退出.

## Usage:

```
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
```
