package structs

import "fmt"

type VNumbers struct {
	Timestamp     string `json:"timestamp"`
	Object        VData  `json:"MAIN.n_object"`
	VampireObject VData  `json:"MAIN.n_vampireobject"`
	ObjectCore    VData  `json:"MAIN.n_objectcore"`
	ObjectHead    VData  `json:"MAIN.n_objecthead"`
	WaitingList   VData  `json:"MAIN.n_waitinglist"`
	Backend       VData  `json:"MAIN.n_backend"`
	Expired       VData  `json:"MAIN.n_expired"`
	LruNuked      VData  `json:"MAIN.n_lru_nuked"`
	LruMoved      VData  `json:"MAIN.n_lru_moved"`
	Vcl           VData  `json:"MAIN.n_vcl"`
	VclAvail      VData  `json:"MAIN.n_vcl_avail"`
	VclDiscard    VData  `json:"MAIN.n_vcl_discard"`
	Purges        VData  `json:"MAIN.n_purges"`
	ObjPurged     VData  `json:"MAIN.n_obj_purged"`
	Gzip          VData  `json:"MAIN.n_gzip"`
	Gunzip        VData  `json:"MAIN.n_gunzip"`
}

func (s VNumbers) String() string {
	return fmt.Sprintf("Timestamp: %v - Number of Objects: %v", s.Timestamp, s.Object.Value)
}
