//go:build !android && with_tailscale && with_gvisor

package tailscale

import "github.com/sagernet/sing-box/experimental/libbox/platform"

func setAndroidProtectFunc(platformInterface platform.Interface) {
}
