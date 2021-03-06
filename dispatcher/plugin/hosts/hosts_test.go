//     Copyright (C) 2020, IrineSistiana
//
//     This file is part of mosdns.
//
//     mosdns is free software: you can redistribute it and/or modify
//     it under the terms of the GNU General Public License as published by
//     the Free Software Foundation, either version 3 of the License, or
//     (at your option) any later version.
//
//     mosdns is distributed in the hope that it will be useful,
//     but WITHOUT ANY WARRANTY; without even the implied warranty of
//     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//     GNU General Public License for more details.
//
//     You should have received a copy of the GNU General Public License
//     along with this program.  If not, see <https://www.gnu.org/licenses/>.

package hosts

import (
	"github.com/IrineSistiana/mosdns/dispatcher/handler"
	"github.com/miekg/dns"
	"net"
	"testing"
)

func Test_hostsContainer_Match(t *testing.T) {
	h := newHostsContainer("test")
	err := h.load("../../testdata/hosts")
	if err != nil {
		t.Fatal(err)
	}

	type args struct {
		name string
		typ  uint16
	}
	tests := []struct {
		name        string
		args        args
		wantMatched bool
		wantErr     bool
		wantAddr    []string
	}{
		{"matched A", args{name: "dns.google.", typ: dns.TypeA}, true, false, []string{"8.8.8.8", "8.8.4.4"}},
		{"matched AAAA", args{name: "dns.google.", typ: dns.TypeAAAA}, true, false, []string{"2001:4860:4860::8844", "2001:4860:4860::8888"}},
		{"not matched A", args{name: "nxdomain.com.", typ: dns.TypeA}, false, false, nil},
	}
	for _, tt := range tests {
		q := new(dns.Msg)
		q.SetQuestion(tt.args.name, tt.args.typ)
		qCtx := handler.NewContext(q)

		t.Run(tt.name, func(t *testing.T) {
			gotMatched, err := h.Match(nil, qCtx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Match() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotMatched != tt.wantMatched {
				t.Errorf("Match() gotMatched = %v, want %v", gotMatched, tt.wantMatched)
			}

			for _, s := range tt.wantAddr {
				wantIP := net.ParseIP(s)
				if wantIP == nil {
					t.Fatal("invalid test case addr")
				}
				found := false
				for _, rr := range qCtx.R.Answer {
					var ip net.IP
					switch rr := rr.(type) {
					case *dns.A:
						ip = rr.A
					case *dns.AAAA:
						ip = rr.AAAA
					default:
						continue
					}
					if ip.Equal(wantIP) {
						found = true
						break
					}
				}
				if !found {
					t.Fatal("wanted ip is not found in response")
				}
			}
		})
	}
}
