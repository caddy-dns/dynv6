package dynv6

import (
	"testing"

	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	libdnsdynv6 "github.com/libdns/dynv6"
)

func TestUnmarshalCaddyfile(t *testing.T) {
	for _, tc := range []struct {
		input   string
		token   string
		wantErr bool
	}{
		{input: "dynv6 abc", token: "abc"},
		{input: "dynv6 {\n\ttoken abc\n}", token: "abc"},
		{input: "dynv6 {\n\ttoken\n}", wantErr: true},
		{input: "dynv6 {\n\ttoken abc def\n}", wantErr: true},
		{input: "dynv6 abc {\n\ttoken def\n}", wantErr: true},
		{input: "dynv6", wantErr: true},
	} {
		p := &Provider{new(libdnsdynv6.Provider)}
		err := p.UnmarshalCaddyfile(caddyfile.NewTestDispenser(tc.input))
		if tc.wantErr {
			if err == nil {
				t.Errorf("%q: want an error, got token %q", tc.input, p.Provider.Token)
			}
			continue
		}
		if err != nil {
			t.Errorf("%q: unexpected error: %v", tc.input, err)
		} else if p.Provider.Token != tc.token {
			t.Errorf("%q: token %q, want %q", tc.input, p.Provider.Token, tc.token)
		}
	}
}
