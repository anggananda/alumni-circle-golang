package models

import "time"

type ListEvent struct {
	IDListevent int64     `json:"id_listevent" form:"id_listevent"`
	IDAlumni    int64     `json:"id_alumni" form:"id_alumni"`
	IDEvent     int64     `json:"id_event" form:"id_event"`
	Tanggal     time.Time `json:"tanggal" form:"tanggal"`
}
