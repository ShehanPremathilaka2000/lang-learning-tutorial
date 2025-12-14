use std::collections::HashMap;

pub fn hash_map() {
    println!("\n========================================");
    println!("       RUST HASHMAPS TUTORIAL");
    println!("========================================\n");

    // ========================================
    // 1. CREATING HASHMAPS
    // ========================================
    println!("--- 1. Creating HashMaps ---");
    println!("Code:");
    println!("  use std::collections::HashMap; // Must import!");
    println!("  let mut scores = HashMap::new();");
    println!("  scores.insert(\"Blue\", 10);");
    println!("  scores.insert(\"Red\", 50);");
    println!("  println!(\"Scores: {{:?}}\", scores);");
    println!("\nOutput:");

    let mut scores = HashMap::new();
    scores.insert("Blue", 10);
    scores.insert("Red", 50);
    println!("  Scores: {:?}\n", scores);

    // ========================================
    // 2. ACCESSING VALUES
    // ========================================
    println!("--- 2. Accessing Values ---");
    println!("Code:");
    println!("  let team_name = \"Blue\";");
    println!("  // .get() returns Option<&V>");
    println!("  match scores.get(team_name) {{");
    println!("      Some(score) => println!(\"Score for {{}}: {{}}\", team_name, score),");
    println!("      None => println!(\"Team not found\"),");
    println!("  }}");
    println!("\nOutput:");

    let team_name = "Blue";
    match scores.get(team_name) {
        Some(score) => println!("  Score for {}: {}", team_name, score),
        None => println!("  Team not found"),
    }
    println!();

    // ========================================
    // 3. UPDATING VALUES
    // ========================================
    println!("--- 3. Updating Values ---");
    println!("Code:");
    println!("  scores.insert(\"Blue\", 25); // Overwrites 10");
    println!("  // Only insert if key has no value");
    println!("  scores.entry(\"Yellow\").or_insert(50);");
    println!("  scores.entry(\"Blue\").or_insert(50); // Won't change (25 exists)");
    println!("  println!(\"Updated: {{:?}}\", scores);");
    println!("\nOutput:");

    scores.insert("Blue", 25);
    scores.entry("Yellow").or_insert(50);
    scores.entry("Blue").or_insert(50);
    println!("  Updated: {:?}\n", scores);

    // ========================================
    // 4. REMOVING VALUES
    // ========================================
    println!("--- 4. Removing Values ---");
    println!("Code:");
    println!("  scores.remove(\"Red\");");
    println!("  println!(\"After Removing Red: {{:?}}\", scores);");
    println!("\nOutput:");

    scores.remove("Red");
    println!("  After Removing Red: {:?}\n", scores);

    // ========================================
    // 5. LOOPING THROUGH HASHMAP
    // ========================================
    println!("--- 5. Looping ---");
    println!("Code:");
    println!("  for (key, value) in &scores {{");
    println!("      println!(\"{{}}: {{}}\", key, value);");
    println!("  }}");
    println!("\nOutput:");

    print!("  Pairs: \n");
    for (key, value) in &scores {
        println!("    {}: {}", key, value);
    }
    println!();

    // ========================================
    // KEY TAKEAWAYS
    // ========================================
    println!("========================================");
    println!("         KEY TAKEAWAYS");
    println!("========================================\n");
    println!("1. IMPORT REQUIRED");
    println!("   - Must use 'use std::collections::HashMap;'. not in prelude.\n");

    println!("2. KEYS AND VALUES");
    println!("   - Stores Key-Value pairs. Keys must be unique.");
    println!("   - Access via .get(&key) which returns Option.\n");

    println!("3. UPDATING");
    println!("   - .insert() overwrites existing values.");
    println!("   - .entry().or_insert() only inserts if missing.\n");

    println!("4. OWNERSHIP");
    println!("   - For owned types like String, HashMap takes ownership of keys/values.");
    println!("   - For Copy types like i32, values are copied.\n");

    println!("========================================");
    println!("       END OF HASHMAPS TUTORIAL");
    println!("========================================\n");
}
