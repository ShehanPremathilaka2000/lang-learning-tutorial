#[allow(dead_code)]
struct StructUser {
    username: String,
    email: String,
    sign_in_count: u64,
    active: bool,
}

impl StructUser {
    fn summarize(&self) {
        println!("{} - {}", self.username, self.email);
    }
}

pub fn structs() {
    println!("\n========================================");
    println!("        RUST STRUCTS TUTORIAL");
    println!("========================================\n");

    // ========================================
    // 1. DEFINING AND INSTANTIATING
    // ========================================
    println!("--- 1. Defining and Instantiating ---");
    println!("Code:");
    println!(
        "  struct StructUser {{ username: String, email: String, sign_in_count: u64, active: bool }}"
    );
    println!("  let user1 = StructUser {{");
    println!("      email: String::from(\"someone@example.com\"),");
    println!("      username: String::from(\"someusername123\"),");
    println!("      active: true,");
    println!("      sign_in_count: 1,");
    println!("  }};");
    println!(
        "  println!(\"User: {{}} ({{}}), Count: {{}}\", user1.username, user1.email, user1.sign_in_count);"
    );
    println!("\nOutput:");

    let user1 = StructUser {
        email: String::from("someone@example.com"),
        username: String::from("someusername123"),
        active: true,
        sign_in_count: 1,
    };
    println!(
        "  User: {} ({}), Count: {}\n",
        user1.username, user1.email, user1.sign_in_count
    );

    // ========================================
    // 2. MUTABILITY
    // ========================================
    println!("--- 2. Mutability ---");
    println!("Code:");
    println!("  let mut user2 = StructUser {{ ... }}; // Must be mutable to change fields");
    println!("  user2.email = String::from(\"newemail@example.com\");");
    println!("  println!(\"Updated Email: {{}}\", user2.email);");
    println!("\nOutput:");

    let mut user2 = StructUser {
        email: String::from("original@example.com"),
        username: String::from("user2"),
        active: true,
        sign_in_count: 1,
    };
    user2.email = String::from("newemail@example.com");
    println!("  Updated Email: {}\n", user2.email);

    // ========================================
    // 3. TUPLE STRUCTS
    // ========================================
    println!("--- 3. Tuple Structs ---");
    println!("Code:");
    println!("  struct Color(i32, i32, i32);");
    println!("  let black = Color(0, 0, 0);");
    println!("  println!(\"Black: ({{}}, {{}}, {{}})\", black.0, black.1, black.2);");
    println!("\nOutput:");

    struct Color(i32, i32, i32);
    let black = Color(0, 0, 0);
    println!("  Black: ({}, {}, {})\n", black.0, black.1, black.2);

    // ========================================
    // 4. METHODS (impl block)
    // ========================================
    println!("--- 4. Methods ---");
    println!("Code:");
    println!("  impl StructUser {{");
    println!("      fn summarize(&self) {{");
    println!("          println!(\"{{}} - {{}}\", self.username, self.email);");
    println!("      }}");
    println!("  }}");
    println!("  user1.summarize();");
    println!("\nOutput:");

    println!("  Calling user1.summarize():");
    print!("  ");
    user1.summarize();
    println!();

    // ========================================
    // KEY TAKEAWAYS
    // ========================================
    println!("========================================");
    println!("         KEY TAKEAWAYS");
    println!("========================================\n");
    println!("1. CUSTOM TYPES");
    println!("   - Structs let you name and package related data together.\n");

    println!("2. MUTABILITY");
    println!("   - Entire instance must be mutable to change any field.");
    println!("   - Rust doesn't allow only some fields to be mutable.\n");

    println!("3. TUPLE STRUCTS");
    println!("   - Named tuples: struct Color(i32, i32, i32).");
    println!("   - Access via index: .0, .1.\n");

    println!("4. METHODS");
    println!("   - Defined in 'impl StructName {{ }}' block.");
    println!("   - Use '&self' to read, '&mut self' to modify.\n");

    println!("========================================");
    println!("       END OF STRUCTS TUTORIAL");
    println!("========================================\n");
}
