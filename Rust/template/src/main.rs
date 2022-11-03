/*
 *
 * Rust Code
 *
 */

fn types() {
    /*
     * number: i8, i16, i32, i64, isize
     *         u8, u16, u32, u63, usize
     *         f32, f64
     * string: &str
     * boolean: true false
     * char
     * unit: ()
     */
    let a = 10;
    let b: i32 = 100;
    let c = 1000i32;
    let d: i64 = 1_000_000_000;
    let e = [
        3.0,
        4.1f32,
        5.21_f32,
    ];
    println!("{} {} {} {} {:.2}", a, b, c, d, e[0]);
}

fn range() {
    for i in 1..10 {
        println!("{}", i);
    }

    for i in 1..=10 {
        println!("{}", i);
    }
}

fn variables() {
    let _s = "hi";

    let mut str = "Hi, Rust!";
    println!("{}", str);
    str = "Hi";
    println!("{}", str);

    let x = 1;
    let x = x + 1;
    {
        let x = x * 3;
        println!("{}", x);
    }
    println!("{}", x);
}

fn lines() {
    println!("===========================================");
}

fn main() {
    lines();
    variables();
    lines();
    types();
    lines();
    range();
    lines();
}
