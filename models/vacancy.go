package models

type Vacancy struct {
	IDVacancy   int64  `json:"id_vacancy" form:"id_vacancy"`
	NamaVacancy string `json:"nama_vacancy" form:"nama_vacancy"`
	Deskripsi   string `json:"deskripsi" form:"deskripsi"`
	Gambar      string `json:"gambar" form:"gambar"`
}
