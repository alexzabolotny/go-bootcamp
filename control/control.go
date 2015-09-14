package main

import (
	"fmt"
	"strings"
)

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie",
		"Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
)

const maxPerPerson = 10

func distributeBitcoins(people []string, stash int) (map[string]int, int) {
	var dRules = map[string]int{"a": 1, "e": 1, "i": 2, "o": 3, "u": 4}
	var distribution = make(map[string]int, len(people))

	for _, person := range people {
		for _, char := range strings.Split(strings.ToLower(person), "") {
			rule, ok := dRules[char]
			if !ok {
				continue
			}
			personal, ok := distribution[person]
			if !ok {
				distribution[person] = 0
				personal = 0
			}
			var change = 0
			if personal+rule > maxPerPerson {
				change = maxPerPerson - personal
				personal = maxPerPerson
			} else {
				personal += rule
				change = rule
			}

			distribution[person] = personal
			stash -= change
		}
	}

	return distribution, stash
}

func main() {
	d, r := distributeBitcoins(users, coins)
	fmt.Println(d)
	fmt.Println("Coins left:", r)
}
