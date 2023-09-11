package models

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductSpecies struct {
	Name           string `json:"name" bson:"name"`
	ScientificName string `json:"scientific_name" bson:"scientific_name"`
	Code           string `json:"code" bson:"code"`
}

type Product struct {
	UUID     string         `json:"uuid" bson:"uuid"`
	Code     string         `json:"code" bson:"code"`
	Name     string         `json:"name" bson:"name"`
	Category string         `json:"category" bson:"category"`
	Species  ProductSpecies `json:"species" bson:"species"`
}

type ProductLogContent struct {
	UUID      string  `json:"uuid" bson:"uuid"`
	Number    int32   `json:"number" bson:"number"`
	Weight    float64 `json:"weight" bson:"weight"`
	Count     int64   `json:"count" bson:"count"`
	Grade     string  `json:"grade" bson:"grade"`
	Products  Product `json:"product" bson:"product"`
	Date      string  `json:"date" bson:"date"`
	CreatedAt int64   `json:"created_at" bson:"created_at"`
}

type ProductLogSupplier struct {
	UUID      string `json:"uuid" bson:"uuid"`
	Name      string `json:"name" bson:"name"`
	Code      string `json:"code" bson:"code"`
	BatchCode string `json:"batch_code" bson:"batch_code"`
}

type ProductLogSpecies struct {
	Name           string `json:"name" bson:"name"`
	ScientificName string `json:"scientific_name" bson:"scientific_name"`
	Code           string `json:"code" bson:"code"`
	LoinsPerFish   int    `json:"loins_per_fish" bson:"loins_per_fish"`
}

type ProductLogType struct {
	Name     string `json:"name" bson:"name"`
	Code     string `json:"code" bson:"code"`
	Whole    bool   `json:"whole" bson:"whole"`
	Gradable bool   `json:"gradable" bson:"gradable"`
	Multiple bool   `json:"multiple" bson:"multiple"`
}

type Ungraded struct {
	Weight float64 `json:"weight" bson:"weight"`
	Count  int64   `json:"count" bson:"count"`
	Name   string  `json:"name" bson:"name"`
	Code   string  `json:"code" bson:"code"`
}

type TypeGroupedContents struct {
	Code   string  `json:"code" bson:"code"`
	Count  int64   `json:"count" bson:"count"`
	Name   string  `json:"name" bson:"name"`
	Weight float64 `json:"weight" bson:"weight"`
}

type CodeGroupedContent struct {
	Types TypeGroupedContents `json:"types" bson:"types"`
}

type ProductLogContainer struct {
	ID              primitive.ObjectID  `json:"_id,omitempty" bson:"_id,omitempty"`
	RefinedILC      string              `json:"refined_ilc" bson:"refined_ilc"`
	ReceivingDate   int64               `json:"receiving_date" bson:"receiving_date"`
	CuttingDate     int64               `json:"cutting_date" bson:"cutting_date"`
	CustomerGroup   string              `json:"customer_group" bson:"customer_group"`
	Species         ProductLogSpecies   `json:"species" bson:"species"`
	Type            ProductLogType      `json:"type" bson:"type"`
	Supplier        ProductLogSupplier  `json:"supplier" bson:"supplier"`
	TotalWeight     float64             `json:"total_weight" bson:"total_weight"`
	TotalCount      int64               `json:"total_count" bson:"total_count"`
	RawTotalCount   int32               `json:"raw_total_count" bson:"raw_total_count"`
	RawTotalWeight  float64             `json:"raw_total_weight" bson:"raw_total_weight"`
	CreatedAt       int64               `json:"created_at" bson:"created_at"`
	UpdatedAt       int64               `json:"updated_at" bson:"updated_at"`
	Contents        []ProductLogContent `json:"contents" bson:"contents"`
	GroupedContents CodeGroupedContent  `json:"grouped_contents" bson:"grouped_contents"`
	TotalCountUsed  int64               `json:"total_count_used" bson:"total_count_used"`
	TotalWeightUsed float64             `json:"total_weight_used" bson:"total_weight_used"`
	UUID            string              `json:"uuid" bson:"uuid"`
}

type PlcWeightandCount struct {
	Grade        string  `json:"grade" bson:"grade"`
	CustomerGrup string  `json:"customer_group" bson:"customer_group"`
	ProductName  string  `json:"product_name" bson:"product_name"`
	Count        int64   `json:"count" bson:"count"`
	Weight       float64 `json:"weight" bson:"weight"`
	TotalCount   int64
	TotalWeight  float64
}

