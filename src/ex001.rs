fn is_divisable(n: i32) -> bool {
    // n が 3 もしくは 5の倍数になっているか否かを返す
    n % 3 == 0 || n % 5 == 0
}

fn main() {
    let mut s = 0;
    for n in 0..1000 {
        if is_divisable(n) {
            s += n
        }
    }

    println!("{}", s);
}
