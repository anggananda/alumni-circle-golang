package models

import "time"

type Feedback struct {
	IDFeedback  int64     `json:"id_feedback" form:"id_feedback"`
	IDAlumni    int64     `json:"id_alumni" form:"id_alumni"`
	IsiFeedback string    `json:"isi_feedback" form:"isi_feedback"`
	Tanggal     time.Time `json:"tanggal" form:"tanggal"`
}
