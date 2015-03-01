package structs

import "fmt"

type VCache struct {
	Timestamp string `json:"timestamp"`
	Hit       VData  `json:"MAIN.cache_hit"`
	HitPass   VData  `json:"MAIN.cache_hitpass"`
	Miss      VData  `json:"MAIN.cache_miss"`
}

func (s VCache) String() string {
	return fmt.Sprintf("Timestamp: %v - Number of Cache Hits: %v", s.Timestamp, s.Hit.Value)
}
