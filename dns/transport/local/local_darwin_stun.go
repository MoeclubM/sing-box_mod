//go:build darwin && !cgo

package local

import (
	"context"

	mDNS "github.com/miekg/dns"
)

func (t *Transport) systemExchange(ctx context.Context, message *mDNS.Msg) (*mDNS.Msg, error) {
	return t.exchange(ctx, message, message.Question[0].Name)
}
