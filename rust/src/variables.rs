pub fn variables() {
    println!("\n========================================");
    println!("       RUST VARIABLES TUTORIAL");
    println!("========================================\n");

    // ========================================
    // 1. IMMUTABLE VARIABLES (Default)
    // ========================================
    println!("--- 1. Immutable Variables ---");
    println!("Code:");
    println!("  let x = 5;");
    println!("  println!(\"x = {{}}\", x);");
    println!("\nOutput:");
    let x = 5;
    println!("  x = {}", x);
    println!("\nNote: Variables are immutable by default in Rust");
    println!("      Trying to reassign would cause: x = 6; // ERROR!\n");

    // ========================================
    // 2. MUTABLE VARIABLES
    // ========================================
    println!("--- 2. Mutable Variables ---");
    println!("Code:");
    println!("  let mut y = 10;");
    println!("  println!(\"y = {{}}\", y);");
    println!("  y = 20;  // Allowed because of 'mut'");
    println!("  println!(\"y = {{}}\", y);");
    println!("  y += 5;");
    println!("  println!(\"y = {{}}\", y);");
    println!("\nOutput:");
    let mut y = 10;
    println!("  y = {}", y);
    y = 20;
    println!("  y = {}", y);
    y += 5;
    println!("  y = {}\n", y);

    // ========================================
    // 3. VARIABLE SHADOWING
    // ========================================
    println!("--- 3. Variable Shadowing ---");
    println!("Code:");
    println!("  let z = 5;");
    println!("  let z = z + 1;  // Shadowing");
    println!("  {{");
    println!("      let z = z * 2;  // Inner scope");
    println!("      println!(\"Inner z = {{}}\", z);");
    println!("  }}");
    println!("  println!(\"Outer z = {{}}\", z);");
    println!("\nOutput:");
    let z = 5;
    let z = z + 1;
    {
        let z = z * 2;
        println!("  Inner z = {}", z);
    }
    println!("  Outer z = {}", z);

    println!("\nCode (Type change with shadowing):");
    println!("  let spaces = \"   \";");
    println!("  let spaces = spaces.len();");
    println!("\nOutput:");
    let spaces = "   ";
    let spaces = spaces.len();
    println!("  Length = {}\n", spaces);

    // ========================================
    // 4. CONSTANTS
    // ========================================
    println!("--- 4. Constants ---");
    println!("Code:");
    println!("  const MAX_POINTS: u32 = 100_000;");
    println!("  const PI: f64 = 3.14159;");
    println!("\nOutput:");
    const MAX_POINTS: u32 = 100_000;
    const PI: f64 = 3.14159;
    println!("  MAX_POINTS = {}", MAX_POINTS);
    println!("  PI = {}", PI);
    println!("\nNote: Constants are always immutable and require type annotations\n");

    // ========================================
    // KEY TAKEAWAYS
    // ========================================
    println!("========================================");
    println!("         KEY TAKEAWAYS");
    println!("========================================\n");

    println!("1. IMMUTABILITY BY DEFAULT");
    println!("   - Variables are immutable by default in Rust");
    println!("   - This prevents accidental modifications and bugs");
    println!("   - Use 'let x = 5;' for immutable variables\n");

    println!("2. MUTABILITY WITH 'mut'");
    println!("   - Add 'mut' keyword to make variables mutable");
    println!("   - Syntax: 'let mut x = 5;'");
    println!("   - Explicitly shows intent to modify the variable\n");

    println!("3. SHADOWING");
    println!("   - Can redeclare a variable with the same name");
    println!("   - Creates a new variable, doesn't modify the original");
    println!("   - Allows changing the type of the value");
    println!("   - Scoped to blocks (inner scope doesn't affect outer)\n");

    println!("4. CONSTANTS");
    println!("   - Declared with 'const' keyword");
    println!("   - Always immutable (can't use 'mut')");
    println!("   - Must have type annotations");
    println!("   - Naming convention: SCREAMING_SNAKE_CASE");
    println!("   - Can be declared in any scope, including global\n");

    println!("5. DIFFERENCES: const vs let");
    println!("   - const: Always immutable, type required, compile-time constant");
    println!("   - let: Immutable by default, type optional, runtime value");
    println!("   - let mut: Mutable, can be changed after declaration\n");

    println!("========================================");
    println!("     END OF VARIABLES TUTORIAL");
    println!("========================================\n");
}
