package main

import (
	"cmp"
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
)

const count = 1_000 //_000

type Stock struct {
	ID          int       `gorm:"primaryKey;autoIncrement"`
	OfferID     string    `gorm:"size:30;index"`
	AccountName string    `gorm:"not null"`
	Warehouse   int       `gorm:"not null"`
	LastStock   int       `gorm:"not null"` // LastStock последний отправленный остаток на площадке
	NewStock    int       `gorm:"not null"` // NewStock самый актуальный остаток
	LastUpdated time.Time `gorm:"not null"` // LastUpdated дата последнего обновления на площадке
	NewUpdated  time.Time `gorm:"not null"` // NewUpdated дата самого актуального остатка
}

func BenchmarkStocks(b *testing.B) {
	stocks := make([]Stock, 0, count)
	for range count {
		stocks = append(stocks, Stock{
			ID:          rand.Int(),
			OfferID:     gofakeit.DigitN(6),
			AccountName: gofakeit.DigitN(35),
			Warehouse:   gofakeit.IntN(10),
			LastStock:   0,
			NewStock:    0,
			LastUpdated: time.Time{},
			NewUpdated:  time.Time{},
		})
	}

	// b.Run("sort", func(b *testing.B) {
	// 	for range b.N {
	// 		sort.Slice(stocks, func(i, j int) bool {
	// 			return stocks[i].OfferID < stocks[j].OfferID
	// 		})
	// 	}
	// })
	// b.Run("slices:basic", func(b *testing.B) {
	// 	for range b.N {
	// 		slices.SortFunc(stocks, func(a, b Stock) int {
	// 			if a.OfferID < b.OfferID {
	// 				return -1
	// 			}
	// 			if a.OfferID > b.OfferID {
	// 				return 1
	// 			}
	//
	// 			return 0
	// 		})
	// 	}
	// })
	// b.Run("slices:cmp", func(b *testing.B) {
	// 	for range b.N {
	// 		slices.SortFunc(stocks, func(a, b Stock) int {
	// 			return cmp.Compare(a.OfferID, b.OfferID)
	// 		})
	// 	}
	// })
	b.Run("slices:cmp3", func(b *testing.B) {
		// sort.SliceStable(stocks, func(i, j int) bool {
		// 	// return cmp.Less(stocks[i].OfferID, stocks[j].OfferID)
		// 	return cmp.Or(
		// 		cmp.Less(stocks[i].OfferID, stocks[j].OfferID),
		// 		cmp.Less(stocks[i].AccountName, stocks[j].AccountName),
		// 		cmp.Less(stocks[i].Warehouse, stocks[j].Warehouse),
		// 	)
		// })
		// slices.SortFunc(stocks, func(a, b Stock) int {
		// 	return cmp.Or(
		// 		cmp.Compare(a.OfferID, b.OfferID),
		// 		cmp.Compare(a.AccountName, b.AccountName),
		// 		cmp.Compare(a.Warehouse, b.Warehouse),
		// 	)
		// })
		sort.SliceStable(stocks, func(i, j int) bool {
			if stocks[i].OfferID != stocks[j].OfferID {
				return stocks[i].OfferID < stocks[j].OfferID
			}

			if stocks[i].AccountName != stocks[j].AccountName {
				return cmp.Less(stocks[i].AccountName, stocks[j].AccountName)
			}

			if stocks[i].Warehouse != stocks[j].Warehouse {
				return stocks[i].Warehouse < stocks[j].Warehouse
			}
			return false
		})

		fmt.Println(stocks)
	})
}
