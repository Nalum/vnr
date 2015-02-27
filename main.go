package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os/exec"
)

const version = "0.1.0"

var show_version bool

func init() {
	flag.BoolVar(&show_version, "version", false, "Show the version of varnish-newrelic and check if there is a newer version available.")
	flag.BoolVar(&show_version, "v", false, "Show the version of varnish-newrelic and check if there is a newer version available.")
}

type config struct {
	key       string
	instances []string
	delay     int
}

type vsdata struct {
	vsType      string
	value       uint64
	flag        string
	description string
}

type stat struct {
	timestamp string
	main      struct {
		uptime                       vsdata
		sess_conn                    vsdata
		sess_drop                    vsdata
		sess_fail                    vsdata
		sess_pipe_overflow           vsdata
		client_req_400               vsdata
		client_req_411               vsdata
		client_req_413               vsdata
		client_req_417               vsdata
		client_req                   vsdata
		cache_hit                    vsdata
		cache_hitpass                vsdata
		cache_miss                   vsdata
		backend_conn                 vsdata
		backend_unhealthy            vsdata
		backend_busy                 vsdata
		backend_fail                 vsdata
		backend_reuse                vsdata
		backend_toolate              vsdata
		backend_recycle              vsdata
		backend_retry                vsdata
		fetch_head                   vsdata
		fetch_length                 vsdata
		fetch_chunked                vsdata
		fetch_eof                    vsdata
		fetch_bad                    vsdata
		fetch_close                  vsdata
		fetch_oldhttp                vsdata
		fetch_zero                   vsdata
		fetch_1xx                    vsdata
		fetch_204                    vsdata
		fetch_304                    vsdata
		fetch_failed                 vsdata
		fetch_no_thread              vsdata
		pools                        vsdata
		threads                      vsdata
		threads_limited              vsdata
		threads_created              vsdata
		threads_destroyed            vsdata
		threads_failed               vsdata
		thread_queue_len             vsdata
		busy_sleep                   vsdata
		busy_wakeup                  vsdata
		sess_queued                  vsdata
		sess_dropped                 vsdata
		n_object                     vsdata
		n_vampireobject              vsdata
		n_objectcore                 vsdata
		n_objecthead                 vsdata
		n_waitinglist                vsdata
		n_backend                    vsdata
		n_expired                    vsdata
		n_lru_nuked                  vsdata
		n_lru_moved                  vsdata
		losthdr                      vsdata
		s_sess                       vsdata
		s_req                        vsdata
		s_pipe                       vsdata
		s_pass                       vsdata
		s_fetch                      vsdata
		s_synth                      vsdata
		s_req_hdrbytes               vsdata
		s_req_bodybytes              vsdata
		s_resp_hdrbytes              vsdata
		s_resp_bodybytes             vsdata
		s_pipe_hdrbytes              vsdata
		s_pipe_in                    vsdata
		s_pipe_out                   vsdata
		sess_closed                  vsdata
		sess_pipeline                vsdata
		sess_readahead               vsdata
		sess_herd                    vsdata
		shm_records                  vsdata
		shm_writes                   vsdata
		shm_flushes                  vsdata
		shm_cont                     vsdata
		shm_cycles                   vsdata
		sms_nreq                     vsdata
		sms_nobj                     vsdata
		sms_nbytes                   vsdata
		sms_balloc                   vsdata
		sms_bfree                    vsdata
		backend_req                  vsdata
		n_vcl                        vsdata
		n_vcl_avail                  vsdata
		n_vcl_discard                vsdata
		bans                         vsdata
		bans_completed               vsdata
		bans_obj                     vsdata
		bans_req                     vsdata
		bans_added                   vsdata
		bans_deleted                 vsdata
		bans_tested                  vsdata
		bans_obj_killed              vsdata
		bans_lurker_tested           vsdata
		bans_tests_tested            vsdata
		bans_lurker_tests_tested     vsdata
		bans_lurker_obj_killed       vsdata
		bans_dups                    vsdata
		bans_lurker_contention       vsdata
		bans_persisted_bytes         vsdata
		bans_persisted_fragmentation vsdata
		n_purges                     vsdata
		n_obj_purged                 vsdata
		exp_mailed                   vsdata
		exp_received                 vsdata
		hcb_nolock                   vsdata
		hcb_lock                     vsdata
		hcb_insert                   vsdata
		esi_errors                   vsdata
		esi_warnings                 vsdata
		vmods                        vsdata
		n_gzip                       vsdata
		n_gunzip                     vsdata
		vsm_free                     vsdata
		vsm_used                     vsdata
		vsm_cooling                  vsdata
		vsm_overflow                 vsdata
		vsm_overflowed               vsdata
	}
}

func main() {
	flag.Parse()

	if show_version {
		fmt.Println("Version: " + version)
		fmt.Println("You are using the most recent release.")
		return
	}

	fmt.Println("Main Started")
	varnishstat := exec.Command("varnishstat", "-j")
	var out bytes.Buffer
	varnishstat.Stdout = &out
	err := varnishstat.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Varnish Stat: %q\n", out.String())
}
