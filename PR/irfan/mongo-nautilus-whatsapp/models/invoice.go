package models

type ContentInvoice struct {
	Uuid_contents string  `json:"uuid", bson:"uuid"`
	Product       Product `json:"product" bson:"product"`
	Price         int64   `json:"price" bson:"price"`
	Count         int64   `json:"count" bson:"count"`
}

type Invoices struct {
	UUID       string           `json:"uuid,omitempty" bson:"uuid"`
	Alias      string           `json:"alias,omitempty" bson:"alias"`
	TotalPrice int64            `json:"total_price,omitempty" bson:"total_price"`
	TotalCount int64            `json:"total_count,omitempty" bson:"total_count"`
	Contents   []ContentInvoice `json:"contents,omitempty" bson:"contents"`
}

var DataInvoices []*Invoices
