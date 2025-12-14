package main

import (
	"fmt"
	"strings"
)

// Package-level constants (can be declared outside functions)
const GlobalConstant int = 100
const Pi = 3.14159
const (
	StatusActive   = "ACTIVE"
	StatusInactive = "INACTIVE"
	StatusPending  = "PENDING"
)

func constants() {
	printConstHeader("GO CONSTANTS TUTORIAL")

	// Section 1: What are Constants?
	printConstSection("1. What are Constants?")
	fmt.Println("   Constants are immutable values that cannot be changed after creation.")
	fmt.Println("   Unlike variables, constants are declared using the 'const' keyword.")
	fmt.Println("   ✅ Constants can be declared at package level (outside functions)\n")

	// Section 2: Package-Level Constants
	printConstSection("2. Package-Level Constants")
	fmt.Printf("   const GlobalConstant int = 100\n")
	fmt.Printf("   → Value: %d (type: %T)\n", GlobalConstant, GlobalConstant)
	fmt.Printf("   const Pi = 3.14159\n")
	fmt.Printf("   → Value: %v (type: %T)\n\n", Pi, Pi)

	// Section 3: Constant with Explicit Type
	printConstSection("3. Constant with Explicit Type")
	const typedConst int = 42
	fmt.Printf("   const typedConst int = 42\n")
	fmt.Printf("   → Value: %d\n", typedConst)
	fmt.Printf("   → Type: %T (explicitly typed)\n\n", typedConst)

	// Section 4: Constant with Type Inference
	printConstSection("4. Constant with Type Inference")
	const inferredConst = "Go is awesome!"
	fmt.Printf("   const inferredConst = \"Go is awesome!\"\n")
	fmt.Printf("   → Value: %s\n", inferredConst)
	fmt.Printf("   → Type: %T (inferred)\n\n", inferredConst)

	// Section 5: Grouped Constants Block
	printConstSection("5. Grouped Constants Block")
	const (
		MaxUsers       = 1000
		MinAge         = 18
		AppName        = "MyGoApp"
		Version string = "1.0.0"
		DebugEnabled   = true
	)
	fmt.Printf("   const (\n")
	fmt.Printf("       MaxUsers       = 1000\n")
	fmt.Printf("       MinAge         = 18\n")
	fmt.Printf("       AppName        = \"MyGoApp\"\n")
	fmt.Printf("       Version string = \"1.0.0\"\n")
	fmt.Printf("       DebugEnabled   = true\n")
	fmt.Printf("   )\n")
	fmt.Printf("   → MaxUsers: %v (type: %T)\n", MaxUsers, MaxUsers)
	fmt.Printf("   → MinAge: %v (type: %T)\n", MinAge, MinAge)
	fmt.Printf("   → AppName: %v (type: %T)\n", AppName, AppName)
	fmt.Printf("   → Version: %v (type: %T)\n", Version, Version)
	fmt.Printf("   → DebugEnabled: %v (type: %T)\n\n", DebugEnabled, DebugEnabled)

	// Section 6: Using Package-Level Constants
	printConstSection("6. Using Package-Level Constants")
	fmt.Printf("   Package-level constants defined at the top:\n")
	fmt.Printf("   → StatusActive: %s\n", StatusActive)
	fmt.Printf("   → StatusInactive: %s\n", StatusInactive)
	fmt.Printf("   → StatusPending: %s\n\n", StatusPending)

	// Section 7: Iota - Auto-incrementing Constants
	printConstSection("7. Iota - Auto-incrementing Constants")
	const (
		Sunday    = iota // 0
		Monday           // 1
		Tuesday          // 2
		Wednesday        // 3
		Thursday         // 4
		Friday           // 5
		Saturday         // 6
	)
	fmt.Printf("   const (\n")
	fmt.Printf("       Sunday = iota  // 0\n")
	fmt.Printf("       Monday         // 1\n")
	fmt.Printf("       Tuesday        // 2\n")
	fmt.Printf("       ...\n")
	fmt.Printf("   )\n")
	fmt.Printf("   → Sunday: %d\n", Sunday)
	fmt.Printf("   → Monday: %d\n", Monday)
	fmt.Printf("   → Tuesday: %d\n", Tuesday)
	fmt.Printf("   → Saturday: %d\n\n", Saturday)

	// Section 8: Iota with Expressions
	printConstSection("8. Iota with Expressions (Powers of 2)")
	const (
		_  = iota             // 0 (ignored with _)
		KB = 1 << (10 * iota) // 1 << 10 = 1024
		MB                    // 1 << 20 = 1048576
		GB                    // 1 << 30 = 1073741824
	)
	fmt.Printf("   const (\n")
	fmt.Printf("       _  = iota             // skip 0\n")
	fmt.Printf("       KB = 1 << (10 * iota) // 1024\n")
	fmt.Printf("       MB                    // 1048576\n")
	fmt.Printf("       GB                    // 1073741824\n")
	fmt.Printf("   )\n")
	fmt.Printf("   → KB: %d bytes\n", KB)
	fmt.Printf("   → MB: %d bytes\n", MB)
	fmt.Printf("   → GB: %d bytes\n\n", GB)

	// Section 9: Constants vs Variables
	printConstSection("9. Constants vs Variables")
	fmt.Println("   ┌─────────────────┬──────────────┬──────────────┐")
	fmt.Println("   │ Feature         │ const        │ var          │")
	fmt.Println("   ├─────────────────┼──────────────┼──────────────┤")
	fmt.Println("   │ Mutability      │ Immutable ✅ │ Mutable      │")
	fmt.Println("   │ Package-level   │ Yes ✅       │ Yes ✅       │")
	fmt.Println("   │ := syntax       │ No ❌        │ Yes (func)   │")
	fmt.Println("   └─────────────────┴──────────────┴──────────────┘\n")

	printConstFooter()
}

func printConstHeader(title string) {
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Printf("  %s\n", title)
	fmt.Println(strings.Repeat("=", 60) + "\n")
}

func printConstSection(title string) {
	fmt.Printf("┌─ %s\n", title)
	fmt.Println("│")
}

func printConstFooter() {
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println("  ✅ Tutorial Complete!")
	fmt.Println("  💡 Key Takeaway: Use constants for values that never change,")
	fmt.Println("     like configuration values, status codes, or mathematical")
	fmt.Println("     constants. Use 'iota' for auto-incrementing enumerations.")
	fmt.Println(strings.Repeat("=", 60) + "\n")
}
