pub fn conditions() {
    println!("\n========================================");
    println!("      RUST CONDITIONS TUTORIAL");
    println!("========================================\n");

    // ========================================
    // 1. BASIC IF / ELSE
    // ========================================
    println!("--- 1. Basic If / Else ---");
    println!("Code:");
    println!("  let age = 18;");
    println!("  if age >= 18 {{");
    println!("      println!(\"You are an adult\");");
    println!("  }} else {{");
    println!("      println!(\"You are a minor\");");
    println!("  }}");
    println!("\nOutput:");

    let age = 18;
    if age >= 18 {
        println!("  You are an adult");
    } else {
        println!("  You are a minor");
    }
    println!();

    // ========================================
    // 2. ELSE IF LADDER
    // ========================================
    println!("--- 2. Else If Ladder ---");
    println!("Code:");
    println!("  let number = 6;");
    println!("  if number % 4 == 0 {{");
    println!("      println!(\"Divisible by 4\");");
    println!("  }} else if number % 3 == 0 {{");
    println!("      println!(\"Divisible by 3\");");
    println!("  }} else {{");
    println!("      println!(\"Not divisible by 4 or 3\");");
    println!("  }}");
    println!("\nOutput:");

    let number = 6;
    print!("  "); // Indent output
    if number % 4 == 0 {
        println!("Divisible by 4");
    } else if number % 3 == 0 {
        println!("Divisible by 3");
    } else {
        println!("Not divisible by 4 or 3");
    }
    println!();

    // ========================================
    // 3. IF AS AN EXPRESSION (Shorthand)
    // ========================================
    println!("--- 3. If as an Expression (Shorthand) ---");
    println!("Code:");
    println!("  let condition = true;");
    println!("  // Like ternary operator in other languages");
    println!("  let number = if condition {{ 5 }} else {{ 6 }};");
    println!("  println!(\"The value is: {{}}\", number);");
    println!("\nOutput:");

    let condition = true;
    let number = if condition { 5 } else { 6 };
    println!("  The value is: {}\n", number);

    // ========================================
    // 4. TYPE CONSISTENCY
    // ========================================
    println!("--- 4. Type Consistency ---");
    println!("Code:");
    println!("  // ERROR: Both arms must return same type");
    println!("  // let x = if true {{ 5 }} else {{ \"six\" }};");
    println!("  println!(\"Rust explicitly requires both if and else blocks\");");
    println!("  println!(\"to yield values of the EXACT same type.\");");
    println!("\nOutput:");
    println!("  (Compiler Error if uncommented: `if` and `else` have incompatible types)\n");

    // ========================================
    // 5. MATCH CONTROL FLOW - BASICS
    // ========================================
    println!("--- 5. Match Control Flow (Basics) ---");
    println!("Code:");
    println!("  let grade = 'B';");
    println!("  match grade {{");
    println!("      'A' => println!(\"Excellent!\"),");
    println!("      'B' => println!(\"Good job!\"),");
    println!("      'C' => println!(\"Keep trying!\"),");
    println!("      _ => println!(\"Invalid grade\"), // Catch-all");
    println!("  }}");
    println!("\nOutput:");

    let grade = 'B';
    print!("  ");
    match grade {
        'A' => println!("Excellent!"),
        'B' => println!("Good job!"),
        'C' => println!("Keep trying!"),
        _ => println!("Invalid grade"),
    }
    println!();

    // ========================================
    // 6. ADVANCED MATCHING
    // ========================================
    println!("--- 6. Advanced Matching ---");
    println!("Code:");
    println!("  let number = 13;");
    println!("  // Match as an expression (returns a value)");
    println!("  let result = match number {{");
    println!("      1 => \"One\",");
    println!("      2 | 3 | 5 | 7 | 11 => \"Prime\", // Multiple values");
    println!("      13..=19 => \"Teen\",             // Range (inclusive)");
    println!("      _ => \"Other\",");
    println!("  }};");
    println!("  println!(\"The number {{}} is {{}}\", number, result);");
    println!("\nOutput:");

    let number = 13;
    let result = match number {
        1 => "One",
        2 | 3 | 5 | 7 | 11 => "Prime",
        13..=19 => "Teen",
        _ => "Other",
    };
    println!("  The number {} is {}\n", number, result);

    // ========================================
    // KEY TAKEAWAYS
    // ========================================
    println!("========================================");
    println!("         KEY TAKEAWAYS");
    println!("========================================\n");
    println!("1. EXPRESSIONS");
    println!("   - 'if' is an expression, meaning it returns a value");
    println!("   - Useful for assigning values based on conditions\n");

    println!("2. TYPES");
    println!("   - Branches of flow MUST return the same type");
    println!("   - Cannot return an int in 'if' and string in 'else'\n");

    println!("3. MATCH BASICS");
    println!("   - Use 'match' when you have many conditions");
    println!("   - Must be exhaustive (handle all possible cases using _)");
    println!("   - Safer than 'switch' in other languages\n");

    println!("4. ADVANCED MATCH");
    println!("   - Can return values (expression)");
    println!("   - Multiple patterns: 1 | 2 | 3");
    println!("   - Ranges: 1..=10 (inclusive)\n");

    println!("========================================");
    println!("     END OF CONDITIONS TUTORIAL");
    println!("========================================\n");
}
