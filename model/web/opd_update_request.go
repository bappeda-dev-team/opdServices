package web

// @Description Update Request OPD
type OpdUpdateRequest struct {
	Id            int    `json:"id" `
	NamaOpd       string `json:"nama_opd" validate:"required"`
	KodeOpd       string `json:"kode_opd" validate:"required,numeric,len=18"`
	Singkatan     string `json:"singkatan" validate:"required"`
	Alamat        string `json:"alamat" validate:"required"`
	Telepon       string `json:"telepon"`
	Fax           string `json:"fax"`
	Email         string `json:"email"`
	Website       string `json:"website" `
	NipKepalaOpd  string `json:"nip_kepala_opd" validate:"required"`
	NamaKepalaOpd string `json:"nama_kepala_opd" validate:"required"`
	PangkatKepala string `json:"pangkat_kepala"`
}
