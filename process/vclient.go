package process

import (
	"github.com/nalum/vnr/structs"
	"log"
)

func VClient(Channels map[string]chan interface{}) {
	var holder interface{}
	var data structs.VClient

	for {
		holder = <-Channels["vclient"]
		data = holder.(structs.VClient)
		log.Println(data)
	}
}
