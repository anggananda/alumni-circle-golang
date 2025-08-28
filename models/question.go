package models

import "time"

type Question struct {
	IDPertanyaan  int64     `json:"id_pertanyaan" form:"id_pertanyaan"`
	IDAlumni      int64     `json:"id_alumni" form:"id_alumni"`
	IsiPertanyaan string    `json:"isi_pertanyaan" form:"isi_pertanyaan"`
	Tanggal       time.Time `json:"tanggal" form:"tanggal"`
}
