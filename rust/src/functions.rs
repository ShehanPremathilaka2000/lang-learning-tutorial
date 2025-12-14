pub fn functions() {
    println!("\n========================================");
    println!("      RUST FUNCTIONS TUTORIAL");
    println!("========================================\n");

    // ========================================
    // 1. BASIC FUNCTION CALL
    // ========================================
    println!("--- 1. Basic Function Call ---");
    println!("Code:");
    println!("  greet();");
    println!("  fn greet() {{");
    println!("      println!(\"  Hello from a function!\");");
    println!("  }}");
    println!("\nOutput:");
    greet();
    println!();

    // ========================================
    // 2. PARAMETERS
    // ========================================
    println!("--- 2. Parameters ---");
    println!("Code:");
    println!("  print_sum(5, 10);");
    println!("  fn print_sum(x: i32, y: i32) {{");
    println!("      println!(\"  The sum of {{}} + {{}} is {{}}\", x, y, x + y);");
    println!("  }}");
    println!("\nOutput:");
    print_sum(5, 10);
    println!();

    // ========================================
    // 3. RETURN VALUES
    // ========================================
    println!("--- 3. Return Values ---");
    println!("Code:");
    println!("  let x = get_five();");
    println!("  let y = add_one(x);");
    println!("  fn get_five() -> i32 {{");
    println!("      5 // Implicit return (no semicolon)");
    println!("  }}");
    println!("  fn add_one(x: i32) -> i32 {{");
    println!("      return x + 1; // Explicit return");
    println!("  }}");
    println!("\nOutput:");

    let x = get_five();
    let y = add_one(x);
    println!("  get_five() returned: {}", x);
    println!("  add_one({}) returned: {}\n", x, y);

    // ========================================
    // 4. SCOPE AND BLOCKS
    // ========================================
    println!("--- 4. Scope and Blocks ---");
    println!("Code:");
    println!("  let outer = 100;");
    println!("  {{");
    println!("      let inner = 200;");
    println!("      println!(\"  Inside block: outer={{}}, inner={{}}\", outer, inner);");
    println!("  }}");
    println!("  // inner is not valid here!");
    println!("  println!(\"  Outside block: outer={{}}\", outer);");
    println!("\nOutput:");

    let outer = 100;
    {
        let inner = 200;
        println!("  Inside block: outer={}, inner={}", outer, inner);
    }
    // println!("Validation: {}", inner); // This would error
    println!("  Outside block: outer={}\n", outer);

    // ========================================
    // KEY TAKEAWAYS
    // ========================================
    println!("========================================");
    println!("         KEY TAKEAWAYS");
    println!("========================================\n");
    println!("1. DECLARATION");
    println!("   - Defined with 'fn keyword'");
    println!("   - Parameters must have types: fn foo(x: i32)");
    println!("   - Snake case convention: my_function_name\n");

    println!("2. RETURN VALUES");
    println!("   - Return type declared with -> Type");
    println!("   - Last expression (no semicolon) is returned implicitly");
    println!("   - 'return' keyword used for early returns\n");

    println!("3. SCOPES");
    println!("   - Variables defined in {{ }} are dropped at }}");
    println!("   - Inner scopes can access outer variables");
    println!("   - Outer scopes CANNOT access inner variables\n");

    println!("========================================");
    println!("     END OF FUNCTIONS TUTORIAL");
    println!("========================================\n");
}

fn greet() {
    println!("  Hello from a function!");
}

fn print_sum(x: i32, y: i32) {
    println!("  The sum of {} + {} is {}", x, y, x + y);
}

fn get_five() -> i32 {
    5
}

fn add_one(x: i32) -> i32 {
    return x + 1;
}
