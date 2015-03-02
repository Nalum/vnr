package process

import (
	"github.com/nalum/vnr/structs"
	"log"
)

func VBackend(Channels map[string]chan interface{}) {
	var holder interface{}
	var data structs.VBackend

	for {
		holder = <-Channels["vbackend"]
		data = holder.(structs.VBackend)
		log.Println(data)
	}
}
