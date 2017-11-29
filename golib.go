// Package golib contains utility functions for learning Go.
package golib

import "fmt"
import "runtime"
import "time"

// A function is exported if its name begins with a capital letter
// Function arguments must have a name and a type
// When consecutive function parameters share a type, the type can be omitted
// from all but the last
// x int, y int => x, y int
func Add(x, y int) int {
	return x + y
}

// Return values can be named.
// Naked return statements should be used only in short functions.
// They can harm readability in longer functions.
func Concat(x, y string) (z string) {
	// Named return values are treated as variables defined at the top of the
	// function.
	z = x + " " + y + "\n"
	// A return statement without arguments returns the named return values.
	// This is known as a "naked" return.
	return
}

// If
func Conditioner(checkMe int) {
	add := 1
	var result string

	// Like `for`, `if` can have a short `init` statement to execute before condition
	// This variable is scoped to the `if` statement
	// if checkMe < 10 {
	if v := checkMe + add; v < 10 {
		result = "less than 10"
	} else {
		result = "not less than 10"
	}

	// `v` is not available here!
	// fmt.Printf(v)

	fmt.Printf("%v is %v (if you add %v)\n", checkMe, result, add)
}

// Echo prints its argument to the console.
func Echo(s string) {
	fmt.Printf(s)
}

// Flow Control
func Looper() {
	// Go has only one looping construct, the for loop.
	sumA := 0

	// init statement: executed before the first iteration, scoped to loop => `i := 0`
	// condition expression: evaluated before every iteration => `i < 10`
	// post statement: executed at the end of every iteration => `i++`
	for i := 0; i < 10; i++ {
		sumA += i
		fmt.Println(sumA)
	}

	// init and post statements are optional
	// This allows Go's `for` loop to act like a `while` loop
	sumB := 1
	for sumB < 1000 {
		sumB += sumB
	}
	fmt.Println(sumB)
}

// A pointer holds the memory address of a value.
func Pointers() {
	// The type *T is a pointer to a T value. Its zero value is nil.
	var p *int

	i, j := 42, 2701

	// The & operator generates a pointer to its operand.
	p = &i 					// point to i

	// The * operator denotes the pointer's underlying value.
	// This is known as "dereferencing" or "indirecting".
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j          // point to j
	*p = *p / 37    // divide j through the pointer
	fmt.Println(j)  // see the new value of j
}

// Reverse returns its argument string reversed rune-wise left to right.
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// Structs
func Structs()  {
	// A `struct` is a collection of fields.
	type Vertex struct {
		X int
		Y int
	}

	// Structs can be constructed with `{}`
	v := Vertex{1, 2}
	fmt.Println(v)

	// Struct fields are accessed using dot notation
	v.X = 4
	fmt.Println(v.X)

	// Struct fields can be accessed through a struct pointer.
	p := &v
	p.X = 1e9
	fmt.Println(v)

	// You can list just a subset of fields by using the Name: syntax.
	// The order of named fields is irrelevant.
	var v2 = Vertex{X: 1}  // Y:0 is implicit
	fmt.Println(v2)
}

// A function can return any number of results.
func Swap(x, y string) (string, string) {
	return y, x
}

func Switcheroo()  {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOS")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s.", os)
	}

	// Switch statements without a condition can be used to cleanly construct
	// long if-then-else chains
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

// The `var` statement declares a list of one or more variables and their type
// A `var` statement can be at package level...
var c, python, java bool
func Variables() {
	// or , function level
	var i int
	// A var declaration can include initializers, one per variable.
	var a, b string = "a", "b"
	// If an initializer is present, the type can be omitted and inferred from the
	// type of the initializer
	var n = 123
	// The `:=` short assignment statement can be used in place of a `var`
	// declaration with implicit type
	// The `:=` construct is ONLY available at function level, NOT at package level
	m := "🎻"
	// Variable declarations can also be factored into blocks, like imports
	var (
		IsOk bool = true
		IsValid bool = false
	)

	// Variables will be initialised (where possible) based on their type
	// This is called a "zero value":
	// 0 for numeric types, false for the boolean, "" (empty string) for strings
	fmt.Println(i, c, python, java)
	// => 0 false false false ""

	fmt.Println(a, b, n, m, IsOk, IsValid)
	// => "a" "b" 123 "🎻" true false

	// A variable's type can be logged
	var v = 1337
	fmt.Printf("v is of type %T\n", v)

	// Constants are declared with the `const` keyword
	// Constants cannot be declared using the := syntax
	const Pi = 3.14
	fmt.Println("Happy", Pi, "Day")
}
