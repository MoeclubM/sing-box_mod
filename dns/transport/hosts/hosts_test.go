package hosts_test

import (
	"net/netip"
	"os"
	"testing"

	"github.com/sagernet/sing-box/dns/transport/hosts"

	"github.com/stretchr/testify/require"
)

func TestHosts(t *testing.T) {
	t.Parallel()
	require.Equal(t, []netip.Addr{netip.AddrFrom4([4]byte{127, 0, 0, 1}), netip.IPv6Loopback()}, hosts.NewFile("testdata/hosts").Lookup("localhost"))
	if hosts.DefaultPath == "" {
		t.Skip("no default hosts path on this platform")
	}
	if _, err := os.Stat(hosts.DefaultPath); err != nil {
		t.Skipf("default hosts file not available: %v", err)
	}
	if result := hosts.NewFile(hosts.DefaultPath).Lookup("localhost"); len(result) == 0 {
		t.Skip("default hosts file does not contain localhost")
	}
}
