// VNR - Plugin for NewRelic which monitors `varnishstat`
// Copyright (C) 2015  Luke Mallon
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package structs

import "fmt"

type VFetch struct {
	Timestamp string `json:"timestamp"`
	Head      VData  `json:"MAIN.fetch_head"`
	Length    VData  `json:"MAIN.fetch_length"`
	Chunked   VData  `json:"MAIN.fetch_chunked"`
	EOF       VData  `json:"MAIN.fetch_eof"`
	Bad       VData  `json:"MAIN.fetch_bad"`
	Close     VData  `json:"MAIN.fetch_close"`
	OldHttp   VData  `json:"MAIN.fetch_oldhttp"`
	Zero      VData  `json:"MAIN.fetch_zero"`
	F1xx      VData  `json:"MAIN.fetch_1xx"`
	F204      VData  `json:"MAIN.fetch_204"`
	F304      VData  `json:"MAIN.fetch_304"`
	Failed    VData  `json:"MAIN.fetch_failed"`
	NoThread  VData  `json:"MAIN.fetch_no_thread"`
}

func (s VFetch) String() string {
	return fmt.Sprintf("Timestamp: %v - Fetch Head: %v", s.Timestamp, s.Head.Value)
}
