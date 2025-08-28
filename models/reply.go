package models

import "time"

type Reply struct {
	IDReply   int64 `json:"id_reply" form:"id_reply"`
	IDAlumni  int64 `json:"id_alumni" form:"id_alumni"`
	IDDiskusi int64 `json:"id_diskusi" form:"id_diskusi"`
  IsiReply string `json:"isi_reply" form:"isi_reply"`
  Tanggal time.Time `json:"tanggal" form:"tanggal"`
}
