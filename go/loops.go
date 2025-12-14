package main

import (
	"fmt"
	"strings"
)

func loops() {
	printLoopHeader("GO LOOPS TUTORIAL")

	// Section 1: What are Loops?
	printLoopSection("1. What are Loops?")
	fmt.Println("   Loops allow you to repeat code multiple times.")
	fmt.Println("   ✅ Go has only ONE loop keyword: 'for'")
	fmt.Println("   ✅ 'for' can be used in multiple ways")
	fmt.Println("   ✅ No 'while' or 'do-while' keywords (use 'for' instead)\n")

	// Section 2: Basic for Loop (3 components)
	printLoopSection("2. Basic for Loop (Classic Style)")
	fmt.Println("   for initialization; condition; post {")
	fmt.Println("       // code to repeat")
	fmt.Println("   }\n")

	fmt.Printf("   for i := 0; i < 5; i++ {\n")
	fmt.Printf("       fmt.Printf(\"Count: %%d \", i)\n")
	fmt.Printf("   }\n\n")
	fmt.Printf("   Output: ")
	for i := 0; i < 5; i++ {
		fmt.Printf("Count: %d ", i)
	}
	fmt.Println("\n")

	// Section 3: for Loop - Counting Down
	printLoopSection("3. for Loop - Counting Down")
	fmt.Printf("   for i := 5; i > 0; i-- {\n")
	fmt.Printf("       fmt.Printf(\"%%d \", i)\n")
	fmt.Printf("   }\n")
	fmt.Printf("   fmt.Println(\"Liftoff!\")\n\n")
	fmt.Printf("   Output: ")
	for i := 5; i > 0; i-- {
		fmt.Printf("%d ", i)
	}
	fmt.Println("Liftoff!\n")

	// Section 4: for Loop - Custom Increment
	printLoopSection("4. for Loop - Custom Increment")
	fmt.Printf("   for i := 0; i <= 10; i += 2 {\n")
	fmt.Printf("       fmt.Printf(\"%%d \", i)\n")
	fmt.Printf("   }\n\n")
	fmt.Printf("   Output: ")
	for i := 0; i <= 10; i += 2 {
		fmt.Printf("%d ", i)
	}
	fmt.Println("\n")

	// Section 5: for as While Loop
	printLoopSection("5. for as While Loop (Only Condition)")
	fmt.Println("   Go doesn't have 'while', but you can use 'for' with only a condition.\n")

	fmt.Printf("   count := 0\n")
	fmt.Printf("   for count < 5 {\n")
	fmt.Printf("       fmt.Printf(\"%%d \", count)\n")
	fmt.Printf("       count++\n")
	fmt.Printf("   }\n\n")
	fmt.Printf("   Output: ")
	count := 0
	for count < 5 {
		fmt.Printf("%d ", count)
		count++
	}
	fmt.Println("\n")

	// Section 6: Infinite Loop
	printLoopSection("6. Infinite Loop")
	fmt.Println("   for {")
	fmt.Println("       // runs forever until break")
	fmt.Println("   }\n")

	fmt.Printf("   counter := 0\n")
	fmt.Printf("   for {\n")
	fmt.Printf("       counter++\n")
	fmt.Printf("       if counter > 3 {\n")
	fmt.Printf("           break\n")
	fmt.Printf("       }\n")
	fmt.Printf("       fmt.Printf(\"%%d \", counter)\n")
	fmt.Printf("   }\n\n")
	fmt.Printf("   Output: ")
	counter := 0
	for {
		counter++
		if counter > 3 {
			break
		}
		fmt.Printf("%d ", counter)
	}
	fmt.Println("\n")

	// Section 7: break Statement
	printLoopSection("7. break Statement (Exit Loop Early)")
	fmt.Println("   'break' immediately exits the loop.\n")

	fmt.Printf("   for i := 1; i <= 10; i++ {\n")
	fmt.Printf("       if i == 5 {\n")
	fmt.Printf("           break  // exit when i is 5\n")
	fmt.Printf("       }\n")
	fmt.Printf("       fmt.Printf(\"%%d \", i)\n")
	fmt.Printf("   }\n\n")
	fmt.Printf("   Output: ")
	for i := 1; i <= 10; i++ {
		if i == 5 {
			break
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println("(stopped at 5)\n")

	// Section 8: continue Statement
	printLoopSection("8. continue Statement (Skip Current Iteration)")
	fmt.Println("   'continue' skips the rest of the current iteration.\n")

	fmt.Printf("   for i := 1; i <= 10; i++ {\n")
	fmt.Printf("       if i%%2 == 0 {\n")
	fmt.Printf("           continue  // skip even numbers\n")
	fmt.Printf("       }\n")
	fmt.Printf("       fmt.Printf(\"%%d \", i)\n")
	fmt.Printf("   }\n\n")
	fmt.Printf("   Output: ")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println("(odd numbers only)\n")

	// Section 9: Nested Loops
	printLoopSection("9. Nested Loops")
	fmt.Printf("   for i := 1; i <= 3; i++ {\n")
	fmt.Printf("       for j := 1; j <= 3; j++ {\n")
	fmt.Printf("           fmt.Printf(\"(%%d,%%d) \", i, j)\n")
	fmt.Printf("       }\n")
	fmt.Printf("       fmt.Println()\n")
	fmt.Printf("   }\n\n")
	fmt.Printf("   Output:\n")
	for i := 1; i <= 3; i++ {
		fmt.Printf("   ")
		for j := 1; j <= 3; j++ {
			fmt.Printf("(%d,%d) ", i, j)
		}
		fmt.Println()
	}
	fmt.Println()

	// Section 10: for-range with Arrays
	printLoopSection("10. for-range with Arrays")
	fmt.Println("   'range' iterates over arrays, slices, maps, and strings.\n")

	numbers := [5]int{10, 20, 30, 40, 50}
	fmt.Printf("   numbers := [5]int{10, 20, 30, 40, 50}\n\n")
	fmt.Printf("   for index, value := range numbers {\n")
	fmt.Printf("       fmt.Printf(\"[%%d]=%%d \", index, value)\n")
	fmt.Printf("   }\n\n")
	fmt.Printf("   Output: ")
	for index, value := range numbers {
		fmt.Printf("[%d]=%d ", index, value)
	}
	fmt.Println("\n")

	// Section 11: for-range with Slices
	printLoopSection("11. for-range with Slices")
	fruits := []string{"Apple", "Banana", "Cherry"}
	fmt.Printf("   fruits := []string{\"Apple\", \"Banana\", \"Cherry\"}\n\n")
	fmt.Printf("   for index, fruit := range fruits {\n")
	fmt.Printf("       fmt.Printf(\"%%d: %%s\\n\", index, fruit)\n")
	fmt.Printf("   }\n\n")
	fmt.Printf("   Output:\n")
	for index, fruit := range fruits {
		fmt.Printf("   %d: %s\n", index, fruit)
	}
	fmt.Println()

	// Section 12: for-range - Index Only
	printLoopSection("12. for-range - Index Only")
	fmt.Printf("   for index := range fruits {\n")
	fmt.Printf("       fmt.Printf(\"%%d \", index)\n")
	fmt.Printf("   }\n\n")
	fmt.Printf("   Output: ")
	for index := range fruits {
		fmt.Printf("%d ", index)
	}
	fmt.Println("\n")

	// Section 13: for-range - Value Only
	printLoopSection("13. for-range - Value Only (Ignore Index)")
	fmt.Printf("   for _, fruit := range fruits {\n")
	fmt.Printf("       fmt.Printf(\"%%s \", fruit)\n")
	fmt.Printf("   }\n\n")
	fmt.Printf("   Output: ")
	for _, fruit := range fruits {
		fmt.Printf("%s ", fruit)
	}
	fmt.Println("\n")

	// Section 14: for-range with Strings
	printLoopSection("14. for-range with Strings (Runes)")
	text := "Go!"
	fmt.Printf("   text := \"Go!\"\n\n")
	fmt.Printf("   for index, char := range text {\n")
	fmt.Printf("       fmt.Printf(\"[%%d]=%%c \", index, char)\n")
	fmt.Printf("   }\n\n")
	fmt.Printf("   Output: ")
	for index, char := range text {
		fmt.Printf("[%d]=%c ", index, char)
	}
	fmt.Println("\n")

	// Section 15: for-range with Maps
	printLoopSection("15. for-range with Maps")
	ages := map[string]int{"Alice": 25, "Bob": 30, "Charlie": 35}
	fmt.Printf("   ages := map[string]int{\"Alice\": 25, \"Bob\": 30, \"Charlie\": 35}\n\n")
	fmt.Printf("   for name, age := range ages {\n")
	fmt.Printf("       fmt.Printf(\"%%s: %%d\\n\", name, age)\n")
	fmt.Printf("   }\n\n")
	fmt.Printf("   Output:\n")
	for name, age := range ages {
		fmt.Printf("   %s: %d\n", name, age)
	}
	fmt.Println()

	// Section 16: Labeled break
	printLoopSection("16. Labeled break (Break Outer Loop)")
	fmt.Printf("   outer:\n")
	fmt.Printf("   for i := 1; i <= 3; i++ {\n")
	fmt.Printf("       for j := 1; j <= 3; j++ {\n")
	fmt.Printf("           if i*j > 4 {\n")
	fmt.Printf("               break outer  // breaks outer loop\n")
	fmt.Printf("           }\n")
	fmt.Printf("           fmt.Printf(\"%%d*%%d=%%d \", i, j, i*j)\n")
	fmt.Printf("       }\n")
	fmt.Printf("   }\n\n")
	fmt.Printf("   Output: ")
outer:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if i*j > 4 {
				break outer
			}
			fmt.Printf("%d*%d=%d ", i, j, i*j)
		}
	}
	fmt.Println("\n")

	// Section 17: Labeled continue
	printLoopSection("17. Labeled continue (Continue Outer Loop)")
	fmt.Printf("   outer:\n")
	fmt.Printf("   for i := 1; i <= 3; i++ {\n")
	fmt.Printf("       for j := 1; j <= 3; j++ {\n")
	fmt.Printf("           if j == 2 {\n")
	fmt.Printf("               continue outer  // continues outer loop\n")
	fmt.Printf("           }\n")
	fmt.Printf("           fmt.Printf(\"(%%d,%%d) \", i, j)\n")
	fmt.Printf("       }\n")
	fmt.Printf("   }\n\n")
	fmt.Printf("   Output: ")
outer2:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if j == 2 {
				continue outer2
			}
			fmt.Printf("(%d,%d) ", i, j)
		}
	}
	fmt.Println("\n")

	// Section 18: Practical Examples
	printLoopSection("18. Practical Examples")

	// Example 1: Sum of numbers
	fmt.Printf("   Example 1: Sum of numbers 1 to 10\n")
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Printf("   Sum: %d\n\n", sum)

	// Example 2: Factorial
	fmt.Printf("   Example 2: Factorial of 5\n")
	factorial := 1
	for i := 1; i <= 5; i++ {
		factorial *= i
	}
	fmt.Printf("   5! = %d\n\n", factorial)

	// Example 3: Fibonacci sequence
	fmt.Printf("   Example 3: First 10 Fibonacci numbers\n")
	fmt.Printf("   Output: ")
	a, b := 0, 1
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", a)
		a, b = b, a+b
	}
	fmt.Println("\n")

	// Example 4: Multiplication table
	fmt.Printf("   Example 4: Multiplication table for 5\n")
	for i := 1; i <= 10; i++ {
		fmt.Printf("   5 x %d = %d\n", i, 5*i)
	}
	fmt.Println()

	// Example 5: Find prime numbers
	fmt.Printf("   Example 5: Prime numbers up to 20\n")
	fmt.Printf("   Output: ")
	for num := 2; num <= 20; num++ {
		isPrime := true
		for i := 2; i*i <= num; i++ {
			if num%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Printf("%d ", num)
		}
	}
	fmt.Println("\n")

	printLoopFooter()
}

func printLoopHeader(title string) {
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Printf("  %s\n", title)
	fmt.Println(strings.Repeat("=", 60) + "\n")
}

func printLoopSection(title string) {
	fmt.Printf("┌─ %s\n", title)
	fmt.Println("│")
}

func printLoopFooter() {
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println("  ✅ Tutorial Complete!")
	fmt.Println("  💡 Key Takeaways:")
	fmt.Println("     • Go has only 'for' loops (no while or do-while)")
	fmt.Println("     • Use 'for condition {}' as a while loop")
	fmt.Println("     • Use 'for {}' for infinite loops")
	fmt.Println("     • Use 'range' to iterate over arrays, slices, maps, strings")
	fmt.Println("     • Use 'break' to exit loops, 'continue' to skip iterations")
	fmt.Println("     • Use labels with break/continue for nested loops")
	fmt.Println(strings.Repeat("=", 60) + "\n")
}
