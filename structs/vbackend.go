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

type VBackend struct {
	Timestamp string `json:"timestamp"`
	Conn      VData  `json:"MAIN.backend_conn"`
	Unhealthy VData  `json:"MAIN.backend_unhealthy"`
	Busy      VData  `json:"MAIN.backend_busy"`
	Fail      VData  `json:"MAIN.backend_fail"`
	Reuse     VData  `json:"MAIN.backend_reuse"`
	TooLate   VData  `json:"MAIN.backend_toolate"`
	Recycle   VData  `json:"MAIN.backend_recycle"`
	Retry     VData  `json:"MAIN.backend_retry"`
	Req       VData  `json:"MAIN.backend_req"`
}

func (s VBackend) String() string {
	return fmt.Sprintf("Timestamp: %v - Number of Connections: %v", s.Timestamp, s.Conn.Value)
}
