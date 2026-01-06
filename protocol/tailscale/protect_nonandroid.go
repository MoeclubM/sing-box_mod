//go:build with_tailscale && !android

package tailscale

import "github.com/sagernet/sing-box/adapter"

func setAndroidProtectFunc(platformInterface adapter.PlatformInterface) {
}
