pub fn tuples() {
    println!("\n========================================");
    println!("        RUST TUPLES TUTORIAL");
    println!("========================================\n");

    // ========================================
    // 1. CREATING TUPLES
    // ========================================
    println!("--- 1. Creating Tuples ---");
    println!("Code:");
    println!("  let person: (&str, i32) = (\"Alice\", 30); // Mixed types");
    println!("  let coordinates = (10.5, 20.1, 8.8); // Inferred");
    println!("  println!(\"Person: {{:?}}\", person);");
    println!("\nOutput:");

    let person: (&str, i32) = ("Alice", 30);
    let coordinates = (10.5, 20.1, 8.8);
    println!("  Person: {:?}", person);
    println!("  Coordinates: {:?}\n", coordinates);

    // ========================================
    // 2. ACCESSING VALUES (Dot Notation)
    // ========================================
    println!("--- 2. Accessing Values (Dot Notation) ---");
    println!("Code:");
    println!("  let name = person.0;");
    println!("  let age = person.1;");
    println!("  println!(\"Name: {{}}, Age: {{}}\", name, age);");
    println!("\nOutput:");

    let name = person.0;
    let age = person.1;
    println!("  Name: {}, Age: {}\n", name, age);

    // ========================================
    // 3. DESTRUCTURING (Unpacking)
    // ========================================
    println!("--- 3. Destructuring (Unpacking) ---");
    println!("Code:");
    println!("  let (x, y, z) = coordinates;");
    println!("  println!(\"x: {{}}, y: {{}}, z: {{}}\", x, y, z);");
    println!("\nOutput:");

    let (x, y, z) = coordinates;
    println!("  x: {}, y: {}, z: {}\n", x, y, z);

    // ========================================
    // 4. RETURNING TUPLES
    // ========================================
    println!("--- 4. Returning Tuples ---");
    println!("Code:");
    println!("  let (quotient, remainder) = divide(10, 3);");
    println!("  println!(\"10 / 3 = {{}} remainder {{}}\", quotient, remainder);");
    println!("\nOutput:");

    let (quotient, remainder) = divide(10, 3);
    println!("  10 / 3 = {} remainder {}\n", quotient, remainder);

    // ========================================
    // KEY TAKEAWAYS
    // ========================================
    println!("========================================");
    println!("         KEY TAKEAWAYS");
    println!("========================================\n");
    println!("1. MIXED TYPES");
    println!("   - Tuples can hold different types: (i32, f64, &str).");
    println!("   - Fixed length once declared.\n");

    println!("2. ACCESSING");
    println!("   - Use dot notation 't.0', 't.1' for direct access.");
    println!("   - Use destructuring 'let (a, b) = t' to unpack.\n");

    println!("3. RETURNING MULTIPLE VALUES");
    println!("   - Tuples are the standard way to return multiple values from a function.\n");

    println!("4. UNIT TYPE ()");
    println!("   - An empty tuple () is called the 'unit type'.");
    println!("   - Functions with no return value implicitly return ().\n");

    println!("========================================");
    println!("       END OF TUPLES TUTORIAL");
    println!("========================================\n");
}

fn divide(dividend: i32, divisor: i32) -> (i32, i32) {
    (dividend / divisor, dividend % divisor)
}
