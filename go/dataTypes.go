package main

import (
	"fmt"
	"strings"
)

func dataTypes() {
	printDataHeader("GO DATA TYPES TUTORIAL")

	// Section 1: Boolean Type
	printDataSection("1. Boolean Type (bool)")
	var isActive bool = true
	var isComplete bool = false
	var defaultBool bool
	fmt.Printf("   var isActive bool = true\n")
	fmt.Printf("   → Value: %t (type: %T)\n", isActive, isActive)
	fmt.Printf("   var isComplete bool = false\n")
	fmt.Printf("   → Value: %t (type: %T)\n", isComplete, isComplete)
	fmt.Printf("   var defaultBool bool  // not initialized\n")
	fmt.Printf("   → Default Value: %t (zero value)\n\n", defaultBool)

	// Section 2: Integer Types
	printDataSection("2. Integer Types")
	var int8Val int8 = 127          // -128 to 127
	var int16Val int16 = 32767      // -32768 to 32767
	var int32Val int32 = 2147483647 // -2147483648 to 2147483647
	var int64Val int64 = 9223372036854775807
	var intVal int = 42 // platform dependent (32 or 64 bit)
	var defaultInt int

	fmt.Printf("   Signed Integers:\n")
	fmt.Printf("   → int8:  %d (range: -128 to 127)\n", int8Val)
	fmt.Printf("   → int16: %d (range: -32,768 to 32,767)\n", int16Val)
	fmt.Printf("   → int32: %d (range: -2.1B to 2.1B)\n", int32Val)
	fmt.Printf("   → int64: %d (range: very large)\n", int64Val)
	fmt.Printf("   → int:   %d (platform dependent, usually 64-bit)\n", intVal)
	fmt.Printf("   → Default: %d (zero value)\n\n", defaultInt)

	// Section 3: Unsigned Integer Types
	printDataSection("3. Unsigned Integer Types")
	var uint8Val uint8 = 255     // 0 to 255
	var uint16Val uint16 = 65535 // 0 to 65535
	var uint32Val uint32 = 4294967295
	var uintVal uint = 100
	var byteVal byte = 'A' // byte is alias for uint8

	fmt.Printf("   Unsigned Integers (only positive):\n")
	fmt.Printf("   → uint8:  %d (range: 0 to 255)\n", uint8Val)
	fmt.Printf("   → uint16: %d (range: 0 to 65,535)\n", uint16Val)
	fmt.Printf("   → uint32: %d (range: 0 to 4.3B)\n", uint32Val)
	fmt.Printf("   → uint:   %d (platform dependent)\n", uintVal)
	fmt.Printf("   → byte:   %d (alias for uint8, char: %c)\n\n", byteVal, byteVal)

	// Section 4: Floating Point Types
	printDataSection("4. Floating Point Types")
	var float32Val float32 = 3.14159
	var float64Val float64 = 3.141592653589793
	var defaultFloat float64

	fmt.Printf("   Decimal Numbers:\n")
	fmt.Printf("   → float32: %.5f (32-bit, ~7 decimal digits)\n", float32Val)
	fmt.Printf("   → float64: %.15f (64-bit, ~15 decimal digits)\n", float64Val)
	fmt.Printf("   → Default: %.1f (zero value)\n\n", defaultFloat)

	// Section 5: String Type
	printDataSection("5. String Type")
	var greeting string = "Hello, Go!"
	var multiline string = `This is a
multi-line string
using backticks`
	var emptyString string
	var runeVal rune = '世' // rune is alias for int32, represents Unicode

	fmt.Printf("   Text Data:\n")
	fmt.Printf("   → string: \"%s\" (type: %T)\n", greeting, greeting)
	fmt.Printf("   → Length: %d characters\n", len(greeting))
	fmt.Printf("   → Multi-line: %q\n", multiline)
	fmt.Printf("   → Default: \"%s\" (empty string)\n", emptyString)
	fmt.Printf("   → rune: %c (Unicode: U+%04X, value: %d)\n\n", runeVal, runeVal, runeVal)

	// Section 6: Complex Number Types
	printDataSection("6. Complex Number Types")
	var complex64Val complex64 = 1 + 2i
	var complex128Val complex128 = 3.14 + 2.71i

	fmt.Printf("   Complex Numbers (real + imaginary):\n")
	fmt.Printf("   → complex64:  %v (type: %T)\n", complex64Val, complex64Val)
	fmt.Printf("   → complex128: %v (type: %T)\n", complex128Val, complex128Val)
	fmt.Printf("   → Real part: %.2f, Imaginary part: %.2f\n\n",
		real(complex128Val), imag(complex128Val))

	// Section 7: Type Conversion
	printDataSection("7. Type Conversion")
	var intNum int = 42
	var floatNum float64 = float64(intNum)
	var stringNum string = fmt.Sprintf("%d", intNum)

	fmt.Printf("   Converting between types:\n")
	fmt.Printf("   → int to float64: %d → %.2f\n", intNum, floatNum)
	fmt.Printf("   → int to string:  %d → \"%s\"\n", intNum, stringNum)
	fmt.Printf("   ⚠️  Note: Go requires explicit type conversion!\n\n")

	// Section 8: Zero Values Summary
	printDataSection("8. Zero Values Summary")
	fmt.Println("   Default values when variables are declared but not initialized:")
	fmt.Println("   ┌──────────────┬──────────────────┐")
	fmt.Println("   │ Type         │ Zero Value       │")
	fmt.Println("   ├──────────────┼──────────────────┤")
	fmt.Println("   │ bool         │ false            │")
	fmt.Println("   │ int/uint     │ 0                │")
	fmt.Println("   │ float        │ 0.0              │")
	fmt.Println("   │ string       │ \"\" (empty)       │")
	fmt.Println("   │ complex      │ 0+0i             │")
	fmt.Println("   │ pointer      │ nil              │")
	fmt.Println("   └──────────────┴──────────────────┘\n")

	printDataFooter()
}

func printDataHeader(title string) {
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Printf("  %s\n", title)
	fmt.Println(strings.Repeat("=", 60) + "\n")
}

func printDataSection(title string) {
	fmt.Printf("┌─ %s\n", title)
	fmt.Println("│")
}

func printDataFooter() {
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println("  ✅ Tutorial Complete!")
	fmt.Println("  💡 Key Takeaway: Go is statically typed. Choose the right")
	fmt.Println("     type for your data. Use int for whole numbers, float64")
	fmt.Println("     for decimals, string for text, and bool for true/false.")
	fmt.Println(strings.Repeat("=", 60) + "\n")
}
