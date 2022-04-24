use std::env;
use std::process::exit;

fn fibo(x: usize, y: u128, z: u128) {
    let mut fibonacci: Vec<u128> = [y, z].to_vec();
    println!("1: {}\n2: {}", fibonacci[0], fibonacci[1]);
    for i in 2..x {
        fibonacci.push(fibonacci[i - 1] + fibonacci[i - 2]);
        println!("{}: {}", i + 1, fibonacci[i]);
    }
}

fn main() {
    let argv: Vec<String> = env::args().collect();
    if argv.len() != 4 {
        println!("You have to provide 3 arguments - first number, secound number, how many numbers to calculate.\nExample:\t./fibo 0 1 10\n");
        exit(1);
    }
    fibo(
        argv[3].parse::<usize>().unwrap(),
        argv[1].parse::<u128>().unwrap(),
        argv[2].parse::<u128>().unwrap(),
    );
    exit(0);
}
