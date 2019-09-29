package main

import (
	"github.com/caddyserver/caddy/caddy/caddymain"
	// Caddy plugins
	_ "github.com/caddyserver/dnsproviders/digitalocean"
)

func main() {
	// optional: disable telemetry
	// caddymain.EnableTelemetry = false
	caddymain.Run()
}