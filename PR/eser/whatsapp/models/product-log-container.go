package models

import (
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
	UUID       string  `json:"uuid" bson:"uuid"`
	Number     int     `json:"number" bson:"number"`
	Weight     float64 `json:"weight" bson:"weight"`
	Count      int     `json:"count" bson:"count"`
	Grade      string  `json:"grade" bson:"grade"`
	Product    Product `json:"product" bson:"product"`
	Date       string  `json:"date" bson:"date"`
	CreatedAt  int64   `json:"created_at" bson:"created_at"`
	DeletedAt  int64   `json:"deleted_at" bson:"deleted_at"`
	CountUsed  int     `json:"count_used" bson:"count_used"`
	WeightUsed float64 `json:"weight_used" bson:"weight_used"`
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

type CodeProduct struct {
	Code map[string]Ungraded
}

type Ungraded struct {
	Weight float64 `json:"weight" bson:"weight"`
	Count  int64   `json:"count" bson:"count"`
	Name   string  `json:"name" bson:"name"`
	Code   string  `json:"code" bson:"code"`
}

type ProductLogContainer struct {
	ID              primitive.ObjectID             `json:"_id,omitempty" bson:"_id,omitempty"`
	RefinedILC      string                         `json:"refined_ilc" bson:"refined_ilc"`
	ReceivingDate   int64                          `json:"receiving_date" bson:"receiving_date"`
	CuttingDate     int64                          `json:"cutting_date" bson:"cutting_date"`
	Supplier        ProductLogSupplier             `json:"supplier" bson:"supplier"`
	CustomerGroup   string                         `json:"customer_group" bson:"customer_group"`
	TotalWeight     float64                        `json:"total_weight" bson:"total_weight"`
	TotalCount      int64                          `json:"total_count" bson:"total_count"`
	RawTotalCount   int64                          `json:"raw_total_count" bson:"raw_total_count"`
	RawTotalWeight  float64                        `json:"raw_total_weight" bson:"raw_total_weight"`
	CreatedAt       int64                          `json:"created_at" bson:"created_at"`
	UpdatedAt       int64                          `json:"updated_at" bson:"updated_at"`
	Contents        []ProductLogContent            `json:"contents" bson:"contents"`
	GroupedContents map[string]map[string]Ungraded `json:"grouped_contents" bson:"grouped_contents"`
	Species         ProductLogSpecies              `json:"species" bson:"species"`
	Type            ProductLogType                 `json:"type" bson:"type"`
	UUID            string                         `json:"uuid" bson:"uuid"`
}

var ProductLogResult []*ProductLogContainer

func (plc *ProductLogContainer) PrintProductLogContainerDataByDate() map[string]Ungraded {

	dataUngrade := make(map[string]Ungraded)

	for j, _ := range plc.GroupedContents {
		for k, v := range plc.GroupedContents[j] {

			if dataUngrade[v.Name].Weight == 0 {
				dataUngrade[v.Name] = Ungraded{
					Weight: v.Weight,
					Count:  v.Count,
					Name:   k,
				}
			} else {
				dataUngrade[k] = Ungraded{
					Weight: dataUngrade[k].Weight + v.Weight,
					Count:  dataUngrade[k].Count + v.Count,
					Name:   v.Name,
				}
			}
		}
	}
	return dataUngrade
	// fmt.Println(dataUngrade)
}

// func (plc *ProductLogContainer) PrintProductLogContainerDataByDate() string {
// 	result := ""

// 	for key, v := range plc.GroupedContents {
// 		result += fmt.Sprintf("            %s            \n", key)
// 		result += "---------------------------\n"

// 		for _, ungraded := range v {
// 			result += fmt.Sprintf("%s\n", ungraded.Name)
// 			result += fmt.Sprintf("%s   %d pcs | %.2f kg\n", ungraded.Code, ungraded.Count, ungraded.Weight)
// 		}

// 		result += "\n"
// 	}

// 	var totalPcs int64 = 0
// 	totalWeight := 0.0

// 	for _, v := range plc.GroupedContents {
// 		for _, ungraded := range v {
// 			totalPcs += ungraded.Count
// 			totalWeight += ungraded.Weight
// 		}
// 	}

// 	result += fmt.Sprintf("Total: %d pcs | %.3f kg\n", totalPcs, totalWeight)

// 	return result
// }
