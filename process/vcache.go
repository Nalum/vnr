package process

import (
	"github.com/nalum/vnr/structs"
	"log"
)

func VCache(Channels map[string]chan interface{}) {
	var holder interface{}
	var data structs.VCache

	for {
		holder = <-Channels["vcache"]
		data = holder.(structs.VCache)
		log.Println(data)
	}
}
