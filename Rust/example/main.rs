mod hust;
use crate::hust::Hust;

fn main() {
    //let hi = "Hi hust.";
    let hi = String::from("Hi hust.");
    let obj = Hust::new(hi.to_string());
    obj.getter();
}
