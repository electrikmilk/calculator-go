package main

import "fmt"

func main() {
	var opt1 int32
	var opt2 int32
	var oper string
	var result int32

	input("Enter first operand", &opt1)
	input_str("Enter the operation [+, -, *, /]", &oper)
	input("Enter second operand", &opt2)

	fmt.Printf("Performing %s on %d and %d...\n", oper, opt1, opt2)
	switch oper {
	case "+":
		add(&opt1, &opt2, &result)
	case "-":
		sub(&opt1, &opt2, &result)
	case "*":
		mul(&opt1, &opt2, &result)
	case "/":
		div(&opt1, &opt2, &result)
	default:
		fmt.Println("Uh oh! Unknown operation.")
		return
	}
	fmt.Printf("= %d\n", result)
}

func input(prompt string, fill *int32) {
	fmt.Printf("%s: ", prompt)
	fmt.Scanln(*&fill)
}

func input_str(prompt string, fill *string) {
	fmt.Printf("%s: ", prompt)
	fmt.Scanln(*&fill)
}

func add(opt1 *int32, opt2 *int32, result *int32) {
	*result = *opt1 + *opt2
}

func sub(opt1 *int32, opt2 *int32, result *int32) {
	*result = *opt1 - *opt2
}

func mul(opt1 *int32, opt2 *int32, result *int32) {
	*result = *opt1 * *opt2
}

func div(opt1 *int32, opt2 *int32, result *int32) {
	*result = *opt1 / *opt2
}
