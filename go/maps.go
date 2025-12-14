package main

import (
	"fmt"
	"strings"
)

func maps() {
	printMapHeader("GO MAPS TUTORIAL")

	// Section 1: What are Maps?
	printMapSection("1. What are Maps?")
	fmt.Println("   Maps are key-value pairs (like dictionaries or hash tables).")
	fmt.Println("   ✅ Store data as key-value associations")
	fmt.Println("   ✅ Fast lookups by key")
	fmt.Println("   ✅ Keys must be unique (no duplicates)")
	fmt.Println("   ✅ Unordered (iteration order is not guaranteed)\n")

	// Section 2: Creating Maps - Using make()
	printMapSection("2. Creating Maps - Using make()")
	fmt.Printf("   ages := make(map[string]int)\n")
	fmt.Printf("   ages[\"Alice\"] = 25\n")
	fmt.Printf("   ages[\"Bob\"] = 30\n\n")
	ages := make(map[string]int)
	ages["Alice"] = 25
	ages["Bob"] = 30
	fmt.Printf("   → %v\n\n", ages)

	// Section 3: Creating Maps - Map Literal
	printMapSection("3. Creating Maps - Map Literal")
	fmt.Printf("   scores := map[string]int{\n")
	fmt.Printf("       \"Math\":    95,\n")
	fmt.Printf("       \"English\": 88,\n")
	fmt.Printf("       \"Science\": 92,\n")
	fmt.Printf("   }\n\n")
	scores := map[string]int{
		"Math":    95,
		"English": 88,
		"Science": 92,
	}
	fmt.Printf("   → %v\n\n", scores)

	// Section 4: Creating Maps - Short Declaration
	printMapSection("4. Creating Maps - Short Declaration")
	fmt.Printf("   cities := map[string]string{\n")
	fmt.Printf("       \"USA\": \"Washington DC\",\n")
	fmt.Printf("       \"UK\":  \"London\",\n")
	fmt.Printf("   }\n\n")
	cities := map[string]string{
		"USA": "Washington DC",
		"UK":  "London",
	}
	fmt.Printf("   → %v\n\n", cities)

	// Section 5: Allowed Key and Value Types
	printMapSection("5. Allowed Key and Value Types")
	fmt.Println("   ┌─────────────────┬──────────────────────────────────┐")
	fmt.Println("   │ Component       │ Allowed Types                    │")
	fmt.Println("   ├─────────────────┼──────────────────────────────────┤")
	fmt.Println("   │ Keys            │ Must be comparable (==, !=)      │")
	fmt.Println("   │                 │ ✅ int, string, bool, pointers   │")
	fmt.Println("   │                 │ ❌ slices, maps, functions       │")
	fmt.Println("   ├─────────────────┼──────────────────────────────────┤")
	fmt.Println("   │ Values          │ Any type (no restrictions)       │")
	fmt.Println("   │                 │ ✅ int, string, struct, slice... │")
	fmt.Println("   └─────────────────┴──────────────────────────────────┘\n")

	// Examples of different key/value types
	fmt.Printf("   Examples:\n")
	intKeys := map[int]string{1: "one", 2: "two"}
	fmt.Printf("   map[int]string:        %v\n", intKeys)

	boolKeys := map[bool]string{true: "yes", false: "no"}
	fmt.Printf("   map[bool]string:       %v\n", boolKeys)

	sliceValues := map[string][]int{"nums": {1, 2, 3}}
	fmt.Printf("   map[string][]int:      %v\n\n", sliceValues)

	// Section 6: Accessing Map Elements
	printMapSection("6. Accessing Map Elements")
	fmt.Printf("   scores := map[string]int{\"Math\": 95, \"English\": 88}\n\n")
	testScores := map[string]int{"Math": 95, "English": 88}

	fmt.Printf("   score := scores[\"Math\"]\n")
	score := testScores["Math"]
	fmt.Printf("   → %d\n\n", score)

	fmt.Printf("   Accessing non-existent key:\n")
	fmt.Printf("   score := scores[\"History\"]  // returns zero value\n")
	missingScore := testScores["History"]
	fmt.Printf("   → %d (zero value for int)\n\n", missingScore)

	// Section 7: Checking if Key Exists
	printMapSection("7. Checking if Key Exists (Comma Ok Idiom)")
	fmt.Printf("   value, exists := scores[\"Math\"]\n")
	fmt.Printf("   if exists {\n")
	fmt.Printf("       fmt.Println(\"Found:\", value)\n")
	fmt.Printf("   }\n\n")

	value, exists := testScores["Math"]
	if exists {
		fmt.Printf("   Output: Found: %d\n\n", value)
	}

	fmt.Printf("   value, exists := scores[\"History\"]\n")
	fmt.Printf("   if exists {\n")
	fmt.Printf("       fmt.Println(\"Found:\", value)\n")
	fmt.Printf("   } else {\n")
	fmt.Printf("       fmt.Println(\"Not found\")\n")
	fmt.Printf("   }\n\n")

	value2, exists2 := testScores["History"]
	if exists2 {
		fmt.Printf("   Output: Found: %d\n\n", value2)
	} else {
		fmt.Printf("   Output: Not found\n\n")
	}

	// Section 8: Adding Elements to Map
	printMapSection("8. Adding Elements to Map")
	fmt.Printf("   colors := make(map[string]string)\n")
	colors := make(map[string]string)
	fmt.Printf("   → %v (empty)\n\n", colors)

	fmt.Printf("   colors[\"red\"] = \"#FF0000\"\n")
	fmt.Printf("   colors[\"green\"] = \"#00FF00\"\n")
	colors["red"] = "#FF0000"
	colors["green"] = "#00FF00"
	fmt.Printf("   → %v\n\n", colors)

	// Section 9: Updating Map Elements
	printMapSection("9. Updating Map Elements")
	fmt.Printf("   Original: %v\n", colors)
	fmt.Printf("   colors[\"red\"] = \"#CC0000\"  // update existing key\n")
	colors["red"] = "#CC0000"
	fmt.Printf("   Updated:  %v\n\n", colors)

	// Section 10: Deleting Elements from Map
	printMapSection("10. Deleting Elements from Map")
	fmt.Printf("   Before delete: %v\n", colors)
	fmt.Printf("   delete(colors, \"green\")\n")
	delete(colors, "green")
	fmt.Printf("   After delete:  %v\n\n", colors)

	fmt.Printf("   Deleting non-existent key (safe, no error):\n")
	fmt.Printf("   delete(colors, \"blue\")\n")
	delete(colors, "blue")
	fmt.Printf("   → No error, map unchanged\n\n")

	// Section 11: Map Length
	printMapSection("11. Map Length")
	fmt.Printf("   scores := map[string]int{\"Math\": 95, \"English\": 88}\n")
	fmt.Printf("   len(scores) = %d\n\n", len(testScores))

	// Section 12: Iterating Over Maps
	printMapSection("12. Iterating Over Maps")
	fmt.Printf("   for key, value := range scores {\n")
	fmt.Printf("       fmt.Printf(\"%%s: %%d\", key, value)\n")
	fmt.Printf("   }\n\n")
	fmt.Printf("   Output:\n")
	for key, value := range testScores {
		fmt.Printf("   %s: %d\n", key, value)
	}
	fmt.Printf("   ⚠️  Order is NOT guaranteed!\n\n")

	// Section 13: Iterating - Keys Only
	printMapSection("13. Iterating - Keys Only")
	fmt.Printf("   for key := range scores {\n")
	fmt.Printf("       fmt.Println(key)\n")
	fmt.Printf("   }\n\n")
	fmt.Printf("   Output:\n")
	for key := range testScores {
		fmt.Printf("   %s\n", key)
	}
	fmt.Println()

	// Section 14: Iterating - Values Only
	printMapSection("14. Iterating - Values Only")
	fmt.Printf("   for _, value := range scores {\n")
	fmt.Printf("       fmt.Println(value)\n")
	fmt.Printf("   }\n\n")
	fmt.Printf("   Output:\n")
	for _, value := range testScores {
		fmt.Printf("   %d\n", value)
	}
	fmt.Println()

	// Section 15: Zero Value of Map (nil)
	printMapSection("15. Zero Value of Map (nil)")
	fmt.Printf("   var m map[string]int  // nil map\n")
	var m map[string]int
	fmt.Printf("   → m == nil: %t\n", m == nil)
	fmt.Printf("   → len(m): %d\n", len(m))
	fmt.Printf("   ⚠️  Cannot add to nil map! Use make() first.\n\n")

	// Section 16: Maps are Reference Types
	printMapSection("16. Maps are Reference Types")
	fmt.Printf("   original := map[string]int{\"a\": 1}\n")
	fmt.Printf("   copy := original\n")
	fmt.Printf("   copy[\"a\"] = 2\n\n")
	original := map[string]int{"a": 1}
	copy := original
	copy["a"] = 2
	fmt.Printf("   original: %v (modified!)\n", original)
	fmt.Printf("   copy:     %v\n", copy)
	fmt.Printf("   ⚠️  Both point to the same underlying data!\n\n")

	// Section 17: Maps with Struct Values
	printMapSection("17. Maps with Struct Values")
	type Person struct {
		name string
		age  int
	}

	fmt.Printf("   type Person struct { name string; age int }\n\n")
	fmt.Printf("   people := map[string]Person{\n")
	fmt.Printf("       \"emp1\": {\"Alice\", 30},\n")
	fmt.Printf("       \"emp2\": {\"Bob\", 25},\n")
	fmt.Printf("   }\n\n")

	people := map[string]Person{
		"emp1": {"Alice", 30},
		"emp2": {"Bob", 25},
	}

	fmt.Printf("   people[\"emp1\"].name = \"%s\"\n", people["emp1"].name)
	fmt.Printf("   people[\"emp2\"].age  = %d\n\n", people["emp2"].age)

	// Section 18: Nested Maps
	printMapSection("18. Nested Maps")
	fmt.Printf("   grades := map[string]map[string]int{\n")
	fmt.Printf("       \"Alice\": {\"Math\": 95, \"English\": 88},\n")
	fmt.Printf("       \"Bob\":   {\"Math\": 82, \"English\": 90},\n")
	fmt.Printf("   }\n\n")

	grades := map[string]map[string]int{
		"Alice": {"Math": 95, "English": 88},
		"Bob":   {"Math": 82, "English": 90},
	}

	fmt.Printf("   grades[\"Alice\"][\"Math\"] = %d\n\n", grades["Alice"]["Math"])

	// Section 19: Practical Examples
	printMapSection("19. Practical Examples")

	// Example 1: Count word frequency
	fmt.Printf("   Example 1: Word Frequency Counter\n")
	words := []string{"apple", "banana", "apple", "cherry", "banana", "apple"}
	frequency := make(map[string]int)
	for _, word := range words {
		frequency[word]++
	}
	fmt.Printf("   Words: %v\n", words)
	fmt.Printf("   Frequency: %v\n\n", frequency)

	// Example 2: Group by category
	fmt.Printf("   Example 2: Group Items by Category\n")
	items := map[string]string{
		"apple":    "fruit",
		"carrot":   "vegetable",
		"banana":   "fruit",
		"broccoli": "vegetable",
	}

	categories := make(map[string][]string)
	for item, category := range items {
		categories[category] = append(categories[category], item)
	}
	fmt.Printf("   Grouped: %v\n\n", categories)

	printMapFooter()
}

func printMapHeader(title string) {
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Printf("  %s\n", title)
	fmt.Println(strings.Repeat("=", 60) + "\n")
}

func printMapSection(title string) {
	fmt.Printf("┌─ %s\n", title)
	fmt.Println("│")
}

func printMapFooter() {
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println("  ✅ Tutorial Complete!")
	fmt.Println("  💡 Key Takeaways:")
	fmt.Println("     • Maps store key-value pairs (unordered)")
	fmt.Println("     • Keys must be unique and comparable")
	fmt.Println("     • Use value, ok := map[key] to check if key exists")
	fmt.Println("     • Use delete(map, key) to remove elements")
	fmt.Println("     • Maps are reference types (modifications affect all refs)")
	fmt.Println("     • Nil maps cannot be written to (use make() first)")
	fmt.Println(strings.Repeat("=", 60) + "\n")
}
