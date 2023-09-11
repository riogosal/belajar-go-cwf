package models

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Costumer struct {
	UUID      string  `bson:"uuid" json:"uuid"`
	Name      string  `bson:"name" json:"name"`
	Code      string  `bson:"code" json:"code"`
	Unit      string  `bson:"unit" json:"unit"`
	Alias     string  `bson:"alias" json:"alias"`
	Phone     Phone   `bson:"phone" json:"phone"`
	Address   Address `bson:"address" json:"address"`
	LicenseId string  `bson:"license_id" json:"license_id"`
}

type Phone struct {
	CountryCode string `bson:"country_code" json:"country_code"`
	Number      string `bson:"number" json:"number"`
}

type Address struct {
	Country  string `bson:"country" json:"country"`
	State    string `bson:"state" json:"state"`
	District string `bson:"district" json:"district"`
	Street   string `bson:"street" json:"street"`
}

type GroupedContentsRefinded struct {
	Count  int     `bson:"count" json:"count"`
	Weight float64 `bson:"weight" json:"weight"`
}

type RefinedMaterialLots struct {
	ID              primitive.ObjectID                 `bson:"_id,omitempty", json:"_id"`
	ILC             string                             `bson:"ilc,omitempty" json:"ilc"`
	RawIlc          string                             `bson:"raw_ilc,omitempty" json:"raw_ilc"`
	Costumer_Group  string                             `bson:"customer_group,omitempty" json:"customer_group"`
	Species         Species                            `bson:"species,omitempty" json:"species"`
	Type            Type                               `bson:"type,omitempty" json:"type"`
	Supplier        Supplier                           `bson:"supplier,omitempty" json:"supplier"`
	ReceivingDate   int                                `bson:"receiving_date,omitempty" json:"receiving_date`
	Created_At      int                                `bson:"created_at,omitempty" json:"Created_at"`
	Updated_At      int                                `bson:"updated_date,omitempty" json:"Updated_date"`
	GroupedContents map[string]GroupedContentsRefinded `bson:"grouped_contents,omitempty" json:"grouped_contents"`
	Total_Weight    float64                            `bson:"total_weight,omitempty" json:"total_weight"`
	Total_Count     int                                `bson:"total_count,omitempty" json:"total_count"`
}

var DataRefinedMaterials []*RefinedMaterialLots

func (r *RefinedMaterialLots) PrintDataRefinedMaterial() string {
	dataRefinedMaterials := ""

	for i, v := range r.GroupedContents {
		dataRefinedMaterials += fmt.Sprintf(`
%s %d Ekor | %2.2f Kg`, i, v.Count, v.Weight)
	}

	return fmt.Sprintf(`
--------------------------------
	%s
--------------------------------
%s`, r.Costumer_Group, dataRefinedMaterials)
}
