package models

import "time"

type ListVacancy struct {
	IDListvc  int64     `json:"id_listvc" form:"id_listvc"`
	IDALumni  int64     `json:"id_alumni" form:"id_alumni"`
	IDVacancy int64     `json:"id_vacancy" form:"id_vacancy"`
	Tanggal   time.Time `json:"tanggal" form:"tanggal"`
}
