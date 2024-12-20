package conf

import (
	"github.com/luoluodaduan/xray-core/app/metrics"
	"github.com/luoluodaduan/xray-core/common/errors"
)

type MetricsConfig struct {
	Tag string `json:"tag"`
}

func (c *MetricsConfig) Build() (*metrics.Config, error) {
	if c.Tag == "" {
		return nil, errors.New("metrics tag can't be empty.")
	}

	return &metrics.Config{
		Tag: c.Tag,
	}, nil
}
