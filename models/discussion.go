package models

import "time"

type Diskusi struct {
	IDDiskusi     int64     `json:"id_duskusi" form:"id_diskusi"`
	IDAlumni      int64     `json:"id_alumni" form:"id_alumni"`
	SubjekDiskusi string    `json:"subjek_diskusi" form:"subjek_diskusi"`
	IsiDiskusi    string    `json:"isi_diskusi" form:"isi_diskusi"`
	Tanggal       time.Time `json:"tanggal" form:"tanggal"`
}

type ListDiskusi struct {
	IDDiskusi     int64     `json:"id_duskusi" form:"id_diskusi"`
	IDAlumni      int64     `json:"id_alumni" form:"id_alumni"`
	SubjekDiskusi string    `json:"subjek_diskusi" form:"subjek_diskusi"`
	IsiDiskusi    string    `json:"isi_diskusi" form:"isi_diskusi"`
	Tanggal       time.Time `json:"tanggal" form:"tanggal"`
	FotoProfile   string    `json:"foto_profile" form:"foto_profile"`
	NamaAlumni    string    `json:"nama_alumni" form:"nama_alumni"`
	Email         string    `json:"email" form:"email"`
}

func (Diskusi) TableName() string {
	return "diskusi"
}
