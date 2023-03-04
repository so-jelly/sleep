package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	help := flag.Bool("help", false, "show usage")
	sleepSeconds := flag.Int("sec", 300, "seconds to wait for")
	flag.Parse()
	if *help {
		flag.Usage()
	}
	time.Sleep(time.Second * time.Duration(*sleepSeconds))
	fmt.Println("well rested")
}
