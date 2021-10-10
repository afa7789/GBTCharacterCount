package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

type TreeNode struct {
	left    *TreeNode // node
	right   *TreeNode // node
	value   *byte     // data it represents
	counter int       // number of times we saw it
}

// start the tree
func startTree() TreeNode {
	return TreeNode{
		value:   nil,
		counter: 0,
		left:    nil,
		right:   nil,
	}
}

// receive the int/char and add it to right place or increase the counter.
func addToTree(a byte, n *TreeNode) {

	// if root node start it
	if !(n.value != nil) {
		n.value = &a
		n.counter = 1
	}

	// if equal increase counter
	if *n.value == a {
		n.counter++

	} else if *n.value > a { // if lesser put it on left
		if n.left != nil { // already has left node do it again on the right
			addToTree(a, n.left)
		} else {
			n.left = &TreeNode{
				value:   &a,
				counter: 1,
				left:    nil,
				right:   nil,
			}
		}
	} else if *n.value < a { // if greater put it on left
		if n.right != nil { // already has right node do it again on the right
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

// TreePrint, print a node as 'character' , character as int , quantidade: number_of_times_read
func TreePrint(t *TreeNode) {
	fmt.Printf(" '%s' ,%v quantidade: %v \n", string(*t.value), *t.value, t.counter)
	// if left node not null call it again
	if t.left != nil {
		TreePrint(t.left)
	}
	// if right node not null call it again
	if t.right != nil {
		TreePrint(t.right)
	}
}

// CoolTreePrint , print the same as above put less verbose and with cool Filled Rectangles
func CoolTreePrint(t *TreeNode, max int) {
	fmt.Printf(" '%s' | id: %4v | qtd: %5v | ", string(*t.value), *t.value, t.counter)

	// casting and conversions
	// turn to float the incoming int to have more exact division
	f := float64(max)
	// the higher this can be is 1
	i := float64(t.counter) / f
	//50 is the max number of rectangles
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

// check error/eof to end programn, Print the node trees.
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

//main duh
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
