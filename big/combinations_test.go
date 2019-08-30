package combinations

import "fmt"

func ExampleAll() {
	combinations := All([]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O"})
	fmt.Println(len(combinations))
	// Output:
	// 32767
}

func ExampleAll_Small() {
	combinations := All([]string{"A", "B", "C"})
	fmt.Println(combinations)
	// Output:
	// [[A] [B] [A B] [C] [A C] [B C] [A B C]]
}
