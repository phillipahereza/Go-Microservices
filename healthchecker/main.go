package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := flag.String("port", "80", "port on localhost to check")
	addr := flag.String("addr", "accountservice", "service on which to perform check")
	flag.Parse()
	resp, err := http.Get(fmt.Sprintf("http://%s:%s/health", *addr, *port))
	if err != nil || resp.StatusCode != 200 {
		os.Exit(1)
	}
	os.Exit(0)
	
}
