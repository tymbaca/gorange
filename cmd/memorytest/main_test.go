package main

import (
	"testing"
	// . "github.com/qrdl/testaroli"
)

// type Human struct {
// 	Name   string
// 	Age    int
// 	Credit int
// 	_align [512]byte
// }

// func Benchmark_Map_vs_Slice(b *testing.B) {
// 	n := 10_000_000
// 	// m := make(map[string]Human, n)
// 	// for i := range n {
// 	// 	name := fake.FullName()
// 	// 	m[name] = Human{
// 	// 		Name:   name,
// 	// 		Age:    i,
// 	// 		Credit: i * 10,
// 	// 	}
// 	// }
// 	// b.Run("map", func(b *testing.B) {
// 	// 	for range b.N {
// 	// 		for range m {
// 	// 		}
// 	// 	}
// 	// })

// 	s := make([]Human, n)
// 	for i := range n {
// 		s[i] = Human{
// 			Name:   fake.FullName(),
// 			Age:    i,
// 			Credit: i * 10,
// 		}
// 	}

// 	b.Run("slice:value", func(b *testing.B) {
// 		for range b.N {
// 			for range s {
// 			}
// 		}
// 	})
// }

type Getter interface {
	Get() int
}

type getterPtr struct {
	val int
}

func (g *getterPtr) Get() int {
	return g.val
}

type getterWithoutPtr struct {
	val int
}

func (g getterWithoutPtr) Get() int {
	return g.val
}

func SumI(slice []Getter) int {
	var total int
	for _, g := range slice {
		total += g.Get()
	}
	return total
}

func SumG[T Getter](slice []T) int {
	var total int
	for _, g := range slice {
		total += g.Get()
	}
	return total
}

func BenchmarkSum(b *testing.B) {
	n := 1_000_000

	b.Run("Interface", func(b *testing.B) {
		b.Run("[]Getter <- &getterPtr", func(b *testing.B) {
			for range b.N {
				iS1 := make([]Getter, n)
				for i := range n {
					iS1[i] = &getterPtr{val: i}
				}
				SumI(iS1)
			}
		})
		b.Run("[]Getter <- &getterWithoutPtr", func(b *testing.B) {
			for range b.N {
				iS2 := make([]Getter, n)
				for i := range n {
					iS2[i] = &getterWithoutPtr{val: i}
				}
				SumI(iS2)
			}
		})
		b.Run("[]Getter <- getterWithoutPtr", func(b *testing.B) {
			for range b.N {
				iS3 := make([]Getter, n)
				for i := range n {
					iS3[i] = getterWithoutPtr{val: i}
				}
				SumI(iS3)
			}
		})
		b.Run("[]Getter <- 1 getterWithoutPtr + ALL getterPtr", func(b *testing.B) {
			for range b.N {
				iS4 := make([]Getter, n)
				for i := range n {
					if i == 0 {
						iS4[i] = getterWithoutPtr{val: i}
					} else {
						iS4[i] = &getterPtr{val: i}
					}
				}
				SumI(iS4)
			}
		})
	})

	b.Run("generic", func(b *testing.B) {
		b.Run("&getterPtr", func(b *testing.B) {
			for range b.N {
				gS1 := make([]*getterPtr, n)
				for i := range n {
					gS1[i] = &getterPtr{val: i}
				}
				SumG(gS1)
			}
		})
		b.Run("&getterWithoutPtr", func(b *testing.B) {
			for range b.N {
				gS2 := make([]*getterWithoutPtr, n)
				for i := range n {
					gS2[i] = &getterWithoutPtr{val: i}
				}
				SumG(gS2)
			}
		})
		b.Run("getterWithoutPtr", func(b *testing.B) {
			for range b.N {
				gS3 := make([]getterWithoutPtr, n)
				for i := range n {
					gS3[i] = getterWithoutPtr{val: i}
				}

				SumG(gS3)
			}
		})
	})
}
