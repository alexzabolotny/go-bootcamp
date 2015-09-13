package main

import "fmt"

var names = []string{"Katrina", "Evan", "Neil", "Adam", "Martin", "Matt",
	"Emma", "Isabella", "Emily", "Madison",
	"Ava", "Olivia", "Sophia", "Abigail",
	"Elizabeth", "Chloe", "Samantha",
	"Addison", "Natalie", "Mia", "Alexis"}

func main() {
	var maxLen int
	for _, n := range names {
		var l = len(n)
		if l > maxLen {
			maxLen = l
		}
	}

	result := make([][]string, maxLen)
	for _, n := range names {
		var l = len(n) - 1
		result[l] = append(result[l], n)
	}

	fmt.Println(result)
}
