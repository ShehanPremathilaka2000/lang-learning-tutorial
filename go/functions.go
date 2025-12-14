package main

import (
	"fmt"
	"strings"
)

func functions() {
	printFuncHeader("GO FUNCTIONS TUTORIAL")

	// Section 1: What are Functions?
	printFuncSection("1. What are Functions?")
	fmt.Println("   Functions are reusable blocks of code that perform a task.")
	fmt.Println("   ✅ Organize code into logical units")
	fmt.Println("   ✅ Avoid code repetition")
	fmt.Println("   ✅ Make code more readable and maintainable\n")

	// Section 2: Function Naming Conventions
	printFuncSection("2. Function Naming Conventions")
	fmt.Println("   ┌────────────────────┬─────────────────────────────────┐")
	fmt.Println("   │ Convention         │ Description                     │")
	fmt.Println("   ├────────────────────┼─────────────────────────────────┤")
	fmt.Println("   │ camelCase          │ Private (lowercase first)       │")
	fmt.Println("   │ PascalCase         │ Public (uppercase first)        │")
	fmt.Println("   │ Descriptive names  │ Use clear, meaningful names     │")
	fmt.Println("   │ Verbs preferred    │ calculateSum, getUserData       │")
	fmt.Println("   └────────────────────┴─────────────────────────────────┘\n")

	fmt.Println("   Examples:")
	fmt.Println("   • calculateTotal()  - private function")
	fmt.Println("   • GetUserName()     - public function (exported)")
	fmt.Println("   • isValid()         - boolean check")
	fmt.Println("   • processData()     - action verb\n")

	// Section 3: Basic Function (No Parameters, No Return)
	printFuncSection("3. Basic Function (No Parameters, No Return)")
	fmt.Printf("   func sayHello() {\n")
	fmt.Printf("       fmt.Println(\"Hello, World!\")\n")
	fmt.Printf("   }\n\n")
	fmt.Printf("   Calling: sayHello()\n")
	fmt.Printf("   Output: ")
	sayHello()
	fmt.Println()

	// Section 4: Function with Parameters
	printFuncSection("4. Function with Parameters")
	fmt.Printf("   func greet(name string) {\n")
	fmt.Printf("       fmt.Printf(\"Hello, %%s!\\n\", name)\n")
	fmt.Printf("   }\n\n")
	fmt.Printf("   Calling: greet(\"Alice\")\n")
	fmt.Printf("   Output: ")
	greet("Alice")
	fmt.Println()

	// Section 5: Multiple Parameters
	printFuncSection("5. Multiple Parameters")
	fmt.Printf("   func greetPerson(firstName string, lastName string, age int) {\n")
	fmt.Printf("       fmt.Printf(\"%%s %%s is %%d years old\\n\", firstName, lastName, age)\n")
	fmt.Printf("   }\n\n")
	fmt.Printf("   Calling: greetPerson(\"John\", \"Doe\", 30)\n")
	fmt.Printf("   Output: ")
	greetPerson("John", "Doe", 30)
	fmt.Println()

	// Section 6: Parameters of Same Type (Shorthand)
	printFuncSection("6. Parameters of Same Type (Shorthand)")
	fmt.Printf("   func addThree(a, b, c int) int {\n")
	fmt.Printf("       return a + b + c\n")
	fmt.Printf("   }\n\n")
	fmt.Printf("   Calling: addThree(5, 10, 15)\n")
	result := addThree(5, 10, 15)
	fmt.Printf("   Output: %d\n\n", result)

	// Section 7: Function with Return Value
	printFuncSection("7. Function with Return Value")
	fmt.Printf("   func add(a int, b int) int {\n")
	fmt.Printf("       return a + b\n")
	fmt.Printf("   }\n\n")
	fmt.Printf("   Calling: result := add(10, 20)\n")
	addResult := add(10, 20)
	fmt.Printf("   Output: %d\n\n", addResult)

	// Section 8: Multiple Return Values
	printFuncSection("8. Multiple Return Values")
	fmt.Printf("   func divide(a, b float64) (float64, error) {\n")
	fmt.Printf("       if b == 0 {\n")
	fmt.Printf("           return 0, errors.New(\"division by zero\")\n")
	fmt.Printf("       }\n")
	fmt.Printf("       return a / b, nil\n")
	fmt.Printf("   }\n\n")
	fmt.Printf("   Calling: result, err := divide(10, 2)\n")
	divResult, err := divide(10, 2)
	if err != nil {
		fmt.Printf("   Error: %v\n\n", err)
	} else {
		fmt.Printf("   Output: %.2f\n\n", divResult)
	}

	// Section 9: Storing Multiple Return Values
	printFuncSection("9. Storing Multiple Return Values")
	fmt.Printf("   func getCoordinates() (int, int) {\n")
	fmt.Printf("       return 10, 20\n")
	fmt.Printf("   }\n\n")

	fmt.Printf("   Option 1: Store both values\n")
	fmt.Printf("   x, y := getCoordinates()\n")
	x, y := getCoordinates()
	fmt.Printf("   → x=%d, y=%d\n\n", x, y)

	fmt.Printf("   Option 2: Ignore one value with _\n")
	fmt.Printf("   x, _ := getCoordinates()  // ignore y\n")
	xOnly, _ := getCoordinates()
	fmt.Printf("   → x=%d\n\n", xOnly)

	// Section 10: Named Return Values
	printFuncSection("10. Named Return Values")
	fmt.Printf("   func rectangle(length, width int) (area, perimeter int) {\n")
	fmt.Printf("       area = length * width\n")
	fmt.Printf("       perimeter = 2 * (length + width)\n")
	fmt.Printf("       return  // naked return\n")
	fmt.Printf("   }\n\n")
	fmt.Printf("   Calling: area, perimeter := rectangle(5, 3)\n")
	area, perimeter := rectangle(5, 3)
	fmt.Printf("   Output: area=%d, perimeter=%d\n\n", area, perimeter)

	// Section 11: Named Return Values - Explicit Return
	printFuncSection("11. Named Return Values - Explicit Return")
	fmt.Printf("   func calculate(a, b int) (sum, product int) {\n")
	fmt.Printf("       sum = a + b\n")
	fmt.Printf("       product = a * b\n")
	fmt.Printf("       return sum, product  // explicit return\n")
	fmt.Printf("   }\n\n")
	fmt.Printf("   Calling: s, p := calculate(4, 5)\n")
	s, p := calculate(4, 5)
	fmt.Printf("   Output: sum=%d, product=%d\n\n", s, p)

	// Section 12: Variadic Functions (Variable Number of Parameters)
	printFuncSection("12. Variadic Functions (...)")
	fmt.Printf("   func sum(numbers ...int) int {\n")
	fmt.Printf("       total := 0\n")
	fmt.Printf("       for _, num := range numbers {\n")
	fmt.Printf("           total += num\n")
	fmt.Printf("       }\n")
	fmt.Printf("       return total\n")
	fmt.Printf("   }\n\n")
	fmt.Printf("   Calling: sum(1, 2, 3, 4, 5)\n")
	total := sum(1, 2, 3, 4, 5)
	fmt.Printf("   Output: %d\n\n", total)

	// Section 13: Passing Slice to Variadic Function
	printFuncSection("13. Passing Slice to Variadic Function")
	fmt.Printf("   nums := []int{10, 20, 30}\n")
	fmt.Printf("   result := sum(nums...)  // unpack slice\n")
	nums := []int{10, 20, 30}
	sliceSum := sum(nums...)
	fmt.Printf("   Output: %d\n\n", sliceSum)

	// Section 14: Function as Value
	printFuncSection("14. Function as Value (First-Class Functions)")
	fmt.Printf("   add := func(a, b int) int {\n")
	fmt.Printf("       return a + b\n")
	fmt.Printf("   }\n")
	fmt.Printf("   result := add(5, 3)\n\n")
	addFunc := func(a, b int) int {
		return a + b
	}
	funcResult := addFunc(5, 3)
	fmt.Printf("   Output: %d\n\n", funcResult)

	// Section 15: Anonymous Functions
	printFuncSection("15. Anonymous Functions (Immediate Execution)")
	fmt.Printf("   func() {\n")
	fmt.Printf("       fmt.Println(\"Anonymous function executed\")\n")
	fmt.Printf("   }()\n\n")
	fmt.Printf("   Output: ")
	func() {
		fmt.Println("Anonymous function executed")
	}()
	fmt.Println()

	// Section 16: Function Returning a Function
	printFuncSection("16. Function Returning a Function (Closure)")
	fmt.Printf("   func multiplier(factor int) func(int) int {\n")
	fmt.Printf("       return func(x int) int {\n")
	fmt.Printf("           return x * factor\n")
	fmt.Printf("       }\n")
	fmt.Printf("   }\n\n")
	fmt.Printf("   double := multiplier(2)\n")
	fmt.Printf("   triple := multiplier(3)\n")
	double := multiplier(2)
	triple := multiplier(3)
	fmt.Printf("   double(5) = %d\n", double(5))
	fmt.Printf("   triple(5) = %d\n\n", triple(5))

	// Section 17: Recursive Functions
	printFuncSection("17. Recursive Functions")
	fmt.Printf("   func factorial(n int) int {\n")
	fmt.Printf("       if n <= 1 {\n")
	fmt.Printf("           return 1\n")
	fmt.Printf("       }\n")
	fmt.Printf("       return n * factorial(n-1)\n")
	fmt.Printf("   }\n\n")
	fmt.Printf("   Calling: factorial(5)\n")
	fact := factorial(5)
	fmt.Printf("   Output: 5! = %d\n\n", fact)

	// Section 18: Defer Statement
	printFuncSection("18. Defer Statement (Execute After Function Returns)")
	fmt.Printf("   func example() {\n")
	fmt.Printf("       defer fmt.Println(\"3. Deferred (runs last)\")\n")
	fmt.Printf("       fmt.Println(\"1. First\")\n")
	fmt.Printf("       fmt.Println(\"2. Second\")\n")
	fmt.Printf("   }\n\n")
	fmt.Printf("   Output:\n   ")
	deferExample()
	fmt.Println()

	// Section 19: Multiple Defer (LIFO - Last In First Out)
	printFuncSection("19. Multiple Defer (Stack Order - LIFO)")
	fmt.Printf("   func multiDefer() {\n")
	fmt.Printf("       defer fmt.Println(\"Third\")\n")
	fmt.Printf("       defer fmt.Println(\"Second\")\n")
	fmt.Printf("       defer fmt.Println(\"First\")\n")
	fmt.Printf("       fmt.Println(\"Main\")\n")
	fmt.Printf("   }\n\n")
	fmt.Printf("   Output:\n   ")
	multiDefer()
	fmt.Println()

	// Section 20: Practical Examples
	printFuncSection("20. Practical Examples")

	// Example 1: Temperature conversion
	fmt.Printf("   Example 1: Temperature Conversion\n")
	celsius := 25.0
	fahrenheit := celsiusToFahrenheit(celsius)
	fmt.Printf("   %.1f°C = %.1f°F\n\n", celsius, fahrenheit)

	// Example 2: String manipulation
	fmt.Printf("   Example 2: String Manipulation\n")
	reversed := reverseString("Hello")
	fmt.Printf("   Reverse of \"Hello\" = \"%s\"\n\n", reversed)

	// Example 3: Min and Max
	fmt.Printf("   Example 3: Find Min and Max\n")
	min, max := findMinMax(3, 7, 2, 9, 1)
	fmt.Printf("   Min: %d, Max: %d\n\n", min, max)

	printFuncFooter()
}

