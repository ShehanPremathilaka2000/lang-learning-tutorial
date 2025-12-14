use std::mem;

pub fn arrays() {
    println!("\n========================================");
    println!("        RUST ARRAYS TUTORIAL");
    println!("========================================\n");

    // ========================================
    // 1. CREATING ARRAYS
    // ========================================
    println!("--- 1. Creating Arrays ---");
    println!("Code:");
    println!("  let numbers: [i32; 5] = [1, 2, 3, 4, 5]; // Type & size explicit");
    println!("  let float_nums = [1.1, 2.2, 3.3]; // Inferred");
    println!("  let zeros = [0; 5]; // Initialize with same value");
    println!("  println!(\"Zeros: {{:?}}\", zeros);");
    println!("\nOutput:");

    let numbers: [i32; 5] = [1, 2, 3, 4, 5];
    // let float_nums = [1.1, 2.2, 3.3];
    let zeros = [0; 5];
    println!("  Numbers: {:?}", numbers);
    println!("  Zeros: {:?}\n", zeros);

    // ========================================
    // 2. ACCESSING & MODIFYING
    // ========================================
    println!("--- 2. Accessing & Modifying ---");
    println!("Code:");
    println!("  let mut arr = [10, 20, 30];");
    println!("  let first = arr[0];");
    println!("  arr[1] = 99; // Modifying (must be mut)");
    println!("  println!(\"First: {{}}, Modified: {{:?}}\", first, arr);");
    println!("\nOutput:");

    let mut arr = [10, 20, 30];
    let first = arr[0];
    arr[1] = 99;
    println!("  First: {}", first);
    println!("  Modified Array: {:?}\n", arr);

    // ========================================
    // 3. ARRAY PROPERTIES (Len & Size)
    // ========================================
    println!("--- 3. Array Properties ---");
    println!("Code:");
    println!("  println!(\"Length: {{}}\", arr.len());");
    println!("  println!(\"Size in bytes: {{}}\", std::mem::size_of_val(&arr));");
    println!("\nOutput:");

    println!("  Length: {}", arr.len());
    println!("  Size in bytes: {}\n", mem::size_of_val(&arr));

    // ========================================
    // 4. ITERATING TO PRINT
    // ========================================
    println!("--- 4. Iterating (Looping) ---");
    println!("Code:");
    println!("  for num in &arr {{ // Iterate over reference");
    println!("      print!(\"{{}} \", num);");
    println!("  }}");
    println!("\nOutput:");

    print!("  ");
    for num in &arr {
        print!("{} ", num);
    }
    println!("\n");

    // ========================================
    // 5. SLICES (Partial Views)
    // ========================================
    println!("--- 5. Slices ---");
    println!("Code:");
    println!("  let slice: &[i32] = &numbers[1..4]; // Indices 1, 2, 3");
    println!("  println!(\"Slice [1..4]: {{:?}}\", slice);");
    println!("\nOutput:");

    let slice: &[i32] = &numbers[1..4];
    println!("  Slice [1..4]: {:?}\n", slice);

    // ========================================
    // KEY TAKEAWAYS
    // ========================================
    println!("========================================");
    println!("         KEY TAKEAWAYS");
    println!("========================================\n");
    println!("1. FIXED SIZE");
    println!("   - Arrays must have a known size at compile time.");
    println!("   - Cannot grow or shrink (use Vectors for that).\n");

    println!("2. SAME TYPE");
    println!("   - All elements must be of the same type (e.g., all i32).\n");

    println!("3. STACK ALLOCATED");
    println!("   - Arrays are stored on the stack (fast access).\n");

    println!("4. DEBUG PRINTING");
    println!("   - Use {{:?}} to print entire arrays (needs Debug trait).\n");

    println!("========================================");
    println!("       END OF ARRAYS TUTORIAL");
    println!("========================================\n");
}
