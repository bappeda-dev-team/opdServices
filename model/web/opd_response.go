package web

// @Description Response OPD
type OpdResponse struct {
	Id            string `json:"id"`
	KodeOpd       string `json:"kode_opd"`
	NamaOpd       string `json:"nama_opd"`
	Singkatan     string `json:"singkatan"`
	Alamat        string `json:"alamat"`
	Telepon       string `json:"telepon"`
	Fax           string `json:"fax"`
	Email         string `json:"email"`
	Website       string `json:"website"`
	NipKepalaOpd  string `json:"nip_kepala_opd"`
	NamaKepalaOpd string `json:"nama_kepala_opd"`
	PangkatKepala string `json:"pangkat_kepala"`
}

type OpdOnlyResponse struct {
	KodeOpd string `json:"kode_opd,omitempty"`
	NamaOpd string `json:"nama_opd,omitempty"`
}
