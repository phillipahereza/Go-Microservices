package main

import (
	"fmt"
	"github.com/phillipahereza/go_microservices/accountservice/dbclient"
	"github.com/phillipahereza/go_microservices/accountservice/service"
)

var appName = "accountservice"

func main() {
	fmt.Printf("Starting %v\n", appName)
	initializeBoltClient()
	service.StartWebserver("8080")
}

func initializeBoltClient() {
	service.DBClient = &dbclient.BoltClient{}
	service.DBClient.OpenBoltDb()
	service.DBClient.Seed()

}