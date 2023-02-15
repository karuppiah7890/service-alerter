package main

import (
	"fmt"
	"log"

	"github.com/karuppiah7890/service-alerter/pkg/config"
	"github.com/karuppiah7890/service-alerter/pkg/servicecheck"
	"github.com/karuppiah7890/service-alerter/pkg/slack"
)

func main() {
	c, err := config.NewConfigFromEnvVars()
	if err != nil {
		log.Fatalf("error occurred while getting configuration from environment variables: %v", err)
	}

	serviceCheckConfig, err := servicecheck.NewConfig(c.GetConfigFilePath())
	if err != nil {
		log.Fatalf("error occurred while loading the config from %s", c.GetConfigFilePath())
	}

	serviceStatuses := servicecheck.RunChecks(serviceCheckConfig)

	for _, serviceStatus := range serviceStatuses {
		total := serviceStatus.Total
		down := serviceStatus.Down
		serviceName := serviceStatus.Name
		if down == total {
			message := fmt.Sprintf("Critical :rotating_light:! All (%d) instances of %s service is down in %s environment :rotating_light:", total, serviceName, c.GetEnvironmentName())
			// TODO: Use Mocks to test the integration with ease for different cases with unit tests
			err := slack.SendMessage(c.GetSlackToken(), c.GetSlackChanel(), message)
			if err != nil {
				log.Fatalf("error occurred while sending slack alert message: %v", err)
			}
		}

		if down > 0 && down < total {
			message := fmt.Sprintf("Warning :warning:! Some (%d out of %d) instances of %s service is down in %s environment :warning:", down, total, serviceName, c.GetEnvironmentName())
			// TODO: Use Mocks to test the integration with ease for different cases with unit tests
			err := slack.SendMessage(c.GetSlackToken(), c.GetSlackChanel(), message)
			if err != nil {
				log.Fatalf("error occurred while sending slack alert message: %v", err)
			}
		}
	}
}