// Helper functions for demonstrations

func sayHello() {
	fmt.Println("Hello, World!")
}

func greet(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

func greetPerson(firstName string, lastName string, age int) {
	fmt.Printf("%s %s is %d years old\n", firstName, lastName, age)
}

func addThree(a, b, c int) int {
	return a + b + c
}

func add(a int, b int) int {
	return a + b
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

func getCoordinates() (int, int) {
	return 10, 20
}

func rectangle(length, width int) (area, perimeter int) {
	area = length * width
	perimeter = 2 * (length + width)
	return // naked return
}

func calculate(a, b int) (sum, product int) {
	sum = a + b
	product = a * b
	return sum, product
}

func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func multiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func deferExample() {
	defer fmt.Println("3. Deferred (runs last)")
	fmt.Println("1. First")
	fmt.Println("2. Second")
}

func multiDefer() {
	defer fmt.Println("Third")
	defer fmt.Println("Second")
	defer fmt.Println("First")
	fmt.Println("Main")
}

func celsiusToFahrenheit(c float64) float64 {
	return (c * 9 / 5) + 32
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func findMinMax(numbers ...int) (min, max int) {
	if len(numbers) == 0 {
		return 0, 0
	}
	min, max = numbers[0], numbers[0]
	for _, num := range numbers {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}
	return
}

func printFuncHeader(title string) {
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Printf("  %s\n", title)
	fmt.Println(strings.Repeat("=", 60) + "\n")
}

func printFuncSection(title string) {
	fmt.Printf("┌─ %s\n", title)
	fmt.Println("│")
}

func printFuncFooter() {
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println("  ✅ Tutorial Complete!")
	fmt.Println("  💡 Key Takeaways:")
	fmt.Println("     • Functions organize code into reusable blocks")
	fmt.Println("     • Use camelCase (private) or PascalCase (public)")
	fmt.Println("     • Functions can return multiple values")
	fmt.Println("     • Use named return values for clarity")
	fmt.Println("     • Use _ to ignore unwanted return values")
	fmt.Println("     • Variadic functions accept variable arguments (...)")
	fmt.Println("     • defer executes code after function returns (LIFO)")
	fmt.Println(strings.Repeat("=", 60) + "\n")
}
