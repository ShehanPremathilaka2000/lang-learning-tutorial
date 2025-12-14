# 🚀 Go Programming Tutorial

An interactive, comprehensive tutorial for learning Go programming language from basics to advanced concepts.

## 📖 About This Project

This is a hands-on, interactive tutorial designed to teach Go programming through 12 comprehensive modules. Each module includes detailed explanations, code examples, and practical demonstrations to help you master Go.

## 🎯 Features

- **Interactive Learning**: Menu-driven interface to choose topics
- **Comprehensive Coverage**: 12 core Go programming topics
- **Practical Examples**: Real-world code demonstrations
- **Beautiful Formatting**: Clean, readable output with visual separators
- **Self-Paced**: Learn at your own speed

## 🚀 Getting Started

### Prerequisites
- Go 1.16 or higher installed on your system

### Running the Tutorial

```bash
cd "d:\personal\go learning"
go run .
```

## 📚 Topics Covered

### 1. Variables
Learn about variable declaration and initialization in Go:
- **Explicit Type Declaration**: `var variable1 string = "Hello"`
- **Type Inference**: `var variable2 = "Auto-typed"`
- **Short Declaration**: `variable3 := "Quick syntax"`
- **Zero Values**: Default values for uninitialized variables
- **Multiple Variables**: Declaring multiple variables at once
- **Grouped Declarations**: Using `var ()` blocks

**Key Concepts**: camelCase naming, := operator (function-only), var keyword (package-level)

---

### 2. Constants
Understanding immutable values in Go:
- **Basic Constants**: `const Pi = 3.14159`
- **Typed Constants**: `const MaxUsers int = 1000`
- **Grouped Constants**: Multiple constants in blocks
- **Iota**: Auto-incrementing constant generator
- **Iota with Expressions**: Powers of 2, bit flags
- **Package-Level Constants**: Exported vs unexported

**Key Concepts**: Compile-time evaluation, no := syntax, iota for enumerations

---

### 3. Data Types
Exploring Go's type system:
- **Boolean**: `bool` (true/false)
- **Integers**: `int`, `int8`, `int16`, `int32`, `int64`
- **Unsigned Integers**: `uint`, `uint8` (byte), `uint16`, `uint32`
- **Floating Point**: `float32`, `float64`
- **Strings**: UTF-8 encoded text
- **Runes**: `rune` (alias for int32, Unicode characters)
- **Complex Numbers**: `complex64`, `complex128`
- **Type Conversion**: Explicit conversion required

**Key Concepts**: Zero values, type safety, explicit conversions

---

### 4. Arrays
Fixed-size collections in Go:
- **Declaration**: `var arr [5]int`
- **Initialization**: `arr := [3]int{1, 2, 3}`
- **Inferred Size**: `arr := [...]int{1, 2, 3, 4}`
- **Accessing Elements**: Zero-indexed `arr[0]`
- **Modifying Elements**: `arr[2] = 10`
- **Array Length**: `len(arr)`
- **Partial Initialization**: Remaining elements are zero values
- **Index-Based Init**: `arr := [5]int{1: 10, 3: 30}`
- **Multi-Dimensional**: `matrix := [2][3]int{...}`
- **Array Comparison**: Can compare with `==`

**Key Concepts**: Fixed size, value types, size is part of type

---

### 5. Slices
Dynamic, flexible arrays in Go:
- **Creation**: `slice := []int{1, 2, 3}` or `make([]int, 5)`
- **Length vs Capacity**: `len()` vs `cap()`
- **Slicing Arrays**: `slice := array[1:4]`
- **Appending**: `slice = append(slice, 4, 5, 6)`
- **Copying**: `copy(dest, src)` for independent copies
- **Range Iteration**: `for i, v := range slice`
- **Nil Slices**: `var slice []int` (nil but usable)
- **Dynamic Growth**: Capacity doubles when needed
- **Reference Types**: Modifications affect all references

**Key Concepts**: Most common collection type, dynamic sizing, underlying arrays

---

### 6. Operators
All Go operators explained:
- **Arithmetic**: `+`, `-`, `*`, `/`, `%`
- **Assignment**: `=`, `+=`, `-=`, `*=`, `/=`, `%=`
- **Increment/Decrement**: `++`, `--` (postfix only)
- **Comparison**: `==`, `!=`, `>`, `<`, `>=`, `<=`
- **Logical**: `&&` (AND), `||` (OR), `!` (NOT)
- **Bitwise**: `&`, `|`, `^`, `~`
- **Bit Shift**: `<<`, `>>`
- **Operator Precedence**: Order of operations

**Key Concepts**: No ternary operator, strict type matching

---

### 7. Conditions
Control flow with if-else and switch:

#### If-Else Statements
- **Simple if**: `if condition { }`
- **if-else**: `if condition { } else { }`
- **if-else if-else**: Multiple conditions
- **Nested if**: Conditions within conditions
- **Short Statement**: `if x := getValue(); x > 0 { }`

#### Switch Statements
- **Basic Switch**: `switch value { case 1: ... }`
- **Multiple Values**: `case 'a', 'e', 'i', 'o', 'u':`
- **No Expression**: `switch { case condition: ... }`
- **Type Switch**: `switch v := i.(type)`
- **Fallthrough**: Explicit fall-through to next case
- **Default Case**: Handles unmatched values

