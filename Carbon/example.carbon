package example api;

fn Hello(name: String) -> String {
    return name;
} 

fn Square(x: i32) -> i32 {
    return x * x;
}

fn Main() -> i32 {
    return Square(12);
}
