package models

import (
	"errors"
)

var (
	ErrInternalServerError = errors.New("internal Server Error")

	ErrNotFound = errors.New("data tidak ditemukan")

	ErrConflict = errors.New("item sudah ada")

	ErrUnauthorized = errors.New("unauthorized")

	ErrPassword = errors.New("username atau Password yang digunakan tidak valid")

	ErrSomethingWrong = errors.New("Terjadi kesalahan")

	ErrorLimitInquiry = errors.New("Pengajuan melebihi limit anda")
)
