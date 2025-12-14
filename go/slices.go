package main

import (
	"fmt"
	"strings"
)

func slices() {
	printSliceHeader("GO SLICES TUTORIAL")

	// Section 1: What are Slices?
	printSliceSection("1. What are Slices?")
	fmt.Println("   Slices are dynamic, flexible views into arrays.")
	fmt.Println("   ✅ Can grow and shrink in size (unlike arrays)")
	fmt.Println("   ✅ Built on top of arrays")
	fmt.Println("   ✅ Have both length (len) and capacity (cap)\n")

	// Section 2: Creating Slices - Literal
	printSliceSection("2. Creating Slices - Slice Literal")
	slice1 := []int{1, 2, 3, 4, 5}
	fmt.Printf("   slice1 := []int{1, 2, 3, 4, 5}\n")
	fmt.Printf("   → Value: %v (type: %T)\n", slice1, slice1)
	fmt.Printf("   → Length: %d (number of elements)\n", len(slice1))
	fmt.Printf("   → Capacity: %d (underlying array size)\n\n", cap(slice1))

	// Section 3: Length vs Capacity
	printSliceSection("3. Understanding Length vs Capacity")
	fmt.Println("   ┌──────────┬─────────────────────────────────────────┐")
	fmt.Println("   │ Property │ Description                             │")
	fmt.Println("   ├──────────┼─────────────────────────────────────────┤")
	fmt.Println("   │ len()    │ Number of elements currently in slice   │")
	fmt.Println("   │ cap()    │ Max elements before reallocation needed │")
	fmt.Println("   └──────────┴─────────────────────────────────────────┘\n")

	// Section 4: Creating Slices from Arrays
	printSliceSection("4. Creating Slices from Arrays")
	var array1 = [5]int{10, 20, 30, 40, 50}
	slice2 := array1[1:4] // elements at index 1, 2, 3
	fmt.Printf("   array1 := [5]int{10, 20, 30, 40, 50}\n")
	fmt.Printf("   slice2 := array1[1:4]  // [start:end] (end is exclusive)\n")
	fmt.Printf("   → Value: %v\n", slice2)
	fmt.Printf("   → Length: %d (elements from index 1 to 3)\n", len(slice2))
	fmt.Printf("   → Capacity: %d (from index 1 to end of array)\n", cap(slice2))
	fmt.Printf("   💡 Capacity = %d because slice can grow to array end\n\n", cap(slice2))

	// Section 5: Slice Syntax Variations
	printSliceSection("5. Slice Syntax Variations")
	numbers := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Printf("   numbers := [6]int{1, 2, 3, 4, 5, 6}\n\n")
	fmt.Printf("   numbers[2:5]  → %v (index 2 to 4)\n", numbers[2:5])
	fmt.Printf("   numbers[:3]   → %v (start to index 2)\n", numbers[:3])
	fmt.Printf("   numbers[3:]   → %v (index 3 to end)\n", numbers[3:])
	fmt.Printf("   numbers[:]    → %v (entire array)\n\n", numbers[:])

	// Section 6: Creating Slices with make()
	printSliceSection("6. Creating Slices with make()")
	slice3 := make([]int, 5)
	fmt.Printf("   slice3 := make([]int, 5)  // length = 5, capacity = 5\n")
	fmt.Printf("   → Value: %v (zero values)\n", slice3)
	fmt.Printf("   → Length: %d, Capacity: %d\n\n", len(slice3), cap(slice3))

	slice4 := make([]int, 5, 10)
	fmt.Printf("   slice4 := make([]int, 5, 10)  // length = 5, capacity = 10\n")
	fmt.Printf("   → Value: %v\n", slice4)
	fmt.Printf("   → Length: %d, Capacity: %d\n", len(slice4), cap(slice4))
	fmt.Printf("   💡 Pre-allocating capacity improves performance!\n\n")

	// Section 7: Accessing Slice Elements
	printSliceSection("7. Accessing Slice Elements")
	fruits := []string{"Apple", "Banana", "Cherry", "Date", "Elderberry"}
	fmt.Printf("   fruits := []string{\"Apple\", \"Banana\", \"Cherry\", \"Date\", \"Elderberry\"}\n")
	fmt.Printf("   → fruits[0]: %s (first element)\n", fruits[0])
	fmt.Printf("   → fruits[2]: %s\n", fruits[2])
	fmt.Printf("   → fruits[4]: %s (last element)\n\n", fruits[4])

	// Section 8: Modifying Slice Elements
	printSliceSection("8. Modifying Slice Elements")
	slice5 := []int{1, 2, 3, 4, 5}
	fmt.Printf("   Original: %v\n", slice5)
	slice5[0] = 100
	slice5[4] = 500
	fmt.Printf("   slice5[0] = 100\n")
	fmt.Printf("   slice5[4] = 500\n")
	fmt.Printf("   → Modified: %v\n\n", slice5)

	// Section 9: Appending Single Element
	printSliceSection("9. Appending Elements to Slice")
	slice6 := []int{1, 2, 3}
	fmt.Printf("   Original: %v (len=%d, cap=%d)\n", slice6, len(slice6), cap(slice6))
	slice6 = append(slice6, 4)
	fmt.Printf("   slice6 = append(slice6, 4)\n")
	fmt.Printf("   → Result: %v (len=%d, cap=%d)\n\n", slice6, len(slice6), cap(slice6))

	// Section 10: Appending Multiple Elements
	printSliceSection("10. Appending Multiple Elements")
	slice7 := []int{1, 2, 3, 4, 5}
	fmt.Printf("   Original: %v (len=%d, cap=%d)\n", slice7, len(slice7), cap(slice7))
	slice7 = append(slice7, 6, 7, 8)
	fmt.Printf("   slice7 = append(slice7, 6, 7, 8)\n")
	fmt.Printf("   → Result: %v (len=%d, cap=%d)\n", slice7, len(slice7), cap(slice7))
	fmt.Printf("   💡 Capacity doubled automatically when needed!\n\n")

	// Section 11: Appending Another Slice
	printSliceSection("11. Appending Another Slice")
	slice8 := []int{1, 2, 3}
	slice9 := []int{4, 5, 6}
	fmt.Printf("   slice8 := %v\n", slice8)
	fmt.Printf("   slice9 := %v\n", slice9)
	slice8 = append(slice8, slice9...)
	fmt.Printf("   slice8 = append(slice8, slice9...)  // ... unpacks slice\n")
	fmt.Printf("   → Result: %v (len=%d, cap=%d)\n\n", slice8, len(slice8), cap(slice8))

	// Section 12: How Capacity Grows
	printSliceSection("12. How Capacity Grows Dynamically")
	demo := []int{1}
	fmt.Printf("   Demonstrating automatic capacity growth:\n\n")
	fmt.Printf("   Initial: %v (len=%d, cap=%d)\n", demo, len(demo), cap(demo))

	for i := 2; i <= 5; i++ {
		demo = append(demo, i)
		fmt.Printf("   After append(%d): len=%d, cap=%d\n", i, len(demo), cap(demo))
	}
	fmt.Printf("\n   💡 Go doubles capacity when reallocation is needed!\n\n")

	// Section 13: Copying Slices
	printSliceSection("13. Copying Slices for Memory Efficiency")
	largeSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	neededPart := largeSlice[2:5]
	fmt.Printf("   largeSlice := %v\n", largeSlice)
	fmt.Printf("   neededPart := largeSlice[2:5]\n")
	fmt.Printf("   → Value: %v (len=%d, cap=%d)\n", neededPart, len(neededPart), cap(neededPart))
	fmt.Printf("   ⚠️  neededPart still references the entire largeSlice!\n\n")

	// Create independent copy
	independentCopy := make([]int, 3)
	copy(independentCopy, neededPart)
	fmt.Printf("   independentCopy := make([]int, 3)\n")
	fmt.Printf("   copy(independentCopy, neededPart)\n")
	fmt.Printf("   → Value: %v (len=%d, cap=%d)\n", independentCopy, len(independentCopy), cap(independentCopy))
	fmt.Printf("   ✅ Now independent! Original can be garbage collected.\n\n")

	// Section 14: Slices are Reference Types
	printSliceSection("14. Slices are Reference Types")
	original := []int{1, 2, 3, 4, 5}
	reference := original
	reference[0] = 999
	fmt.Printf("   original := []int{1, 2, 3, 4, 5}\n")
	fmt.Printf("   reference := original\n")
	fmt.Printf("   reference[0] = 999\n")
	fmt.Printf("   → original: %v (changed!)\n", original)
	fmt.Printf("   → reference: %v\n", reference)
	fmt.Printf("   ⚠️  Both point to the same underlying array!\n\n")

	// Section 15: Iterating Over Slices
	printSliceSection("15. Iterating Over Slices")
	colors := []string{"Red", "Green", "Blue"}
	fmt.Printf("   colors := []string{\"Red\", \"Green\", \"Blue\"}\n\n")
	fmt.Printf("   Using for-range loop:\n")
	for index, value := range colors {
		fmt.Printf("   → Index %d: %s\n", index, value)
	}
	fmt.Printf("\n   Using range with value only:\n")
	for _, color := range colors {
		fmt.Printf("   → %s\n", color)
	}
	fmt.Println()

	// Section 16: Nil Slices
	printSliceSection("16. Nil Slices vs Empty Slices")
	var nilSlice []int
	emptySlice := []int{}
	fmt.Printf("   var nilSlice []int\n")
	fmt.Printf("   → Value: %v, len=%d, cap=%d, nil=%t\n",
		nilSlice, len(nilSlice), cap(nilSlice), nilSlice == nil)
	fmt.Printf("   emptySlice := []int{}\n")
	fmt.Printf("   → Value: %v, len=%d, cap=%d, nil=%t\n\n",
		emptySlice, len(emptySlice), cap(emptySlice), emptySlice == nil)

	printSliceFooter()
}

func printSliceHeader(title string) {
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Printf("  %s\n", title)
	fmt.Println(strings.Repeat("=", 60) + "\n")
}

func printSliceSection(title string) {
	fmt.Printf("┌─ %s\n", title)
	fmt.Println("│")
}

func printSliceFooter() {
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println("  ✅ Tutorial Complete!")
	fmt.Println("  💡 Key Takeaways:")
	fmt.Println("     • Slices are dynamic and flexible (unlike fixed arrays)")
	fmt.Println("     • Use make() to pre-allocate capacity for performance")
	fmt.Println("     • Use copy() to create independent slices")
	fmt.Println("     • Slices are reference types - modifications affect all refs")
	fmt.Println(strings.Repeat("=", 60) + "\n")
}
