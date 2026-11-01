package conf

import (
	"strings"

	"github.com/luoluodaduan/xray-core/app/commander"
	loggerservice "github.com/luoluodaduan/xray-core/app/log/command"
	observatoryservice "github.com/luoluodaduan/xray-core/app/observatory/command"
	handlerservice "github.com/luoluodaduan/xray-core/app/proxyman/command"
	routerservice "github.com/luoluodaduan/xray-core/app/router/command"
	statsservice "github.com/luoluodaduan/xray-core/app/stats/command"
	"github.com/luoluodaduan/xray-core/common/errors"
	"github.com/luoluodaduan/xray-core/common/serial"
)

type APIConfig struct {
	Tag      string   `json:"tag"`
	Listen   string   `json:"listen"`
	Services []string `json:"services"`
}

func (c *APIConfig) Build() (*commander.Config, error) {
	if c.Tag == "" {
		return nil, errors.New("API tag can't be empty.")
	}

	services := make([]*serial.TypedMessage, 0, 16)
	for _, s := range c.Services {
		switch strings.ToLower(s) {
		case "reflectionservice":
			services = append(services, serial.ToTypedMessage(&commander.ReflectionConfig{}))
		case "handlerservice":
			services = append(services, serial.ToTypedMessage(&handlerservice.Config{}))
		case "loggerservice":
			services = append(services, serial.ToTypedMessage(&loggerservice.Config{}))
		case "statsservice":
			services = append(services, serial.ToTypedMessage(&statsservice.Config{}))
		case "observatoryservice":
			services = append(services, serial.ToTypedMessage(&observatoryservice.Config{}))
		case "routingservice":
			services = append(services, serial.ToTypedMessage(&routerservice.Config{}))
		}
	}

	return &commander.Config{
		Tag:     c.Tag,
		Listen:  c.Listen,
		Service: services,
	}, nil
}
