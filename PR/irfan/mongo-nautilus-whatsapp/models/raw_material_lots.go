package models

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Species struct {
	Name            string `json:"name" bson:"name"`
	Scientific_name string `json:"scientific_name" bson:"scientific_name"`
	Code            string `json:"code" bson:"code"`
	LoinsPerFish    int    `json:"loins_per_fish" bson:"loins_per_fish"`
}

type Type struct {
	Name     string `json:"name" bson:"name"`
	Code     string `json:"code" bson:"code"`
	Whole    bool   `json:"whole" bson:"whole"`
	Gradable bool   `json:"gradable" bson:"gradable"`
	Multiple bool   `json:"multiple" bson:"multiple"`
}

type Supplier struct {
	UUID      string `json:"uuid" bson:"uuid"`
	Name      string `json:"name" bson:"name"`
	Code      string `json:"code" bson:"code"`
	BatchCode string `json:"batch_code" bson:"batch_code"`
}

type Contents struct {
	UUID         string  `json:"uuid" bson:"uuid"`
	Number       int     `json:"number" bson:"number"`
	Grade        string  `json:"grade" bson:"grade"`
	Count        int     `json:"count" bson:"count"`
	Weight       float64 `json:"weight" bson:"weight"`
	Conteiner_Id int     `json:"conteiner_id" bson:"conteiner_id"`
	Date         string  `json:"date" bson:"date"`
	Created_At   int64   `json:"created_at" bson:"created_at"`
	Updated_At   int     `json:"updated_at" bson:"updated_at"`
}

type GroupedContents struct {
	Weight float64 `json:"weight" bson:"weight"`
	Count  int     `json:"count" bson:"count"`
}

type RawMaterials struct {
	ID               primitive.ObjectID         `bson:"_id,omitempty" json:"-"`
	ILC              string                     `json:"ilc" bson:"ilc"`
	Delivery_Number  string                     `json:"delivery_number" bson:"delivery_number"`
	Species          Species                    `json:"species" bson:"species"`
	Type             Type                       `json:"type" bson:"type"`
	Supplier         Supplier                   `json:"supplier" bson:"supplier"`
	Total_Weight     float64                    `json:"total_weight" bson:"total_weight"`
	Total_count      int                        `json:"total_count" bson:"total_count"`
	Created_At       int                        `json:"created_at" bson:"created_at"`
	Updated_At       int                        `json:"updated_at" bson:"updated_at"`
	Contents         []Contents                 `json:"contents" bson:"contents"`
	Grouped_Contents map[string]GroupedContents `json:"grouped_contents" bson:"grouped_contents"`
}

type Certificates struct {
	UUID       string `json:"uuid" bson:"uuid"`
	Code       string `json:"code" bson:"code"`
	Created_at int    `json:"created_at" bson:"created_at"`
	Disabled   bool   `json:"disabled" bson:"disabled"`
	Name       string `json:"name" bson:"name"`
	Number     string `json:"number" bson:"number"`
	Updated_at int    `json:"updated_at" bson:"updated_at"`
}

var DataRawMaterials []*RawMaterials

func DataCountsLoins() {

}

func (r *RawMaterials) CetakStrRwaMatetirialLot() string {
	DataGradeWeightAndCounts := ""

	for l, v := range r.Grouped_Contents {
		DataGradeWeightAndCounts += fmt.Sprintf("%s %d Ekor | %2.f2 Kg\n", l, v.Count, v.Weight)
	}

	if r.Type.Code != "GG" {
		return fmt.Sprintf(`
-------------------------
   %s
   %s (%s,%s)
-------------------------
Tot %d Loin | %2.2f Kg
`, r.ILC, r.Supplier.Name, r.Supplier.Code, r.Type.Code, r.Total_count, r.Total_Weight)
	} else {
		return fmt.Sprintf(`
-------------------------
   %s
   %s (%s, %s)
-------------------------
%s
Tot %d Ekor | %2.2f Kg
`, r.ILC, r.Supplier.Name, r.Species.Code, r.Type.Code, DataGradeWeightAndCounts, r.Total_count, r.Total_Weight)
	}

}
