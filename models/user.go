package models

type User struct {
	ID           int    `db:"id" json:"id"`
	Nik          string `db:"nik" json:"nik"`
	FullName     string `db:"full_name" json:"full_name"`
	LegalName    string `db:"legal_ame" json:"legal_ame"`
	TempatLahir  string `db:"tempat_lahir" json:"tempat_lahir"`
	TanggalLahir string `db:"tanggal_lahir" json:"tanggal_lahir"`
	Gaji         int    `db:"gaji" json:"gaji"`
	FotoKTP      string `db:"foto_ktp" json:"foto_ktp"`
	FotoSelfie   string `db:"foto_selfie" json:"foto_selfie"`
}

type UserPayload struct {
	Nik          string `db:"nik" json:"nik" validate:"required"`
	FullName     string `db:"full_name" json:"full_name" validate:"required"`
	LegalName    string `db:"legal_name" json:"legal_ame" validate:"required"`
	TempatLahir  string `db:"tempat_lahir" json:"tempat_lahir" validate:"required"`
	TanggalLahir string `db:"tanggal_lahir" json:"tanggal_lahir" validate:"required"`
	Gaji         int    `db:"gaji" json:"gaji" validate:"required"`
	FotoKTP      string `db:"foto_ktp" json:"foto_ktp" validate:"required"`
	FotoSelfie   string `db:"foto_selfie" json:"foto_selfie" validate:"required"`
}
