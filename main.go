package main

import (
	"time"

	"github.com/seferen/logger"
)

var log = logger.GetLoger("main")

var (
	quite  = make(chan bool)
	target = make(chan int, 2)
)

func main() {
	for i := 0; i < 2; i++ {
		go func() {
			for {
				select {
				case <-quite:
					return
				case x := <-target:
					time.Sleep(4 * time.Second)
					log.Info(x)
				}
			}
		}()
	}

	for i := 0; i < 50; i++ {
		target <- i

	}

}
