package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	sleepSeconds := int64(86400)
	if len(os.Args) > 1 {
		arg, err := strconv.ParseInt(os.Args[1], 10, 0)
		if err != nil {
			fmt.Println("Sleep time must be an integer")
			os.Exit(1)
		}
		sleepSeconds = arg
	}
	time.Sleep(time.Second * time.Duration(sleepSeconds))
	fmt.Println("well rested")
}
