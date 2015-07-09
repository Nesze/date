package main

import (
	"fmt"
        "os"
        "os/signal"
	"time"
)

func main() {
        ctrlc := make(chan os.Signal)
        signal.Notify(ctrlc, os.Interrupt, os.Kill)

	ticker := time.Tick(time.Second)
	go func() {
            for {
		<-ticker
		fmt.Printf("\r%s", time.Now().Format(time.UnixDate))
            }
	}()

        <-ctrlc
        fmt.Println()
}
