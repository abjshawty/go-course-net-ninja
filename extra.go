package main

import "fmt"

// points is a sample slice of integers used for demonstration purposes.
var points = []int{20, 90, 100, 45, 70}

// tarrifs prints a message indicating tariffs were applied to the given item.
func tarrifs(n string) {
	fmt.Println("Tarrifs were applied on: ", n)
}

// curse prints a message about the given person being defeated by the strongest sorcerer.
func curse(n string) {
	fmt.Printf("%s was beaten by the strongest sorcerer of today, %s", n, strongest)
}
