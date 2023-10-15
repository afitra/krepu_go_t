package models

type User struct {
	ID           int    `db:"id" json:"-"`
	Nik          string `db:"nik" json:"nik"`
	Role         string `db:"role" json:"role"`
	UserName     string `db:"user_name" json:"user_name"`
	Password     string `db:"password" json:"password"`
	FullName     string `db:"full_name" json:"full_name"`
	LegalName    string `db:"legal_name" json:"legal_name"`
	TempatLahir  string `db:"tempat_lahir" json:"tempat_lahir"`
	TanggalLahir string `db:"tanggal_lahir" json:"tanggal_lahir"`
	Gaji         int    `db:"gaji" json:"gaji"`
	FotoKTP      string `db:"foto_ktp" json:"foto_ktp"`
	FotoSelfie   string `db:"foto_selfie" json:"foto_selfie"`
	TenorSatu    int    `db:"tenor_satu" json:"tenor_satu"`
	TenorDua     int    `db:"tenor_dua" json:"tenor_dua"`
	TenorTiga    int    `db:"tenor_tiga" json:"tenor_tiga"`
	TenorEmpat   int    `db:"tenor_empat" json:"tenor_empat"`
}

type UserPayload struct {
	Nik          string `json:"nik" validate:"required"`
	UserName     string `json:"user_name" validate:"required,min=6"`
	Password     string `json:"password" validate:"required,min=6"`
	FullName     string `json:"full_name" validate:"required"`
	LegalName    string `json:"legal_name" validate:"required"`
	TempatLahir  string `json:"tempat_lahir" validate:"required"`
	TanggalLahir string `json:"tanggal_lahir" validate:"required"`
	Gaji         int    `json:"gaji" validate:"required"`
	FotoKTP      string `json:"foto_ktp" validate:"required"`
	FotoSelfie   string `json:"foto_selfie" validate:"required"`
	TenorSatu    int    `json:"tenor_satu"`
	TenorDua     int    `json:"tenor_dua"`
	TenorTiga    int    `json:"tenor_tiga"`
	TenorEmpat   int    `json:"tenor_empat"`
}

type LoginPayload struct {
	UserName string `json:"user_name" validate:"required,min=6"`
	Password string `json:"password" validate:"required,min=6"`
}
