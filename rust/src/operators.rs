pub fn operators() {
    println!("\n========================================");
    println!("      RUST OPERATORS TUTORIAL");
    println!("========================================\n");

    // ========================================
    // 1. ARITHMETIC OPERATORS
    // ========================================
    println!("--- 1. Arithmetic Operators ---");
    println!("Code:");
    println!("  let a = 10;");
    println!("  let b = 3;");
    println!("  let sum = a + b;");
    println!("  let difference = a - b;");
    println!("  let product = a * b;");
    println!("  let quotient = a / b;       // Integer division");
    println!("  let remainder = a % b;");
    println!("  let float_quotient = 10.0 / 3.0; // Float division");
    println!("\nOutput:");

    let a = 10;
    let b = 3;
    println!("  a = {}, b = {}", a, b);
    println!("  a + b = {}", a + b);
    println!("  a - b = {}", a - b);
    println!("  a * b = {}", a * b);
    println!("  a / b = {} (Integer division)", a / b);
    println!("  a % b = {}", a % b);
    println!("  10.0 / 3.0 = {:.2} (Float division)\n", 10.0 / 3.0);

    // ========================================
    // 2. COMPARISON OPERATORS
    // ========================================
    println!("--- 2. Comparison Operators ---");
    println!("Code:");
    println!("  let c = 5;");
    println!("  let d = 10;");
    println!("  println!(\"{{}} == {{}} is {{}}\", c, d, c == d);");
    println!("  println!(\"{{}} != {{}} is {{}}\", c, d, c != d);");
    println!("  println!(\"{{}} > {{}} is {{}}\", c, d, c > d);");
    println!("  println!(\"{{}} < {{}} is {{}}\", c, d, c < d);");
    println!("\nOutput:");

    let c = 5;
    let d = 10;
    println!("  {} == {} is {}", c, d, c == d);
    println!("  {} != {} is {}", c, d, c != d);
    println!("  {} > {} is {}", c, d, c > d);
    println!("  {} < {} is {}", c, d, c < d);
    println!("  {} >= {} is {}", c, d, c >= d);
    println!("  {} <= {} is {}\n", c, d, c <= d);

    // ========================================
    // 3. LOGICAL OPERATORS
    // ========================================
    println!("--- 3. Logical Operators ---");
    println!("Code:");
    println!("  let t = true;");
    println!("  let f = false;");
    println!("  println!(\"true AND false (&&) is {{}}\", t && f);");
    println!("  println!(\"true OR false (||) is {{}}\", t || f);");
    println!("  println!(\"NOT true (!) is {{}}\", !t);");
    println!("\nOutput:");

    let t = true;
    let f = false;
    println!("  true && false is {}", t && f);
    println!("  true || false is {}", t || f);
    println!("  !true is {}\n", !t);

    // ========================================
    // 4. ASSIGNMENT OPERATORS
    // ========================================
    println!("--- 4. Assignment Operators ---");
    println!("Code:");
    println!("  let mut score = 10;");
    println!("  score += 5; // score = score + 5");
    println!("  score -= 2; // score = score - 2");
    println!("  score *= 2; // score = score * 2");
    println!("  score /= 3; // score = score / 3");
    println!("\nOutput:");

    let mut score = 10;
    println!("  Initial score: {}", score);
    score += 5;
    println!("  After += 5: {}", score);
    score -= 2;
    println!("  After -= 2: {}", score);
    score *= 2;
    println!("  After *= 2: {}", score);
    score /= 2;
    println!("  After /= 2: {}\n", score);

    // ========================================
    // KEY TAKEAWAYS
    // ========================================
    println!("========================================");
    println!("         KEY TAKEAWAYS");
    println!("========================================\n");
    println!("1. ARITHMETIC");
    println!("   - Integer division truncates (5 / 2 = 2)");
    println!("   - Use floating point numbers for precise division");
    println!("   - No ++ or -- operators in Rust (use += 1)\n");

    println!("2. COMPARISON");
    println!("   - Always returns a boolean (true/false)");
    println!("   - Can compare numbers, chars, etc.\n");

    println!("3. LOGICAL");
    println!("   - && (AND): true only if both are true");
    println!("   - || (OR): true if at least one is true");
    println!("   - ! (NOT): inverts the boolean value\n");

    println!("========================================");
    println!("     END OF OPERATORS TUTORIAL");
    println!("========================================\n");
}
