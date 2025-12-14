pub fn data_types() {
    println!("\n========================================");
    println!("      RUST DATA TYPES TUTORIAL");
    println!("========================================\n");

    // ========================================
    // 1. SCALAR TYPES - INTEGERS
    // ========================================
    println!("--- 1. Scalar Types: Integers ---");
    println!("Code:");
    println!("  // Signed integers (can be negative)");
    println!("  let i8_val: i8 = -128;");
    println!("  let i16_val: i16 = -32_768;");
    println!("  let i32_val: i32 = -2_147_483_648;");
    println!("  let i64_val: i64 = -9_223_372_036_854_775_808;");
    println!("  let i128_val: i128 = 100;");
    println!("  let isize_val: isize = -100;  // Pointer-sized");
    println!("\n  // Unsigned integers (only positive)");
    println!("  let u8_val: u8 = 255;");
    println!("  let u16_val: u16 = 65_535;");
    println!("  let u32_val: u32 = 4_294_967_295;");
    println!("  let u64_val: u64 = 100;");
    println!("\nOutput:");

    let i8_val: i8 = -128;
    let i16_val: i16 = -32_768;
    let i32_val: i32 = -2_147_483_648;
    let i64_val: i64 = -9_223_372_036_854_775_808;
    let i128_val: i128 = 100;
    let isize_val: isize = -100;

    let u8_val: u8 = 255;
    let u16_val: u16 = 65_535;
    let u32_val: u32 = 4_294_967_295;
    let u64_val: u64 = 100;

    println!(
        "  Signed: i8={}, i16={}, i32={}, i64={}",
        i8_val, i16_val, i32_val, i64_val
    );
    println!("  i128={}, isize={}", i128_val, isize_val);
    println!(
        "  Unsigned: u8={}, u16={}, u32={}, u64={}\n",
        u8_val, u16_val, u32_val, u64_val
    );

    // ========================================
    // 2. SCALAR TYPES - FLOATING POINT
    // ========================================
    println!("--- 2. Scalar Types: Floating Point ---");
    println!("Code:");
    println!("  let f32_val: f32 = 3.14159;");
    println!("  let f64_val: f64 = 2.718281828459045;");
    println!("  let default_float = 3.0;  // Defaults to f64");
    println!("\nOutput:");

    let f32_val: f32 = 3.14159;
    let f64_val: f64 = 2.718281828459045;
    let default_float = 3.0;

    println!("  f32 = {}", f32_val);
    println!("  f64 = {}", f64_val);
    println!("  default_float = {} (type: f64)\n", default_float);

    // ========================================
    // 3. SCALAR TYPES - BOOLEAN
    // ========================================
    println!("--- 3. Scalar Types: Boolean ---");
    println!("Code:");
    println!("  let is_active: bool = true;");
    println!("  let is_greater: bool = 10 > 5;");
    println!("  let result = !is_active;  // NOT operator");
    println!("\nOutput:");

    let is_active: bool = true;
    let is_greater: bool = 10 > 5;
    let result = !is_active;

    println!("  is_active = {}", is_active);
    println!("  is_greater = {}", is_greater);
    println!("  !is_active = {}\n", result);

    // ========================================
    // 4. SCALAR TYPES - CHARACTER
    // ========================================
    println!("--- 4. Scalar Types: Character ---");
    println!("Code:");
    println!("  let letter: char = 'A';");
    println!("  let digit: char = '5';");
    println!("  let emoji: char = '🦀';  // Rust mascot!");
    println!("  let heart: char = '❤';");
    println!("\nOutput:");

    let letter: char = 'A';
    let digit: char = '5';
    let emoji: char = '🦀';
    let heart: char = '❤';

    println!("  letter = {}", letter);
    println!("  digit = {}", digit);
    println!("  emoji = {}", emoji);
    println!("  heart = {}\n", heart);

    // ========================================
    // 5. COMPOUND TYPES - TUPLES
    // ========================================
    println!("--- 5. Compound Types: Tuples ---");
    println!("Code:");
    println!("  let person: (&str, i32, f64) = (\"Alice\", 30, 5.6);");
    println!("  let (name, age, height) = person;  // Destructuring");
    println!("  let first = person.0;  // Index access");
    println!("\nOutput:");

    let person: (&str, i32, f64) = ("Alice", 30, 5.6);
    let (name, age, height) = person;
    let first = person.0;

    println!("  person = {:?}", person);
    println!(
        "  Destructured: name={}, age={}, height={}",
        name, age, height
    );
    println!("  person.0 = {}\n", first);

    // ========================================
    // 6. COMPOUND TYPES - ARRAYS
    // ========================================
    println!("--- 6. Compound Types: Arrays ---");
    println!("Code:");
    println!("  let numbers: [i32; 5] = [1, 2, 3, 4, 5];");
    println!("  let zeros = [0; 3];  // [0, 0, 0]");
    println!("  let first = numbers[0];");
    println!("  let length = numbers.len();");
    println!("\nOutput:");

    let numbers: [i32; 5] = [1, 2, 3, 4, 5];
    let zeros = [0; 3];
    let first_num = numbers[0];
    let length = numbers.len();

    println!("  numbers = {:?}", numbers);
    println!("  zeros = {:?}", zeros);
    println!("  first element = {}", first_num);
    println!("  length = {}\n", length);

    // ========================================
    // 7. TYPE INFERENCE
    // ========================================
    println!("--- 7. Type Inference ---");
    println!("Code:");
    println!("  let inferred_int = 42;  // i32 by default");
    println!("  let inferred_float = 3.14;  // f64 by default");
    println!("  let inferred_bool = true;");
    println!("  let inferred_char = 'R';");
    println!("\nOutput:");

    let inferred_int = 42;
    let inferred_float = 3.14;
    let inferred_bool = true;
    let inferred_char = 'R';

    println!("  inferred_int = {} (type: i32)", inferred_int);
    println!("  inferred_float = {} (type: f64)", inferred_float);
    println!("  inferred_bool = {} (type: bool)", inferred_bool);
    println!("  inferred_char = {} (type: char)\n", inferred_char);

    // ========================================
    // 8. EXPLICIT TYPE ANNOTATIONS
    // ========================================
    println!("--- 8. Explicit Type Annotations ---");
    println!("Code:");
    println!("  let explicit_u8: u8 = 255;");
    println!("  let explicit_f32: f32 = 3.14;");
    println!("  let explicit_tuple: (i32, &str) = (100, \"Rust\");");
    println!("  let explicit_array: [f64; 3] = [1.1, 2.2, 3.3];");
    println!("\nOutput:");

    let explicit_u8: u8 = 255;
    let explicit_f32: f32 = 3.14;
    let explicit_tuple: (i32, &str) = (100, "Rust");
    let explicit_array: [f64; 3] = [1.1, 2.2, 3.3];

    println!("  explicit_u8 = {}", explicit_u8);
    println!("  explicit_f32 = {}", explicit_f32);
    println!("  explicit_tuple = {:?}", explicit_tuple);
    println!("  explicit_array = {:?}\n", explicit_array);

    // ========================================
    // 9. CONSTANTS WITH TYPES
    // ========================================
    println!("--- 9. Constants (Type Required) ---");
    println!("Code:");
    println!("  const MAX_SCORE: u32 = 100;");
    println!("  const PI: f64 = 3.14159265359;");
    println!("  const APP_NAME: &str = \"RustApp\";");
    println!("  const IS_DEBUG: bool = false;");
    println!("\nOutput:");

    const MAX_SCORE: u32 = 100;
    const PI: f64 = 3.14159265359;
    const APP_NAME: &str = "RustApp";
    const IS_DEBUG: bool = false;

    println!("  MAX_SCORE = {}", MAX_SCORE);
    println!("  PI = {}", PI);
    println!("  APP_NAME = {}", APP_NAME);
    println!("  IS_DEBUG = {}\n", IS_DEBUG);

    // ========================================
    // 10. TYPE CASTING
    // ========================================
    println!("--- 10. Type Casting ---");
    println!("Code:");
    println!("  let int_val: i32 = 42;");
    println!("  let float_val = int_val as f64;");
    println!("  let byte_val = int_val as u8;");
    println!("  let char_val = 65u8 as char;  // ASCII 'A'");
    println!("\nOutput:");

    let int_val: i32 = 42;
    let float_val = int_val as f64;
    let byte_val = int_val as u8;
    let char_val = 65u8 as char;

    println!("  int_val = {} (i32)", int_val);
    println!("  float_val = {} (f64)", float_val);
    println!("  byte_val = {} (u8)", byte_val);
    println!("  char_val = {} (char)\n", char_val);

    // ========================================
    // 11. STRING TYPES
    // ========================================
    println!("--- 11. String Types ---");
    println!("Code:");
    println!("  let string_slice: &str = \"Hello\";  // Immutable");
    println!("  let string_object: String = String::from(\"World\");");
    println!("  let concatenated = format!(\"{{}} {{}}\", string_slice, string_object);");
    println!("\nOutput:");

    let string_slice: &str = "Hello";
    let string_object: String = String::from("World");
    let concatenated = format!("{} {}", string_slice, string_object);

    println!("  string_slice = {}", string_slice);
    println!("  string_object = {}", string_object);
    println!("  concatenated = {}\n", concatenated);

    // ========================================
    // 12. NUMERIC LITERALS
    // ========================================
    println!("--- 12. Numeric Literals ---");
    println!("Code:");
    println!("  let decimal = 98_222;");
    println!("  let hex = 0xff;");
    println!("  let octal = 0o77;");
    println!("  let binary = 0b1111_0000;");
    println!("  let byte = b'A';  // u8 only");
    println!("\nOutput:");

    let decimal = 98_222;
    let hex = 0xff;
    let octal = 0o77;
    let binary = 0b1111_0000;
    let byte = b'A';

    println!("  decimal = {}", decimal);
    println!("  hex = {}", hex);
    println!("  octal = {}", octal);
    println!("  binary = {}", binary);
    println!("  byte = {}\n", byte);

    // ========================================
    // KEY TAKEAWAYS
    // ========================================
    println!("========================================");
    println!("         KEY TAKEAWAYS");
    println!("========================================\n");

    println!("1. SCALAR TYPES (Single Values)");
    println!("   - Integers: i8, i16, i32, i64, i128, isize (signed)");
    println!("   - Integers: u8, u16, u32, u64, u128, usize (unsigned)");
    println!("   - Floats: f32, f64 (default is f64)");
    println!("   - Boolean: bool (true/false)");
    println!("   - Character: char (4 bytes, Unicode)\n");

    println!("2. COMPOUND TYPES (Multiple Values)");
    println!("   - Tuple: Fixed size, mixed types (i32, f64, char)");
    println!("   - Array: Fixed size, same type [i32; 5]\n");

    println!("3. TYPE ANNOTATIONS");
    println!("   - Optional for 'let' (Rust infers types)");
    println!("   - Required for 'const'");
    println!("   - Syntax: let name: type = value;\n");

    println!("4. TYPE INFERENCE");
    println!("   - Rust automatically determines types");
    println!("   - Default integer: i32");
    println!("   - Default float: f64\n");

    println!("5. CONSTANTS");
    println!("   - Must have explicit type annotation");
    println!("   - Always immutable");
    println!("   - Evaluated at compile time");
    println!("   - Use SCREAMING_SNAKE_CASE naming\n");

    println!("6. TYPE CASTING");
    println!("   - Use 'as' keyword for explicit conversion");
    println!("   - Example: let x = 5i32 as f64;\n");

    println!("7. STRING TYPES");
    println!("   - &str: String slice (immutable, borrowed)");
    println!("   - String: Owned, growable string\n");

    println!("========================================");
    println!("     END OF DATA TYPES TUTORIAL");
    println!("========================================\n");
}
