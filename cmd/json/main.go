package main

import (
	"encoding/json"
	"fmt"

	"gorm.io/datatypes"
)

type OzonAuth struct {
	ClientID string `json:"clientID"`
	ApiKey   string `json:"apiKey"`
}

type MegamarketJsonFeed struct {
	FileAttributes FileAttributes
	Outlets        []Outlet
}

type FileAttributes struct {
	MerchantID int
	Type       string
	Datetime   string
}

type Outlet struct {
	ID     int     `json:"id"`
	Offers []Offer `json:"offers"`
}

type Offer struct {
	ID       int
	Price    int
	Quantity int
	OldPrice int
}

func main() {
	data := `["fsdfsd", "beasfa"]`
	msg := datatypes.JSON(json.RawMessage(data))
	fmt.Printf("msg.String(): %s\n", msg.String())

	// var ozonAuth OzonAuth
	// json.Unmarshal([]byte(`{
	//                 "some-shit": "tymbaca",
	//                 "another stuff": "mykey"
	//         }`), &ozonAuth)
	//
	// fmt.Printf("%#v\n", ozonAuth)

	// feed := MegamarketJsonFeed{
	// 	FileAttributes: FileAttributes{
	// 		MerchantID: 125012,
	// 		Type:       "shit",
	// 		Datetime:   time.Now().Format(time.RFC3339),
	// 	},
	// 	Outlets: []Outlet{
	// 		{ID: 15515, Offers: []Offer{
	// 			{ID: 1252136, Price: 1351, Quantity: 236, OldPrice: 3262},
	// 		}},
	// 	},
	// }

	// data, err := json.MarshalIndent(Outlet{
	// 	ID:     10,
	// 	Offers: []Offer{},
	// }, "", "  ")
	// if err != nil {
	// 	panic(err)
	// }
	//
	// fmt.Println(string(data))
}
