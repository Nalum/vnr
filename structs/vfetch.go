package structs

import "fmt"

type VFetch struct {
	Timestamp string `json:"timestamp"`
	Head      VData  `json:"MAIN.fetch_head"`
	Length    VData  `json:"MAIN.fetch_length"`
	Chunked   VData  `json:"MAIN.fetch_chunked"`
	EOF       VData  `json:"MAIN.fetch_eof"`
	Bad       VData  `json:"MAIN.fetch_bad"`
	Close     VData  `json:"MAIN.fetch_close"`
	OldHttp   VData  `json:"MAIN.fetch_oldhttp"`
	Zero      VData  `json:"MAIN.fetch_zero"`
	F1xx      VData  `json:"MAIN.fetch_1xx"`
	F204      VData  `json:"MAIN.fetch_204"`
	F304      VData  `json:"MAIN.fetch_304"`
	Failed    VData  `json:"MAIN.fetch_failed"`
	NoThread  VData  `json:"MAIN.fetch_no_thread"`
}

func (s VFetch) String() string {
	return fmt.Sprintf("Timestamp: %v - Fetch Head: %v", s.Timestamp, s.Head.Value)
}
