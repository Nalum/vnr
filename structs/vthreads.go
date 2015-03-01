package structs

import "fmt"

type VThreads struct {
	Timestamp string `json:"timestamp"`
	Pools     VData  `json:"MAIN.pools"`
	Count     VData  `json:"MAIN.threads"`
	Limited   VData  `json:"MAIN.threads_limited"`
	Created   VData  `json:"MAIN.threads_created"`
	Destroyed VData  `json:"MAIN.threads_destroyed"`
	Failed    VData  `json:"MAIN.threads_failed"`
	QueueLen  VData  `json:"MAIN.thread_queue_len"`
}

func (s VThreads) String() string {
	return fmt.Sprintf("Timestamp: %v - Number of Threads: %v", s.Timestamp, s.Count.Value)
}
