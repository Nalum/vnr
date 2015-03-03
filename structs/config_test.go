package structs

import (
	"testing"
)

func TestConfigString(t *testing.T) {
	config := Config{
		Key:       "123",
		Instances: []string{"a", "b"},
		Interval:  3,
	}
	want := "Key: 123 - Instances: [a b] - Interval: 3"

	if result := config.String(); result != want {
		t.Errorf("Config Stringer function didn't return expected result.\nWant: \"%v\"\nResult: \"%v\"", want, result)
	}
}
