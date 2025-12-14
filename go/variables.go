package main

import (
	"fmt"
	"strings"
)

func variables() {
	printHeader("GO VARIABLES TUTORIAL")

	// Section 1: Basic Variable Declaration with Type
	printSection("1. Variable Declaration with Explicit Type")
	var variable1 string = "Hello, Go!"
	fmt.Printf("   var variable1 string = \"Hello, Go!\"\n")
	fmt.Printf("   → Value: %s\n", variable1)
	fmt.Printf("   → Type: %T\n\n", variable1)

	// Section 2: Type Inference with var
	printSection("2. Type Inference with 'var' Keyword")
	var variable2 = "Type inferred automatically"
	fmt.Printf("   var variable2 = \"Type inferred automatically\"\n")
	fmt.Printf("   → Value: %s\n", variable2)
	fmt.Printf("   → Type: %T (inferred)\n\n", variable2)

	// Section 3: Short Declaration (:=)
	printSection("3. Short Declaration Operator (:=)")
	variable3 := "Quick and concise!"
	fmt.Printf("   variable3 := \"Quick and concise!\"\n")
	fmt.Printf("   → Value: %s\n", variable3)
	fmt.Printf("   → Type: %T (inferred)\n", variable3)
	fmt.Printf("   ⚠️  Note: := can only be used inside functions\n\n")

	// Section 4: Default Values
	printSection("4. Default Values (Zero Values)")
	var variable4 string
	var number int
	var boolean bool
	fmt.Printf("   var variable4 string  // not initialized\n")
	fmt.Printf("   → Value: \"%s\" (empty string)\n", variable4)
	fmt.Printf("   var number int\n")
	fmt.Printf("   → Value: %d (zero)\n", number)
	fmt.Printf("   var boolean bool\n")
	fmt.Printf("   → Value: %t (false)\n\n", boolean)

	// Section 5: Delayed Assignment
	printSection("5. Declare First, Assign Later")
	var variable5 string
	variable5 = "Assigned after declaration"
	fmt.Printf("   var variable5 string\n")
	fmt.Printf("   variable5 = \"Assigned after declaration\"\n")
	fmt.Printf("   → Value: %s\n\n", variable5)

	// Section 6: Multiple Variables (Same Type)
	printSection("6. Multiple Variables of Same Type")
	var variable6, variable7 string = "First", "Second"
	fmt.Printf("   var variable6, variable7 string = \"First\", \"Second\"\n")
	fmt.Printf("   → variable6: %s\n", variable6)
	fmt.Printf("   → variable7: %s\n\n", variable7)

	// Section 7: Multiple Variables (Different Types)
	printSection("7. Multiple Variables of Different Types")
	var variable8, variable9 = 42, "Mixed types!"
	fmt.Printf("   var variable8, variable9 = 42, \"Mixed types!\"\n")
	fmt.Printf("   → variable8: %v (type: %T)\n", variable8, variable8)
	fmt.Printf("   → variable9: %v (type: %T)\n\n", variable9, variable9)

	// Section 8: Short Declaration with Multiple Variables
	printSection("8. Short Declaration with Multiple Variables")
	variable10, variable11 := 100, "Short form"
	fmt.Printf("   variable10, variable11 := 100, \"Short form\"\n")
	fmt.Printf("   → variable10: %v (type: %T)\n", variable10, variable10)
	fmt.Printf("   → variable11: %v (type: %T)\n\n", variable11, variable11)

	// Section 9: Grouped Variable Declaration
	printSection("9. Grouped Variable Declaration Block")
	var (
		variable12        = 999
		variable13        = "Grouped declaration"
		variable14 string = "With explicit type"
	)
	fmt.Printf("   var (\n")
	fmt.Printf("       variable12 = 999\n")
	fmt.Printf("       variable13 = \"Grouped declaration\"\n")
	fmt.Printf("       variable14 string = \"With explicit type\"\n")
	fmt.Printf("   )\n")
	fmt.Printf("   → variable12: %v (type: %T)\n", variable12, variable12)
	fmt.Printf("   → variable13: %v (type: %T)\n", variable13, variable13)
	fmt.Printf("   → variable14: %v (type: %T)\n\n", variable14, variable14)

	// Summary
	printFooter()
}

func printHeader(title string) {
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Printf("  %s\n", title)
	fmt.Println(strings.Repeat("=", 60) + "\n")
}

func printSection(title string) {
	fmt.Printf("┌─ %s\n", title)
	fmt.Println("│")
}

func printFooter() {
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println("  ✅ Tutorial Complete!")
	fmt.Println("  💡 Key Takeaway: Use := for quick declarations inside functions,")
	fmt.Println("     and var for package-level or when you need explicit types.")
	fmt.Println(strings.Repeat("=", 60) + "\n")
}
