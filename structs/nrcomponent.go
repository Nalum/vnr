package structs

import "fmt"

type NRComponent struct {
	Name     string
	GUID     string
	Duration uint64
	Metrics  struct {
		ComponentMainUpTimeSeconds       uint64
		ComponentMainConnectedSessions   uint64
		ComponentMainDroppedSessions     uint64
		ComponentMainFailedSessions      uint64
		ComponentMainPipeSessionOverflow uint64
		ComponentMainCacheHit            uint64
		ComponentMainCacheMiss           uint64
		ComponentMainBackendConnections  uint64
		ComponentMainBackendUnhealthy    uint64
		ComponentMainBackendBusy         uint64
	}
}

func (nrc NRComponent) String() string {
	return fmt.Sprintf("Name: %v - GUID: %v - Duration: %v", nrc.Name, nrc.GUID, nrc.Duration)
}
