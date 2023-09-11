package models

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Grade struct {
	Code   string  `bson:"code" json:"code"`
	Count  int     `bson:"count" json:"count"`
	Name   string  `bson:"name" json:"name"`
	Weight float64 `bson:"weight" json:"weight"`
}

type DataGrades struct {
	Counts int     `bson:"counts" json:"counts"`
	Weight float64 `bson:"weight" json:"weight"`
	Name   string  `bson:"name" json:"name"`
}

type TotalCountAndWeight struct {
	Count  int64   `bson:"count" json:"count"`
	Weight float64 `bson:"weight" json:"weight"`
}

type Product struct {
	UUID     string  `bson:"uuid" json:"uuid"`
	Code     string  `bson:"code" json:"code"`
	Name     string  `bson:"name" json:"name"`
	Category string  `bson:"category" json:"category"`
	Species  Species `bson:"species" json:"species"`
}

type Content struct {
	UUID       string  `bson:"uuid" json:"uuid"`
	Number     int     `bson:"number,omitempty" json:"number"`
	Weight     float64 `bson:"weight,omitempty" json:"weight"`
	Count      int64   `bson:"count,omitempty" json:"count"`
	Grade      string  `bson:"grade,omitempty" json:"grade"`
	Product    Product `bson:"product,omitempty" json:"product"`
	Date       string  `bson:"date" json:"date"`
	CreatedAt  int64   `bson:"created_at" json:"created_at"`
	CountUsed  int     `bson:"count_used" json:"count_used"`
	WeightUsed float64 `bson:"weight_used" json:"weight_used"`
}

type ProductLogContainers struct {
	ID               primitive.ObjectID          `bson:"_id,omitempty" json:"-"`
	RefinedIlc       string                      `bson:"refined_ilc,omitempty" json:"refined_ilc"`
	ReceivingDate    int                         `bson:"receiving_date,omitempty" json:"receiving_date"`
	Cuttingdate      int                         `bson:"cutting_date,omitempty" json:"cutting_date"`
	Supplier         Supplier                    `bson:"supplier,omitempty" json:"supplier"`
	CustomerGroup    string                      `bson:"customer_group,omitempty" json:"customer_group"`
	Total_Weight     float64                     `bson:"total_weight,omitempty" json:"total_weight"`
	Total_count      int                         `bson:"total_count,omitempty" json:"total_count"`
	RawTotalCount    int                         `bson:"raw_total_count,omitempty" json:"raw_total_count"`
	RawTotalWeight   float32                     `bson:"raw_total_weight,omiteempty" json:"raw_total_weight"`
	Created_At       int64                       `bson:"created_at,omitempty" json:"created_at"`
	Updated_At       int64                       `bson:"updated_at,omitempty" json:"updated_at"`
	Contents         []Content                   `bson:"contents" json:"contents"`
	Grouped_Contents map[string]map[string]Grade `bson:"grouped_contents" json:"grouped_contents"`
}

var DataProducts []*ProductLogContainers

func (p *ProductLogContainers) PrintProductLogContainer() {
	var dataGrade = map[string]map[string]DataGrades{}

	for _, v := range p.Contents {
		if dataGrade[v.Product.Name] == nil {
			dataGrade[v.Product.Name] = make(map[string]DataGrades)

		}
		if dataGrade[v.Product.Name][v.Grade].Counts == 0 {

			dataGrade[v.Product.Name][v.Grade] = DataGrades{}
		}

		x := dataGrade[v.Product.Name][v.Grade]
		x.Counts += int(v.Count)
		x.Weight += v.Weight

		dataGrade[v.Product.Name][v.Grade] = x

	}

	fmt.Println(dataGrade)

}

func DataProductsdate(productDate []*ProductLogContainers) {
	dataProducts := map[string]map[string]map[string]DataGrades{}
	dataTotals := map[string]TotalCountAndWeight{}

	for _, p := range productDate {
		if dataProducts[p.CustomerGroup] == nil && dataTotals[p.CustomerGroup].Count == 0 {
			dataProducts[p.CustomerGroup] = make(map[string]map[string]DataGrades)
			dataTotals[p.CustomerGroup] = TotalCountAndWeight{
				Count:  int64(p.Total_count),
				Weight: p.Total_Weight,
			}
		} else {
			dataTotals[p.CustomerGroup] = TotalCountAndWeight{
				Count:  dataTotals[p.CustomerGroup].Count + int64(p.Total_count),
				Weight: dataTotals[p.CustomerGroup].Weight + p.Total_Weight,
			}
		}
		for _, v := range p.Contents {
			if dataProducts[p.CustomerGroup][v.Product.Name] == nil {
				dataProducts[p.CustomerGroup][v.Product.Name] = make(map[string]DataGrades)
				if dataProducts[p.CustomerGroup][v.Product.Name][v.Grade].Counts == 0 {
					dataProducts[p.CustomerGroup][v.Product.Name][v.Grade] = DataGrades{}
				}

			}

			x := dataProducts[p.CustomerGroup][v.Product.Name][v.Grade]
			x.Counts += int(v.Count)
			x.Weight += v.Weight
			// x.Name = v.Product.Name

			dataProducts[p.CustomerGroup][v.Product.Name][v.Grade] = x
		}

	}

	for i, _ := range dataProducts {
		fmt.Printf(`
-------------------------
	%s
-------------------------`, i)
		for j, v := range dataProducts[i] {
			fmt.Printf(`
%s`, j)
			for k, val := range v {
				if k == "" {
					fmt.Printf(`
Ungraded %d Pcs | %2.2f Kg
					`, val.Counts, val.Weight)
				} else {
					fmt.Printf(`
%s %d Pcs | %2.2f Kg
					`, k, val.Counts, val.Weight)
				}
			}

		}
		fmt.Printf(`
Total: %d Pcs | %2.2f Kg
		`, dataTotals[i].Count, dataTotals[i].Weight)
	}
	// fmt.Println(dataTotals)
}
