package dynv6

import (
	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/libdns/dynv6"
)

// Provider wraps the provider implementation as a Caddy module.
type Provider struct{ *dynv6.Provider }

func init() {
	caddy.RegisterModule(Provider{})
}

// CaddyModule returns the Caddy module information.
func (Provider) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "dns.providers.dynv6",
		New: func() caddy.Module { return &Provider{new(dynv6.Provider)} },
	}
}

// Provision implements caddy.Provisioner.
// Before using the provider config, resolve placeholders in the API token.
func (p *Provider) Provision(ctx caddy.Context) error {
	p.Provider.Token = caddy.NewReplacer().ReplaceAll(p.Provider.Token, "")
	return nil
}

// UnmarshalCaddyfile sets up the DNS provider from Caddyfile tokens. Syntax:
//
// dynv6 [<token>] {
//     token <token>
// }
//
// Expansion of placeholders in the API token is left to the JSON config caddy.Provisioner (above).
func (p *Provider) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		if d.NextArg() {
			p.Provider.Token = d.Val()
		}
		if d.NextArg() {
			return d.ArgErr()
		}
		for nesting := d.Nesting(); d.NextBlock(nesting); {
			switch d.Val() {
			case "token":
				if p.Provider.Token != "" {
					return d.Err("API token already set")
				}
				p.Provider.Token = d.Val()
				if d.NextArg() {
					return d.ArgErr()
				}
			default:
				return d.Errf("unrecognized subdirective '%s'", d.Val())
			}
		}
	}
	if p.Provider.Token == "" {
		return d.Err("missing API token")
	}
	return nil
}

// Interface guards
var (
	_ caddyfile.Unmarshaler = (*Provider)(nil)
	_ caddy.Provisioner     = (*Provider)(nil)
)
