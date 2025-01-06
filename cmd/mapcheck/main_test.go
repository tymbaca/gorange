package mapcheck

import (
	"strconv"
	"testing"
	"time"
)

const (
	pricesPerOffer = 45
	offerCount     = 250_000
)

func BenchmarkCheck(b *testing.B) {
	m := make(map[string]Payload, offerCount)

	t1 := time.Now()
	t2 := time.Now().Add(-20 * time.Minute)

	for i := range 100_000 {
		key := strconv.Itoa(i)
		prices := make([]Price, 0, pricesPerOffer)
		for i := range pricesPerOffer {
			price := Price{
				NewPrice:   100,
				LastPrice:  100,
				NewUpdate:  t1,
				LastUpdate: t1,
			}

			// половина записей будет отличаться
			if i%2 == 0 {
				price.LastUpdate = t2
			}

			prices = append(prices, price)
		}

		m[key] = Payload{
			Prices: prices,
		}
	}

	for range b.N {
		events := make([]Event, 0, offerCount*pricesPerOffer)
		for offer, payload := range m {
			for _, price := range payload.Prices {
				if !price.LastUpdate.Equal(price.NewUpdate) {
					events = append(events, Event{
						OfferID: offer,
						Price:   price.NewPrice,
					})
				}
			}
		}

		_ = events
	}
}
