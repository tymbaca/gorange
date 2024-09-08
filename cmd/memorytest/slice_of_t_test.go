package main

import (
	"math/rand/v2"
	"net/http"
	"testing"
	"time"
)

type Human struct {
	Name      string
	Age       int
	BirthDate time.Time
	Credit    int
	// !!! this is not a slice !!! if it would be the slice - his inner memory will not be stored in T directly
	align [1024]byte
}

// T is 56 bytes (without align)
// Benchmark_Map_vs_Slice/slice:value-8                 162           7440204 ns/op         5603330 B/op          1 allocs/op
// Benchmark_Map_vs_Slice/slice:ptr-8                   129           9351051 ns/op         7202860 B/op     100001 allocs/op

// T is 1 KB and less - (align [1024]byte)
// Benchmark_Map_vs_Slice/slice:value-8                  79          14778411 ns/op        108003380 B/op         1 allocs/op
// Benchmark_Map_vs_Slice/slice:ptr-8                    42          27221760 ns/op        116002851 B/op    100001 allocs/op

// T is 3 KB - (align [1024 * 3]byte)
// Benchmark_Map_vs_Slice/slice:value-8                  39          30324247 ns/op        312803340 B/op         1 allocs/op
// Benchmark_Map_vs_Slice/slice:ptr-8                    33          34614981 ns/op        320802863 B/op    100001 allocs/op

// T is 10 KB - (align [1024 * 10]byte)
// Benchmark_Map_vs_Slice/slice:value-8                  12          99623101 ns/op        1029603328 B/op        1 allocs/op
// Benchmark_Map_vs_Slice/slice:ptr-8                    25          46180567 ns/op        1088802885 B/op   100001 allocs/op

func Benchmark_Map_vs_Slice(b *testing.B) {
	n := 100_000
	http.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {})

	b.Run("slice:value", func(b *testing.B) {
		for range b.N {
			humans := getHumansByValue(n)
			humans = updateHumanValues(humans)
			averageAge := getAverageAgeByValues(humans)
			_ = averageAge
		}
	})
	b.Run("slice:ptr", func(b *testing.B) {
		for range b.N {
			humans := getHumansByPtr(n)
			humans = updateHumanPtr(humans)
			averageAge := getAverageAgeByPtr(humans)
			_ = averageAge
		}
	})
}

func getHumansByValue(count int) []Human {
	result := make([]Human, 0, count)
	for range count {
		result = append(result, Human{
			Name:      "Some name",
			Age:       0, // will be updated
			BirthDate: randDate(1935, 2024),
			Credit:    rand.Int(),
		})
	}
	return result
}

func getHumansByPtr(count int) []*Human {
	result := make([]*Human, 0, count)
	for range count {
		result = append(result, &Human{
			Name:      "Some name",
			Age:       0, // will be updated
			BirthDate: randDate(1935, 2024),
			Credit:    rand.Int(),
		})
	}
	return result
}

func updateHumanValues(humans []Human) []Human {
	for i, human := range humans {
		birthYear := human.BirthDate.Year()
		currentYear := time.Now().Year()
		humans[i].Age = currentYear - birthYear // i know this is not accurate
	}
	return humans
}

func updateHumanPtr(humans []*Human) []*Human {
	for i, human := range humans {
		birthYear := human.BirthDate.Year()
		currentYear := time.Now().Year()
		humans[i].Age = currentYear - birthYear // i know this is not accurate
	}
	return humans
}

func getAverageAgeByValues(humans []Human) float64 {
	var total int
	for _, human := range humans {
		total += human.Age
	}
	return float64(total) / float64(len(humans))
}

func getAverageAgeByPtr(humans []*Human) float64 {
	var total int
	for _, human := range humans {
		total += human.Age
	}
	return float64(total) / float64(len(humans))
}

func randDate(minYear, maxYear int) time.Time {
	min := time.Date(minYear, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(maxYear, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := rand.Int64N(delta) + min
	return time.Unix(sec, 0)
}
