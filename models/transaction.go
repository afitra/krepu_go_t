package models

type Transaction struct {
	ID        int    `db:"id" json:"id"`
	NoKontrak string `db:"no_kontrak" json:"no_kontrak"`
	OTR       int    `db:"otr" json:"otr"`
	AdminFee  int    `db:"admin_fee" json:"admin_fee"`
	Cicilan   int    `db:"cicilan" json:"cicilan"`
	Bunga     int    `db:"bunga" json:"bunga"`
	NamaAsset string `db:"nama_asset" json:"nama_asset"`
}
