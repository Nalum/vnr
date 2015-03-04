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

func TestVBansString(t *testing.T) {
	data := VBans{
		Timestamp: "timestamp",
		Count: VData{
			VSType:      "string",
			Value:       42,
			Flag:        "a",
			Description: "A description of the item",
		},
		Completed: VData{
			VSType:      "string",
			Value:       42,
			Flag:        "a",
			Description: "A description of the item",
		},
		Obj: VData{
			VSType:      "string",
			Value:       42,
			Flag:        "a",
			Description: "A description of the item",
		},
		Req: VData{
			VSType:      "string",
			Value:       42,
			Flag:        "a",
			Description: "A description of the item",
		},
		Added: VData{
			VSType:      "string",
			Value:       42,
			Flag:        "a",
			Description: "A description of the item",
		},
		Deleted: VData{
			VSType:      "string",
			Value:       42,
			Flag:        "a",
			Description: "A description of the item",
		},
		Tested: VData{
			VSType:      "string",
			Value:       42,
			Flag:        "a",
			Description: "A description of the item",
		},
		ObjKilled: VData{
			VSType:      "string",
			Value:       42,
			Flag:        "a",
			Description: "A description of the item",
		},
		LurkerTested: VData{
			VSType:      "string",
			Value:       42,
			Flag:        "a",
			Description: "A description of the item",
		},
		TestsTested: VData{
			VSType:      "string",
			Value:       42,
			Flag:        "a",
			Description: "A description of the item",
		},
		LurkerTestsTested: VData{
			VSType:      "string",
			Value:       42,
			Flag:        "a",
			Description: "A description of the item",
		},
		LurkerObjKilled: VData{
			VSType:      "string",
			Value:       42,
			Flag:        "a",
			Description: "A description of the item",
		},
		Dups: VData{
			VSType:      "string",
			Value:       42,
			Flag:        "a",
			Description: "A description of the item",
		},
		LurkerContention: VData{
			VSType:      "string",
			Value:       42,
			Flag:        "a",
			Description: "A description of the item",
		},
		PersistedBytes: VData{
			VSType:      "string",
			Value:       42,
			Flag:        "a",
			Description: "A description of the item",
		},
		PersistedFragmentation: VData{
			VSType:      "string",
			Value:       42,
			Flag:        "a",
			Description: "A description of the item",
		},
	}
	want := "Timestamp: timestamp - Number of Bans: 42"

	if result := data.String(); result != want {
		t.Errorf("Stringer function didn't return expected result.\nWant: \"%v\"\nResult: \"%v\"", want, result)
	}
}