**Key Concepts**: Opening brace on same line, no break needed in switch, automatic break

---

### 8. Loops
Iteration in Go (only `for` loop):
- **Classic For**: `for i := 0; i < 10; i++ { }`
- **While-Style**: `for condition { }`
- **Infinite Loop**: `for { }`
- **Range Over Arrays/Slices**: `for i, v := range arr { }`
- **Range Over Maps**: `for key, value := range map { }`
- **Range Over Strings**: Iterates over runes
- **Break**: Exit loop early
- **Continue**: Skip to next iteration
- **Labeled Break/Continue**: For nested loops

**Key Concepts**: No while/do-while keywords, range for iteration

---

### 9. Functions
Reusable code blocks:
- **Basic Functions**: `func name() { }`
- **Parameters**: `func greet(name string) { }`
- **Return Values**: `func add(a, b int) int { }`
- **Multiple Returns**: `func divide(a, b float64) (float64, error)`
- **Named Returns**: `func calc() (sum, product int) { }`
- **Variadic Functions**: `func sum(nums ...int) int`
- **Anonymous Functions**: `func() { }()`
- **Closures**: Functions returning functions
- **Recursion**: Functions calling themselves
- **Defer**: Execute after function returns (LIFO)

**Naming**: camelCase (private), PascalCase (public/exported)

**Key Concepts**: First-class functions, multiple returns common, defer for cleanup

---

### 10. Structs
Custom data types grouping fields:
- **Definition**: `type Person struct { name string; age int }`
- **Creation**: Named fields, positional, partial initialization
- **Accessing Fields**: `person.name`
- **Modifying Fields**: `person.age = 30`
- **Anonymous Structs**: One-time use structs
- **Nested Structs**: Structs within structs
- **Pointers to Structs**: `&person`, auto-dereferencing
- **Pass by Value**: Functions receive copies
- **Pass by Pointer**: Modify original struct
- **Methods**: `func (p Person) method() { }`
- **Pointer Receivers**: `func (p *Person) modify() { }`
- **Struct Tags**: Metadata for JSON, validation
- **Comparison**: Can compare with `==`

**Key Concepts**: No inheritance, composition over inheritance, value vs pointer receivers

---

### 11. Maps
Key-value pairs (hash tables):
- **Creation**: `make(map[string]int)` or `map[string]int{}`
- **Adding Elements**: `map[key] = value`
- **Accessing Elements**: `value := map[key]`
- **Checking Existence**: `value, exists := map[key]`
- **Updating**: Same as adding
- **Deleting**: `delete(map, key)`
- **Iteration**: `for key, value := range map`
- **Length**: `len(map)`
- **Nil Maps**: Cannot write to nil maps
- **Reference Types**: Modifications affect all references
- **Allowed Keys**: Must be comparable (no slices, maps, functions)
- **Allowed Values**: Any type

**Key Concepts**: Unordered, unique keys, fast lookups

---

### 12. Defer
Delayed execution for cleanup:
- **Basic Defer**: `defer fmt.Println("cleanup")`
- **Execution Time**: Just before function returns
- **Argument Evaluation**: Evaluated immediately
- **LIFO Order**: Last deferred runs first
- **With Panics**: Still executes on panic
- **Anonymous Functions**: Capture by reference
- **Named Returns**: Can modify return values
- **Common Patterns**: 
  - File closing: `defer file.Close()`
  - Mutex unlocking: `defer mu.Unlock()`
  - Timing: `defer func() { log(time.Since(start)) }()`

**Key Concepts**: Guaranteed cleanup, stack-based execution, panic-safe

---

## 🎨 Project Structure

```
go learning/
├── helloWorld.go      # Main interactive menu
├── variables.go       # Variables tutorial
├── constants.go       # Constants tutorial
├── dataTypes.go       # Data types tutorial
├── arrays.go          # Arrays tutorial
├── slices.go          # Slices tutorial
├── operators.go       # Operators tutorial
├── conditions.go      # Conditions & switch tutorial
├── loops.go           # Loops tutorial
├── functions.go       # Functions tutorial
├── structs.go         # Structs tutorial
├── maps.go            # Maps tutorial
├── defer.go           # Defer tutorial
└── README.md          # This file
```

## 💡 Learning Tips

1. **Start Sequential**: Begin with Variables and progress through topics in order
2. **Practice**: Try modifying the examples and observe the results
3. **Experiment**: Break things to understand error messages
4. **Review**: Revisit topics as needed - the interactive menu makes it easy
5. **Build Projects**: Apply what you learn in small projects

## 🎯 What You'll Learn

By completing this tutorial, you'll understand:
- ✅ Go's type system and variable declarations
- ✅ Control flow (if, switch, for loops)
- ✅ Data structures (arrays, slices, maps, structs)
- ✅ Functions and methods
- ✅ Resource management with defer
- ✅ Go's unique features and idioms

## 🤝 Contributing

This is a personal learning project. Feel free to fork and customize for your own learning journey!

## 📝 License

This project is for educational purposes.

## 🙏 Acknowledgments

Created as a comprehensive, interactive learning resource for Go programming fundamentals.

---

**Happy Learning! 🚀**

*Keep coding and have fun with Go!*
