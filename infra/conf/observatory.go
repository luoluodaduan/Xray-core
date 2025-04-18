package conf

import (
	"github.com/luoluodaduan/xray-core/app/observatory"
	"github.com/luoluodaduan/xray-core/app/observatory/burst"
	"github.com/luoluodaduan/xray-core/infra/conf/cfgcommon/duration"
	"google.golang.org/protobuf/proto"
)

type ObservatoryConfig struct {
	SubjectSelector   []string          `json:"subjectSelector"`
	ProbeURL          string            `json:"probeURL"`
	ProbeInterval     duration.Duration `json:"probeInterval"`
	EnableConcurrency bool              `json:"enableConcurrency"`
}

func (o *ObservatoryConfig) Build() (proto.Message, error) {
	return &observatory.Config{SubjectSelector: o.SubjectSelector, ProbeUrl: o.ProbeURL, ProbeInterval: int64(o.ProbeInterval), EnableConcurrency: o.EnableConcurrency}, nil
}

type BurstObservatoryConfig struct {
	SubjectSelector []string `json:"subjectSelector"`
	// health check settings
	HealthCheck *healthCheckSettings `json:"pingConfig,omitempty"`
}

func (b BurstObservatoryConfig) Build() (proto.Message, error) {
	if result, err := b.HealthCheck.Build(); err == nil {
		return &burst.Config{SubjectSelector: b.SubjectSelector, PingConfig: result.(*burst.HealthPingConfig)}, nil
	} else {
		return nil, err
	}
}
