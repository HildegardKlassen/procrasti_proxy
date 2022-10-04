package main

import (
	"fmt"
	"log"
	"net/http"
	"procrasti_proxy/prcproxy"
)

func main() {

	pP := prcproxy.NewPrcProxy(uint(2222), false, true, []string{}, 0)
	router := pP.CreateRouter()

	log.Fatal(http.ListenAndServe(fmt.Sprintf("localhost:%d", pP.Port), router))

}
