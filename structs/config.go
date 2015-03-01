package structs

import "fmt"

type Config struct {
	Key       string
	Instances []string
	Interval  int
}

func (c Config) String() string {
	return fmt.Sprintf("Key: %v - Instances: %v - Interval: %v", c.Key, c.Instances, c.Interval)
}
