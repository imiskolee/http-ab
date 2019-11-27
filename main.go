package main

import (
	"flag"
	"fmt"
)

func main() {
	configStr := flag.String("config","./config.json","config...")
	flag.Parse()
	config,err := NewConfigFromFile(*configStr)
	if err != nil {
		fmt.Println(err)
	}
	r1,err := DoHttpRequest(config.Mainly)
	if err != nil {
		fmt.Println("can't do mainly request:" + err.Error())
	}
	r2,err := DoHttpRequest(config.Secondly)
	if err != nil {
		fmt.Println("can't do secondly request:" + err.Error())
	}
	CompareResponse(config,r1,r2)
}
