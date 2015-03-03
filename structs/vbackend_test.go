package structs

import (
	"testing"
)

func TestVBackendString(t *testing.T) {
	data := VBackend{
		Timestamp: "timestamp",
		Conn: VData{
			VSType:      "string",
			Value:       64,
			Flag:        "a",
			Description: "A description of the item",
		},
		Unhealthy: VData{
			VSType:      "string",
			Value:       64,
			Flag:        "a",
			Description: "A description of the item",
		},
		Busy: VData{
			VSType:      "string",
			Value:       64,
			Flag:        "a",
			Description: "A description of the item",
		},
		Fail: VData{
			VSType:      "string",
			Value:       64,
			Flag:        "a",
			Description: "A description of the item",
		},
		Reuse: VData{
			VSType:      "string",
			Value:       64,
			Flag:        "a",
			Description: "A description of the item",
		},
		TooLate: VData{
			VSType:      "string",
			Value:       64,
			Flag:        "a",
			Description: "A description of the item",
		},
		Recycle: VData{
			VSType:      "string",
			Value:       64,
			Flag:        "a",
			Description: "A description of the item",
		},
		Retry: VData{
			VSType:      "string",
			Value:       64,
			Flag:        "a",
			Description: "A description of the item",
		},
		Req: VData{
			VSType:      "string",
			Value:       64,
			Flag:        "a",
			Description: "A description of the item",
		},
	}
	want := "Timestamp: timestamp - Number of Connections: 64"

	if result := data.String(); result != want {
		t.Errorf("Config Stringer function didn't return expected result.\nWant: \"%v\"\nResult: \"%v\"", want, result)
	}
}
