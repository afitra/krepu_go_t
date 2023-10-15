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
	Tenor        int    `db:"tenor" json:"tenor"`
}

type UserPayload struct {
	Nik          string `db:"nik" json:"nik" validate:"required"`
	UserName     string `db:"user_name" json:"user_name" validate:"required,min=6"`
	Password     string `db:"password" json:"password" validate:"required,min=6"`
	FullName     string `db:"full_name" json:"full_name" validate:"required"`
	LegalName    string `db:"legal_name" json:"legal_name" validate:"required"`
	TempatLahir  string `db:"tempat_lahir" json:"tempat_lahir" validate:"required"`
	TanggalLahir string `db:"tanggal_lahir" json:"tanggal_lahir" validate:"required"`
	Gaji         int    `db:"gaji" json:"gaji" validate:"required"`
	FotoKTP      string `db:"foto_ktp" json:"foto_ktp" validate:"required"`
	FotoSelfie   string `db:"foto_selfie" json:"foto_selfie" validate:"required"`
	Tenor        int    `db:"tenor" json:"tenor" `
}

type LoginPayload struct {
	UserName string `json:"user_name" validate:"required,min=6"`
	Password string `json:"password" validate:"required,min=6"`
}
