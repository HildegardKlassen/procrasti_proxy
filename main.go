package main

import (
	"fmt"
	"log"

	"github.com/HildegardKlassen/procrasti_proxy/prcproxy"
)

func main() {
	fmt.Println("The procrasti proxy is started...")
	err := prcproxy.Run()
	log.Fatalf(err.Error())
}
