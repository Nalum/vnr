package structs

import "fmt"

type VBackend struct {
	Timestamp string `json:"timestamp"`
	Conn      VData  `json:"MAIN.backend_conn"`
	Unhealthy VData  `json:"MAIN.backend_unhealthy"`
	Busy      VData  `json:"MAIN.backend_busy"`
	Fail      VData  `json:"MAIN.backend_fail"`
	Reuse     VData  `json:"MAIN.backend_reuse"`
	TooLate   VData  `json:"MAIN.backend_toolate"`
	Recycle   VData  `json:"MAIN.backend_recycle"`
	Retry     VData  `json:"MAIN.backend_retry"`
	Req       VData  `json:"MAIN.backend_req"`
}

func (s VBackend) String() string {
	return fmt.Sprintf("Timestamp: %v - Number of Connections: %v", s.Timestamp, s.Conn.Value)
}
