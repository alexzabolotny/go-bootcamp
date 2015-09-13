package main

import "golang.org/x/tour/tree"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	recWalk(t, ch)

	close(ch)
}

func recWalk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		recWalk(t.Left, ch)
	}

	ch <- t.Value

	if t.Right != nil {
		recWalk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		x1, ok1 := <-ch1
		x2, ok2 := <-ch2
		switch {
		case ok1 != ok2:
			return false
		case !ok1:
			return true
		case x1 != x2:
			return false
		}
	}
}

func main() {
	ch := make(chan int)

	Walk(tree.New(1), ch)
}
