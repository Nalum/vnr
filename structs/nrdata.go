package structs

import "fmt"

type NRData struct {
	Agent struct {
		Host    string
		Version string
		PID     uint64
	}
	Components []NRComponent
}

func (nrd NRData) String() string {
	return fmt.Sprintf("Agent Host: %v - Agent Version: %v - Agent PID: %v", nrd.Agent.Host, nrd.Agent.Version, nrd.Agent.PID)
}
