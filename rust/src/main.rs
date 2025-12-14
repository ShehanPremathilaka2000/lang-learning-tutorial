use std::io::{self, Write};
mod arrays;
mod borrowing;
mod conditions;
mod data_types;
mod enums;
mod functions;
mod hash_map;
mod loops;
mod operators;
mod ownership;
mod strings;
mod structs;
mod tuples;
mod variables;
mod vectors;

fn main() {
    print_welcome_banner();

    loop {
        print_menu();

        let mut input = String::new();
        io::stdout().flush().expect("Failed to flush stdout"); // Ensure prompt appears before input
        io::stdin()
            .read_line(&mut input)
            .expect("Failed to read line");

        let choice: i32 = match input.trim().parse() {
            Ok(num) => num,
            Err(_) => {
                println!("\n❌ Invalid input! Please enter a number.");
                continue;
            }
        };

        if choice == 0 {
            print_goodbye();
            break;
        }

        execute_choice(choice);

        // Add a small divider before showing menu again
        println!("\n{}", "─".repeat(60));
    }
}

fn print_welcome_banner() {
    println!("\n{}", "═".repeat(60));
    println!("╔════════════════════════════════════════════════════════╗");
    println!("║                                                        ║");
    println!("║          🚀  WELCOME TO RUST TUTORIAL  🚀             ║");
    println!("║                                                        ║");
    println!("║        Learn Rust Programming Step by Step            ║");
    println!("║                                                        ║");
    println!("╚════════════════════════════════════════════════════════╝");
    println!("{}\n", "═".repeat(60));
}

fn print_menu() {
    println!("\n{}", "═".repeat(60));
    println!("📚  AVAILABLE TOPICS");
    println!("{}", "═".repeat(60));
    println!();

    let topics = [
        "Variables",
        "Data Types",
        "Operators",
        "Conditions",
        "Loops",
        "Functions",
        "Strings",
        "Ownership",
        "Borrowing",
        "Arrays",
        "Vectors",
        "Tuples",
        "Hash Maps",
        "Structs",
        "Enums",
    ];

    for (i, topic) in topics.iter().enumerate() {
        print!("  {:2}. {:<20}", i + 1, topic);
        if (i + 1) % 2 == 0 {
            println!();
        }
    }
    // Handle odd number of topics nicely if needed, but above loop handles newlines well.
    // If the last one didn't print a newline, we print one.
    if topics.len() % 2 != 0 {
        println!();
    }

    println!();
    println!("{}", "─".repeat(60));
    println!("   0. Exit Tutorial");
    println!("{}", "═".repeat(60));
    print!("\n👉 Enter your choice: ");
}

fn execute_choice(choice: i32) {
    println!(); // Spacing
    match choice {
        1 => variables::variables(),
        2 => data_types::data_types(),
        3 => operators::operators(),
        4 => conditions::conditions(),
        5 => loops::loops(),
        6 => functions::functions(),
        7 => strings::strings(),
        8 => ownership::ownership(),
        9 => borrowing::borrowing(),
        10 => arrays::arrays(),
        11 => vectors::vectors(),
        12 => tuples::tuples(),
        13 => hash_map::hash_map(),
        14 => structs::structs(),
        15 => enums::enums(),
        _ => println!("❌ Invalid choice! Please enter a number between 0 and 15."),
    }
}

fn print_goodbye() {
    println!("\n{}", "═".repeat(60));
    println!("╔════════════════════════════════════════════════════════╗");
    println!("║                                                        ║");
    println!("║          ✨  Thank You for Learning Rust!  ✨         ║");
    println!("║                                                        ║");
    println!("║              Keep Coding and Have Fun! 🦀             ║");
    println!("║                                                        ║");
    println!("╚════════════════════════════════════════════════════════╝");
    println!("{}\n", "═".repeat(60));
}
