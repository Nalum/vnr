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

package process

import (
	"github.com/nalum/vnr/structs"
	"log"
)

func VSession(Channels map[string]chan interface{}) {
	var holder interface{}
	var data structs.VSession

	for {
		holder = <-Channels["vsession"]
		data = holder.(structs.VSession)
		log.Println(data)
	}
}
