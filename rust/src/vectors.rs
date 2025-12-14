pub fn vectors() {
    println!("\n========================================");
    println!("        RUST VECTORS TUTORIAL");
    println!("========================================\n");

    // ========================================
    // 1. CREATING VECTORS
    // ========================================
    println!("--- 1. Creating Vectors ---");
    println!("Code:");
    println!("  let mut v1: Vec<i32> = Vec::new(); // Empty");
    println!("  v1.push(1);");
    println!("  let v2 = vec![10, 20, 30]; // Macro (Common way)");
    println!("  println!(\"v1: {{:?}}, v2: {{:?}}\", v1, v2);");
    println!("\nOutput:");

    let mut v1: Vec<i32> = Vec::new();
    v1.push(1);
    let v2 = vec![10, 20, 30];
    println!("  v1: {:?}, v2: {:?}\n", v1, v2);

    // ========================================
    // 2. ACCESSING ELEMENTS
    // ========================================
    println!("--- 2. Accessing Elements ---");
    println!("Code:");
    println!("  let v = vec![100, 200, 300];");
    println!("  let second = v[1]; // Direct indexing (Panic if out of bounds)");
    println!("  // Safe access using .get()");
    println!("  match v.get(5) {{");
    println!("      Some(val) => println!(\"Found {{}}\", val),");
    println!("      None => println!(\"Index out of bounds\"),");
    println!("  }}");
    println!("\nOutput:");

    let v = vec![100, 200, 300];
    println!("  Index 1: {}", v[1]);
    print!("  Index 5 (via get): ");
    match v.get(5) {
        Some(val) => println!("Found {}", val),
        None => println!("Index out of bounds"),
    }
    println!();

    // ========================================
    // 3. MODIFYING (Add/Remove)
    // ========================================
    println!("--- 3. Modifying (Push, Pop, Insert, Remove) ---");
    println!("Code:");
    println!("  let mut nums = vec![1, 2, 3];");
    println!("  nums.push(4);    // Add to BACK (Fast O(1))");
    println!("  nums.pop();      // Remove from BACK (Fast O(1))");
    println!("  nums.insert(0, 0); // Add to FRONT (Slow O(n) - shifting)");
    println!("  nums.remove(1);    // Remove at index (Slow O(n) - shifting)");
    println!("  println!(\"Modified: {{:?}}\", nums);");
    println!("\nOutput:");

    let mut nums = vec![1, 2, 3];
    nums.push(4);
    nums.pop();
    nums.insert(0, 0); // Becomes [0, 1, 2, 3]
    nums.remove(1); // Removes '1', becomes [0, 2, 3]
    println!("  Modified: {:?}\n", nums);

    // ========================================
    // 4. PERFORMANCE NOTE
    // ========================================
    println!("--- 4. Efficiency Note ---");
    println!("- Adding/Removing at the END is Fast (O(1)).");
    println!("- Adding/Removing at the START/MIDDLE is Slow (O(n)) because");
    println!("  all subsequent elements must be SHIFTED in memory.\n");

    // ========================================
    // 5. LOOPING & OWNERSHIP
    // ========================================
    println!("--- 5. Looping & Ownership ---");
    println!("Code:");
    println!("  let vec = vec![10, 20, 30];");
    println!("  // Iterate by reference (&vec) to keep ownership");
    println!("  for val in &vec {{");
    println!("      print!(\"{{}} \", val);");
    println!("  }}");
    println!("  // vec is still valid here");
    println!("  println!(\"\\nVec len: {{}}\", vec.len());");
    println!("\nOutput:");

    let vec = vec![10, 20, 30];
    print!("  Items: ");
    for val in &vec {
        print!("{} ", val);
    }
    println!("\n  Vec len: {}\n", vec.len());

    // ========================================
    // KEY TAKEAWAYS
    // ========================================
    println!("========================================");
    println!("         KEY TAKEAWAYS");
    println!("========================================\n");
    println!("1. HEAP STORAGE");
    println!("   - Vectors store data on the heap. Can grow/shrink.");
    println!("   - Arrays are stack-based and fixed size.\n");

    println!("2. SAFE ACCESS");
    println!("   - Use .get() to handle out-of-bounds safely (returns Option).");
    println!("   - Use [] if you want to panic on failure.\n");

    println!("3. REFERENCES IN LOOPS");
    println!("   - Use `for x in &v` to read without moving ownership.");
    println!("   - Use `for x in &mut v` to modify elements.");
    println!("   - Use `for x in v` to CONSUME the vector (move).\n");

    println!("4. PERFORMANCE");
    println!("   - Push/Pop at back = Fast.");
    println!("   - Insert/Remove at front = Slow.\n");

    println!("========================================");
    println!("       END OF VECTORS TUTORIAL");
    println!("========================================\n");
}
