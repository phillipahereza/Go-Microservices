package main

import (
	"flag"
	"fmt"
	"github.com/phillipahereza/go_microservices/accountservice/config"
	"github.com/phillipahereza/go_microservices/accountservice/dbclient"
	"github.com/phillipahereza/go_microservices/accountservice/service"
	"github.com/spf13/viper"
)

var appName = "accountservice"

func init() {
	profile := flag.String("profile", "test", "Environment profile, something similar to spring profiles")
	configServerUrl := flag.String("configServerUrl", "http://configserver:8888", "Address to the config server")
	configBranch := flag.String("configBranch", "master", "git branch from which to fetch configuration")
	flag.Parse()

	viper.Set("profile", *profile)
	viper.Set("configServerUrl", *configServerUrl)
	viper.Set("configBranch", *configBranch)
}

func main() {
	fmt.Printf("Starting %v\n", appName)
	config.LoadConfigurationFromBranch(
		viper.GetString("configServerUrl"),
		appName,
		viper.GetString("profile"),
		viper.GetString("configBranch"))

	initializeBoltClient()
	service.StartWebserver("6767")
}

func initializeBoltClient() {
	service.DBClient = &dbclient.BoltClient{}
	service.DBClient.OpenBoltDb()
	service.DBClient.Seed()

}