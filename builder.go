package zendesk

import (
	"fmt"
	"io/ioutil"

	resty "gopkg.in/resty.v0"
	yaml "gopkg.in/yaml.v2"
)

// FromToken func loads Client basedon Auth Basic access
func FromToken(config Configuration) Client {
	username := fmt.Sprintf("%s/token", config.Email)

	restyClient := resty.SetBasicAuth(username, config.Token)
	restyClient.SetHeader("Accept", "application/json")
	restyClient.SetHeader("Content-Type", "application/json")

	return Client{
		domain:     config.Domain,
		apiVersion: config.APIVersion,
		client:     restyClient,
	}
}

// LoadConfiguration func reads cnfiguration yml file
func LoadConfiguration(path string) Configuration {
	zendeskConfiguration := Configuration{}
	zendeskConfigurationFile, err := ioutil.ReadFile(path)

	if err != nil {
		panic(fmt.Errorf("error: %v", err))
	}

	err = yaml.Unmarshal(zendeskConfigurationFile, &zendeskConfiguration)

	if err != nil {
		panic(fmt.Errorf("error: %v", err))
	}

	return zendeskConfiguration
}
