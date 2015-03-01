package structs

import "fmt"

type VClient struct {
	Timestamp string `json:"timestamp"`
	Req400    VData  `json:"MAIN.client_req_400"`
	Req411    VData  `json:"MAIN.client_req_411"`
	Req413    VData  `json:"MAIN.client_req_413"`
	Req417    VData  `json:"MAIN.client_req_417"`
	Req       VData  `json:"MAIN.client_req"`
}

func (s VClient) String() string {
	return fmt.Sprintf("Timestamp: %v - Number of Client Requests: %v", s.Timestamp, s.Req.Value)
}
