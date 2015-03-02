package process

import (
	"github.com/nalum/vnr/structs"
	"log"
)

func VUpTime(Channels map[string]chan interface{}) {
	var holder interface{}
	var data structs.VUpTime

	for {
		holder = <-Channels["vuptime"]
		data = holder.(structs.VUpTime)
		log.Println(data)
	}
}
