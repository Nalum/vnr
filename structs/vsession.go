package structs

import "fmt"

type VSession struct {
	Timestamp    string `json:"timestamp"`
	Conn         VData  `json:"MAIN.sess_conn"`
	Drop         VData  `json:"MAIN.sess_drop"`
	Fail         VData  `json:"MAIN.sess_fail"`
	PipeOverflow VData  `json:"MAIN.sess_pipe_overflow"`
	Closed       VData  `json:"MAIN.sess_closed"`
	PipeLine     VData  `json:"MAIN.sess_pipeline"`
	ReadAhead    VData  `json:"MAIN.sess_readahead"`
	Herd         VData  `json:"MAIN.sess_herd"`
	Queued       VData  `json:"MAIN.sess_queued"`
	Dropped      VData  `json:"MAIN.sess_dropped"`
}

func (s VSession) String() string {
	return fmt.Sprintf("Timestamp: %v - Number of Session Connections: %v", s.Timestamp, s.Conn.Value)
}
