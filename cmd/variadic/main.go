package main

type Fooer interface {
	Foo()
}

type A struct{}

func (a A) Foo() {
}

type B struct{}

func (a B) Foo() {
}

type C struct{}

func (a C) Foo() {
}

func main() {
	// runFooers(Fooer(A{}), Fooer(B{}), []Fooer{C{}, C{}}...)
}

func runFooers(fooers ...Fooer) {
	for _, fooer := range fooers {
		fooer.Foo()
	}
}
