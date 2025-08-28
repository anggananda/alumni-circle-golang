package models

import "time"

type Event struct {
	IDEvent      int64     `json:"id_event" form:"id_event"`
	IDKategori   int64     `json:"id_kategori" form:"id_kategori"`
	NamaEvent    string    `json:"nama_event" form:"nama_event"`
	TanggalEvent time.Time `json:"tanggal_event" form:"tanggal_event"`
	Lokasi       string    `json:"lokasi" form:"lokasi"`
	Deskripsi    string    `json:"deskripsi" form:"deskripsi"`
	Gambar       string    `json:"gambar" form:"gambar"`
	Latitude     float64   `json:"latitude" form:"latitude"`
	Longitude    float64   `json:"longidute" form:"longitude"`
}

type EventWithCategory struct {
	IDEvent      int64     `json:"id_event" form:"id_event"`
	IDKategori   int64     `json:"id_kategori" form:"id_kategori"`
	NamaEvent    string    `json:"nama_event" form:"nama_event"`
	TanggalEvent time.Time `json:"tanggal_event" form:"tanggal_event"`
	Lokasi       string    `json:"lokasi" form:"lokasi"`
	Deskripsi    string    `json:"deskripsi" form:"deskripsi"`
	Gambar       string    `json:"gambar" form:"gambar"`
	Latitude     float64   `json:"latitude" form:"latitude"`
	Longitude    float64   `json:"longidute" form:"longitude"`
	Kategori     string    `json:"kategori" form:"kategori"`
}

func (Event) TableName() string {
	return "event"
}
