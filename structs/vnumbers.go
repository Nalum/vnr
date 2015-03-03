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

type VNumbers struct {
	Timestamp     string `json:"timestamp"`
	Object        VData  `json:"MAIN.n_object"`
	VampireObject VData  `json:"MAIN.n_vampireobject"`
	ObjectCore    VData  `json:"MAIN.n_objectcore"`
	ObjectHead    VData  `json:"MAIN.n_objecthead"`
	WaitingList   VData  `json:"MAIN.n_waitinglist"`
	Backend       VData  `json:"MAIN.n_backend"`
	Expired       VData  `json:"MAIN.n_expired"`
	LruNuked      VData  `json:"MAIN.n_lru_nuked"`
	LruMoved      VData  `json:"MAIN.n_lru_moved"`
	Vcl           VData  `json:"MAIN.n_vcl"`
	VclAvail      VData  `json:"MAIN.n_vcl_avail"`
	VclDiscard    VData  `json:"MAIN.n_vcl_discard"`
	Purges        VData  `json:"MAIN.n_purges"`
	ObjPurged     VData  `json:"MAIN.n_obj_purged"`
	Gzip          VData  `json:"MAIN.n_gzip"`
	Gunzip        VData  `json:"MAIN.n_gunzip"`
}

func (s VNumbers) String() string {
	return fmt.Sprintf("Timestamp: %v - Number of Objects: %v", s.Timestamp, s.Object.Value)
}
