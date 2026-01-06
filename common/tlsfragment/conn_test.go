package tf_test

import (
	"context"
	"crypto/tls"
	"net"
	"testing"
	"time"

	tf "github.com/sagernet/sing-box/common/tlsfragment"

	"github.com/stretchr/testify/require"
)

func TestTLSFragment(t *testing.T) {
	t.Parallel()
	if testing.Short() {
		t.Skip("skipping network test in -short mode")
	}
	tcpConn, err := net.DialTimeout("tcp", "1.1.1.1:443", 5*time.Second)
	if err != nil {
		t.Skipf("skipping network test: %v", err)
	}
	defer tcpConn.Close()
	tlsConn := tls.Client(tf.NewConn(tcpConn, context.Background(), true, false, 0), &tls.Config{
		ServerName: "www.cloudflare.com",
	})
	require.NoError(t, tlsConn.Handshake())
}

func TestTLSRecordFragment(t *testing.T) {
	t.Parallel()
	if testing.Short() {
		t.Skip("skipping network test in -short mode")
	}
	tcpConn, err := net.DialTimeout("tcp", "1.1.1.1:443", 5*time.Second)
	if err != nil {
		t.Skipf("skipping network test: %v", err)
	}
	defer tcpConn.Close()
	tlsConn := tls.Client(tf.NewConn(tcpConn, context.Background(), false, true, 0), &tls.Config{
		ServerName: "www.cloudflare.com",
	})
	require.NoError(t, tlsConn.Handshake())
}

func TestTLS2Fragment(t *testing.T) {
	t.Parallel()
	if testing.Short() {
		t.Skip("skipping network test in -short mode")
	}
	tcpConn, err := net.DialTimeout("tcp", "1.1.1.1:443", 5*time.Second)
	if err != nil {
		t.Skipf("skipping network test: %v", err)
	}
	defer tcpConn.Close()
	tlsConn := tls.Client(tf.NewConn(tcpConn, context.Background(), true, true, 0), &tls.Config{
		ServerName: "www.cloudflare.com",
	})
	require.NoError(t, tlsConn.Handshake())
}
