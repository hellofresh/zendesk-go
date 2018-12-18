package zendesk

import (
	"fmt"
	"io/ioutil"
	"os"

	resty "gopkg.in/resty.v0"
	yaml "gopkg.in/yaml.v2"
)

func FromToken(config ZendeskConfiguration) Client {
	username := fmt.Sprintf("%s/token", config.Email)

	restyClient := resty.SetBasicAuth(username, config.Token)
	restyClient.SetHeader("Accept", "application/json")
	restyClient.SetHeader("Content-Type", "application/json")

	return Client{
		domain:     config.Domain,
		apiVersion: config.ApiVersion,
		client:     restyClient,
	}
}

func FromEnv() Client {
	zendeskConfiguration := ZendeskConfiguration{}
	zendeskConfiguration.ApiVersion = os.Getenv("ZENDESK_API_VERSION")
	zendeskConfiguration.Domain = os.Getenv("ZENDESK_DOMAIN")
	zendeskConfiguration.Email = os.Getenv("ZENDESK_EMAIL")
	zendeskConfiguration.Password = os.Getenv("ZENDESK_PASSWORD")
	zendeskConfiguration.Token = os.Getenv("ZENDESK_TOKEN")

	fmt.Println(zendeskConfiguration)
	return FromToken(zendeskConfiguration)
}

func LoadConfiguration(path string) ZendeskConfiguration {
	zendeskConfiguration := ZendeskConfiguration{}
	zendeskConfigurationFile, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Errorf("error: %v", err)
	}

	err = yaml.Unmarshal(zendeskConfigurationFile, &zendeskConfiguration)

	if err != nil {
		fmt.Errorf("error: %v", err)
	}

	return zendeskConfiguration
}
