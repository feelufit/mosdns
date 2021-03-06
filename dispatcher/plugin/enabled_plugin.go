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

package plugin

import (
	// import all plugins
	_ "github.com/IrineSistiana/mosdns/dispatcher/plugin/cache"
	_ "github.com/IrineSistiana/mosdns/dispatcher/plugin/executable/blackhole"
	_ "github.com/IrineSistiana/mosdns/dispatcher/plugin/executable/ecs"
	_ "github.com/IrineSistiana/mosdns/dispatcher/plugin/executable/fallback"
	_ "github.com/IrineSistiana/mosdns/dispatcher/plugin/executable/fast_forward"
	_ "github.com/IrineSistiana/mosdns/dispatcher/plugin/executable/forward"
	_ "github.com/IrineSistiana/mosdns/dispatcher/plugin/executable/ipset"
	_ "github.com/IrineSistiana/mosdns/dispatcher/plugin/executable/pipeline"
	_ "github.com/IrineSistiana/mosdns/dispatcher/plugin/executable/sequence"
	_ "github.com/IrineSistiana/mosdns/dispatcher/plugin/hosts"
	_ "github.com/IrineSistiana/mosdns/dispatcher/plugin/http_server"
	_ "github.com/IrineSistiana/mosdns/dispatcher/plugin/logger"
	_ "github.com/IrineSistiana/mosdns/dispatcher/plugin/matcher/domain_matcher"
	_ "github.com/IrineSistiana/mosdns/dispatcher/plugin/matcher/ip_matcher"
	_ "github.com/IrineSistiana/mosdns/dispatcher/plugin/matcher/qtype_matcher"
	_ "github.com/IrineSistiana/mosdns/dispatcher/plugin/matcher/simple_matcher"
	_ "github.com/IrineSistiana/mosdns/dispatcher/plugin/plain_server"
)
