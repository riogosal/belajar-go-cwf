package models

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Phone struct {
	CountryCode string `json:"country_code" bson:"country_code"`
	Number      string `json:"number" bson:"number"`
}

type Address struct {
	Country  string `json:"country" bson:"country"`
	State    string `json:"state" bson:"state"`
	District string `json:"district" bson:"district"`
	Street   string `json:"street" bson:"street"`
}

type PackingCustomer struct {
	UUID      string  `json:"uuid" bson:"uuid"`
	Name      string  `json:"name" bson:"name"`
	Code      string  `json:"code" bson:"code"`
	Unit      string  `json:"unit" bson:"unit"`
	Alias     string  `json:"alias" bson:"alias"`
	Phone     Phone   `json:"phone" bson:"phone"`
	Address   Address `json:"address" bson:"address"`
	LicenseID string  `json:"license_id" bson:"license_id"`
}

type PackingSpecies struct {
	Name           string `json:"name" bson:"name"`
	ScientificName string `json:"scientific_name" bson:"scientific_name"`
	Code           string `json:"code" bson:"code"`
}

type PackingProduct struct {
	UUID                    string         `json:"uuid" bson:"uuid"`
	Name                    string         `json:"name" bson:"name"`
	Code                    string         `json:"code" bson:"code"`
	Category                string         `json:"category" bson:"category"`
	Species                 PackingSpecies `json:"species" bson:"species"`
	SapCode                 string         `json:"sap_code" bson:"sap_code"`
	CustomerProductCode     string         `json:"customer_product_code" bson:"customer_product_code"`
	CustomerProductCategory string         `json:"customer_product_category" bson:"customer_product_category"`
}

type ContentSpecies struct {
	Name           string `json:"name" bson:"name"`
	ScientificName string `json:"scientific_name" bson:"scientific_name"`
	Code           string `json:"code" bson:"code"`
}

type ProductContent struct {
	UUID                    string         `json:"uuid" bson:"uuid"`
	Name                    string         `json:"name" bson:"name"`
	Code                    string         `json:"code" bson:"code"`
	Category                string         `json:"category" bson:"category"`
	Species                 ContentSpecies `json:"species" bson:"species"`
	SapCode                 string         `json:"sap_code" bson:"sap_code"`
	CustomerProductCode     string         `json:"customer_product_code" bson:"customer_product_code"`
	CustomerProductCategory string         `json:"customer_product_category" bson:"customer_product_category"`
	Grade                   string         `json:"grade" bson:"grade"`
}

type PackingContent struct {
	UUID           string         `json:"uuid" bson:"uuid"`
	ProductLogUUID string         `json:"product_log_uuid" bson:"product_log_uuid"`
	Ilc            string         `json:"ilc" bson:"ilc"`
	CustomerGroup  string         `json:"customer_group" bson:"customer_group"`
	Product        ProductContent `json:"product" bson:"product"`
	Weight         float64        `json:"weight" bson:"weight"`
	Count          int            `json:"count" bson:"count"`
}

type Packing struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UUID         string             `json:"uuid" bson:"uuid"`
	PreviousUUID string             `json:"previous_uuid" bson:"previous_uuid"`
	Customer     PackingCustomer    `json:"customer" bson:"customer"`
	Product      PackingProduct     `json:"product" bson:"product"`
	Box          string             `json:"box" bson:"box"`
	CreatedAt    int64              `json:"created_at" bson:"created_at"`
	TotalWeight  float64            `json:"total_weight" bson:"total_weight"`
	TotalCount   int                `json:"total_count" bson:"total_count"`
	Contents     []PackingContent   `json:"contents" bson:"contents"`
	ContentsILC  []string           `json:"contents_ilc" bson:"contents_ilc"`
}

type WeightAndCount struct {
	Count        int64
	Weight       float64
	CustNameGrup string
}

var Packings []*Packing

func PrintPackingDataByDate(resultPacking []*Packing) {
	dataCollecPacking := make(map[string]WeightAndCount)

	for _, dataPacking := range Packings {
		if dataCollecPacking[dataPacking.Product.Name].Count == 0 {
			dataCollecPacking[dataPacking.Product.Name] = WeightAndCount{
				Count:        int64(dataPacking.TotalCount),
				Weight:       dataPacking.TotalWeight,
				CustNameGrup: dataPacking.Customer.Name,
			}
		} else {
			dataCollecPacking[dataPacking.Product.Name] = WeightAndCount{
				Count:        dataCollecPacking[dataPacking.Product.Name].Count + int64(dataPacking.TotalCount),
				Weight:       dataCollecPacking[dataPacking.Product.Name].Weight + dataPacking.TotalWeight,
				CustNameGrup: dataPacking.Customer.Name,
			}
		}
	}
	// fmt.Println(dataCollecPacking)r
	for k, v := range dataCollecPacking {
		fmt.Printf(`
--------------------------------------
%s 
--------------------------------------
%s %d pcs | %2.2f Kg
		`, v.CustNameGrup, k, v.Count, v.Weight)
	}
}
