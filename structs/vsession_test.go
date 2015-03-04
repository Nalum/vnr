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

import (
	"testing"
)

func TestVSessionString(t *testing.T) {
	data := VSession{
		Timestamp: "timestamp",
		Conn: VData{
			Type:        "string",
			Value:       42,
			Flag:        "a",
			Description: "A description of the item",
		},
		Drop: VData{
			Type:        "string",
			Value:       42,
			Flag:        "a",
			Description: "A description of the item",
		},
		Fail: VData{
			Type:        "string",
			Value:       42,
			Flag:        "a",
			Description: "A description of the item",
		},
		PipeOverflow: VData{
			Type:        "string",
			Value:       42,
			Flag:        "a",
			Description: "A description of the item",
		},
		Closed: VData{
			Type:        "string",
			Value:       42,
			Flag:        "a",
			Description: "A description of the item",
		},
		PipeLine: VData{
			Type:        "string",
			Value:       42,
			Flag:        "a",
			Description: "A description of the item",
		},
		ReadAhead: VData{
			Type:        "string",
			Value:       42,
			Flag:        "a",
			Description: "A description of the item",
		},
		Herd: VData{
			Type:        "string",
			Value:       42,
			Flag:        "a",
			Description: "A description of the item",
		},
		Queued: VData{
			Type:        "string",
			Value:       42,
			Flag:        "a",
			Description: "A description of the item",
		},
		Dropped: VData{
			Type:        "string",
			Value:       42,
			Flag:        "a",
			Description: "A description of the item",
		},
	}
	want := "Timestamp: timestamp - Number of Session Connections: 42"

	if result := data.String(); result != want {
		t.Errorf("Stringer function didn't return expected result.\nWant: \"%v\"\nResult: \"%v\"", want, result)
	}
}
