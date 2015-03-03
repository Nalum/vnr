package structs

import (
	"testing"
)

func TestVSessionString(t *testing.T) {
	data := VSession{
		Timestamp: "timestamp",
		Conn: VData{
			VSType:      "string",
			Value:       42,
			Flag:        "a",
			Description: "A description of the item",
		},
		Drop: VData{
			VSType:      "string",
			Value:       42,
			Flag:        "a",
			Description: "A description of the item",
		},
		Fail: VData{
			VSType:      "string",
			Value:       42,
			Flag:        "a",
			Description: "A description of the item",
		},
		PipeOverflow: VData{
			VSType:      "string",
			Value:       42,
			Flag:        "a",
			Description: "A description of the item",
		},
		Closed: VData{
			VSType:      "string",
			Value:       42,
			Flag:        "a",
			Description: "A description of the item",
		},
		PipeLine: VData{
			VSType:      "string",
			Value:       42,
			Flag:        "a",
			Description: "A description of the item",
		},
		ReadAhead: VData{
			VSType:      "string",
			Value:       42,
			Flag:        "a",
			Description: "A description of the item",
		},
		Herd: VData{
			VSType:      "string",
			Value:       42,
			Flag:        "a",
			Description: "A description of the item",
		},
		Queued: VData{
			VSType:      "string",
			Value:       42,
			Flag:        "a",
			Description: "A description of the item",
		},
		Dropped: VData{
			VSType:      "string",
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
