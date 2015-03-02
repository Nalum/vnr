package process

import (
	"github.com/nalum/vnr/structs"
	"log"
)

func VFetch(Channels map[string]chan interface{}) {
	var holder interface{}
	var data structs.VFetch

	for {
		holder = <-Channels["vfetch"]
		data = holder.(structs.VFetch)
		log.Println(data)
	}
}
