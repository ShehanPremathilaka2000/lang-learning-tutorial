// Define Enum
#[derive(Debug)] // Enable debug printing
#[allow(dead_code)] // Suppress unused variant warnings
enum Messages {
    Quit,
    Move { x: i32, y: i32 }, // Anonymous struct variant
    Write(String),           // Tuple variant
    ChangeColor(i32, i32, i32),
}

impl Messages {
    fn call(&self) {
        println!("  Msg Called: {:?}", self);
    }
}

pub fn enums() {
    println!("\n========================================");
    println!("        RUST ENUMS TUTORIAL");
    println!("========================================\n");

    // ========================================
    // 1. DEFINING ENUMS
    // ========================================
    println!("--- 1. Defining and Using Enums ---");
    println!("- Enums allow you to define a type by enumerating its possible variants.");
    println!("Code:");
    println!("  enum Messages {{");
    println!("      Quit,");
    println!("      Move {{ x: i32, y: i32 }},");
    println!("      Write(String),");
    println!("      ChangeColor(i32, i32, i32),");
    println!("  }}");
    println!("  let msg = Messages::Write(String::from(\"hello\"));");
    println!("  println!(\"{{:?}}\", msg);");
    println!("\nOutput:");

    let msg = Messages::Write(String::from("hello"));
    println!("  {:?}\n", msg);

    // ========================================
    // 2. METHODS ON ENUMS
    // ========================================
    println!("--- 2. Methods on Enums ---");
    println!("Code:");
    println!("  impl Messages {{ fn call(&self) {{ ... }} }}");
    println!("  msg.call();");
    println!("\nOutput:");

    msg.call();
    println!();

    // ========================================
    // 3. MATCH CONTROL FLOW
    // ========================================
    println!("--- 3. Match Control Flow ---");
    println!("Code:");
    println!("  match msg {{");
    println!("      Messages::Quit => println!(\"Quitting\"),");
    println!("      Messages::Write(text) => println!(\"Writing: {{}}\", text),");
    println!("      _ => println!(\"Other action\"),");
    println!("  }}");
    println!("\nOutput:");

    match &msg {
        Messages::Quit => println!("  Quitting"),
        Messages::Write(text) => println!("  Writing: {}", text),
        _ => println!("  Other action"),
    }
    println!();

    // ========================================
    // 4. OPTION ENUM
    // ========================================
    println!("--- 4. The Option Enum ---");
    println!("- Rust doesn't have NULL. It has Option<T>.");
    println!("- enum Option<T> {{ Some(T), None }}");
    println!("Code:");
    println!("  let some_number = Some(5);");
    println!("  let absent_number: Option<i32> = None;");
    println!("  match some_number {{");
    println!("      Some(n) => println!(\"Found {{}}\", n),");
    println!("      None => println!(\"Missing\"),");
    println!("  }}");
    println!("\nOutput:");

    let some_number = Some(5);
    let absent_number: Option<i32> = None;

    match some_number {
        Some(n) => println!("  Found {}", n),
        None => println!("  Missing"),
    }

    match absent_number {
        Some(n) => println!("  Found {}", n),
        None => println!("  Missing (None case)"),
    }
    println!();

    // ========================================
    // KEY TAKEAWAYS
    // ========================================
    println!("========================================");
    println!("         KEY TAKEAWAYS");
    println!("========================================\n");
    println!("1. VARIANTS WITH DATA");
    println!("   - Enums can store data (Tuples, Structs, or Unit).");
    println!("   - Each variant can type independently.\n");

    println!("2. PATTERN MATCHING");
    println!("   - 'match' is powerful for destructuring Enums.");
    println!("   - Must be exhaustive (cover all cases).\n");

    println!("3. OPTION<T>");
    println!("   - Replaces Null.");
    println!("   - Forces you to handle the None case explicitly.\n");

    println!("========================================");
    println!("       END OF ENUMS TUTORIAL");
    println!("========================================\n");
}
