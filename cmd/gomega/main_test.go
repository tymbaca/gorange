package main

import (
	"testing"
	"time"

	. "github.com/onsi/gomega"
)

type Event struct {
	Value int
	Time  time.Time
}

func TestMatcher(t *testing.T) {
	RegisterTestingT(t)

	expected := []Event{{Value: 10}, {Value: 20}}
	actual := []Event{{Value: 20, Time: time.Now()}, {Value: 10, Time: time.Now()}}

	Expect(actual).To(WithTransform(func(es []Event) []Event {
		for i := range es {
			es[i].Time = time.Time{}
		}
		return es
	}, ConsistOf(expected)))
}
