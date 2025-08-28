package models

import "time"

type Alumni struct {
	IDAlumni        int64     `json:"id_alumni" form:"id_alumni"`
	NamaAlumni      string    `json:"nama_alumni" form:"nama_alumni"`
	JenisKelamin    string    `json:"jenis_kelamin" form:"jenis_kelamin"`
	Alamat          string    `json:"alamat" form:"alamat"`
	Email           string    `json:"email" form:"email"`
	TanggalLulus    time.Time `json:"tanggal_lulus" form:"tanggal_lulus"`
	Angkatan        string    `json:"angkatan" form:"angkatan"`
	StatusPekerjaan string    `json:"status_pekerjaan" form:"status_pekerjaan"`
	Username        string    `json:"username" form:"username"`
	Password        string    `json:"password" form:"password"`
	Roles           string    `json:"roles" form:"roles"`
	FotoProfile     string    `json:"foto_profile" form:"foto_profile"`
}

func (Alumni) TableName() string {
	return "alumni" // sesuaikan dengan table MySQL
}
