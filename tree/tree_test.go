package main

import (
	"fmt"
	"testing"

	"golang.org/x/tour/tree"
)

func TestWalk(t *testing.T) {
	cases := []struct {
		in   *tree.Tree
		want [10]int
	}{
		{tree.New(1), [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{tree.New(2), [10]int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}},
	}

	for _, c := range cases {
		ch := make(chan int)
		go Walk(c.in, ch)

		var walked [10]int

		i := 0
		for v := range ch {
			walked[i] = v
			i++
		}

		if walked != c.want {
			t.Errorf("got: %v", walked)
		} else {
			fmt.Print(".")
		}
	}
}

func TestSame(t *testing.T) {
	cases := []struct {
		t1   *tree.Tree
		t2   *tree.Tree
		want bool
	}{
		{tree.New(1), tree.New(1), true},
		{tree.New(1), tree.New(2), false},
	}

	for _, c := range cases {
		if c.want != Same(c.t1, c.t2) {
			t.Errorf("%v is equal to %v", c.t1, c.t2)
		} else {
			fmt.Print(".")
		}
	}
}
