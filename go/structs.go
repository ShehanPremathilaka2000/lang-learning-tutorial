package main

import (
	"fmt"
	"strings"
)

func structs() {
	printStructHeader("GO STRUCTS TUTORIAL")

	// Section 1: What are Structs?
	printStructSection("1. What are Structs?")
	fmt.Println("   Structs are user-defined types that group related data together.")
	fmt.Println("   ✅ Collection of fields with different data types")
	fmt.Println("   ✅ Similar to classes in other languages (but no inheritance)")
	fmt.Println("   ✅ Used to create custom data structures\n")

	// Section 2: Defining a Struct
	printStructSection("2. Defining a Struct")
	fmt.Printf("   type Person struct {\n")
	fmt.Printf("       name    string\n")
	fmt.Printf("       age     int\n")
	fmt.Printf("       city    string\n")
	fmt.Printf("   }\n\n")

	// Section 3: Creating Struct Instances
	printStructSection("3. Creating Struct Instances")

	// Method 1: Using field names
	fmt.Printf("   Method 1: Using field names\n")
	fmt.Printf("   person1 := Person{name: \"Alice\", age: 30, city: \"NYC\"}\n")
	person1 := Person{name: "Alice", age: 30, city: "NYC"}
	fmt.Printf("   → %+v\n\n", person1)

	// Method 2: Positional values
	fmt.Printf("   Method 2: Positional values (order matters)\n")
	fmt.Printf("   person2 := Person{\"Bob\", 25, \"LA\"}\n")
	person2 := Person{"Bob", 25, "LA"}
	fmt.Printf("   → %+v\n\n", person2)

	// Method 3: Partial initialization
	fmt.Printf("   Method 3: Partial initialization (rest are zero values)\n")
	fmt.Printf("   person3 := Person{name: \"Charlie\"}\n")
	person3 := Person{name: "Charlie"}
	fmt.Printf("   → %+v\n\n", person3)

	// Method 4: Zero value struct
	fmt.Printf("   Method 4: Zero value struct\n")
	fmt.Printf("   var person4 Person\n")
	var person4 Person
	fmt.Printf("   → %+v\n\n", person4)

	// Section 4: Accessing Struct Fields
	printStructSection("4. Accessing Struct Fields")
	fmt.Printf("   person := Person{name: \"David\", age: 35, city: \"Chicago\"}\n\n")
	person := Person{name: "David", age: 35, city: "Chicago"}

	fmt.Printf("   Reading fields:\n")
	fmt.Printf("   person.name = \"%s\"\n", person.name)
	fmt.Printf("   person.age  = %d\n", person.age)
	fmt.Printf("   person.city = \"%s\"\n\n", person.city)

	// Section 5: Modifying Struct Fields
	printStructSection("5. Modifying Struct Fields")
	fmt.Printf("   Original: %+v\n", person)
	person.age = 36
	person.city = "Boston"
	fmt.Printf("   person.age = 36\n")
	fmt.Printf("   person.city = \"Boston\"\n")
	fmt.Printf("   Modified: %+v\n\n", person)

	// Section 6: Anonymous Structs
	printStructSection("6. Anonymous Structs (No Type Name)")
	fmt.Printf("   point := struct {\n")
	fmt.Printf("       x int\n")
	fmt.Printf("       y int\n")
	fmt.Printf("   }{x: 10, y: 20}\n\n")
	point := struct {
		x int
		y int
	}{x: 10, y: 20}
	fmt.Printf("   → point: %+v\n\n", point)

	// Section 7: Nested Structs
	printStructSection("7. Nested Structs")
	fmt.Printf("   type Address struct {\n")
	fmt.Printf("       street  string\n")
	fmt.Printf("       city    string\n")
	fmt.Printf("       zipCode string\n")
	fmt.Printf("   }\n\n")
	fmt.Printf("   type Employee struct {\n")
	fmt.Printf("       name    string\n")
	fmt.Printf("       age     int\n")
	fmt.Printf("       address Address\n")
	fmt.Printf("   }\n\n")

	emp := Employee{
		name: "Eve",
		age:  28,
		address: Address{
			street:  "123 Main St",
			city:    "Seattle",
			zipCode: "98101",
		},
	}
	fmt.Printf("   emp := Employee{...}\n")
	fmt.Printf("   → %+v\n", emp)
	fmt.Printf("   → emp.address.city = \"%s\"\n\n", emp.address.city)

	// Section 8: Pointers to Structs
	printStructSection("8. Pointers to Structs")
	fmt.Printf("   p1 := Person{name: \"Frank\", age: 40, city: \"Miami\"}\n")
	fmt.Printf("   p2 := &p1  // pointer to p1\n\n")
	p1 := Person{name: "Frank", age: 40, city: "Miami"}
	p2 := &p1

	fmt.Printf("   p1: %+v\n", p1)
	fmt.Printf("   p2: %+v\n", *p2)
	fmt.Printf("   p2.age = 41  // Go auto-dereferences\n")
	p2.age = 41
	fmt.Printf("   p1: %+v (modified via pointer)\n\n", p1)

	// Section 9: Passing Structs to Functions (By Value)
	printStructSection("9. Passing Structs to Functions (By Value)")
	fmt.Printf("   func printPerson(p Person) {\n")
	fmt.Printf("       fmt.Printf(\"%%s is %%d years old\", p.name, p.age)\n")
	fmt.Printf("   }\n\n")

	testPerson := Person{name: "Grace", age: 32, city: "Denver"}
	fmt.Printf("   Original: %+v\n", testPerson)
	fmt.Printf("   Calling: printPerson(testPerson)\n")
	fmt.Printf("   Output: ")
	printPerson(testPerson)
	fmt.Printf("   Note: Function receives a COPY of the struct\n\n")

	// Section 10: Passing Structs by Pointer
	printStructSection("10. Passing Structs by Pointer (Modify Original)")
	fmt.Printf("   func updateAge(p *Person, newAge int) {\n")
	fmt.Printf("       p.age = newAge\n")
	fmt.Printf("   }\n\n")

	fmt.Printf("   Before: %+v\n", testPerson)
	fmt.Printf("   Calling: updateAge(&testPerson, 33)\n")
	updateAge(&testPerson, 33)
	fmt.Printf("   After:  %+v (modified!)\n\n", testPerson)

	// Section 11: Struct Methods (Receiver Functions)
	printStructSection("11. Struct Methods (Receiver Functions)")
	fmt.Printf("   func (p Person) introduce() {\n")
	fmt.Printf("       fmt.Printf(\"Hi, I'm %%s\", p.name)\n")
	fmt.Printf("   }\n\n")

	methodPerson := Person{name: "Henry", age: 45, city: "Austin"}
	fmt.Printf("   Calling: methodPerson.introduce()\n")
	fmt.Printf("   Output: ")
	methodPerson.introduce()
	fmt.Println()

	// Section 12: Pointer Receiver Methods
	printStructSection("12. Pointer Receiver Methods (Can Modify)")
	fmt.Printf("   func (p *Person) haveBirthday() {\n")
	fmt.Printf("       p.age++\n")
	fmt.Printf("   }\n\n")

	fmt.Printf("   Before: %+v\n", methodPerson)
	fmt.Printf("   Calling: methodPerson.haveBirthday()\n")
	methodPerson.haveBirthday()
	fmt.Printf("   After:  %+v (age incremented!)\n\n", methodPerson)

	// Section 13: Struct Comparison
	printStructSection("13. Struct Comparison")
	p3 := Person{name: "Ivy", age: 28, city: "Portland"}
	p4 := Person{name: "Ivy", age: 28, city: "Portland"}
	p5 := Person{name: "Jack", age: 30, city: "Phoenix"}

	fmt.Printf("   p3 := Person{name: \"Ivy\", age: 28, city: \"Portland\"}\n")
	fmt.Printf("   p4 := Person{name: \"Ivy\", age: 28, city: \"Portland\"}\n")
	fmt.Printf("   p5 := Person{name: \"Jack\", age: 30, city: \"Phoenix\"}\n\n")
	fmt.Printf("   p3 == p4: %t (same values)\n", p3 == p4)
	fmt.Printf("   p3 == p5: %t (different values)\n\n", p3 == p5)

	// Section 14: Struct Tags (Metadata)
	printStructSection("14. Struct Tags (Metadata for JSON, etc.)")
	fmt.Printf("   type User struct {\n")
	fmt.Printf("       Name  string `json:\"name\"`\n")
	fmt.Printf("       Email string `json:\"email\"`\n")
	fmt.Printf("       Age   int    `json:\"age,omitempty\"`\n")
	fmt.Printf("   }\n\n")
	fmt.Printf("   Tags are used for JSON encoding/decoding, validation, etc.\n\n")

	// Section 15: Empty Struct
	printStructSection("15. Empty Struct (Zero Memory)")
	fmt.Printf("   type Empty struct{}\n\n")
	fmt.Printf("   var e Empty\n")
	var e Empty
	fmt.Printf("   → Size: 0 bytes (useful for sets, signals)\n")
	fmt.Printf("   → Value: %+v\n\n", e)

	// Section 16: Returning Structs from Functions
	printStructSection("16. Returning Structs from Functions")
	fmt.Printf("   func createPerson(name string, age int) Person {\n")
	fmt.Printf("       return Person{name: name, age: age, city: \"Unknown\"}\n")
	fmt.Printf("   }\n\n")

	newPerson := createPerson("Kate", 27)
	fmt.Printf("   newPerson := createPerson(\"Kate\", 27)\n")
	fmt.Printf("   → %+v\n\n", newPerson)

	// Section 17: Practical Example - Rectangle
	printStructSection("17. Practical Example - Rectangle")
	rect := Rectangle{width: 10, height: 5}
	fmt.Printf("   rect := Rectangle{width: 10, height: 5}\n")
	fmt.Printf("   rect.area() = %d\n", rect.area())
	fmt.Printf("   rect.perimeter() = %d\n\n", rect.perimeter())

	printStructFooter()
}

