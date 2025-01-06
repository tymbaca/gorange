package main

import (
	"fmt"
	"testing"
)

// Define all provided structs

// BenchmarkExtractOfferIDs-8           831           1271685 ns/op         2400256 B/op          1 allocs/op

type ResponseOffers struct {
	Offers     []Offer
	Categories []Category
	Market     Market
	Warehouse  Warehouse
	WorkScheme WorkScheme
}

type Offer struct {
	OfferID           string
	StaticAttributes  StaticAttributesDTO
	DynamicAttributes DynamicAttributesDTO
	DeliveryOptions   DeliveryOptionsDTO
	ShipmentOptions   ShipmentOptionsDTO
	OzonID            string
}

type StaticAttributesDTO struct {
	Title                   string
	Description             string
	TechnicalSpecifications TechnicalSpecificationsDTO
	CategoryID              int64
	Pictures                []string
	Vendor                  string
	URL                     string
	CategoryName            string
	ProductName             string
	EAN                     string
}

type TechnicalSpecificationsDTO struct {
	Weight               string
	Height               string
	Length               string
	TechnicalDescription string
	Gabarity             string
	Country              string
	Warranty             string
}

type DynamicAttributesDTO struct {
	Price     PriceDTO
	Stocks    StockDTO
	Available bool
}

type PriceDTO struct {
	CostCopecs    int64
	OldCostCopecs int64
}

type StockDTO struct {
	Warehouse Warehouse
	Stock     int64
}

type DeliveryOptionsDTO struct {
	Days        int64
	Cost        int64
	OrderBefore int64
}

type ShipmentOptionsDTO struct {
	Days        int64
	OrderBefore int64
}

// Mock types for additional structs
type (
	Market     string
	Warehouse  string
	WorkScheme string
)

type Category struct {
	ID       string
	ParentID string
	Name     string
}

// Helper function to generate mock offers with realistic data
func generateOffers(numOffers int) []Offer {
	offers := make([]Offer, numOffers)
	for i := 0; i < numOffers; i++ {
		offers[i] = Offer{
			OfferID: fmt.Sprintf("OfferID-%d", i),
			StaticAttributes: StaticAttributesDTO{
				Title:       fmt.Sprintf("Product Title %d", i),
				Description: fmt.Sprintf("Description for product %d", i),
				TechnicalSpecifications: TechnicalSpecificationsDTO{
					Weight:               "1kg",
					Height:               "10cm",
					Length:               "20cm",
					TechnicalDescription: "High quality",
					Gabarity:             "Standard",
					Country:              "Country XYZ",
					Warranty:             "1 year",
				},
				CategoryID:   int64(i),
				Pictures:     []string{"http://example.com/pic1.jpg", "http://example.com/pic2.jpg"},
				Vendor:       "Vendor XYZ",
				URL:          fmt.Sprintf("http://example.com/product%d", i),
				CategoryName: fmt.Sprintf("Category %d", i),
				ProductName:  fmt.Sprintf("Product Name %d", i),
				EAN:          fmt.Sprintf("EAN-%d", i),
			},
			DynamicAttributes: DynamicAttributesDTO{
				Price: PriceDTO{
					CostCopecs:    int64(i * 1000),
					OldCostCopecs: int64(i * 1100),
				},
				Stocks: StockDTO{
					Warehouse: "западный",
					Stock:     int64(i * 10),
				},
				Available: true,
			},
			DeliveryOptions: DeliveryOptionsDTO{
				Days:        int64(i % 5),
				Cost:        int64(i * 50),
				OrderBefore: 12,
			},
			ShipmentOptions: ShipmentOptionsDTO{
				Days:        int64(i % 3),
				OrderBefore: 18,
			},
			OzonID: fmt.Sprintf("Ozon-%d", i),
		}
	}
	return offers
}

func BenchmarkExtractOfferIDs(b *testing.B) {
	numOffers := 150_000 // Adjust for different dataset sizes
	offers := generateOffers(numOffers)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = extractOfferIDs(offers)
	}
}

// Function to extract OfferIDs from []*Offer
func extractOfferIDs(offers []Offer) []string {
	ids := make([]string, len(offers))
	for i, offer := range offers {
		ids[i] = offer.OfferID
	}
	return ids
}
