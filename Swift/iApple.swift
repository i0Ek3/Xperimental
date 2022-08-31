// Swift 5.5.1 example

func Hi() {
    let date = "2021-10-31"

    let version = "5.5.1"
    let msg = "This is Swift \(version) example."

    print("\(msg) Created on \(date).")
}

class Test {
    var str = "Class"
    var name: String

    init(name: String) {
        self.name = name
    }

    func show() -> String {
        return "This is a \(str) test."
    }
}

// like interface in Go
protocol Protocol {
    var desc: String { get }
    mutating func show()
}

class Example: Protocol {
    var desc: String = "Example class,"

    func show() {
        desc += " just for Swift Protocol."
        print(desc)
    }
}

func runit() {
    // basic one
    Hi()

    // Class
    let t = Test(name: "Test")
    print(t.show())

    // Protocol
    let e = Example()
    e.show()

    // 
}

runit()
