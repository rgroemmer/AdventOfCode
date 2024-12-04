use std::fs::read_to_string;

fn main() {
    let filename = "input";
    let mut col1 = Vec::new();
    let mut col2 = Vec::new();

    for line in read_to_string(filename).unwrap().lines() {
        let splitted: Vec<&str> = line.split_whitespace().collect();
        if splitted.len() == 2 {
            col1.push(splitted[0].to_string());
            col2.push(splitted[1].to_string());
        }
    }
    col1.sort();
    col2.sort();

    let mut result_one: i32 = 0;
    let mut result_two: i32 = 0;

    for i in 0..col1.len() {
        let val1: i32 = col1[i].parse().expect("Failed to parse string of col1.");
        let val2: i32 = col2[i].parse().expect("Failed to parse string of col2.");

        if val2 > val1 {
            result_one += val2 - val1;
            continue;
        }
        result_one += val1 - val2;
    }
    println!("Result part 1: {}", result_one);

    // TODO: find out what exactly does bottow mean (&col1?)
    for s1 in &col1 {
        let mut count = 0;
        for s2 in &col2 {
            if s1 == s2 {
                count += 1;
            }
        }
        let val1: i32 = s1.parse().expect("Failed to parse string of col1.");
        result_two += val1 * count;
    }
    println!("Result part 2: {}", result_two);
}
