//go:build darwin && !cgo

package local

import (
	"context"

	C "github.com/sagernet/sing-box/constant"
	"github.com/sagernet/sing-box/dns"

	mDNS "github.com/miekg/dns"
)

func (t *Transport) Exchange(ctx context.Context, message *mDNS.Msg) (*mDNS.Msg, error) {
	question := message.Question[0]
	if t.hosts != nil && (question.Qtype == mDNS.TypeA || question.Qtype == mDNS.TypeAAAA) {
		addresses := t.hosts.Lookup(dns.FqdnToDomain(question.Name))
		if len(addresses) > 0 {
			return dns.FixedResponse(message.Id, question, addresses, C.DefaultDNSTTL), nil
		}
	}
	if t.fallback && t.dhcpTransport != nil {
		dhcpServers := t.dhcpTransport.Fetch()
		if len(dhcpServers) > 0 {
			return t.dhcpTransport.Exchange0(ctx, message, dhcpServers)
		}
	}
	// Non-cgo Darwin builds cannot call libresolv, so fall back to querying system DNS servers directly.
	return t.exchange(ctx, message, question.Name)
}
