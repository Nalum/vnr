package structs

import "testing"

func TestString(t *testing.T) {
	config := Config{
		Key:       "123",
		Instances: []string{"a", "b"},
		Interval:  3,
	}

	if result := config.String(); result != "Key: 123 - Instances: [a b] - Interval: 3" {
		t.Errorf("Config Stringer function didn't return expected result. \"%v\"", result)
	}
}
