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

type VClient struct {
	Timestamp string `json:"timestamp"`
	Req400    VData  `json:"MAIN.client_req_400"`
	Req411    VData  `json:"MAIN.client_req_411"`
	Req413    VData  `json:"MAIN.client_req_413"`
	Req417    VData  `json:"MAIN.client_req_417"`
	Req       VData  `json:"MAIN.client_req"`
}

func (s VClient) String() string {
	return fmt.Sprintf("Timestamp: %v - Number of Client Requests: %v", s.Timestamp, s.Req.Value)
}
