pub fn ownership() {
    println!("\n========================================");
    println!("      RUST OWNERSHIP TUTORIAL");
    println!("========================================\n");

    // ========================================
    // 1. STACK VS HEAP DATA
    // ========================================
    println!("--- 1. Stack vs Heap Data ---");
    println!("STACK DATA (Fixed Size, Fast, Copy):");
    println!("  - Integers (i32, u8, etc.)");
    println!("  - Booleans (bool)");
    println!("  - Characters (char)");
    println!("  - Floating point (f64)");
    println!("  - Arrays/Tuples (if containing only stack data)\n");

    println!("HEAP DATA (Unknown Size, Slower, Move):");
    println!("  - String (not &str)");
    println!("  - Vectors (Vec<T>)");
    println!("  - Box<T>");
    println!("  - Custom structs (usually)\n");

    // ========================================
    // 2. MOVE SEMANTICS (Ownership Transfer)
    // ========================================
    println!("--- 2. Move Semantics (Ownership Transfer) ---");
    println!("Code:");
    println!("  let s1 = String::from(\"Hello\");");
    println!("  let s2 = s1; // OWNERSHIP MOVED to s2");
    println!("  // println!(\"{{}}\", s1); // ERROR: s1 is invalid!");
    println!("  println!(\"s2 owns the data: {{}}\", s2);");
    println!("\nOutput:");

    let s1 = String::from("Hello");
    let s2 = s1; // Move occurs here
    // println!("{}", s1); // This would cause a compile error
    println!("  s2 owns the data: {}\n", s2);

    // ========================================
    // 3. COPY TRAIT (Stack-Only Data)
    // ========================================
    println!("--- 3. Copy Trait (Stack Data) ---");
    println!("Code:");
    println!("  let x = 5;");
    println!("  let y = x; // COPY occurs (integers are Copy)");
    println!("  println!(\"x: {{}}, y: {{}} // Both are valid\", x, y);");
    println!("\nOutput:");

    let x = 5;
    let y = x; // Copy occurs
    println!("  x: {}, y: {}\n", x, y);

    // ========================================
    // 4. CLONE (Deep Copy)
    // ========================================
    println!("--- 4. Clone (Deep Copy) ---");
    println!("Code:");
    println!("  let s3 = String::from(\"World\");");
    println!("  let s4 = s3.clone(); // Explicit deep copy");
    println!("  println!(\"s3: {{}}, s4: {{}} // Both are valid\", s3, s4);");
    println!("\nOutput:");

    let s3 = String::from("World");
    let s4 = s3.clone();
    println!("  s3: {}, s4: {}\n", s3, s4);

    // ========================================
    // 5. REFERENCES & BORROWING
    // ========================================
    println!("--- 5. References & Borrowing ---");
    println!("Code:");
    println!("  let s5 = String::from(\"Rust\");");
    println!("  let len = calculate_length(&s5); // Pass by Reference (Borrow)");
    println!("  println!(\"The length of '{{}}' is {{}}.\", s5, len);");
    println!("  // s5 is still valid because we borrowed it, didn't move it");
    println!("\nOutput:");

    let s5 = String::from("Rust");
    let len = calculate_length(&s5);
    println!("  The length of '{}' is {}.\n", s5, len);

    // ========================================
    // KEY TAKEAWAYS
    // ========================================
    println!("========================================");
    println!("         KEY TAKEAWAYS");
    println!("========================================\n");
    println!("1. STACK MEMORY");
    println!("   - Fast access, fixed size at compile time.");
    println!("   - Data is COPIED automatically (Copy Trait).");
    println!("   - Examples: Integers, bool, char, Arrays.\n");

    println!("2. HEAP MEMORY");
    println!("   - Slower access, growable size at runtime.");
    println!("   - Data is MOVED automatically (Ownership Transfer).");
    println!("   - Examples: String, Vec<T>, Box<T>.\n");

    println!("3. OWNERSHIP RULES");
    println!("   - One owner at a time.");
    println!("   - When owner goes out of scope, value is dropped.");
    println!("   - Use .clone() for deep copies of Heap data.\n");

    println!("4. BORROWING (&)");
    println!("   - Access data without taking ownership.");
    println!("   - &T (Immutable): Many allowed.");
    println!("   - &mut T (Mutable): Only one allowed at a time.\n");

    println!("========================================");
    println!("     END OF OWNERSHIP TUTORIAL");
    println!("========================================\n");
}

fn calculate_length(s: &String) -> usize {
    s.len()
}
