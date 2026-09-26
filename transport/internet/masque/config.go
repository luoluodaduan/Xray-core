package masque

import (
	"github.com/luoluodaduan/xray-core/common"
	"github.com/luoluodaduan/xray-core/transport/internet"
)

const protocolName = "masque"

const DefaultPath = "/.well-known/masque/ip/*/*/"

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return &Config{
			Path: DefaultPath,
		}
	}))
}
