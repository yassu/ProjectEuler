fn max_factorial(n: i64) -> i64 {
    let mut n = n;
    let mut max_fact = 1;
    for num in 2 .. (n+1) {
        if n % num == 0 {
            max_fact = num;
            while n % num == 0 {
                n /= num;
            }
        }
        if n == 1 {
            break;
        }
    }
    max_fact
}


fn main() {
    println!("{}", max_factorial(600851475143));
}
