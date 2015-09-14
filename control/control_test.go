package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDistributeBitcoins(t *testing.T) {
	var (
		people = []string{
			"Matthew", "Sarah", "Augustus", "Heidi", "Emilie",
			"Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
		}
		stash = 50

		dist = map[string]int{
			"Matthew":   2,
			"Peter":     2,
			"Giana":     4,
			"Adriano":   7,
			"Elizabeth": 5,
			"Sarah":     2,
			"Augustus":  10,
			"Heidi":     5,
			"Emilie":    6,
			"Aaron":     5,
		}
		left = 2
	)

	gotDist, gotLeft := distributeBitcoins(people, stash)
	if !reflect.DeepEqual(gotDist, dist) {
		t.Error(fmt.Printf("Distribution is not as expected: %v instead of %v\n", gotDist, dist))
	}
	if gotLeft != left {
		t.Error(fmt.Printf("Reminder is not as expected: %v instead of %v\n", gotLeft, left))
	}
}
