package splithttp

import (
	"context"

	"github.com/luoluodaduan/xray-core/common"
)

//go:generate go run github.com/luoluodaduan/xray-core/common/errors/errorgen

const protocolName = "splithttp"

func init() {
	common.Must(common.RegisterConfig((*Config)(nil), func(ctx context.Context, config interface{}) (interface{}, error) {
		return nil, newError("splithttp is a transport protocol.")
	}))
}
