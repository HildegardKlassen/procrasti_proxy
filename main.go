package main

import (
	"fmt"
	"log"
	"procrasti_proxy/prcproxy"
)

func main() {
	fmt.Println("The procrasti proxy is started...")
	err := prcproxy.Run()
	log.Fatalf(err.Error())
}
