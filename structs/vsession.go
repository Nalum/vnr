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

type VSession struct {
	Timestamp    string `json:"timestamp"`
	Conn         VData  `json:"MAIN.sess_conn"`
	Drop         VData  `json:"MAIN.sess_drop"`
	Fail         VData  `json:"MAIN.sess_fail"`
	PipeOverflow VData  `json:"MAIN.sess_pipe_overflow"`
	Closed       VData  `json:"MAIN.sess_closed"`
	PipeLine     VData  `json:"MAIN.sess_pipeline"`
	ReadAhead    VData  `json:"MAIN.sess_readahead"`
	Herd         VData  `json:"MAIN.sess_herd"`
	Queued       VData  `json:"MAIN.sess_queued"`
	Dropped      VData  `json:"MAIN.sess_dropped"`
}

func (s VSession) String() string {
	return fmt.Sprintf("Timestamp: %v - Number of Session Connections: %v", s.Timestamp, s.Conn.Value)
}
