// reference standard libaray
use std::f64::consts::PI;

// import module
mod hust_mod;

// enable debug mode
#[derive(Debug)]

// define struct
struct Triangle {
    a: u32,
    b: u32,
    c: u32,
}

fn h(hi: String) {
    println!("{}", hi);
}

use std::io::stdin;
fn inputer() {
    let mut strs = String::new();
    stdin().read_line(&mut strs).expect("Cannot to read line.");
    println!("The input content is: {}", strs);
}

use std::fs;
fn reader(filepath: String) {
    let file = fs::read_to_string(filepath).unwrap();
    println!("{}", file)
}

use std::io::prelude::*;
use std::fs::File;
fn writer(filepath: String, content: String) {
    let mut file = File::create(filepath).unwrap();
    file::write(content).unwrap();
}

// use keyword for module
use crate::hust_mod::hust::_1::_1f;
use crate::hust_mod::hust::_1::_1f as f1;

fn main() {
    let hi = "Hi rust";
    let _i: i32 = 0;
    for _i in 0..7 {
        h(hi.to_string());
    }

    let str = String::from(hi);
    let slice = &str[..2];
    println!("{}", slice);

    let t = Triangle{ a: 10, b: 20, c: 30 };
    println!("t is {:?}", t);

    //println!("{}", hust_mod::hust::_1::_1f())
    println!("{}", _1f());
    println!("{}", f1());

    println!("{}", (PI * 4.0 / 3.0).cos());

    // cmd args
    let args = std::env::args();
    for arg in args {
        println!("{}", arg);
    }

    // input string
    println!("Please input something:");
    inputer();

    // read string from file
    reader("./hust".to_string());
}
