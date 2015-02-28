package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"time"
)

const version = "0.1.0"

var showVersion bool
var initSystem string

func init() {
	flag.BoolVar(&showVersion, "version", false, "Show the version of varnish-newrelic and check if there is a newer version available.")
	flag.BoolVar(&showVersion, "v", false, "Show the version of varnish-newrelic and check if there is a newer version available.")
}

type Config struct {
	Key       string
	Instances []string
	Interval  int
}

func (c Config) String() string {
	return fmt.Sprintf("Key: %v\nInstances: %v\nInterval: %v", c.Key, c.Instances, c.Interval)
}

var VNRConfig Config

type VSData struct {
	VSType      string `json:"type"`
	Value       uint64 `json:"value"`
	Flag        string `json:"flag"`
	Description string `json:"description"`
}

type Stat struct {
	Timestamp                      string `json:"timestamp"`
	MainUpTime                     VSData `json:"MAIN.uptime"`
	MainSessConn                   VSData `json:"MAIN.sess_conn"`
	MainSessDrop                   VSData `json:"MAIN.sess_drop"`
	MainSessFail                   VSData `json:"MAIN.sess_fail"`
	MainSessPipeOverflow           VSData `json:"MAIN.sess_pipe_overflow"`
	MainClientReq400               VSData `json:"MAIN.client_req_400"`
	MainClientReq411               VSData `json:"MAIN.client_req_411"`
	MainClientReq413               VSData `json:"MAIN.client_req_413"`
	MainClientReq417               VSData `json:"MAIN.client_req_417"`
	MainClientReq                  VSData `json:"MAIN.client_req"`
	MainCacheHit                   VSData `json:"MAIN.cache_hit"`
	MainCacheHitPass               VSData `json:"MAIN.cache_hitpass"`
	MainCacheMiss                  VSData `json:"MAIN.cache_miss"`
	MainBackendConn                VSData `json:"MAIN.backend_conn"`
	MainBackendUnhealthy           VSData `json:"MAIN.backend_unhealthy"`
	MainBackendBusy                VSData `json:"MAIN.backend_busy"`
	MainBackendFail                VSData `json:"MAIN.backend_fail"`
	MainBackendReuse               VSData `json:"MAIN.backend_reuse"`
	MainBackendTooLate             VSData `json:"MAIN.backend_toolate"`
	MainBackendRecycle             VSData `json:"MAIN.backend_recycle"`
	MainBackendRetry               VSData `json:"MAIN.backend_retry"`
	MainFetchHead                  VSData `json:"MAIN.fetch_head"`
	MainFetchLength                VSData `json:"MAIN.fetch_length"`
	MainFetchChunked               VSData `json:"MAIN.fetch_chunked"`
	MainFetchEOF                   VSData `json:"MAIN.fetch_eof"`
	MainFetchBad                   VSData `json:"MAIN.fetch_bad"`
	MainFetchClose                 VSData `json:"MAIN.fetch_close"`
	MainFetchOldHttp               VSData `json:"MAIN.fetch_oldhttp"`
	MainFetchZero                  VSData `json:"MAIN.fetch_zero"`
	MainFetch1xx                   VSData `json:"MAIN.fetch_1xx"`
	MainFetch204                   VSData `json:"MAIN.fetch_204"`
	MainFetch304                   VSData `json:"MAIN.fetch_304"`
	MainFetchFailed                VSData `json:"MAIN.fetch_failed"`
	MainFetchNoThread              VSData `json:"MAIN.fetch_no_thread"`
	MainPools                      VSData `json:"MAIN.pools"`
	MainThreads                    VSData `json:"MAIN.threads"`
	MainThreadsLimited             VSData `json:"MAIN.threads_limited"`
	MainThreadsCreated             VSData `json:"MAIN.threads_created"`
	MainThreadsDestroyed           VSData `json:"MAIN.threads_destroyed"`
	MainThreadsFailed              VSData `json:"MAIN.threads_failed"`
	MainThreadQueueLen             VSData `json:"MAIN.thread_queue_len"`
	MainBusySleep                  VSData `json:"MAIN.busy_sleep"`
	MainBusyWakeup                 VSData `json:"MAIN.busy_wakeup"`
	MainSessQueued                 VSData `json:"MAIN.sess_queued"`
	MainSessDropped                VSData `json:"MAIN.sess_dropped"`
	MainNObject                    VSData `json:"MAIN.n_object"`
	MainNVampireObject             VSData `json:"MAIN.n_vampireobject"`
	MainNObjectCore                VSData `json:"MAIN.n_objectcore"`
	MainNObjectHead                VSData `json:"MAIN.n_objecthead"`
	MainNWaitingList               VSData `json:"MAIN.n_waitinglist"`
	MainNBackend                   VSData `json:"MAIN.n_backend"`
	MainNExpired                   VSData `json:"MAIN.n_expired"`
	MainNLruNuked                  VSData `json:"MAIN.n_lru_nuked"`
	MainNLruMoved                  VSData `json:"MAIN.n_lru_moved"`
	MainLostHdr                    VSData `json:"MAIN.losthdr"`
	MainSSess                      VSData `json:"MAIN.s_sess"`
	MainSReq                       VSData `json:"MAIN.s_req"`
	MainSPipe                      VSData `json:"MAIN.s_pipe"`
	MainSPass                      VSData `json:"MAIN.s_pass"`
	MainSFetch                     VSData `json:"MAIN.s_fetch"`
	MainSSynth                     VSData `json:"MAIN.s_synth"`
	MainSReqHdrBytes               VSData `json:"MAIN.s_req_hdrbytes"`
	MainSReqBodyBytes              VSData `json:"MAIN.s_req_bodybytes"`
	MainSRespHdrBytes              VSData `json:"MAIN.s_resp_hdrbytes"`
	MainSRespBodyBytes             VSData `json:"MAIN.s_resp_bodybytes"`
	MainSPipeHdrBytes              VSData `json:"MAIN.s_pipe_hdrbytes"`
	MainSPipeIn                    VSData `json:"MAIN.s_pipe_in"`
	MainSPipeOut                   VSData `json:"MAIN.s_pipe_out"`
	MainSessClosed                 VSData `json:"MAIN.sess_closed"`
	MainSessPipeLine               VSData `json:"MAIN.sess_pipeline"`
	MainSessReadAhead              VSData `json:"MAIN.sess_readahead"`
	MainSessHerd                   VSData `json:"MAIN.sess_herd"`
	MainShmRecords                 VSData `json:"MAIN.shm_records"`
	MainShmWrites                  VSData `json:"MAIN.shm_writes"`
	MainShmFlushes                 VSData `json:"MAIN.shm_flushes"`
	MainShmCont                    VSData `json:"MAIN.shm_cont"`
	MainShmCycles                  VSData `json:"MAIN.shm_cycles"`
	MainSmsNReq                    VSData `json:"MAIN.sms_nreq"`
	MainSmsNObj                    VSData `json:"MAIN.sms_nobj"`
	MainSmsNBytes                  VSData `json:"MAIN.sms_nbytes"`
	MainSmsBAlloc                  VSData `json:"MAIN.sms_balloc"`
	MainSmsBFree                   VSData `json:"MAIN.sms_bfree"`
	MainBackendReq                 VSData `json:"MAIN.backend_req"`
	MainNVcl                       VSData `json:"MAIN.n_vcl"`
	MainNVclAvail                  VSData `json:"MAIN.n_vcl_avail"`
	MainNVclDiscard                VSData `json:"MAIN.n_vcl_discard"`
	MainBans                       VSData `json:"MAIN.bans"`
	MainBansCompleted              VSData `json:"MAIN.bans_completed"`
	MainBansObj                    VSData `json:"MAIN.bans_obj"`
	MainBansReq                    VSData `json:"MAIN.bans_req"`
	MainBansAdded                  VSData `json:"MAIN.bans_added"`
	MainBansDeleted                VSData `json:"MAIN.bans_deleted"`
	MainBansTested                 VSData `json:"MAIN.bans_tested"`
	MainBansObjKilled              VSData `json:"MAIN.bans_obj_killed"`
	MainBansLurkerTested           VSData `json:"MAIN.bans_lurker_tested"`
	MainBansTestsTested            VSData `json:"MAIN.bans_tests_tested"`
	MainBansLurkerTestsTested      VSData `json:"MAIN.bans_lurker_tests_tested"`
	MainBansLurkerObjKilled        VSData `json:"MAIN.bans_lurker_obj_killed"`
	MainBansDups                   VSData `json:"MAIN.bans_dups"`
	MainBansLurkerContention       VSData `json:"MAIN.bans_lurker_contention"`
	MainBansPersistedBytes         VSData `json:"MAIN.bans_persisted_bytes"`
	MainBansPersistedFragmentation VSData `json:"MAIN.bans_persisted_fragmentation"`
	MainNPurges                    VSData `json:"MAIN.n_purges"`
	MainNObjPurged                 VSData `json:"MAIN.n_obj_purged"`
	MainExpMailed                  VSData `json:"MAIN.exp_mailed"`
	MainExpReceived                VSData `json:"MAIN.exp_received"`
	MainHcbNoLock                  VSData `json:"MAIN.hcb_nolock"`
	MainHcbLock                    VSData `json:"MAIN.hcb_lock"`
	MainHcbInsert                  VSData `json:"MAIN.hcb_insert"`
	MainEsiErrors                  VSData `json:"MAIN.esi_errors"`
	MainEsiWarnings                VSData `json:"MAIN.esi_warnings"`
	MainVMods                      VSData `json:"MAIN.vmods"`
	MainNGzip                      VSData `json:"MAIN.n_gzip"`
	MainNGunzip                    VSData `json:"MAIN.n_gunzip"`
	MainVsmFree                    VSData `json:"MAIN.vsm_free"`
	MainVsmUsed                    VSData `json:"MAIN.vsm_used"`
	MainVsmCooling                 VSData `json:"MAIN.vsm_cooling"`
	MainVsmOverflow                VSData `json:"MAIN.vsm_overflow"`
	MainVsmOverflowed              VSData `json:"MAIN.vsm_overflowed"`
}

func (s Stat) String() string {
	return fmt.Sprintf("Timestamp: %v\nUp Time: %v", s.Timestamp, s.MainUpTime.Value)
}

var StatData Stat

func main() {
	flag.Parse()

	if showVersion {
		fmt.Println("Version: " + version)
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

	if err := json.Unmarshal(out.Bytes(), &StatData); err != nil {
		panic(err)
	}

	log.Println(StatData)
}
