package main

import (
	"fmt"
	"strings"
)

func operators() {
	printOpHeader("GO OPERATORS TUTORIAL")

	// Section 1: Arithmetic Operators
	printOpSection("1. Arithmetic Operators")
	a, b := 15, 4
	fmt.Printf("   a := 15, b := 4\n\n")
	fmt.Printf("   Addition:       a + b = %d\n", a+b)
	fmt.Printf("   Subtraction:    a - b = %d\n", a-b)
	fmt.Printf("   Multiplication: a * b = %d\n", a*b)
	fmt.Printf("   Division:       a / b = %d (integer division)\n", a/b)
	fmt.Printf("   Modulus:        a %% b = %d (remainder)\n\n", a%b)

	// Float division
	x, y := 15.0, 4.0
	fmt.Printf("   x := 15.0, y := 4.0\n")
	fmt.Printf("   Division:       x / y = %.2f (float division)\n\n", x/y)

	// Section 2: Assignment Operators
	printOpSection("2. Assignment Operators")
	num := 10
	fmt.Printf("   num := 10\n\n")

	num += 5
	fmt.Printf("   num += 5   → num = %d (same as num = num + 5)\n", num)
	num -= 3
	fmt.Printf("   num -= 3   → num = %d (same as num = num - 3)\n", num)
	num *= 2
	fmt.Printf("   num *= 2   → num = %d (same as num = num * 2)\n", num)
	num /= 4
	fmt.Printf("   num /= 4   → num = %d (same as num = num / 4)\n", num)
	num %= 5
	fmt.Printf("   num %%= 5   → num = %d (same as num = num %% 5)\n\n", num)

	// Section 3: Increment and Decrement
	printOpSection("3. Increment and Decrement Operators")
	counter := 5
	fmt.Printf("   counter := 5\n")
	counter++
	fmt.Printf("   counter++  → counter = %d (increment by 1)\n", counter)
	counter--
	fmt.Printf("   counter--  → counter = %d (decrement by 1)\n", counter)
	fmt.Printf("   ⚠️  Note: ++counter and --counter are NOT valid in Go!\n\n")

	// Section 4: Comparison Operators
	printOpSection("4. Comparison Operators (Return bool)")
	p, q := 10, 20
	fmt.Printf("   p := 10, q := 20\n\n")
	fmt.Printf("   Equal to:              p == q → %t\n", p == q)
	fmt.Printf("   Not equal to:          p != q → %t\n", p != q)
	fmt.Printf("   Greater than:          p > q  → %t\n", p > q)
	fmt.Printf("   Less than:             p < q  → %t\n", p < q)
	fmt.Printf("   Greater than or equal: p >= q → %t\n", p >= q)
	fmt.Printf("   Less than or equal:    p <= q → %t\n\n", p <= q)

	// Section 5: Logical Operators
	printOpSection("5. Logical Operators (Boolean Logic)")
	t, f := true, false
	fmt.Printf("   t := true, f := false\n\n")
	fmt.Printf("   Logical AND: t && t → %t (both must be true)\n", t && t)
	fmt.Printf("                t && f → %t\n", t && f)
	fmt.Printf("                f && f → %t\n\n", f && f)

	fmt.Printf("   Logical OR:  t || t → %t (at least one must be true)\n", t || t)
	fmt.Printf("                t || f → %t\n", t || f)
	fmt.Printf("                f || f → %t\n\n", f || f)

	fmt.Printf("   Logical NOT: !t → %t (negation)\n", !t)
	fmt.Printf("                !f → %t\n\n", !f)

	// Real-world example
	age := 25
	hasLicense := true
	fmt.Printf("   Real-world example:\n")
	fmt.Printf("   age := 25, hasLicense := true\n")
	fmt.Printf("   Can drive: (age >= 18) && hasLicense → %t\n\n", (age >= 18) && hasLicense)

	// Section 6: Bitwise Operators
	printOpSection("6. Bitwise Operators (Bit Manipulation)")
	m, n := 12, 10 // 12 = 1100, 10 = 1010 in binary
	fmt.Printf("   m := 12 (binary: 1100), n := 10 (binary: 1010)\n\n")
	fmt.Printf("   Bitwise AND: m & n  = %d (binary: %04b)\n", m&n, m&n)
	fmt.Printf("   Bitwise OR:  m | n  = %d (binary: %04b)\n", m|n, m|n)
	fmt.Printf("   Bitwise XOR: m ^ n  = %d (binary: %04b)\n", m^n, m^n)
	fmt.Printf("   Bitwise NOT: ^m     = %d (inverts all bits)\n\n", ^m)

	// Section 7: Bit Shift Operators
	printOpSection("7. Bit Shift Operators")
	val := 8 // 1000 in binary
	fmt.Printf("   val := 8 (binary: 1000)\n\n")
	fmt.Printf("   Left shift:  val << 1 = %d (binary: %05b) [multiply by 2]\n", val<<1, val<<1)
	fmt.Printf("   Left shift:  val << 2 = %d (binary: %06b) [multiply by 4]\n", val<<2, val<<2)
	fmt.Printf("   Right shift: val >> 1 = %d (binary: %03b) [divide by 2]\n", val>>1, val>>1)
	fmt.Printf("   Right shift: val >> 2 = %d (binary: %02b) [divide by 4]\n\n", val>>2, val>>2)

	// Section 8: Operator Precedence
	printOpSection("8. Operator Precedence (Order of Operations)")
	result := 2 + 3*4
	fmt.Printf("   2 + 3 * 4 = %d (multiplication first)\n", result)
	result = (2 + 3) * 4
	fmt.Printf("   (2 + 3) * 4 = %d (parentheses first)\n\n", result)

	fmt.Println("   Precedence (highest to lowest):")
	fmt.Println("   1. Parentheses:        ( )")
	fmt.Println("   2. Unary:              +, -, !, ^")
	fmt.Println("   3. Multiplicative:     *, /, %, <<, >>, &, &^")
	fmt.Println("   4. Additive:           +, -, |, ^")
	fmt.Println("   5. Comparison:         ==, !=, <, <=, >, >=")
	fmt.Println("   6. Logical AND:        &&")
	fmt.Println("   7. Logical OR:         ||\n")

	// Section 9: Compound Bitwise Assignment
	printOpSection("9. Compound Bitwise Assignment Operators")
	bits := 12
	fmt.Printf("   bits := 12 (binary: %04b)\n\n", bits)

	bits &= 10
	fmt.Printf("   bits &= 10  → bits = %d (binary: %04b)\n", bits, bits)
	bits = 12
	bits |= 10
	fmt.Printf("   bits |= 10  → bits = %d (binary: %04b)\n", bits, bits)
	bits = 12
	bits ^= 10
	fmt.Printf("   bits ^= 10  → bits = %d (binary: %04b)\n", bits, bits)
	bits = 8
	bits <<= 2
	fmt.Printf("   bits <<= 2  → bits = %d (binary: %06b)\n", bits, bits)
	bits = 8
	bits >>= 1
	fmt.Printf("   bits >>= 1  → bits = %d (binary: %03b)\n\n", bits, bits)

	// Section 10: Practical Examples
	printOpSection("10. Practical Examples")

	// Check if number is even
	number := 42
	isEven := number%2 == 0
	fmt.Printf("   Check if %d is even: %d %% 2 == 0 → %t\n", number, number, isEven)

	// Check if number is power of 2
	num2 := 16
	isPowerOf2 := (num2 > 0) && (num2&(num2-1)) == 0
	fmt.Printf("   Check if %d is power of 2: %t (using bitwise)\n", num2, isPowerOf2)

	// Swap two numbers using XOR
	c, d := 5, 10
	fmt.Printf("   Swap using XOR: c=%d, d=%d\n", c, d)
	c = c ^ d
	d = c ^ d
	c = c ^ d
	fmt.Printf("   After swap: c=%d, d=%d\n\n", c, d)

	printOpFooter()
}

func printOpHeader(title string) {
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Printf("  %s\n", title)
	fmt.Println(strings.Repeat("=", 60) + "\n")
}

func printOpSection(title string) {
	fmt.Printf("┌─ %s\n", title)
	fmt.Println("│")
}

func printOpFooter() {
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println("  ✅ Tutorial Complete!")
	fmt.Println("  💡 Key Takeaways:")
	fmt.Println("     • Use arithmetic operators for math calculations")
	fmt.Println("     • Use comparison operators for conditions")
	fmt.Println("     • Use logical operators to combine boolean expressions")
	fmt.Println("     • Use bitwise operators for low-level bit manipulation")
	fmt.Println("     • Remember operator precedence (use parentheses!)")
	fmt.Println(strings.Repeat("=", 60) + "\n")
}
