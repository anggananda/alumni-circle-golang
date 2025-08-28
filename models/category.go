package models

type Kategori struct {
	IDKategori int64  `json:"id_kategori" form:"id_kategori"`
	Kategori   string `json:"kategori" form:"kategori"`
	Deskripsi  string `json:"deskripsi" form:"deskripsi"`
	Gambar     string `json:"gambar" form:"gambar"`
}
