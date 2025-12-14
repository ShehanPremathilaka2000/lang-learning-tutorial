package main

import (
	"fmt"
	"strings"
)

func defers() {
	printDeferHeader("GO DEFER STATEMENT TUTORIAL")

	// Section 1: What is defer?
	printDeferSection("1. What is the defer Statement?")
	fmt.Println("   defer schedules a function call to execute just BEFORE")
	fmt.Println("   the surrounding function returns.")
	fmt.Println("   ✅ Guarantees cleanup code runs")
	fmt.Println("   ✅ Prevents resource leaks")
	fmt.Println("   ✅ Executes even if function panics\n")

	// Section 2: Core Rules of defer
	printDeferSection("2. Core Rules of defer")
	fmt.Println("   ┌──────────────────────┬─────────────────────────────────┐")
	fmt.Println("   │ Rule                 │ Description                     │")
	fmt.Println("   ├──────────────────────┼─────────────────────────────────┤")
	fmt.Println("   │ Execution Time       │ Just before function returns    │")
	fmt.Println("   │ Argument Evaluation  │ Evaluated IMMEDIATELY           │")
	fmt.Println("   │ Execution Order      │ LIFO (Last-In, First-Out)       │")
	fmt.Println("   └──────────────────────┴─────────────────────────────────┘\n")

	// Section 3: Basic defer Example
	printDeferSection("3. Basic defer Example")
	fmt.Printf("   func example() {\n")
	fmt.Printf("       fmt.Println(\"1. Start\")\n")
	fmt.Printf("       defer fmt.Println(\"3. Deferred (runs last)\")\n")
	fmt.Printf("       fmt.Println(\"2. Middle\")\n")
	fmt.Printf("   }\n\n")
	fmt.Printf("   Output:\n   ")
	basicDeferExample()
	fmt.Println()

	// Section 4: Argument Evaluation vs Function Execution
	printDeferSection("4. Argument Evaluation vs Function Execution")
	fmt.Println("   ┌─────────────────────┬──────────────────────────────────┐")
	fmt.Println("   │ Phase               │ When it Happens                  │")
	fmt.Println("   ├─────────────────────┼──────────────────────────────────┤")
	fmt.Println("   │ Argument Evaluation │ IMMEDIATELY when defer is hit    │")
	fmt.Println("   │ Function Execution  │ Just BEFORE function returns     │")
	fmt.Println("   └─────────────────────┴──────────────────────────────────┘\n")

	fmt.Printf("   func argumentEvaluation() {\n")
	fmt.Printf("       i := 1\n")
	fmt.Printf("       defer fmt.Println(\"Result:\", i)  // i evaluated NOW\n")
	fmt.Printf("       i = 2  // This does NOT affect deferred argument\n")
	fmt.Printf("       fmt.Println(\"i is now:\", i)\n")
	fmt.Printf("   }\n\n")
	fmt.Printf("   Output:\n   ")
	argumentEvaluationExample()
	fmt.Println()

	// Section 5: LIFO (Stack) Execution Order
	printDeferSection("5. LIFO (Stack) Execution Order")
	fmt.Printf("   func lifoExample() {\n")
	fmt.Printf("       defer fmt.Println(\"First\")   // Scheduled 1st, Executes 3rd\n")
	fmt.Printf("       defer fmt.Println(\"Second\")  // Scheduled 2nd, Executes 2nd\n")
	fmt.Printf("       defer fmt.Println(\"Third\")   // Scheduled 3rd, Executes 1st\n")
	fmt.Printf("       fmt.Println(\"Main\")\n")
	fmt.Printf("   }\n\n")
	fmt.Printf("   Output:\n   ")
	lifoExample()
	fmt.Println()

	// Section 6: Common Use Case - Resource Cleanup
	printDeferSection("6. Common Use Case - Resource Cleanup")
	fmt.Printf("   Typical pattern for file operations:\n\n")
	fmt.Printf("   func processFile(filename string) error {\n")
	fmt.Printf("       file, err := os.Open(filename)\n")
	fmt.Printf("       if err != nil {\n")
	fmt.Printf("           return err\n")
	fmt.Printf("       }\n")
	fmt.Printf("       defer file.Close()  // Guaranteed to run!\n\n")
	fmt.Printf("       // ... process file ...\n")
	fmt.Printf("       return nil\n")
	fmt.Printf("   }\n\n")
	fmt.Printf("   💡 defer ensures file.Close() runs even if errors occur!\n\n")

	// Section 7: Multiple defer Statements
	printDeferSection("7. Multiple defer Statements (Stack Behavior)")
	fmt.Printf("   Opening and closing resources in reverse order:\n\n")
	fmt.Printf("   func multipleResources() {\n")
	fmt.Printf("       defer fmt.Println(\"Close Database\")\n")
	fmt.Printf("       defer fmt.Println(\"Close File\")\n")
	fmt.Printf("       defer fmt.Println(\"Close Connection\")\n")
	fmt.Printf("       fmt.Println(\"Opening resources...\")\n")
	fmt.Printf("   }\n\n")
	fmt.Printf("   Output:\n   ")
	multipleResourcesExample()
	fmt.Printf("   💡 Resources close in reverse order (LIFO)!\n\n")

	// Section 8: defer with Anonymous Functions
	printDeferSection("8. defer with Anonymous Functions")
	fmt.Printf("   func anonymousDefer() {\n")
	fmt.Printf("       x := 10\n")
	fmt.Printf("       defer func() {\n")
	fmt.Printf("           fmt.Println(\"x is:\", x)  // Captures x by reference\n")
	fmt.Printf("       }()\n")
	fmt.Printf("       x = 20  // This WILL affect the deferred function\n")
	fmt.Printf("   }\n\n")
	fmt.Printf("   Output:\n   ")
	anonymousDeferExample()
	fmt.Printf("   💡 Anonymous functions capture variables by reference!\n\n")

	// Section 9: defer in Loops
	printDeferSection("9. defer in Loops (Be Careful!)")
	fmt.Printf("   ⚠️  defer in loops can cause issues:\n\n")
	fmt.Printf("   func deferInLoop() {\n")
	fmt.Printf("       for i := 1; i <= 3; i++ {\n")
	fmt.Printf("           defer fmt.Println(i)\n")
	fmt.Printf("       }\n")
	fmt.Printf("   }\n\n")
	fmt.Printf("   Output (reverse order):\n   ")
	deferInLoopExample()
	fmt.Printf("   ⚠️  All defers execute at function end, not loop end!\n\n")

	// Section 10: defer and Return Values
	printDeferSection("10. defer and Named Return Values")
	fmt.Printf("   defer can modify named return values:\n\n")
	fmt.Printf("   func increment() (result int) {\n")
	fmt.Printf("       defer func() { result++ }()\n")
	fmt.Printf("       return 5\n")
	fmt.Printf("   }\n\n")
	fmt.Printf("   Calling increment():\n")
	fmt.Printf("   → Result: %d (defer modified it!)\n\n", incrementExample())

	// Section 11: defer and Panics
	printDeferSection("11. defer and Panics")
	fmt.Printf("   defer runs even if function panics:\n\n")
	fmt.Printf("   func panicExample() {\n")
	fmt.Printf("       defer fmt.Println(\"Cleanup runs even on panic!\")\n")
	fmt.Printf("       panic(\"Something went wrong\")\n")
	fmt.Printf("   }\n\n")
	fmt.Printf("   💡 This enables cleanup during crashes!\n\n")

	// Section 12: Practical Example - Timer
	printDeferSection("12. Practical Example - Measuring Execution Time")
	fmt.Printf("   func measureTime() {\n")
	fmt.Printf("       start := time.Now()\n")
	fmt.Printf("       defer func() {\n")
	fmt.Printf("           fmt.Println(\"Took:\", time.Since(start))\n")
	fmt.Printf("       }()\n")
	fmt.Printf("       // ... function logic ...\n")
	fmt.Printf("   }\n\n")

	// Section 13: Practical Example - Lock/Unlock
	printDeferSection("13. Practical Example - Mutex Lock/Unlock")
	fmt.Printf("   func safeOperation() {\n")
	fmt.Printf("       mu.Lock()\n")
	fmt.Printf("       defer mu.Unlock()  // Guaranteed unlock!\n\n")
	fmt.Printf("       // ... critical section ...\n")
	fmt.Printf("   }\n\n")
	fmt.Printf("   💡 Prevents deadlocks from forgetting to unlock!\n\n")

	// Section 14: Common Patterns
	printDeferSection("14. Common defer Patterns")
	fmt.Println("   1. File Operations:")
	fmt.Println("      defer file.Close()")
	fmt.Println()
	fmt.Println("   2. Database Connections:")
	fmt.Println("      defer db.Close()")
	fmt.Println()
	fmt.Println("   3. HTTP Response Bodies:")
	fmt.Println("      defer resp.Body.Close()")
	fmt.Println()
	fmt.Println("   4. Mutex Locks:")
	fmt.Println("      defer mu.Unlock()")
	fmt.Println()
	fmt.Println("   5. Timing Functions:")
	fmt.Println("      defer func() { fmt.Println(time.Since(start)) }()\n")

	// Section 15: Summary Table
	printDeferSection("15. Summary")
	fmt.Println("   ┌─────────────────────┬──────────────────────────────────┐")
	fmt.Println("   │ Feature             │ Description                      │")
	fmt.Println("   ├─────────────────────┼──────────────────────────────────┤")
	fmt.Println("   │ Execution Timing    │ Just before function returns     │")
	fmt.Println("   │ Argument Evaluation │ Evaluated when defer runs        │")
	fmt.Println("   │ Order of Execution  │ LIFO (Last-In, First-Out)        │")
	fmt.Println("   │ Works with Panic    │ Yes, still executes              │")
	fmt.Println("   │ Main Use Case       │ Resource cleanup & management    │")
	fmt.Println("   └─────────────────────┴──────────────────────────────────┘\n")

	printDeferFooter()
}

