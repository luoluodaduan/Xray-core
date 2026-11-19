package masque

import (
	"time"

	"github.com/luoluodaduan/xray-core/common"
	"github.com/luoluodaduan/xray-core/transport/internet"
)

const protocolName = "masque"

const DefaultPath = "/.well-known/masque/ip/*/*/"

func (c *Config) keepAlivePeriod() time.Duration {
	if c.Xmux == nil {
		return 0
	}
	return time.Duration(c.Xmux.HKeepAlivePeriod) * time.Second
}

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return &Config{
			Path: DefaultPath,
		}
	}))
}
