package structs

import (
	"testing"
)

func TestVThreadsString(t *testing.T) {
	data := VThreads{
		Timestamp: "timestamp",
		Pools: VData{
			VSType:      "string",
			Value:       42,
			Flag:        "a",
			Description: "A description of the item",
		},
		Count: VData{
			VSType:      "string",
			Value:       42,
			Flag:        "a",
			Description: "A description of the item",
		},
		Limited: VData{
			VSType:      "string",
			Value:       42,
			Flag:        "a",
			Description: "A description of the item",
		},
		Created: VData{
			VSType:      "string",
			Value:       42,
			Flag:        "a",
			Description: "A description of the item",
		},
		Destroyed: VData{
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
		QueueLen: VData{
			VSType:      "string",
			Value:       42,
			Flag:        "a",
			Description: "A description of the item",
		},
	}
	want := "Timestamp: timestamp - Number of Threads: 42"

	if result := data.String(); result != want {
		t.Errorf("Stringer function didn't return expected result.\nWant: \"%v\"\nResult: \"%v\"", want, result)
	}
}
