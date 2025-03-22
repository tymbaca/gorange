package main

import (
	"fmt"

	"github.com/rotisserie/eris"
)

func main() {
	err := eris.Errorf("something bad: %d", 342)

	err = eris.Wrapf(err, "something really bad: %d", 62)

	// fmt.Printf("%#v", eris.Unpack(err))
	// fmt.Printf("%s", eris.Cause(err).Error())
	fmt.Printf("%s", err.Error())
}




{{ yml_catalog} 2025-03-11T17:21:59+03:00 {{HOLODILNIK.RU - Москва и область Холодильник.Ру https://www.holodilnik.ru/ Holodilnik Marketplaces FG {[{RUR 1}]} true true {[{0  office supplies} {1  home appliances} {2 1 home security systems} {3 2 fitness equipment} {4 3 stationery} {5 4 pet food} {6 5 coffee and tea products} {7 6 laptop bags and cases} {8 7 smart home devices} {9 8 board games}]}} [{1 Bamboo Lime Lamp vendor.model true 0000001 http://www.forwardexploit.biz/roi/b2c/generate 0 office supplies Dabo Health Bamboo Lime Lamp   3190288401811 true [] [] For travelers and musicians who need precision, the Bold Clock Prime turbo delivers with gps-enabled, yellow design, and durable marble construction. Ideal for camping. RUR 5622379120438507 42973004893993269 []  <nil> <nil> {в наличии} <nil>} {2 Compact Hair Dryer Prime vendor.model false 0000002 https://www.seniorproactive.org/holistic/brand 6 coffee and tea products Socrata Compact Hair Dryer Prime   2729087305534 true [] [] For those who need aesthetic appeal, this angry product, made of aluminum, offers wireless and is perfect for baking. RUR 80507009890224279 20448164850063980 []  <nil> <nil> {} <nil>} {3 Sleek Gadget Zoom vendor.model false 0000003 https://www.principaldisintermediate.biz/24-365 4 stationery Honest Buildings Sleek Gadget Zoom   7321880612878 true [] [] Enjoy enhanced relaxation with this product, featuring high-performance and made with sleepy silver, delivering sustainability. RUR 5607999346857631 75355085231233060 []  <nil> <nil> {} <nil>} {4 Navy Thermostat Spark vendor.model true 0000004 https://www.productclicks-and-mortar.io/e-enable/transparent/repurpose 0 office supplies SAP Navy Thermostat Spark   7357219062614 true [] [] Designed for families and children, this product includes voice-controlled and a mysterious bronze construction, ensuring ergonomic design during pet care. RUR 59305753964370115 58166074320612821 []  <nil> <nil> {в наличии} <nil>} {5 Black Monitor Blaze vendor.model false 0000005 https://www.principalaggregate.org/methodologies 3 fitness equipment InfoCommerce Group Black Monitor Blaze   3122334964960 true [] [] Enjoy enhanced cooking with this product, featuring gps-enabled and made with prickling plastic, delivering comfort. RUR 7096983293205887 78175137887003315 []  <nil> <nil> {} <nil>}]}}
{{ }            2025-03-11T17:21:59+03:00 {{HOLODILNIK.RU - Москва и область Холодильник.Ру https://www.holodilnik.ru/ Holodilnik Marketplaces FG {[{RUR 1}]} true true {[{0  office supplies} {1  home appliances} {2 1 home security systems} {3 2 fitness equipment} {4 3 stationery} {5 4 pet food} {6 5 coffee and tea products} {7 6 laptop bags and cases} {8 7 smart home devices} {9 8 board games}]}} [{1 Bamboo Lime Lamp vendor.model true 0000001 http://www.forwardexploit.biz/roi/b2c/generate 0 office supplies Dabo Health Bamboo Lime Lamp   3190288401811 true [] [] For travelers and musicians who need precision, the Bold Clock Prime turbo delivers with gps-enabled, yellow design, and durable marble construction. Ideal for camping. RUR 5622379120438507 42973004893993269 []  <nil> <nil> {в наличии} <nil>} {2 Compact Hair Dryer Prime vendor.model false 0000002 https://www.seniorproactive.org/holistic/brand 6 coffee and tea products Socrata Compact Hair Dryer Prime   2729087305534 true [] [] For those who need aesthetic appeal, this angry product, made of aluminum, offers wireless and is perfect for baking. RUR 80507009890224279 20448164850063980 []  <nil> <nil> {} <nil>} {3 Sleek Gadget Zoom vendor.model false 0000003 https://www.principaldisintermediate.biz/24-365 4 stationery Honest Buildings Sleek Gadget Zoom   7321880612878 true [] [] Enjoy enhanced relaxation with this product, featuring high-performance and made with sleepy silver, delivering sustainability. RUR 5607999346857631 75355085231233060 []  <nil> <nil> {} <nil>} {4 Navy Thermostat Spark vendor.model true 0000004 https://www.productclicks-and-mortar.io/e-enable/transparent/repurpose 0 office supplies SAP Navy Thermostat Spark   7357219062614 true [] [] Designed for families and children, this product includes voice-controlled and a mysterious bronze construction, ensuring ergonomic design during pet care. RUR 59305753964370115 58166074320612821 []  <nil> <nil> {в наличии} <nil>} {5 Black Monitor Blaze vendor.model false 0000005 https://www.principalaggregate.org/methodologies 3 fitness equipment InfoCommerce Group Black Monitor Blaze   3122334964960 true [] [] Enjoy enhanced cooking with this product, featuring gps-enabled and made with prickling plastic, delivering comfort. RUR 7096983293205887 78175137887003315 []  <nil> <nil> {} <nil>}]}}

<smm_fbs.Shop>: {
          Shop: {
              Name: "HOLODILNIK.RU - Москва и область",
              Company: "Холодильник.Ру",
              URL: "https://www.holodilnik.ru/",
              Platform: "Holodilnik Marketplaces FG",
              Currencies: {
                  Currency: [
                      {ID: "RUR", Rate: "1"},
                  ],
              },
              Delivery: true,
              Pickup: true,
              Categories: {
                  Category: [
                      {ID: "0", ParentID: "", Name: "books"},
                      {
                          ID: "1",
                          ParentID: "",
                          Name: "home security systems",
                      },
                      {ID: "2", ParentID: "1", Name: "pet food"},
                      {ID: "3", ParentID: "2", Name: "art supplies"},
                      {ID: "4", ParentID: "3", Name: "stationery"},
                      {
                          ID: "5",
                          ParentID: "4",
                          Name: "musical instruments",
                      },
                      {ID: "6", ParentID: "5", Name: "cookware"},
                      {
                          ID: "7",
                          ParentID: "6",
                          Name: "fitness equipment",
                      },
                      {ID: "8", ParentID: "7", Name: "jewelry"},
                      {
                          ID: "9",
                          ParentID: "8",
                          Name: "sneakers and athletic shoes",
                      },
                  ],
              },
          },
          Offers: [
              {
                  ID: "1",
                  Name: "Sleek Smart Speaker Nova",
                  Type: "vendor.model",
                  Available: false,
                  XmlCode: "0000001",
                  URL: "http://www.productb2c.com/mesh/empower/leverage",
                  CategoryID: "1",
                  TypePrefix: "home security systems",
                  Vendor: "REI Systems",
                  Model: "Sleek Smart Speaker Nova",
                  ManufacturerWarranty: "",
                  WarrantyDays: "",
                  Barcode: "9175231480550",
                  Delivery: true,
                  DeliveryOptions: nil,
                  ShipmentOptions: nil,
                  Description: "This dizzying product is perfect for seniors who need time-saving. Built from leather and featuring advanced, it's ideal for collaboration.",
                  CurrencyID: "RUR",
                  Price: 58711226387602082,
                  OldPrice: 16791316442002443,
                  Picture: nil,
                  CountryOfOrigin: "",
                  Dimensions: nil,
                  Weight: nil,
                  OrderingTime: {Ordering: ""},
                  Outlets: nil,
              },
          ],
      }
  to equal
      <smm_fbs.Shop>: {
          Shop: {
              Name: "HOLODILNIK.RU - Москва и область",
              Company: "Холодильник.Ру",
              URL: "https://www.holodilnik.ru/",
              Platform: "Holodilnik Marketplaces FG",
              Currencies: {
                  Currency: [
                      {ID: "RUR", Rate: "1"},
                  ],
              },
              Delivery: true,
              Pickup: true,
              Categories: {
                  Category: [
                      {ID: "0", ParentID: "", Name: "books"},
                      {
                          ID: "1",
                          ParentID: "",
                          Name: "home security systems",
                      },
                      {ID: "2", ParentID: "1", Name: "pet food"},
                      {ID: "3", ParentID: "2", Name: "art supplies"},
                      {ID: "4", ParentID: "3", Name: "stationery"},
                      {
                          ID: "5",
                          ParentID: "4",
                          Name: "musical instruments",
                      },
                      {ID: "6", ParentID: "5", Name: "cookware"},
                      {
                          ID: "7",
                          ParentID: "6",
                          Name: "fitness equipment",
                      },
                      {ID: "8", ParentID: "7", Name: "jewelry"},
                      {
                          ID: "9",
                          ParentID: "8",
                          Name: "sneakers and athletic shoes",
                      },
                  ],
              },
          },
          Offers: [
              {
                  ID: "1",
                  Name: "Sleek Smart Speaker Nova",
                  Type: "vendor.model",
                  Available: false,
                  XmlCode: "0000001",
                  URL: "http://www.productb2c.com/mesh/empower/leverage",
                  CategoryID: "1",
                  TypePrefix: "home security systems",
                  Vendor: "REI Systems",
                  Model: "Sleek Smart Speaker Nova",
                  ManufacturerWarranty: "",
                  WarrantyDays: "",
                  Barcode: "9175231480550",
                  Delivery: true,
                  DeliveryOptions: nil,
                  ShipmentOptions: nil,
                  Description: "This dizzying product is perfect for seniors who need time-saving. Built from leather and featuring advanced, it's ideal for collaboration.",
                  CurrencyID: "RUR",
                  Price: 58711226387602082,
                  OldPrice: 16791316442002443,
                  Picture: [],
                  CountryOfOrigin: "",
                  Dimensions: nil,
                  Weight: nil,
                  OrderingTime: {Ordering: ""},
                  Outlets: nil,
              },
          ],
      }
