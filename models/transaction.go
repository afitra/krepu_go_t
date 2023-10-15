package models

type Transaction struct {
	ID        int    `db:"id" json:"id"`
	UserId    int    `db:"user_id" json:"user_id"`
	NoKontrak string `db:"no_kontrak" json:"no_kontrak"`
	OTR       int    `db:"otr" json:"otr"`
	AdminFee  int    `db:"admin_fee" json:"admin_fee"`
	Cicilan   int    `db:"cicilan" json:"cicilan"`
	Bunga     int    `db:"bunga" json:"bunga"`
	NamaAsset string `db:"nama_asset" json:"nama_asset"`
	Status    string `db:"status" json:"status"`
	Tenor     int    `db:"tenor" json:"tenor"`
	Pengajuan int    `db:"pengajuan" json:"pengajuan"`
}

type PayloadInquiry struct {
	Otr       int    `json:"otr" validate:"required"`
	AdminFee  int    `json:"admin_fee" validate:"required"`
	Cicilan   int    `json:"cicilan" validate:"required"`
	Bunga     int    `json:"bunga" validate:"required"`
	NamaAsset string `json:"nama_asset" validate:"required"`
	Tenor     int    `json:"tenor"  validate:"required"`
	Pengajuan int    `json:"pengajuan"  validate:"required"`
}

type PayloadPay struct {
	NoKontrak string `json:"no_kontrak"`
}
