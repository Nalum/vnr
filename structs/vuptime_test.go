package structs

import (
	"testing"
)

func TestVUpTimeString(t *testing.T) {
	data := VUpTime{
		Timestamp: "timestamp",
		Count: VData{
			VSType:      "string",
			Value:       42,
			Flag:        "a",
			Description: "A description of the item",
		},
	}
	want := "Timestamp: timestamp - Up Time: 42"

	if result := data.String(); result != want {
		t.Errorf("Stringer function didn't return expected result.\nWant: \"%v\"\nResult: \"%v\"", want, result)
	}
}
