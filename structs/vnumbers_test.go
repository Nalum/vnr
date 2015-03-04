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

func TestVNumbersString(t *testing.T) {
	data := VNumbers{
		Timestamp: "timestamp",
		Object: VData{
			Type:        "string",
			Value:       42,
			Flag:        "a",
			Description: "A description of the item",
		},
		VampireObject: VData{
			Type:        "string",
			Value:       42,
			Flag:        "a",
			Description: "A description of the item",
		},
		ObjectCore: VData{
			Type:        "string",
			Value:       42,
			Flag:        "a",
			Description: "A description of the item",
		},
		ObjectHead: VData{
			Type:        "string",
			Value:       42,
			Flag:        "a",
			Description: "A description of the item",
		},
		WaitingList: VData{
			Type:        "string",
			Value:       42,
			Flag:        "a",
			Description: "A description of the item",
		},
		Backend: VData{
			Type:        "string",
			Value:       42,
			Flag:        "a",
			Description: "A description of the item",
		},
		Expired: VData{
			Type:        "string",
			Value:       42,
			Flag:        "a",
			Description: "A description of the item",
		},
		LruNuked: VData{
			Type:        "string",
			Value:       42,
			Flag:        "a",
			Description: "A description of the item",
		},
		LruMoved: VData{
			Type:        "string",
			Value:       42,
			Flag:        "a",
			Description: "A description of the item",
		},
		Vcl: VData{
			Type:        "string",
			Value:       42,
			Flag:        "a",
			Description: "A description of the item",
		},
		VclAvail: VData{
			Type:        "string",
			Value:       42,
			Flag:        "a",
			Description: "A description of the item",
		},
		VclDiscard: VData{
			Type:        "string",
			Value:       42,
			Flag:        "a",
			Description: "A description of the item",
		},
		Purges: VData{
			Type:        "string",
			Value:       42,
			Flag:        "a",
			Description: "A description of the item",
		},
		ObjPurged: VData{
			Type:        "string",
			Value:       42,
			Flag:        "a",
			Description: "A description of the item",
		},
		Gzip: VData{
			Type:        "string",
			Value:       42,
			Flag:        "a",
			Description: "A description of the item",
		},
		Gunzip: VData{
			Type:        "string",
			Value:       42,
			Flag:        "a",
			Description: "A description of the item",
		},
	}
	want := "Timestamp: timestamp - Number of Objects: 42"

	if result := data.String(); result != want {
		t.Errorf("Stringer function didn't return expected result.\nWant: \"%v\"\nResult: \"%v\"", want, result)
	}
}
