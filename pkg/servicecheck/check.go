package servicecheck

import (
	"context"
	"fmt"
	"net/http"
)

type ServiceStatuses []*ServiceStatus

type ServiceStatus struct {
	Name  string
	Total int
	Down  int
}

func RunChecks(config *Config) ServiceStatuses {
	serviceStatuses := make(ServiceStatuses, 0)

	// iterate through each service
	services := config.HttpServices

	for _, service := range services {
		total := len(service.Instances)
		serviceName := service.Name
		down := 0

		// iterate through each service instance
		for _, instance := range service.Instances {
			// ignore errors for now
			// TODO: handle these errors
			isUp, _ := checkIsUp(instance)

			if !isUp {
				down++
			}
		}

		serviceStatuses = append(serviceStatuses, &ServiceStatus{
			Name:  serviceName,
			Total: total,
			Down:  down,
		})
	}

	return serviceStatuses
}

func checkIsUp(instance Instance) (bool, error) {
	req, err := http.NewRequestWithContext(context.TODO(), "GET", instance.StatusUrl, nil)
	if err != nil {
		return false, fmt.Errorf("error occurred while creating request: %v", err)
	}

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return false, nil
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return true, nil
	}

	return false, nil
}
