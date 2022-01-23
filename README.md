dynv6 module for Caddy
===========================

This package contains a DNS provider module for [Caddy](https://github.com/caddyserver/caddy). It can be used to manage DNS records with dynv6 accounts.

## Caddy module name

```
dns.providers.dynv6
```

## Config examples

To use this module for the ACME DNS challenge, [configure the ACME issuer in your Caddy JSON](https://caddyserver.com/docs/json/apps/tls/automation/policies/issuer/acme/) like so:

```json
{
	"module": "acme",
	"challenges": {
		"dns": {
			"provider": {
				"name": "dynv6",
				"api_token": "{env.DYNV6_API_TOKEN}"
			}
		}
	}
}
```

or with the Caddyfile:

```
tls {
	dns dynv6 {env.DYNV6_API_TOKEN}
}
```

You can replace `{env.DYNV6_API_TOKEN}` with the actual auth token if you prefer to put it directly in your config instead of an environment variable.


## Authenticating

See [the associated README in the libdns package](https://github.com/libdns/dynv6) for important information about credentials.