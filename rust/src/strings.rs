pub fn strings() {
    println!("\n========================================");
    println!("      RUST STRINGS TUTORIAL");
    println!("========================================\n");

    // ========================================
    // 1. &str vs String
    // ========================================
    println!("--- 1. &str vs String ---");
    println!("Code:");
    println!("  let s_slice: &str = \"Hello\"; // String Literal (immutable, fixed size)");
    println!("  let s_obj: String = String::from(\"World\"); // Heap allocated, growable");
    println!("  println!(\"Slice: {{}}, Object: {{}}\", s_slice, s_obj);");
    println!("\nOutput:");

    let s_slice: &str = "Hello";
    let s_obj: String = String::from("World");
    println!("  Slice: {}, Object: {}\n", s_slice, s_obj);

    // ========================================
    // 2. CREATING STRINGS
    // ========================================
    println!("--- 2. Creating Strings ---");
    println!("Code:");
    println!("  let s1 = String::from(\"Using from()\");");
    println!("  let s2 = \"Using to_string()\".to_string();");
    println!("  let empty = String::new();");
    println!("\nOutput:");

    let s1 = String::from("Using from()");
    let s2 = "Using to_string()".to_string();
    let empty = String::new();
    println!("  s1: {}", s1);
    println!("  s2: {}", s2);
    println!("  empty: '{}' (len: {})\n", empty, empty.len());

    // ========================================
    // 3. MUTABILITY & MODIFICATION
    // ========================================
    println!("--- 3. Mutability & Modification ---");
    println!("Code:");
    println!("  let mut s = String::from(\"Hello\");");
    println!("  s.push_str(\", World\"); // Append string");
    println!("  s.push('!');            // Append char");
    println!("  println!(\"Modified: {{}}\", s);");
    println!("\nOutput:");

    let mut s = String::from("Hello");
    s.push_str(", World");
    s.push('!');
    println!("  Modified: {}\n", s);

    // ========================================
    // 4. CONCATENATION
    // ========================================
    println!("--- 4. Concatenation ---");
    println!("Code:");
    println!("  let s1 = String::from(\"Hello\");");
    println!("  let s2 = String::from(\"Rust\");");
    println!("  // Note: s1 is moved and cannot be used after this!");
    println!("  let s3 = s1 + \", \" + &s2; // using + operator");
    println!("  let s4 = format!(\"{{}} is cool\", s2); // using format! (no move)");
    println!("\nOutput:");

    let s1 = String::from("Hello");
    let s2 = String::from("Rust");
    let s3 = s1 + ", " + &s2; 
    let s4 = format!("{} is cool", s2);
    println!("  s3: {}", s3);
    println!("  s4: {}\n", s4);

    // ========================================
    // 5. OTHER STRING CONCEPTS
    // ========================================
    println!("--- 5. Length & Capacity ---");
    println!("Code:");
    println!("  let text = \"Rust\";");
    println!("  println!(\"Length: {{}}\", text.len());");
    println!("  println!(\"Is empty: {{}}\", text.is_empty());");
    println!("  println!(\"Contains 'us': {{}}\", text.contains(\"us\"));");
    println!("\nOutput:");

    let text = "Rust";
    println!("  Length: {}", text.len());
    println!("  Is empty: {}", text.is_empty());
    println!("  Contains 'us': {}\n", text.contains("us"));

    // ========================================
    // KEY TAKEAWAYS
    // ========================================
    println!("========================================");
    println!("         KEY TAKEAWAYS");
    println!("========================================\n");
    println!("1. TYPES");
    println!("   - &str: String slice, often a reference to string logic/literals. Fast, fixed.");
    println!("   - String: Owned, heap-allocated, growable buffer.");
    
    println!("2. MODIFICATION");
    println!("   - Must be 'mut' to modify.");
    println!("   - .push_str() appends a string, .push() appends a char.");
    
    println!("3. CONCATENATION");
    println!("   - '+': Takes ownership of the first string (s1 + &s2).");
    println!("   - 'format!': Doesn't take ownership, works like println! but returns a String.");

    println!("========================================");
    println!("     END OF STRINGS TUTORIAL");
    println!("========================================\n");
}