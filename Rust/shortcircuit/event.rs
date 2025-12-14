struct Event<T> {
    evt1: T,
    evt2: T
}

impl<T: std::cmp::PartialEq + std::cmp::PartialOrd> Event<T> {
    fn get(&self) -> &T {
        if self.evt1 == self.evt2 || self.evt1 > self.evt2 {
            &self.evt1
        } else {
            &self.evt2
        }
    }
}

fn main() {
    // simple example with array, next step we should use different event to indicate that
    let arr = [1, 2];
    let e = Event { evt1: arr[0], evt2: arr[1] };
    println!("e.evt = {}", e.get());
}
