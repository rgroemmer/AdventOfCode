use std::fs::read_to_string;

fn main() {
    let filename = "input";

    for line in read_to_string(filename).unwrap().lines() {
        let splitted: Vec<&str> = line.split_whitespace().collect();
        let mut increasing: Option<&bool> = None;

        for i in 1..splitted.len() {
            // Check if report is safe
            let before = splitted[i - 1];
            let act = splitted[i];

            // Check if increasing
            if before < act {
                match increasing {
                    Some(value) => {
                        if *value == false {
                            // It was descreasing, now its increasing: Unsafe!!
                            break;
                        }
                    }
                    // Its nil, set to increasing
                    None => increasing = Some(&true),
                }
            }
            if before > act {
                match increasing {
                    Some(value) => {
                        if *value == true {
                            // It was increasing, not its descreasing: Unsafe!!
                            break;
                        }
                    }
                    None => increasing = Some(&false),
                }
            }
        }
    }

    let mut result_one: i32 = 0;
    let mut result_two: i32 = 0;

    println!("Result part 1: {}", result_one);

    println!("Result part 2: {}", result_two);
}