// Struct definitions for demonstrations

type Person struct {
	name string
	age  int
	city string
}

type Address struct {
	street  string
	city    string
	zipCode string
}

type Employee struct {
	name    string
	age     int
	address Address
}

type Rectangle struct {
	width  int
	height int
}

type Empty struct{}

// Helper functions for demonstrations

func printPerson(p Person) {
	fmt.Printf("%s is %d years old\n", p.name, p.age)
}

func updateAge(p *Person, newAge int) {
	p.age = newAge
}

func createPerson(name string, age int) Person {
	return Person{name: name, age: age, city: "Unknown"}
}

// Methods for Person struct

func (p Person) introduce() {
	fmt.Printf("Hi, I'm %s from %s\n", p.name, p.city)
}

func (p *Person) haveBirthday() {
	p.age++
}

// Methods for Rectangle struct

func (r Rectangle) area() int {
	return r.width * r.height
}

func (r Rectangle) perimeter() int {
	return 2 * (r.width + r.height)
}

// Print helper functions

func printStructHeader(title string) {
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Printf("  %s\n", title)
	fmt.Println(strings.Repeat("=", 60) + "\n")
}

func printStructSection(title string) {
	fmt.Printf("┌─ %s\n", title)
	fmt.Println("│")
}

func printStructFooter() {
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println("  ✅ Tutorial Complete!")
	fmt.Println("  💡 Key Takeaways:")
	fmt.Println("     • Structs group related data of different types")
	fmt.Println("     • Access fields with dot notation: person.name")
	fmt.Println("     • Pass by value (copy) or by pointer (reference)")
	fmt.Println("     • Use pointer receivers to modify struct in methods")
	fmt.Println("     • Structs can be compared with == if all fields comparable")
	fmt.Println("     • Use struct tags for metadata (JSON, validation, etc.)")
	fmt.Println(strings.Repeat("=", 60) + "\n")
}
