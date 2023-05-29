package entity

type PatronRepo struct {
	Id          string `db:"id" validate:"required"`
	Barcode     string `db:"barcode" validate:"required"`
	Fullname    string `db:"last_name"`
	PtypeCode   string `db:"ptype_code"`
	Description string `db:"description"`
}

type Patron struct {
	Id          string `json:"Id" validate:"required"`
	Barcode     string `json:"Barcode" validate:"required"`
	Fullname    string `json:"Fullname"`
	PtypeCode   string `json:"PtypeCode"`
	Description string `json:"Description"`
}
