package document

import (
	"fmt"
)

func ExampleAdd() {
	result := Add(4, 2)
	fmt.Println("4 + 2 =", result)

	// Output:
	// 4 + 2 = 6
}

func ExampleSub() {
	result := Sub(4, 2)
	fmt.Println("4 - 2 =", result)

	// Output:
	// 4 - 2 = 2
}

func ExampleMul() {
	result := Mul(4, 2)
	fmt.Println("4 * 2 =", result)

	// Output:
	// 4 * 2 = 8
}

func ExampleDiv() {
	result := Div(4, 2)
	fmt.Println("4 / 2 =", result)

	// Output:
	// 4 / 2 = 2
}
