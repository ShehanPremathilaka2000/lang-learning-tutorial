package main

import (
	"fmt"
	"strings"
)

func conditions() {
	printCondHeader("GO CONDITIONS TUTORIAL")

	// Section 1: What are Conditions?
	printCondSection("1. What are Conditions?")
	fmt.Println("   Conditions allow your program to make decisions.")
	fmt.Println("   ✅ Execute code based on boolean expressions")
	fmt.Println("   ✅ Control the flow of your program")
	fmt.Println("   ✅ Use comparison and logical operators\n")

	// Section 2: Simple if Statement
	printCondSection("2. Simple if Statement")
	age := 20
	fmt.Printf("   age := 20\n\n")
	fmt.Printf("   if age >= 18 {\n")
	fmt.Printf("       fmt.Println(\"You are an adult\")\n")
	fmt.Printf("   }\n\n")
	fmt.Printf("   Output: ")
	if age >= 18 {
		fmt.Println("You are an adult")
	}
	fmt.Println()

	// Section 3: if-else Statement
	printCondSection("3. if-else Statement")
	temperature := 15
	fmt.Printf("   temperature := 15\n\n")
	fmt.Printf("   if temperature > 20 {\n")
	fmt.Printf("       fmt.Println(\"It's warm\")\n")
	fmt.Printf("   } else {\n")
	fmt.Printf("       fmt.Println(\"It's cold\")\n")
	fmt.Printf("   }\n\n")
	fmt.Printf("   Output: ")
	if temperature > 20 {
		fmt.Println("It's warm")
	} else {
		fmt.Println("It's cold")
	}
	fmt.Println()

	// Section 4: if-else if-else Statement
	printCondSection("4. if-else if-else Statement")
	score := 75
	fmt.Printf("   score := 75\n\n")
	fmt.Printf("   if score >= 90 {\n")
	fmt.Printf("       fmt.Println(\"Grade: A\")\n")
	fmt.Printf("   } else if score >= 80 {\n")
	fmt.Printf("       fmt.Println(\"Grade: B\")\n")
	fmt.Printf("   } else if score >= 70 {\n")
	fmt.Printf("       fmt.Println(\"Grade: C\")\n")
	fmt.Printf("   } else {\n")
	fmt.Printf("       fmt.Println(\"Grade: F\")\n")
	fmt.Printf("   }\n\n")
	fmt.Printf("   Output: ")
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else if score >= 70 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: F")
	}
	fmt.Println()

	// Section 5: Nested if Statements
	printCondSection("5. Nested if Statements")
	userAge := 25
	hasLicense := true
	fmt.Printf("   userAge := 25\n")
	fmt.Printf("   hasLicense := true\n\n")
	fmt.Printf("   if userAge >= 18 {\n")
	fmt.Printf("       if hasLicense {\n")
	fmt.Printf("           fmt.Println(\"You can drive\")\n")
	fmt.Printf("       } else {\n")
	fmt.Printf("           fmt.Println(\"Get a license first\")\n")
	fmt.Printf("       }\n")
	fmt.Printf("   } else {\n")
	fmt.Printf("       fmt.Println(\"Too young to drive\")\n")
	fmt.Printf("   }\n\n")
	fmt.Printf("   Output: ")
	if userAge >= 18 {
		if hasLicense {
			fmt.Println("You can drive")
		} else {
			fmt.Println("Get a license first")
		}
	} else {
		fmt.Println("Too young to drive")
	}
	fmt.Println()

	// Section 6: if with Short Statement
	printCondSection("6. if with Short Statement (Variable Declaration)")
	fmt.Printf("   if num := 10; num > 5 {\n")
	fmt.Printf("       fmt.Println(\"num is greater than 5\")\n")
	fmt.Printf("   }\n")
	fmt.Printf("   // num is only available inside the if block\n\n")
	fmt.Printf("   Output: ")
	if num := 10; num > 5 {
		fmt.Println("num is greater than 5")
	}
	fmt.Printf("   💡 Variable 'num' is scoped to the if block only!\n\n")

	// Section 7: Multiple Conditions with Logical Operators
	printCondSection("7. Multiple Conditions with Logical Operators")
	username := "admin"
	password := "secret123"
	fmt.Printf("   username := \"admin\"\n")
	fmt.Printf("   password := \"secret123\"\n\n")
	fmt.Printf("   if username == \"admin\" && password == \"secret123\" {\n")
	fmt.Printf("       fmt.Println(\"Login successful\")\n")
	fmt.Printf("   } else {\n")
	fmt.Printf("       fmt.Println(\"Login failed\")\n")
	fmt.Printf("   }\n\n")
	fmt.Printf("   Output: ")
	if username == "admin" && password == "secret123" {
		fmt.Println("Login successful")
	} else {
		fmt.Println("Login failed")
	}
	fmt.Println()

	// Section 8: IMPORTANT - Go Brace Syntax Rules
	printCondSection("8. ⚠️  IMPORTANT: Go Brace Syntax Rules")
	fmt.Println("   In Go, the opening brace { MUST be on the same line!")
	fmt.Println("   The closing brace } and else MUST be on the same line!\n")

	fmt.Println("   ✅ CORRECT:")
	fmt.Println("   if condition {")
	fmt.Println("       // code")
	fmt.Println("   } else {")
	fmt.Println("       // code")
	fmt.Println("   }\n")

	fmt.Println("   ❌ WRONG (will cause compile error):")
	fmt.Println("   if condition")
	fmt.Println("   {")
	fmt.Println("       // code")
	fmt.Println("   }")
	fmt.Println("   else")
	fmt.Println("   {")
	fmt.Println("       // code")
	fmt.Println("   }\n")

	fmt.Println("   💡 This is enforced by Go's automatic semicolon insertion!\n")

	// Section 9: Comparing Different Types
	printCondSection("9. Comparing Different Types")
	str1 := "hello"
	str2 := "world"
	fmt.Printf("   str1 := \"hello\"\n")
	fmt.Printf("   str2 := \"world\"\n\n")
	fmt.Printf("   if str1 == str2 {\n")
	fmt.Printf("       fmt.Println(\"Strings are equal\")\n")
	fmt.Printf("   } else {\n")
	fmt.Printf("       fmt.Println(\"Strings are different\")\n")
	fmt.Printf("   }\n\n")
	fmt.Printf("   Output: ")
	if str1 == str2 {
		fmt.Println("Strings are equal")
	} else {
		fmt.Println("Strings are different")
	}
	fmt.Println()

	// Section 10: Checking for Empty/Zero Values
	printCondSection("10. Checking for Empty/Zero Values")
	var emptyString string
	var zeroNum int
	var nilSlice []int

	fmt.Printf("   var emptyString string\n")
	fmt.Printf("   var zeroNum int\n")
	fmt.Printf("   var nilSlice []int\n\n")

	fmt.Printf("   if emptyString == \"\" {\n")
	fmt.Printf("       fmt.Println(\"String is empty\")\n")
	fmt.Printf("   }\n")
	if emptyString == "" {
		fmt.Printf("   Output: String is empty\n\n")
	}

	fmt.Printf("   if zeroNum == 0 {\n")
	fmt.Printf("       fmt.Println(\"Number is zero\")\n")
	fmt.Printf("   }\n")
	if zeroNum == 0 {
		fmt.Printf("   Output: Number is zero\n\n")
	}

	fmt.Printf("   if nilSlice == nil {\n")
	fmt.Printf("       fmt.Println(\"Slice is nil\")\n")
	fmt.Printf("   }\n")
	if nilSlice == nil {
		fmt.Printf("   Output: Slice is nil\n\n")
	}

	// Section 11: Practical Examples
	printCondSection("11. Practical Examples")

	// Example 1: Even or Odd
	number := 17
	fmt.Printf("   Example 1: Check if number is even or odd\n")
	fmt.Printf("   number := 17\n\n")
	if number%2 == 0 {
		fmt.Printf("   %d is even\n\n", number)
	} else {
		fmt.Printf("   %d is odd\n\n", number)
	}

	// Example 2: Leap Year
	year := 2024
	fmt.Printf("   Example 2: Check if year is a leap year\n")
	fmt.Printf("   year := 2024\n\n")
	if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
		fmt.Printf("   %d is a leap year\n\n", year)
	} else {
		fmt.Printf("   %d is not a leap year\n\n", year)
	}

	// Example 3: Range Check
	value := 45
	fmt.Printf("   Example 3: Check if value is in range\n")
	fmt.Printf("   value := 45\n\n")
	if value >= 0 && value <= 100 {
		fmt.Printf("   %d is within range [0-100]\n\n", value)
	} else {
		fmt.Printf("   %d is outside range [0-100]\n\n", value)
	}

	// Section 12: Switch Statement - Basic
	printCondSection("12. Switch Statement - Basic")
	fmt.Println("   Switch is a cleaner way to write multiple if-else statements.")
	fmt.Println("   ✅ No break needed (automatic in Go)")
	fmt.Println("   ✅ Only the matching case executes\n")

	day := 3
	fmt.Printf("   day := 3\n\n")
	fmt.Printf("   switch day {\n")
	fmt.Printf("   case 1:\n")
	fmt.Printf("       fmt.Println(\"Monday\")\n")
	fmt.Printf("   case 2:\n")
	fmt.Printf("       fmt.Println(\"Tuesday\")\n")
	fmt.Printf("   case 3:\n")
	fmt.Printf("       fmt.Println(\"Wednesday\")\n")
	fmt.Printf("   case 4:\n")
	fmt.Printf("       fmt.Println(\"Thursday\")\n")
	fmt.Printf("   case 5:\n")
	fmt.Printf("       fmt.Println(\"Friday\")\n")
	fmt.Printf("   }\n\n")
	fmt.Printf("   Output: ")
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	}
	fmt.Println()

	// Section 13: Switch with Default Case
	printCondSection("13. Switch with Default Case")
	fmt.Println("   The 'default' case runs when no other case matches.\n")

	dayNum := 7
	fmt.Printf("   dayNum := 7\n\n")
	fmt.Printf("   switch dayNum {\n")
	fmt.Printf("   case 1:\n")
	fmt.Printf("       fmt.Println(\"Monday\")\n")
	fmt.Printf("   case 2:\n")
	fmt.Printf("       fmt.Println(\"Tuesday\")\n")
	fmt.Printf("   default:\n")
	fmt.Printf("       fmt.Println(\"Weekend or invalid day\")\n")
	fmt.Printf("   }\n\n")
	fmt.Printf("   Output: ")
	switch dayNum {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	default:
		fmt.Println("Weekend or invalid day")
	}
	fmt.Println()

	// Section 14: Switch with Multiple Values per Case
	printCondSection("14. Switch with Multiple Values per Case")
	fmt.Println("   You can match multiple values in a single case.\n")

	char := 'e'
	fmt.Printf("   char := 'e'\n\n")
	fmt.Printf("   switch char {\n")
	fmt.Printf("   case 'a', 'e', 'i', 'o', 'u':\n")
	fmt.Printf("       fmt.Println(\"Vowel\")\n")
	fmt.Printf("   case 'y':\n")
	fmt.Printf("       fmt.Println(\"Sometimes a vowel\")\n")
	fmt.Printf("   default:\n")
	fmt.Printf("       fmt.Println(\"Consonant\")\n")
	fmt.Printf("   }\n\n")
	fmt.Printf("   Output: ")
	switch char {
	case 'a', 'e', 'i', 'o', 'u':
		fmt.Println("Vowel")
	case 'y':
		fmt.Println("Sometimes a vowel")
	default:
		fmt.Println("Consonant")
	}
	fmt.Println()

	// Section 15: Switch with Short Statement
	printCondSection("15. Switch with Short Statement")
	fmt.Println("   Like if, switch can have a short statement before the condition.\n")

	fmt.Printf("   switch grade := 85; {\n")
	fmt.Printf("   case grade >= 90:\n")
	fmt.Printf("       fmt.Println(\"A\")\n")
	fmt.Printf("   case grade >= 80:\n")
	fmt.Printf("       fmt.Println(\"B\")\n")
	fmt.Printf("   case grade >= 70:\n")
	fmt.Printf("       fmt.Println(\"C\")\n")
	fmt.Printf("   default:\n")
	fmt.Printf("       fmt.Println(\"F\")\n")
	fmt.Printf("   }\n\n")
	fmt.Printf("   Output: ")
	switch grade := 85; {
	case grade >= 90:
		fmt.Println("A")
	case grade >= 80:
		fmt.Println("B")
	case grade >= 70:
		fmt.Println("C")
	default:
		fmt.Println("F")
	}
	fmt.Println()

	// Section 16: Switch without Expression (like if-else)
	printCondSection("16. Switch without Expression")
	fmt.Println("   Switch without an expression is like a clean if-else chain.\n")

	time := 14
	fmt.Printf("   time := 14\n\n")
	fmt.Printf("   switch {\n")
	fmt.Printf("   case time < 12:\n")
	fmt.Printf("       fmt.Println(\"Good morning\")\n")
	fmt.Printf("   case time < 17:\n")
	fmt.Printf("       fmt.Println(\"Good afternoon\")\n")
	fmt.Printf("   default:\n")
	fmt.Printf("       fmt.Println(\"Good evening\")\n")
	fmt.Printf("   }\n\n")
	fmt.Printf("   Output: ")
	switch {
	case time < 12:
		fmt.Println("Good morning")
	case time < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}
	fmt.Println()

	// Section 17: Switch on Type
	printCondSection("17. Switch on Type")
	fmt.Println("   You can switch on the type of an interface variable.\n")

	var i interface{} = "hello"
	fmt.Printf("   var i interface{} = \"hello\"\n\n")
	fmt.Printf("   switch v := i.(type) {\n")
	fmt.Printf("   case int:\n")
	fmt.Printf("       fmt.Printf(\"Integer: %%d\", v)\n")
	fmt.Printf("   case string:\n")
	fmt.Printf("       fmt.Printf(\"String: %%s\", v)\n")
	fmt.Printf("   case bool:\n")
	fmt.Printf("       fmt.Printf(\"Boolean: %%t\", v)\n")
	fmt.Printf("   default:\n")
	fmt.Printf("       fmt.Printf(\"Unknown type\")\n")
	fmt.Printf("   }\n\n")
	fmt.Printf("   Output: ")
	switch v := i.(type) {
	case int:
		fmt.Printf("Integer: %d\n", v)
	case string:
		fmt.Printf("String: %s\n", v)
	case bool:
		fmt.Printf("Boolean: %t\n", v)
	default:
		fmt.Printf("Unknown type\n")
	}
	fmt.Println()

	// Section 18: Fallthrough Keyword
	printCondSection("18. Fallthrough Keyword")
	fmt.Println("   By default, Go switch doesn't fall through to next case.")
	fmt.Println("   Use 'fallthrough' to explicitly continue to next case.\n")

	num := 1
	fmt.Printf("   num := 1\n\n")
	fmt.Printf("   switch num {\n")
	fmt.Printf("   case 1:\n")
	fmt.Printf("       fmt.Println(\"One\")\n")
	fmt.Printf("       fallthrough\n")
	fmt.Printf("   case 2:\n")
	fmt.Printf("       fmt.Println(\"Two or after one\")\n")
	fmt.Printf("   case 3:\n")
	fmt.Printf("       fmt.Println(\"Three\")\n")
	fmt.Printf("   }\n\n")
	fmt.Printf("   Output:\n   ")
	switch num {
	case 1:
		fmt.Println("One")
		fallthrough
	case 2:
		fmt.Println("Two or after one")
	case 3:
		fmt.Println("Three")
	}
	fmt.Printf("   ⚠️  fallthrough executes next case unconditionally!\n\n")

	// Section 19: Practical Switch Examples
	printCondSection("19. Practical Switch Examples")

	// Example 1: Month days
	month := "February"
	fmt.Printf("   Example 1: Days in month\n")
	fmt.Printf("   month := \"February\"\n\n")
	switch month {
	case "January", "March", "May", "July", "August", "October", "December":
		fmt.Printf("   %s has 31 days\n\n", month)
	case "April", "June", "September", "November":
		fmt.Printf("   %s has 30 days\n\n", month)
	case "February":
		fmt.Printf("   %s has 28 or 29 days\n\n", month)
	default:
		fmt.Printf("   Invalid month\n\n")
	}

	// Example 2: HTTP Status Code
	statusCode := 404
	fmt.Printf("   Example 2: HTTP Status Code\n")
	fmt.Printf("   statusCode := 404\n\n")
	switch statusCode {
	case 200:
		fmt.Println("   OK")
	case 404:
		fmt.Println("   Not Found")
	case 500:
		fmt.Println("   Internal Server Error")
	default:
		fmt.Println("   Unknown Status")
	}
	fmt.Println()

	printCondFooter()
}

func printCondHeader(title string) {
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Printf("  %s\n", title)
	fmt.Println(strings.Repeat("=", 60) + "\n")
}

func printCondSection(title string) {
	fmt.Printf("┌─ %s\n", title)
	fmt.Println("│")
}

func printCondFooter() {
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println("  ✅ Tutorial Complete!")
	fmt.Println("  💡 Key Takeaways:")
	fmt.Println("     • Use if-else for decision making in your code")
	fmt.Println("     • Opening brace { must be on the same line as if/else")
	fmt.Println("     • } else must be on the same line (not separate lines)")
	fmt.Println("     • Use logical operators (&&, ||) for multiple conditions")
	fmt.Println("     • Use switch for cleaner multiple condition checks")
	fmt.Println("     • Switch doesn't need break (automatic in Go)")
	fmt.Println("     • Use default case for unmatched values")
	fmt.Println("     • Multiple values per case: case 1, 2, 3:")
	fmt.Println(strings.Repeat("=", 60) + "\n")
}
