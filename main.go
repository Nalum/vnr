package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/nalum/vnr/structs"
	"io/ioutil"
	"log"
	"os/exec"
	"time"
)

var showVersion bool

func init() {
	flag.BoolVar(&showVersion, "version", false, "Show the version of varnish-newrelic and check if there is a newer version available.")
}

var VNRConfig structs.Config
var UpTime structs.VUpTime
var Session structs.VSession
var Cache structs.VCache

func main() {
	flag.Parse()

	if showVersion {
		fmt.Println("Version: " + Version)
		fmt.Println("You are using the most recent release.")
		return
	}

	CFGJson, err := ioutil.ReadFile("/etc/varnish-newrelic.json")

	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(CFGJson, &VNRConfig); err != nil {
		panic(err)
	}

	log.Println(VNRConfig)

	tick := time.NewTicker(time.Second * time.Duration(VNRConfig.Interval))

	for _ = range tick.C {
		log.Println("Tick")
		go varnishStat()
	}
}

func varnishStat() {
	varnishstat := exec.Command("varnishstat", "-j")
	var out bytes.Buffer
	varnishstat.Stdout = &out
	err := varnishstat.Run()

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Command Executed")
	}

	if err := json.Unmarshal(out.Bytes(), &UpTime); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(out.Bytes(), &Session); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(out.Bytes(), &Cache); err != nil {
		panic(err)
	}

	log.Println(UpTime)
	log.Println(Session)
	log.Println(Cache)
}
