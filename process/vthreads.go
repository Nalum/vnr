package process

import (
	"github.com/nalum/vnr/structs"
	"log"
)

func VThreads(Channels map[string]chan interface{}) {
	var holder interface{}
	var data structs.VThreads

	for {
		holder = <-Channels["vthreads"]
		data = holder.(structs.VThreads)
		log.Println(data)
	}
}
