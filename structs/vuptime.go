package structs

import "fmt"

type VUpTime struct {
	Timestamp string `json:"timestamp"`
	Count     VData  `json:"MAIN.uptime"`
}

func (s VUpTime) String() string {
	return fmt.Sprintf("Timestamp: %v - Up Time: %v", s.Timestamp, s.Count.Value)
}
