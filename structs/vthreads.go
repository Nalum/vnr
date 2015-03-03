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

type VThreads struct {
	Timestamp string `json:"timestamp"`
	Pools     VData  `json:"MAIN.pools"`
	Count     VData  `json:"MAIN.threads"`
	Limited   VData  `json:"MAIN.threads_limited"`
	Created   VData  `json:"MAIN.threads_created"`
	Destroyed VData  `json:"MAIN.threads_destroyed"`
	Failed    VData  `json:"MAIN.threads_failed"`
	QueueLen  VData  `json:"MAIN.thread_queue_len"`
}

func (s VThreads) String() string {
	return fmt.Sprintf("Timestamp: %v - Number of Threads: %v", s.Timestamp, s.Count.Value)
}
