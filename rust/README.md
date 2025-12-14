# 🦀 Rust Programming Tutorial

An interactive, comprehensive tutorial for learning Rust programming language from basics to advanced concepts.

## 📖 About This Project

This is a hands-on, interactive tutorial designed to teach Rust programming through 15 comprehensive modules. Each module includes detailed explanations, code examples, and practical demonstrations to help you master Rust's unique features like ownership, borrowing, and safe concurrency.

## 🎯 Features

- **Interactive Learning**: Menu-driven interface to choose topics
- **Comprehensive Coverage**: 15 core Rust programming topics
- **Practical Examples**: Real-world code demonstrations
- **Beautiful Formatting**: Clean, readable output with visual separators
- **Self-Paced**: Learn at your own speed
- **Safety First**: Focus on Rust's memory safety guarantees

## 🚀 Getting Started

### Prerequisites
- Rust and Cargo installed (via [rustup](https://rustup.rs/))

### Running the Tutorial

```bash
cd "d:\personal\langLearning\rust"
cargo run
```

## 📚 Topics Covered

### 1. Variables
Learn about variable declaration and mutability:
- **Immutability by Default**: `let x = 5;`
- **Mutability**: `let mut x = 5;`
- **Shadowing**: Re-declaring variables with `let`
- **Constants**: `const MAX_POINTS: u32 = 100_000;`

### 2. Data Types
Exploring Rust's static type system:
- **Scalar Types**: Integers, Floating-point, Booleans, Characters
- **Compound Types**: Tuples, Arrays
- **Type Inference**: Compiler deductions vs explicit annotations

### 3. Operators
Arithmetic and logical operations:
- **Arithmetic**: `+`, `-`, `*`, `/`, `%`
- **Comparison**: `==`, `!=`, `>`, `<`, `>=`, `<=`
- **Logical**: `&&` (AND), `||` (OR), `!` (NOT)
- **Assignment**: `+=`, `-=`, etc.

### 4. Conditions
Control flow with `if`:
- **If Expressions**: Returns values (e.g., `let number = if condition { 5 } else { 6 };`)
- **Else If**: Handling multiple conditions
- **Type Consistency**: Branches must return same types

### 5. Loops
Iteration mechanisms:
- **loop**: Infinite loops with `break` values
- **while**: Conditional looping
- **for**: Iterating over collections and ranges
- **Loop Labels**: Breaking outer loops

### 6. Functions
Reusable code blocks:
- **Parameters**: Typed arguments
- **Return Values**: Expressions (no semicolon) vs Statements
- **Scopes**: Block-based variable lifetimes

### 7. Strings
Text handling:
- **String**: Heap-allocated, growable
- **&str**: String slice, immutable reference
- **String Creation**: `String::from()`, `.to_string()`
- **Modification**: `push_str()`, format macro

### 8. Ownership
Rust's core memory management feature:
- **Ownership Rules**: Single owner per value
- **Move Semantics**: Transferring ownership
- **Copy Trait**: Stack-only data copying
- **Clone**: Deep copying heap data

### 9. Borrowing
Accessing data without taking ownership:
- **References**: `&T`
- **Mutable References**: `&mut T` (One at a time)
- **Borrowing Rules**: Multiple readers OR one writer
- **Dangling References**: Compile-time prevention

### 10. Arrays
Fixed-size stack collections:
- **Declaration**: `[T; N]`
- **Slices**: Views into arrays
- **Iteration**: Looping over arrays

### 11. Vectors
Growable heap collections:
- **Creation**: `vec![]`, `Vec::new()`
- **Modification**: `push`, `pop`, `insert`, `remove`
- **Safe Access**: `.get()` vs `[]` panic
- **Performance**: Front vs Back operations

### 12. Tuples
Groupings of mixed types:
- **Creation**: `(i32, f64, u8)`
- **Access**: `t.0`, `t.1`
- **Destructuring**: `let (x, y, z) = t;`
- **Unit Type**: `()`

### 13. Hash Maps
Key-value storage:
- **Creation**: `HashMap::new()`
- **Entry API**: `.entry().or_insert()`
- **Ownership**: Key ownership transfer

### 14. Structs
Custom data types:
- **Named Fields**: Standard structs
- **Tuple Structs**: Named tuples
- **Methods**: `impl` blocks & `self`
- **Associated Functions**: Constructors (e.g., `new`)

### 15. Enums
Enumerated types and pattern matching:
- **Variants**: Simple, Tuple, Struct variants
- **Option Enum**: `Some(T)` vs `None`
- **Match Expressions**: Exhaustive pattern matching

---

## 🎨 Project Structure

```
rust/
├── src/
│   ├── main.rs        # Main interactive menu
│   ├── variables.rs   # Variables tutorial
│   ├── data_types.rs  # Data types tutorial
│   ├── operators.rs   # Operators tutorial
│   ├── conditions.rs  # Conditions tutorial
│   ├── loops.rs       # Loops tutorial
│   ├── functions.rs   # Functions tutorial
│   ├── strings.rs     # Strings tutorial
│   ├── ownership.rs   # Ownership tutorial
│   ├── borrowing.rs   # Borrowing tutorial
│   ├── arrays.rs      # Arrays tutorial
│   ├── vectors.rs     # Vectors tutorial
│   ├── tuples.rs      # Tuples tutorial
│   ├── hash_map.rs    # Hash Maps tutorial
│   ├── structs.rs     # Structs tutorial
│   └── enums.rs       # Enums tutorial
└── README.md          # This file
```

## 💡 Learning Tips

1. **Understand Ownership**: It's unique to Rust. Spend extra time here.
2. **Read Compiler Errors**: rustc is very helpful; read the suggestions!
3. **Experiment**: Try removing `mut`, cloning vs borrowing, etc.
4. **Use cargo check**: Fast syntax checking during development.

## 🎯 What You'll Learn

By completing this tutorial, you'll understand:
- ✅ Rust's memory safety without garbage collection
- ✅ Ownership, Borrowing, and Lifetimes
- ✅ Strong static typing and Type Inference
- ✅ Modern error handling with Option and Result
- ✅ Data structures and Control Flow in Rust

## 🤝 Contributing

This is a personal learning project. Feel free to copy and modify for your own learning!

## 📝 License

This project is for educational purposes.

## 🙏 Acknowledgments

Created as a comprehensive, interactive learning resource for Rust programming fundamentals.

---

**Happy Learning! 🦀**

*Keep coding and have fun with Rust!*
