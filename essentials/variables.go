package essentials

import (
	"fmt"
)

// Constants in Go

const (
	PI       float32 = 3.14
	BASE_URL string  = "https://api.example.com"
)

// undefined const
const DENEME = 87987 // bu da sabit bir değer

func Sum(a int, b int) int {
	return a + b
}

func DiffrentVariables() {

	var ii, jj string = "Hello", "World"

	fmt.Print(ii, "\n", jj)

	var a int = 10
	var b float64 = 20.5
	var c string = "Hello"
	var d bool = true

	b = 55.56

	deneme := 1241 // biraz farklıymış, pek beğenmedim
	boolDeneme := false

	boolDeneme = true

	// formatting float
	fmt.Printf("Float: %.3f\n", b)
	fmt.Printf("String: %s\n", c)

	fmt.Println("boolDeneme:", boolDeneme)
	fmt.Println("Deneme:", deneme)
	fmt.Println("Integer:", a)
	fmt.Println("Float:", b)
	fmt.Println("String:", c)
	fmt.Println("Boolean:", d)

	fmt.Println("---------- Multiple Variables --------------")

	// bu da c sharp'dan çok farklı
	var e, f, g int = 1, 2, 3
	var h, i, j = 4.5, "Hello", true

	fmt.Println("e:", e)
	fmt.Println("f:", f)
	fmt.Println("g:", g)
	fmt.Println("h:", h)
	fmt.Println("i:", i)
	fmt.Println("j:", j)
	fmt.Println("Sum of e and f:", Sum(e, f))
	fmt.Println("Sum of e, f, and g:", Sum(Sum(e, f), g))

	// grouped variable declaration
	fmt.Println("---------- Grouped Variable Declaration --------------")
	var (
		x int     = 10
		y float64 = 20.5
		z string  = "Hello"
		w bool    = true
	)

	// print all
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
	fmt.Println("w:", w)

	//------------------------------------

	fmt.Printf("PI: %.2f\n", PI)
	fmt.Printf("Base URL: %s, Type: %T\n", BASE_URL, BASE_URL)
	fmt.Printf("Deneme: %d\n", DENEME)
}

// Verb	Description
// %v	Prints the value in the default format
// %#v	Prints the value in Go-syntax format
// %T	Prints the type of the value
// %%	Prints the % sign
func FormattingValues() {

	var a int = 10
	var b float64 = 20.5
	var c string = "Hello"
	var d bool = true

	fmt.Printf("Integer: %d\n", a)
	fmt.Printf("Float: %.2f\n", b)
	fmt.Printf("String: %s\n", c)
	fmt.Printf("Boolean: %t\n", d)

	fmt.Printf("Default format: %v\n", a)
	fmt.Printf("Go-syntax format: %#v\n", a)
	fmt.Printf("Type of a: %T\n", a)

	fmt.Printf("%% sign: %%\n") // prints the % sign
	fmt.Printf("Base URL: %s, Type: %T\n", BASE_URL, BASE_URL)
	fmt.Printf("Deneme: %d\n", DENEME)
}

// Integer Formatting Verbs
// The following verbs can be used with the integer data type:

// Verb	Description
// %b	Base 2
// %d	Base 10
// %+d	Base 10 and always show sign
// %o	Base 8
// %O	Base 8, with leading 0o
// %x	Base 16, lowercase
// %X	Base 16, uppercase
// %#x	Base 16, with leading 0x
// %4d	Pad with spaces (width 4, right justified)
// %-4d	Pad with spaces (width 4, left justified)
// %04d	Pad with zeroes (width 4

func IntegerFormatting() {
	fmt.Println("Integer Formatting:")
	fmt.Printf("Base 2: %b\n", 10)                                         // 1010
	fmt.Printf("Base 10: %d\n", 10)                                        // 10
	fmt.Printf("Base 8: %o\n", 10)                                         // 12
	fmt.Printf("Base 16 (lowercase): %x\n", 10)                            // a
	fmt.Printf("Base 16 (uppercase): %X\n", 10)                            // A
	fmt.Printf("Base 16 with leading 0x: %#x\n", 10)                       // 0xa
	fmt.Printf("Padded with spaces (width 4, right justified): %4d\n", 10) // "  10"
	fmt.Printf("Padded with spaces (width 4, left justified): %-4d\n", 10) // "10  "
	fmt.Printf("Padded with zeroes (width 4): %04d\n", 10)                 // "0010"
}


// String Formatting Verbs
// The following verbs can be used with the string data type:

// Verb	Description
// %s	Prints the value as plain string
// %q	Prints the value as a double-quoted string
// %8s	Prints the value as plain string (width 8, right justified)
// %-8s	Prints the value as plain string (width 8, left justified)
// %x	Prints the value as hex dump of byte values
// % x	Prints the value as hex dump with spaces
func StringFormatting() {
	fmt.Println("String Formatting:")
	fmt.Printf("Plain string: %s\n", "Hello")                          // Hello
	fmt.Printf("Double-quoted string: %q\n", "Hello")                  // "Hello"
	fmt.Printf("Width 8, right justified: %8s\n", "Hello")            // "   Hello"
	fmt.Printf("Width 8, left justified: %-8s\n", "Hello")            // "Hello   "
	fmt.Printf("Hex dump: %x\n", "Hello")                              // 48656c6c6f
	fmt.Printf("Hex dump with spaces: % x\n", "Hello")                 // 48 65 6c 6c 6f
}