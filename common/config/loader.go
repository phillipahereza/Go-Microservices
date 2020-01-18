package config

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
)

// Loads config from for example http://configserver:8888/accountservice/test/P8
func LoadConfigurationFromBranch(configServerUrl string, appName string, profile string, branch string) {
	url := fmt.Sprintf("%s/%s/%s/%s", configServerUrl, appName, profile, branch)
	logrus.Infof("Loading config from %s\n", url)
	body, err := fetchConfiguration(url)
	if err != nil {
		logrus.Errorf("Couldn't load configuration, cannot start. Terminating. Error: %v", err.Error())
		panic("Couldn't load configuration, cannot start. Terminating. Error: " + err.Error())
	}
	parseConfiguration(body)
}

func fetchConfiguration(url string) ([]byte, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	logrus.Printf("Getting config from %v\n", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		logrus.Errorf("Couldn't load configuration, cannot start. Terminating. Error: %v", err.Error())
		panic("Couldn't load configuration, cannot start. Terminating. Error: " + err.Error())
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Errorln("Error reading configuration: " + err.Error())
		panic("Error reading configuration: " + err.Error())
	}
	return body, err
}

func parseConfiguration(body []byte) {
	var cloudConfig springCloudConfig
	err := json.Unmarshal(body, &cloudConfig)
	if err != nil {
		logrus.Errorln("Cannot parse configuration, message: " + err.Error())
		panic("Cannot parse configuration, message: " + err.Error())
	}

	for key, value := range cloudConfig.PropertySources[0].Source {
		viper.Set(key, value)
		logrus.Infoln("Loading config property %v => %v\n", key, value)
	}
	if viper.IsSet("server_name") {
		logrus.Infoln("Successfully loaded configuration for service %s\n", viper.GetString("server_name"))
	}
}

type springCloudConfig struct {
	Name            string           `json:"name"`
	Profiles        []string         `json:"profiles"`
	Label           string           `json:"label"`
	Version         string           `json:"version"`
	PropertySources []propertySource `json:"propertySources"`
}

type propertySource struct {
	Name   string                 `json:"name"`
	Source map[string]interface{} `json:"source"`
}
