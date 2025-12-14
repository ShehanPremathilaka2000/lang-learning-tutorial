package main

import (
	"fmt"
	"strings"
)

func arrays() {
	printArrayHeader("GO ARRAYS TUTORIAL")

	// Section 1: What are Arrays?
	printArraySection("1. What are Arrays?")
	fmt.Println("   Arrays are fixed-size collections of elements of the same type.")
	fmt.Println("   ✅ Size is part of the type (cannot be changed after creation)")
	fmt.Println("   ✅ Elements are stored in contiguous memory")
	fmt.Println("   ✅ Zero-indexed (first element is at index 0)\n")

	// Section 2: Array Declaration with Size
	printArraySection("2. Array Declaration with Explicit Size")
	var arr1 = [3]int{1, 2, 3}
	arr2 := [5]string{"Go", "is", "awesome", "and", "fun"}
	fmt.Printf("   var arr1 = [3]int{1, 2, 3}\n")
	fmt.Printf("   → Value: %v (type: %T)\n", arr1, arr1)
	fmt.Printf("   → Size: %d elements\n", len(arr1))
	fmt.Printf("   arr2 := [5]string{\"Go\", \"is\", \"awesome\", \"and\", \"fun\"}\n")
	fmt.Printf("   → Value: %v\n", arr2)
	fmt.Printf("   → Size: %d elements\n\n", len(arr2))

	// Section 3: Array Declaration with Inferred Size
	printArraySection("3. Array Declaration with Inferred Size (...)")
	var arr3 = [...]int{10, 20, 30, 40}
	arr4 := [...]float64{3.14, 2.71, 1.41}
	fmt.Printf("   var arr3 = [...]int{10, 20, 30, 40}\n")
	fmt.Printf("   → Value: %v (type: %T)\n", arr3, arr3)
	fmt.Printf("   → Size: %d (inferred from elements)\n", len(arr3))
	fmt.Printf("   arr4 := [...]float64{3.14, 2.71, 1.41}\n")
	fmt.Printf("   → Value: %v\n", arr4)
	fmt.Printf("   → Size: %d (inferred)\n\n", len(arr4))

	// Section 4: Accessing Array Elements
	printArraySection("4. Accessing Array Elements")
	fruits := [4]string{"Apple", "Banana", "Cherry", "Date"}
	fmt.Printf("   fruits := [4]string{\"Apple\", \"Banana\", \"Cherry\", \"Date\"}\n")
	fmt.Printf("   → fruits[0]: %s (first element)\n", fruits[0])
	fmt.Printf("   → fruits[1]: %s\n", fruits[1])
	fmt.Printf("   → fruits[2]: %s\n", fruits[2])
	fmt.Printf("   → fruits[3]: %s (last element)\n\n", fruits[3])

	// Section 5: Modifying Array Elements
	printArraySection("5. Modifying Array Elements")
	numbers := [3]int{1, 2, 3}
	fmt.Printf("   Original: %v\n", numbers)
	numbers[0] = 100
	numbers[2] = 300
	fmt.Printf("   numbers[0] = 100\n")
	fmt.Printf("   numbers[2] = 300\n")
	fmt.Printf("   → Modified: %v\n\n", numbers)

	// Section 6: Array Length
	printArraySection("6. Array Length")
	arr5 := [7]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("   arr5 := [7]int{1, 2, 3, 4, 5, 6, 7}\n")
	fmt.Printf("   → len(arr5): %d\n", len(arr5))
	fmt.Printf("   ⚠️  Note: Array size is fixed and part of its type!\n\n")

	// Section 7: Default Values (Zero Values)
	printArraySection("7. Default Values (Zero Values)")
	var intArray [3]int
	var stringArray [2]string
	var boolArray [4]bool
	fmt.Printf("   var intArray [3]int     // not initialized\n")
	fmt.Printf("   → Value: %v (zeros)\n", intArray)
	fmt.Printf("   var stringArray [2]string\n")
	fmt.Printf("   → Value: %v (empty strings)\n", stringArray)
	fmt.Printf("   var boolArray [4]bool\n")
	fmt.Printf("   → Value: %v (false values)\n\n", boolArray)

	// Section 8: Empty Initialization
	printArraySection("8. Empty Initialization")
	arr6 := [3]int{}
	arr7 := [4]string{}
	fmt.Printf("   arr6 := [3]int{}\n")
	fmt.Printf("   → Value: %v (all zeros)\n", arr6)
	fmt.Printf("   arr7 := [4]string{}\n")
	fmt.Printf("   → Value: %v (all empty strings)\n\n", arr7)

	// Section 9: Partial Initialization
	printArraySection("9. Partial Initialization")
	arr8 := [5]int{1, 2}
	fmt.Printf("   arr8 := [5]int{1, 2}  // only first 2 elements\n")
	fmt.Printf("   → Value: %v\n", arr8)
	fmt.Printf("   → Remaining elements are zero values\n\n")

	// Section 10: Index-Based Initialization
	printArraySection("10. Index-Based Initialization")
	arr9 := [5]int{1: 10, 3: 30}
	arr10 := [4]string{0: "first", 3: "last"}
	fmt.Printf("   arr9 := [5]int{1: 10, 3: 30}\n")
	fmt.Printf("   → Value: %v\n", arr9)
	fmt.Printf("   → Index 1 = 10, Index 3 = 30, others = 0\n")
	fmt.Printf("   arr10 := [4]string{0: \"first\", 3: \"last\"}\n")
	fmt.Printf("   → Value: %v\n\n", arr10)

	// Section 11: Iterating Over Arrays
	printArraySection("11. Iterating Over Arrays")
	colors := [3]string{"Red", "Green", "Blue"}
	fmt.Printf("   colors := [3]string{\"Red\", \"Green\", \"Blue\"}\n\n")
	fmt.Printf("   Using for loop with index:\n")
	for i := 0; i < len(colors); i++ {
		fmt.Printf("   → colors[%d] = %s\n", i, colors[i])
	}
	fmt.Printf("\n   Using for-range loop:\n")
	for index, value := range colors {
		fmt.Printf("   → Index %d: %s\n", index, value)
	}
	fmt.Println()

	// Section 12: Multi-Dimensional Arrays
	printArraySection("12. Multi-Dimensional Arrays (2D Arrays)")
	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Printf("   matrix := [2][3]int{\n")
	fmt.Printf("       {1, 2, 3},\n")
	fmt.Printf("       {4, 5, 6},\n")
	fmt.Printf("   }\n")
	fmt.Printf("   → Full matrix: %v\n", matrix)
	fmt.Printf("   → matrix[0][0]: %d (first element)\n", matrix[0][0])
	fmt.Printf("   → matrix[1][2]: %d (last element)\n\n", matrix[1][2])

	// Section 13: Array Comparison
	printArraySection("13. Array Comparison")
	a := [3]int{1, 2, 3}
	b := [3]int{1, 2, 3}
	c := [3]int{1, 2, 4}
	fmt.Printf("   a := [3]int{1, 2, 3}\n")
	fmt.Printf("   b := [3]int{1, 2, 3}\n")
	fmt.Printf("   c := [3]int{1, 2, 4}\n")
	fmt.Printf("   → a == b: %t (same values)\n", a == b)
	fmt.Printf("   → a == c: %t (different values)\n\n", a == c)

	printArrayFooter()
}

func printArrayHeader(title string) {
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Printf("  %s\n", title)
	fmt.Println(strings.Repeat("=", 60) + "\n")
}

func printArraySection(title string) {
	fmt.Printf("┌─ %s\n", title)
	fmt.Println("│")
}

func printArrayFooter() {
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println("  ✅ Tutorial Complete!")
	fmt.Println("  💡 Key Takeaway: Arrays have fixed size. For dynamic size,")
	fmt.Println("     use slices instead! Arrays are useful when you know the")
	fmt.Println("     exact number of elements at compile time.")
	fmt.Println(strings.Repeat("=", 60) + "\n")
}
