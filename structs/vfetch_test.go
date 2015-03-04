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

func TestVFetchString(t *testing.T) {
	data := VFetch{
		Timestamp: "timestamp",
		Head: VData{
			VSType:      "string",
			Value:       42,
			Flag:        "a",
			Description: "A description of the item",
		},
		Length: VData{
			VSType:      "string",
			Value:       42,
			Flag:        "a",
			Description: "A description of the item",
		},
		Chunked: VData{
			VSType:      "string",
			Value:       42,
			Flag:        "a",
			Description: "A description of the item",
		},
		EOF: VData{
			VSType:      "string",
			Value:       42,
			Flag:        "a",
			Description: "A description of the item",
		},
		Bad: VData{
			VSType:      "string",
			Value:       42,
			Flag:        "a",
			Description: "A description of the item",
		},
		Close: VData{
			VSType:      "string",
			Value:       42,
			Flag:        "a",
			Description: "A description of the item",
		},
		OldHttp: VData{
			VSType:      "string",
			Value:       42,
			Flag:        "a",
			Description: "A description of the item",
		},
		Zero: VData{
			VSType:      "string",
			Value:       42,
			Flag:        "a",
			Description: "A description of the item",
		},
		F1xx: VData{
			VSType:      "string",
			Value:       42,
			Flag:        "a",
			Description: "A description of the item",
		},
		F204: VData{
			VSType:      "string",
			Value:       42,
			Flag:        "a",
			Description: "A description of the item",
		},
		F304: VData{
			VSType:      "string",
			Value:       42,
			Flag:        "a",
			Description: "A description of the item",
		},
		Failed: VData{
			VSType:      "string",
			Value:       42,
			Flag:        "a",
			Description: "A description of the item",
		},
		NoThread: VData{
			VSType:      "string",
			Value:       42,
			Flag:        "a",
			Description: "A description of the item",
		},
	}
	want := "Timestamp: timestamp - Fetch Head: 42"

	if result := data.String(); result != want {
		t.Errorf("Stringer function didn't return expected result.\nWant: \"%v\"\nResult: \"%v\"", want, result)
	}
}
