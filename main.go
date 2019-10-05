package main

import (
	"github.com/caddyserver/caddy/caddy/caddymain"
	// Caddy plugins
	_ "github.com/caddyserver/dnsproviders/digitalocean"
	//_ "github.com/abiosoft/caddy-git"
    _ "github.com/awoodbeck/caddy-git"
)

func main() {
	// optional: disable telemetry
	// caddymain.EnableTelemetry = false
	caddymain.Run()
}