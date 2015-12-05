package goservicediscover


import (
	"fmt"
	"github.com/hashicorp/consul/api"
)

type Service struct {
	Name    string
	Address string
	Id      string
	Port    int
}

type ServiceDiscover interface {
	GetServices(string) ([]Service, error)
}

type ConsulService struct {
	client *api.Client
}

func MakeConsulService() (*ConsulService, error) {
	consulService := ConsulService{}
	config := api.Config{Address: "consulserver:8500"}
	client, err := api.NewClient(&config)
	consulService.client = client
	return &consulService, err
}

func (c *ConsulService) GetServices(serviceName string) ([]Service, error) {
	var services []Service
	if c.client != nil {
		health := c.client.Health()
		serviceEntries, _, err := health.Service(serviceName, "", true, nil)
		if len(serviceEntries) > 0 {
			for i := range serviceEntries {
				service := Service{}
				service.Name = serviceEntries[i].Service.Service
				service.Address = serviceEntries[i].Service.Address
				service.Port = serviceEntries[i].Service.Port
				service.Id = serviceEntries[i].Service.ID
				services = append(services, service)
			}

		}
		return services, err
	}
	return services, fmt.Errorf("You must call MakeConsulService() to get an instance of ConsulService")
}

func Init() {

}