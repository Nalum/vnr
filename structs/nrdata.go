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

type NRData struct {
	Agent struct {
		Host    string
		Version string
		PID     uint64
	}
	Components []NRComponent
}

func (nrd NRData) String() string {
	return fmt.Sprintf("Agent Host: %v - Agent Version: %v - Agent PID: %v", nrd.Agent.Host, nrd.Agent.Version, nrd.Agent.PID)
}
