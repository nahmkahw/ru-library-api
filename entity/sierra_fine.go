package entity

type Fine struct {
	Id            string `json:"Id" validate:"required"`
	Barcode       string `json:"Barcode" validate:"required"`
	Amount        string `json:"Amount"`
	InvoiceNum    string `json:"InvoiceNum"`
	ItemCharge    string `json:"ItemCharge"`
	ProcessingFee string `json:"ProcessingFee"`
	BillingFee    string `json:"BillingFee"`
	ChargeCode    string `json:"ChargeCode"`
	PaidAmount    string `json:"PaidAmount"`
	ItemId        string `json:"ItemId"`
	LocationCode  string `json:"LocationCode"`
	Title         string `json:"Title"`
}

type FineRepo struct {
	Id            string `db:"id" validate:"required"`
	Barcode       string `db:"barcode" validate:"required"`
	Amount        string `db:"owed_amt"`
	InvoiceNum    string `db:"invoice_num"`
	ItemCharge    string `db:"item_charge_amt"`
	ProcessingFee string `db:"processing_fee_amt"`
	BillingFee    string `db:"billing_fee_amt"`
	ChargeCode    string `db:"charge_code"`
	PaidAmount    string `db:"paid_amt"`
	ItemId        string `db:"item_id"`
	LocationCode  string `db:"location_code"`
	Title         string `db:"title"`
}
