package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main()  {
	go TimeTicker()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	select {
	case <-c:
		break
	}
}

// 定时任务
func TimeTicker()  {
	ticker :=time.NewTicker(3*time.Second)
	go func() {
		for i :=range ticker.C{
			fmt.Printf("========= %s =======\n",i)
		}
	}()
}
