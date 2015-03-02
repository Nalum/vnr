package process

import (
	"github.com/nalum/vnr/structs"
	"log"
)

func VSession(Channels map[string]chan interface{}) {
	var holder interface{}
	var data structs.VSession

	for {
		holder = <-Channels["vsession"]
		data = holder.(structs.VSession)
		log.Println(data)
	}
}
