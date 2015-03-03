// VNR - Plugin for NewRelic which monitors `varnishstat`
// Copyright (C) 2015  Luke Mallon
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/nalum/vnr/process"
	"github.com/nalum/vnr/structs"
	"io/ioutil"
	"log"
	"os/exec"
	"time"
)

var showVersion bool
var VNRConfig structs.Config
var Channels = map[string]chan interface{}{
	"vbackend": make(chan interface{}),
	"vbans":    make(chan interface{}),
	"vcache":   make(chan interface{}),
	"vclient":  make(chan interface{}),
	"vfetch":   make(chan interface{}),
	"vnumbers": make(chan interface{}),
	"vsession": make(chan interface{}),
	"vthreads": make(chan interface{}),
	"vuptime":  make(chan interface{}),
}

func init() {
	flag.BoolVar(&showVersion, "version", false, "Show the version of varnish-newrelic and check if there is a newer version available.")
}

func main() {
	flag.Parse()

	if showVersion {
		fmt.Println("Version: " + Version)
		fmt.Println("You are using the most recent release.")
		return
	}

	CFGJson, err := ioutil.ReadFile("/etc/varnish-newrelic.json")

	if err != nil {
		log.Fatalln(err)
	}

	if err := json.Unmarshal(CFGJson, &VNRConfig); err != nil {
		log.Fatalln(err)
	}

	log.Println(VNRConfig)

	tick := time.NewTicker(time.Second * time.Duration(VNRConfig.Interval))

	go process.VBackend(Channels)
	go process.VBans(Channels)
	go process.VCache(Channels)
	go process.VClient(Channels)
	go process.VFetch(Channels)
	go process.VNumbers(Channels)
	go process.VSession(Channels)
	go process.VThreads(Channels)
	go process.VUpTime(Channels)

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

	var Backend structs.VBackend
	var Bans structs.VBans
	var Cache structs.VCache
	var Client structs.VClient
	var Fetch structs.VFetch
	var Numbers structs.VNumbers
	var Session structs.VSession
	var Threads structs.VThreads
	var UpTime structs.VUpTime

	if err := json.Unmarshal(out.Bytes(), &Backend); err != nil {
		log.Fatalln(err)
	}

	Channels["vbackend"] <- Backend

	if err := json.Unmarshal(out.Bytes(), &Bans); err != nil {
		log.Fatalln(err)
	}

	Channels["vbans"] <- Bans

	if err := json.Unmarshal(out.Bytes(), &Cache); err != nil {
		log.Fatalln(err)
	}

	Channels["vcache"] <- Cache

	if err := json.Unmarshal(out.Bytes(), &Client); err != nil {
		log.Fatalln(err)
	}

	Channels["vclient"] <- Client

	if err := json.Unmarshal(out.Bytes(), &Fetch); err != nil {
		log.Fatalln(err)
	}

	Channels["vfetch"] <- Fetch

	if err := json.Unmarshal(out.Bytes(), &Numbers); err != nil {
		log.Fatalln(err)
	}

	Channels["vnumbers"] <- Numbers

	if err := json.Unmarshal(out.Bytes(), &Session); err != nil {
		log.Fatalln(err)
	}

	Channels["vsession"] <- Session

	if err := json.Unmarshal(out.Bytes(), &Threads); err != nil {
		log.Fatalln(err)
	}

	Channels["vthreads"] <- Threads

	if err := json.Unmarshal(out.Bytes(), &UpTime); err != nil {
		log.Fatalln(err)
	}

	Channels["vuptime"] <- UpTime
}
