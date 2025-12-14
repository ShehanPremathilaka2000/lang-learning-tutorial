pub fn borrowing() {
    println!("\n========================================");
    println!("      RUST BORROWING TUTORIAL");
    println!("========================================\n");

    // ========================================
    // 1. IMMUTABLE REFERENCES (Borrowing)
    // ========================================
    println!("--- 1. Immutable References (Shared Borrow) ---");
    println!("Code:");
    println!("  let s1 = String::from(\"Hello\");");
    println!("  let len = calculate_length(&s1); // &s1 creates a reference");
    println!("  println!(\"The length of '{{}}' is {{}}.\", s1, len);");
    println!("\nOutput:");

    let s1 = String::from("Hello");
    let len = calculate_length(&s1);
    println!("  The length of '{}' is {}.\n", s1, len);

    // ========================================
    // 2. MUTABLE REFERENCES (Exclusive Borrow)
    // ========================================
    println!("--- 2. Mutable References (Exclusive Borrow) ---");
    println!("Code:");
    println!("  let mut s = String::from(\"Hello\");");
    println!("  change(&mut s); // Pass mutable reference");
    println!("  println!(\"Modified string: {{}}\", s);");
    println!("\nOutput:");

    let mut s = String::from("Hello");
    change(&mut s);
    println!("  Modified string: {}\n", s);

    // ========================================
    // 3. THE RULES OF REFERENCES
    // ========================================
    println!("--- 3. The Rules of References ---");
    println!("Code:");
    println!("  let mut s = String::from(\"Hello\");");
    println!("  let r1 = &s; // OK");
    println!("  let r2 = &s; // OK: Multiple immutable refs allowed");
    println!("  // let r3 = &mut s; // ERROR: Cannot borrow mutably while immutable borrowed");
    println!("  println!(\"r1: {{}}, r2: {{}}\", r1, r2);");
    println!("  // r1 and r2 scope ends here (non-lexical lifetimes), so r3 is now safe");
    println!("  let r3 = &mut s;");
    println!("  r3.push_str(\" World\");");
    println!("  println!(\"r3: {{}}\", r3);");
    println!("\nOutput:");

    let mut s = String::from("Hello");
    let r1 = &s;
    let r2 = &s;
    println!("  r1: {}, r2: {}", r1, r2);
    // println!("{}, {}", r1, r2); // If used here, r1/r2 would extend, blocking r3

    let r3 = &mut s;
    r3.push_str(" World");
    println!("  r3: {}\n", r3);

    // ========================================
    // 4. CLONE VS BORROW
    // ========================================
    println!("--- 4. Clone vs Borrow ---");
    println!("- CLONE: Expensive deep copy. Both variables own independent data.");
    println!("- BORROW: Cheap pointer copy. Original variable retains ownership.");
    println!("Code:");
    println!("  let original = String::from(\"Data\");");
    println!("  let borrowed = &original;      // Fast, refers to original");
    println!("  let cloned = original.clone(); // Slow, duplicates data");
    println!(
        "  println!(\"Addr: {{:p}} (orig), {{:p}} (borrowed referent), {{:p}} (cloned)\", &original, borrowed, &cloned);"
    );
    println!("\nOutput:");

    let original = String::from("Data");
    let borrowed = &original;
    let cloned = original.clone();
    println!("  Original: '{}'", original);
    println!("  Borrowed: '{}' (points to Original)", borrowed);
    println!("  Cloned:   '{}' (new separate data)\n", cloned);

    // ========================================
    // KEY TAKEAWAYS
    // ========================================
    println!("========================================");
    println!("         KEY TAKEAWAYS");
    println!("========================================\n");
    println!("1. BORROWING (&T)");
    println!("   - Allows reading data without taking ownership.");
    println!("   - Default is immutable.");
    println!("   - Can have INFINITE immutable references at once.\n");

    println!("2. MUTABLE REFERENCES (&mut T)");
    println!("   - Allows modifying data without taking ownership.");
    println!("   - Can only have ONE mutable reference at a time.");
    println!("   - Prevents 'data races' at compile time.\n");

    println!("3. THE DATA RACE RULE");
    println!("   - Cannot have mutable reference while immutable ones exist.");
    println!("   - (Unless the immutable ones are no longer used).\n");

    println!("4. DANGLING REFERENCES");
    println!("   - Rust prevents pointing to invalid memory.");
    println!("   - Compiler ensures references never outlive the data they point to.\n");

    println!("========================================");
    println!("     END OF BORROWING TUTORIAL");
    println!("========================================\n");
}

fn calculate_length(s: &String) -> usize {
    s.len()
}

fn change(some_string: &mut String) {
    some_string.push_str(", world");
}
