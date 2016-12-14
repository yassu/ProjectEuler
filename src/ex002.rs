fn fib(n: i32) -> i32 {
    let mut a = 1; let mut b = 1;
    for _ in 0..(n-1) {
        let c = a;
        a = b;
        b = b + c;
    }
    a
}

fn main() {
    let mut s = 0;
    let mut n = 1;

    loop {
        if n % 3 == 0 {     // f % 2 == 0 <=> n % 3 == 0
            println!("{}", n);
            let f = fib(n);
            if f > 4000000 {
                break;
            }
            s += f;
        }
        n += 1;
    }
    println!("{}", s);
}
