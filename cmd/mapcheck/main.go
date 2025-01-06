package mapcheck

import "time"

type Payload struct {
	Prices []Price
}

type Price struct {
	NewPrice   int
	LastPrice  int
	NewUpdate  time.Time
	LastUpdate time.Time
	Type       int
}

type Event struct {
	OfferID string
	Price   int
}
