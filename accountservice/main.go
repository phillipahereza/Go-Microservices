package main

import (
	"fmt"
	"github.com/phillipahereza/go-microservices/accountservice/service"
)

var appName = "accountservice"

func main() {
	fmt.Printf("Starting %v\n", appName)
	service.StartWebserver("8080")
}