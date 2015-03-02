package process

import (
	"github.com/nalum/vnr/structs"
	"log"
)

func VNumbers(Channels map[string]chan interface{}) {
	var holder interface{}
	var data structs.VNumbers

	for {
		holder = <-Channels["vnumbers"]
		data = holder.(structs.VNumbers)
		log.Println(data)
	}
}
