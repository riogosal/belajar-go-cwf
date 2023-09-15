package models

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Packings struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	UUID         string             `bson:"uuid,omitempty" json:"uuid"`
	PreviousUUID string             `bson:"previous_uuid,omitempty" json:"previous_uuid"`
	Customer     Customer           `bson:"customer,omitempty" json:"customer"`
	Product      Product            `bson:"product,omitempty" json:"product"`
	Box          string             `bson:"box,omitempty" json:"box"`
	Created_At   int64              `bson:"created_at,omitempty" json:"created_at"`
	TotalWeight  float64            `bson:"total_weight,omitempty" json:"total_weight"`
	TotalCount   int8               `bson:"total_count,omitempty" json:"total_count"`
	Contents     []ContentsPacking  `bson:"contents,omitempty" json:"contents"`
	ContentIlcs  []string           `bson:"content_ilcs,omitempty" json:"content_ilcs,omitempty"`
}

type ContentsPacking struct {
	UUID           string  `bson:"uuid" json:"uuid"`
	ProductLogUUID string  `bson:"product_log_uuid" json:"product_log_uuid"`
	ILC            string  `bson:"ilc" json:"ilc"`
	Category       string  `bson:"category" json:"category"`
	Species        Species `bson:"species" json:"species"`
	Weight         float64 `bson:"weight" json:"weight"`
	Count          int     `bson:"count" json:"count"`
	Created_At     int     `bson:"created_at" json:created_at"`
}

type Customer struct {
	UUID  string `bson:"uuid,omitempty" json:"uuid"`
	Name  string `bson:"name,omitempty" json:"name"`
	Code  string `bson:"code,omitempty" json:"code"`
	Unit  string `bson:"unit,omitempty" json:"unit"`
	Alias string `bson:"alias,omitempty" json:"alias"`
}

type TotalPackingByDate struct {
	Count  int8
	Weight float64
	Box    int
}

var DataPackings []*Packings

func PrintDataPackings(dataPackings []*Packings) {
	dataPackingsdate := map[string]TotalPackingByDate{}
	for _, v := range dataPackings {
		if dataPackingsdate[v.Customer.Name].Count == 0 {
			dataPackingsdate[v.Customer.Name] = TotalPackingByDate{
				Count:  v.TotalCount,
				Weight: v.TotalWeight,
				Box:    1,
			}
		} else {
			// x := dataPackingsdate[v.Customer.Name]
			// x.Count += v.TotalCount
			// x.Weight += v.TotalWeight
			// x.Box += 1
			dataPackingsdate[v.Customer.Name] = TotalPackingByDate{
				Count:  dataPackingsdate[v.Customer.Name].Count + v.TotalCount,
				Weight: dataPackingsdate[v.Customer.Name].Weight + v.TotalWeight,
				Box:    dataPackingsdate[v.Customer.Name].Box + 1,
			}
		}
	}

	for i, val := range dataPackingsdate {
		fmt.Printf(`
--------------------------------
%s
--------------------------------
%d Pcs | %2.2f Kg
Box : %d `, i, val.Count, val.Weight, val.Box)
	}
	fmt.Printf(`
Total Mc : %d`, len(dataPackings))

}
