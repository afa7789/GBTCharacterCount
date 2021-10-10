package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

type TreeNode struct {
	left    *TreeNode
	right   *TreeNode
	value   *byte
	counter int
}

func startTree() TreeNode {
	return TreeNode{
		value:   nil,
		counter: 0,
		left:    nil,
		right:   nil,
	}
}

func addToTree(a byte, n *TreeNode) {

	if !(n.value != nil) {
		n.value = &a
		n.counter = 1
	}

	if *n.value == a {
		n.counter++
	} else if *n.value > a {
		if n.left != nil {
			addToTree(a, n.left)
		} else {
			n.left = &TreeNode{
				value:   &a,
				counter: 1,
				left:    nil,
				right:   nil,
			}
		}
	} else if *n.value < a {
		if n.right != nil {
			addToTree(a, n.right)
		} else {
			n.right = &TreeNode{
				value:   &a,
				counter: 1,
				left:    nil,
				right:   nil,
			}
		}
	}
}

// maior contador da arvore.
func MaxCounterOnTree(t *TreeNode, max int) int {
	aux := max
	if t.left != nil {
		aux = MaxCounterOnTree(t.left, max)
		if aux > max {
			max = aux
		}
	}
	if t.right != nil {
		aux = MaxCounterOnTree(t.right, max)
		if aux > max {
			max = aux
		}
	}
	if t.counter > max {
		max = t.counter
	}
	return max
}

func TreePrint(t *TreeNode) {
	fmt.Printf(" '%s' ,%v quantidade: %v \n", string(*t.value), *t.value, t.counter)
	if t.left != nil {
		TreePrint(t.left)
	}
	if t.right != nil {
		TreePrint(t.right)
	}
}

func CoolTreePrint(t *TreeNode, max int) {
	fmt.Printf(" '%s' | id: %4v | qtd: %5v | ", string(*t.value), *t.value, t.counter)
	f := float64(max)
	i := float64(t.counter) / f
	for j := 0; j < int(i*50); j++ {
		fmt.Printf("▮")
	}
	fmt.Printf("\n")
	if t.left != nil {
		CoolTreePrint(t.left, max)
	}
	if t.right != nil {
		CoolTreePrint(t.right, max)
	}
}

func check(e error, f *os.File, t *TreeNode) {
	if e != nil {
		if e == io.EOF {
			max := MaxCounterOnTree(t, 0)
			print("Record of Repeats: " + fmt.Sprint(max) + "\n")
			CoolTreePrint(t, max)
			f.Close()
			os.Exit(0)
		} else {
			panic(e)
		}
	}
}

func main() {

	filename := flag.String("file", "example.txt", "file to characters")
	t := startTree()
	f, err := os.Open(*filename)
	check(err, f, &t)

	b1 := make([]byte, 1)
	for {
		n1, err := f.Read(b1)
		check(err, f, &t)
		if n1 > 0 {
			addToTree(b1[n1-1], &t)
		}
	}
}
