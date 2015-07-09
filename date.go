package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.Tick(time.Second)
	for {
		<-ticker
		fmt.Printf("\r%s", time.Now().Format(time.UnixDate))
	}
}
