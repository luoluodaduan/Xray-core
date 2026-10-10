package tcp

import (
	"github.com/luoluodaduan/xray-core/common"
	"github.com/luoluodaduan/xray-core/transport/internet"
)

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