type TotalCountAndWeight struct {
	Count  int64
	Weight float64
}

// map[string]map[string]Ungraded

var ProductLogResult []*ProductLogContainer

func PrintProductLogContainerDataByDate(productDate []*ProductLogContainer) {
	dataProducts := map[string]map[string]map[string]Ungraded{}
	dataTotals := map[string]TotalCountAndWeight{}

	for _, p := range productDate {
		if dataProducts[p.CustomerGroup] == nil && dataTotals[p.CustomerGroup].Count == 0 {
			dataProducts[p.CustomerGroup] = make(map[string]map[string]Ungraded)
			dataTotals[p.CustomerGroup] = TotalCountAndWeight{
				Count:  int64(p.TotalCount),
				Weight: p.TotalWeight,
			}
		} else {
			dataTotals[p.CustomerGroup] = TotalCountAndWeight{
				Count:  dataTotals[p.CustomerGroup].Count + int64(p.TotalCount),
				Weight: dataTotals[p.CustomerGroup].Weight + p.TotalWeight,
			}
		}
		for _, v := range p.Contents {
			if dataProducts[p.CustomerGroup][v.Products.Name] == nil {
				dataProducts[p.CustomerGroup][v.Products.Name] = make(map[string]Ungraded)
				if dataProducts[p.CustomerGroup][v.Products.Name][v.Grade].Count == 0 {
					dataProducts[p.CustomerGroup][v.Products.Name][v.Grade] = Ungraded{}
				}

			}

			data := dataProducts[p.CustomerGroup][v.Products.Name][v.Grade]
			data.Count += v.Count
			data.Weight += v.Weight
			// data.Name = v.Product.Name

			dataProducts[p.CustomerGroup][v.Products.Name][v.Grade] = data
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
						`, val.Count, val.Weight)
				} else {
					fmt.Printf(`
%s %d Pcs | %2.2f Kg
						`, k, val.Count, val.Weight)
				}
			}

		}
		fmt.Printf(`
Total: %d Pcs | %2.2f Kg
			`, dataTotals[i].Count, dataTotals[i].Weight)
	}
}

// fmt.Println(dataTotals)
// for _, dataPlc := range resultPLC {

// 	if dataGrup[dataPlc.CustomerGroup] == nil {
// 		dataGrup[dataPlc.CustomerGroup] = make(map[string]PlcWeightandCount)
// 	}

// 	for _, content := range dataPlc.Contents {
// 		// Cek apakah produk sudah ada dalam dataGrup
// 		if _, ok := dataGrup[content.Products.Name]; !ok {
// 			dataGrup[content.Products.Name] = make(map[string]PlcWeightandCount)
// 		}
// 		// Cek apakah grade sudah ada dalam dataGrup
// 		if _, ok := dataGrup[content.Products.Name][content.Grade]; !ok {
// 			dataGrup[content.Products.Name][content.Grade] = PlcWeightandCount{
// 				Grade:        content.Grade,
// 				CustomerGrup: dataPlc.CustomerGroup,
// 				ProductName:  content.Products.Name,
// 				Count:        content.Count,
// 				Weight:       content.Weight,
// 			}
// 		} else {
// 			// Jika gradenya ada, + jumlah dan beratnya
// 			data := dataGrup[content.Products.Name][content.Grade]
// 			data.Count += content.Count
// 			data.Weight += content.Weight
// 			dataGrup[content.Products.Name][content.Grade] = data
// 		}
// 	}
// }

// for productName, gradeData := range dataGrup {
// 	//  apakah ada grade yang benar(gardenya tidak kosong)
// 	dataGrade := false
// 	for grade := range gradeData {
// 		if grade != "ProductName" && grade != "" {
// 			dataGrade = true
// 			break
// 		}
// 	}
// 	// Jika tidak ada grade yang valid, lewati produk ini
// 	if !dataGrade {
// 		continue
// 	}

// 	fmt.Printf(`
// -----------------------
// %s
// -----------------------
// 	`, productName)
// 	for grade, v := range gradeData {
// 		if grade == "ProductName" || grade == "" {
// 			continue // tena gradenya, sekip mi
// 		}
// 		fmt.Printf(`
// %s %d pcs | %6.2f kg`, grade, v.Count, v.Weight)
// 	}
// }

// }
