package dns_test

import (
	"context"
	"testing"
	"time"

	. "github.com/luoluodaduan/xray-core/app/dns"
	"github.com/luoluodaduan/xray-core/common"
	"github.com/luoluodaduan/xray-core/common/net"
	"github.com/luoluodaduan/xray-core/features/dns"
)

func TestLocalNameServer(t *testing.T) {
	s := NewLocalNameServer(QueryStrategy_USE_IP)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	ips, _, err := s.QueryIP(ctx, "google.com", net.IP{}, dns.IPOption{
		IPv4Enable: true,
		IPv6Enable: true,
		FakeEnable: false,
	}, false)
	cancel()
	common.Must(err)
	if len(ips) == 0 {
		t.Error("expect some ips, but got 0")
	}
}