// Example functions

func basicDeferExample() {
	fmt.Println("1. Start")
	defer fmt.Println("3. Deferred (runs last)")
	fmt.Println("2. Middle")
}

func argumentEvaluationExample() {
	i := 1
	defer fmt.Println("Result:", i) // i evaluated NOW (i = 1)
	i = 2
	fmt.Println("i is now:", i)
}

func lifoExample() {
	defer fmt.Println("First")
	defer fmt.Println("Second")
	defer fmt.Println("Third")
	fmt.Println("Main")
}

func multipleResourcesExample() {
	defer fmt.Println("Close Database")
	defer fmt.Println("Close File")
	defer fmt.Println("Close Connection")
	fmt.Println("Opening resources...")
}

func anonymousDeferExample() {
	x := 10
	defer func() {
		fmt.Println("x is:", x) // Captures x by reference
	}()
	x = 20
}

func deferInLoopExample() {
	for i := 1; i <= 3; i++ {
		defer fmt.Println(i)
	}
}

func incrementExample() (result int) {
	defer func() { result++ }()
	return 5
}

// Print helper functions

func printDeferHeader(title string) {
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Printf("  %s\n", title)
	fmt.Println(strings.Repeat("=", 60) + "\n")
}

func printDeferSection(title string) {
	fmt.Printf("┌─ %s\n", title)
	fmt.Println("│")
}

func printDeferFooter() {
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println("  ✅ Tutorial Complete!")
	fmt.Println("  💡 Key Takeaways:")
	fmt.Println("     • defer executes just before function returns")
	fmt.Println("     • Arguments are evaluated immediately, not at execution")
	fmt.Println("     • Multiple defers execute in LIFO (stack) order")
	fmt.Println("     • Use defer for cleanup (files, locks, connections)")
	fmt.Println("     • defer runs even if function panics")
	fmt.Println("     • Anonymous functions in defer capture by reference")
	fmt.Println(strings.Repeat("=", 60) + "\n")
}
