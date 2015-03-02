package process

import (
	"github.com/nalum/vnr/structs"
	"log"
)

func VBans(Channels map[string]chan interface{}) {
	var holder interface{}
	var data structs.VBans

	for {
		holder = <-Channels["vbans"]
		data = holder.(structs.VBans)
		log.Println(data)
	}
}
