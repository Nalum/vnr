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

type VBans struct {
	Timestamp              string `json:"timestamp"`
	Count                  VData  `json:"MAIN.bans"`
	Completed              VData  `json:"MAIN.bans_completed"`
	Obj                    VData  `json:"MAIN.bans_obj"`
	Req                    VData  `json:"MAIN.bans_req"`
	Added                  VData  `json:"MAIN.bans_added"`
	Deleted                VData  `json:"MAIN.bans_deleted"`
	Tested                 VData  `json:"MAIN.bans_tested"`
	ObjKilled              VData  `json:"MAIN.bans_obj_killed"`
	LurkerTested           VData  `json:"MAIN.bans_lurker_tested"`
	TestsTested            VData  `json:"MAIN.bans_tests_tested"`
	LurkerTestsTested      VData  `json:"MAIN.bans_lurker_tests_tested"`
	LurkerObjKilled        VData  `json:"MAIN.bans_lurker_obj_killed"`
	Dups                   VData  `json:"MAIN.bans_dups"`
	LurkerContention       VData  `json:"MAIN.bans_lurker_contention"`
	PersistedBytes         VData  `json:"MAIN.bans_persisted_bytes"`
	PersistedFragmentation VData  `json:"MAIN.bans_persisted_fragmentation"`
}

func (s VBans) String() string {
	return fmt.Sprintf("Timestamp: %v - Number of Bans: %v", s.Timestamp, s.Count.Value)
}
