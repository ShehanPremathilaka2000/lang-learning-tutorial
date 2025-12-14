pub fn loops() {
    println!("\n========================================");
    println!("        RUST LOOPS TUTORIAL");
    println!("========================================\n");

    // ========================================
    // 1. INFINITE LOOP & RETURNING VALUES
    // ========================================
    println!("--- 1. Infinite Loop & Returning Values ---");
    println!("Code:");
    println!("  let mut counter = 0;");
    println!("  // 'loop' creates an infinite loop");
    println!("  let result = loop {{");
    println!("      counter += 1;");
    println!("      if counter == 10 {{");
    println!("          break counter * 2; // Return value!");
    println!("      }}");
    println!("  }};");
    println!("  println!(\"The result is {{}}\", result);");
    println!("\nOutput:");

    let mut counter = 0;
    let result = loop {
        counter += 1;
        if counter == 10 {
            break counter * 2;
        }
    };
    println!("  The result is {}\n", result);

    // ========================================
    // 2. WHILE LOOP
    // ========================================
    println!("--- 2. While Loop ---");
    println!("Code:");
    println!("  let mut number = 3;");
    println!("  while number != 0 {{");
    println!("      print!(\"{{}}! \", number);");
    println!("      number -= 1;");
    println!("  }}");
    println!("  println!(\"LIFTOFF!\");");
    println!("\nOutput:");

    let mut number = 3;
    print!("  ");
    while number != 0 {
        print!("{}! ", number);
        number -= 1;
    }
    println!("LIFTOFF!\n");

    // ========================================
    // 3. FOR LOOP & RANGES
    // ========================================
    println!("--- 3. For Loop & Ranges ---");
    println!("Code:");
    println!("  // 1..4 is exclusive (1, 2, 3)");
    println!("  for i in 1..4 {{");
    println!("      print!(\"Exclusive: {{}} \", i);");
    println!("  }}");
    println!("  // 1..=4 is inclusive (1, 2, 3, 4)");
    println!("  for i in 1..=4 {{");
    println!("      print!(\"Inclusive: {{}} \", i);");
    println!("  }}");
    println!("\nOutput:");

    print!("  ");
    for i in 1..4 {
        print!("Exclusive: {} ", i);
    }
    println!();
    print!("  ");
    for i in 1..=4 {
        print!("Inclusive: {} ", i);
    }
    println!("\n");

    // ========================================
    // 4. FOR LOOP WITH COLLECTIONS
    // ========================================
    println!("--- 4. For Loop with Collections ---");
    println!("Code:");
    println!("  let array = [10, 20, 30];");
    println!("  for element in array {{");
    println!("      print!(\"Value: {{}} \", element);");
    println!("  }}");
    println!("\nOutput:");

    let array = [10, 20, 30];
    print!("  ");
    for element in array {
        print!("Value: {} ", element);
    }
    println!("\n");

    // ========================================
    // 5. LOOP LABELS (Nested Loops)
    // ========================================
    println!("--- 5. Loop Labels ---");
    println!("Code:");
    println!("  'outer: loop {{");
    println!("      println!(\"Entered outer loop\");");
    println!("      'inner: loop {{");
    println!("          println!(\"Entered inner loop, breaking outer\");");
    println!("          break 'outer; // Breaks the OUTER loop");
    println!("      }}");
    println!("  }}");
    println!("\nOutput:");

    print!("  ");
    'outer: loop {
        println!("Entered outer loop");
        print!("  ");
        #[allow(unused_labels)]
        'inner: loop {
            println!("Entered inner loop, breaking outer");
            break 'outer;
        }
    }
    println!();

    // ========================================
    // KEY TAKEAWAYS
    // ========================================
    println!("========================================");
    println!("         KEY TAKEAWAYS");
    println!("========================================\n");
    println!("1. LOOP TYPES");
    println!("   - 'loop': Infinite, can return values using 'break value'");
    println!("   - 'while': Runs while condition is true");
    println!("   - 'for': Best for ranges and collections\n");

    println!("2. RANGES");
    println!("   - start..end  : Exclusive (end is NOT included)");
    println!("   - start..=end : Inclusive (end IS included)\n");

    println!("3. CONTROL");
    println!("   - 'break': Exit loop");
    println!("   - 'continue': Skip rest of iteration");
    println!("   - Loop Labels ('label): Use to break explicit nested loops\n");

    println!("========================================");
    println!("         END OF LOOPS TUTORIAL");
    println!("========================================\n");
}
