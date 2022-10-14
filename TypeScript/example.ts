// TS types
function typeInTS(): void { // void
    console.log("this is type set in TypeScript")
    let ok: boolean = true
    let num: number = 18
    let name: string = "i0Ek3"
    let un: undefined = undefined
    let nil: null = null
    let unsure: any = 4 // any just like := in Go, but unsecure
    let unsure2: unknown = 0 // recommend
    let nvr: never
    let list: number[] = [1, 2, 3]
    let tuple: [number, string] = [18, 'i0Ek3']

    // just like lambada in Python or C++
    let sub = (a: number, b: number): number => {
        return a - b
    }

    let fn = function() {}

    // union type, like ∩
    let multi: number | string | boolean
    multi = 8
    multi = '8'
    multi = true
    console.log(multi)

    // cross type, use & to denote, like U in math
    interface A {
        name: string
        age: number
    }
    type B = A & { sex: boolean }

    // type alias
    type Name = string
    // => func type just like anonymous function
    // function an(fn: (a: string) => void) {
    //     fn("")
    // }
    type NameFunc = () => string
    type NameOrSth = Name | NameFunc

    // what's the difference between type and interface?
    // type can declare primitive types, for example union types, cross types, tuples, but not interface.
    // An interface can combine repeated declarations, but not a type.
    interface AA {
        a: string
    }

    interface AA {
        b: number
    }

    const aa: AA = {
        a: 'i0Ek3',
        b: 18.
    }
    console.log(aa.a, aa.b)

    // But, which one should we use?
    // When using combination or cross types, use type
    // When using extends or implements of class, use interface
}

// function parameter
// ? means options, must in last position
// c is default parameter, 
function parameter(a: number, b: number, c:number = 2, d?: number): number {
    return a + b + c
}

// function override
function add(x: number[]): number
function add(x: string[]): string
function add(x: number[], y: number[]): number
function add(x: string[], y: string[]): string
function add(x: any[], y?: any[]): any {
    if (Array.isArray(y) && typeof y[0] === 'number') {
        return x.reduce((acc, cur) => acc+cur) + y.reduce((acc, cur) => acc+cur)
    }
    if (Array.isArray(y) && typeof y[0] === 'string') {
        return x.join() + ',' + y.join()
    }
    if (typeof x[0] === 'string') {
        return x.join()
    } else if (typeof x[0] === 'number') {
        return x.reduce((acc, cur) => acc+cur)
    }
}

console.log(add([1, 1, 1]))
console.log(add(['i0Ek3', '18']))
console.log(add([1, 1, 1], [1, 2, 3]))
console.log(add(['i0Ek3', '18'], ['he\'s', 'good enough']))

// interface, just like interface in Go
interface Example {
    name: string
    readonly age: number
    unsure?: unknown
}

const e: Example = {
    name: 'i0Ek3',
    age: 18,
    unsure: 18,
}

// use interface describe function type
interface Sum{
    (a: number, b: number): number
}

const sum:Sum = (x, y) => {
    return x + y
}

interface Custom {
    [elem: string]: string
}

const obj: Custom = {
    a: 'a',
    b: 'b',
    c: 'c',
}

// interface also is so-called duck typing
interface FuncWithField {
    (x: number): number
    name: string
}

const fn: FuncWithField = (x) => {
    return x
}

fn.name = 'FuncWithField'

// class in TS
// all method are public by default, so public keyword can omit
class Person {
    protected name: string
    constructor(name: string) {
        this.name = name
    }

    speak() {
        console.log(`${this.name} can speak English`)
    }
}

const p = new Person('i0Ek3')
//console.log(p.name)
p.speak()

class Coder extends Person {
    static worktime = 996
    mofish: boolean
    private constructor(name: string, mofish: boolean) {
        super(name)
        this.mofish = mofish
    }
    public coding() {
        console.log(`${this.name} can code the whole day`)
    }
    // speak also is public 
    speak() {
        return `Coder ${super.speak()}`
    }
}

//const c = new Coder('icu', true)
//c.coding()

// abstract class just same with C++
abstract class Animal {
    constructor(name: string) {
        this.name = name
    }
    public name: string
    public abstract eat(): void
}

class Dog extends Animal {
    constructor(name: string) {
        super(name)
    }
    public eat() {
        console.log('dog eat bones')
    }
}

class Cat extends Animal {
    constructor(name: string) {
        super(name)
    }
    public eat() {
        console.log('cat eat rats')
    }
}

class ChainsInvoke {
    step1() {
        console.log('step 1')
        return this
    }
    step2() {
        console.log('step 2')
        return this
    }
}

const ci = new ChainsInvoke()
ci.step1().step2()

// the relationship between interface and class
// class can implement one or more interface
interface Generic {
    gen(): void
}

class Generator implements Generic {
    gen() {}
}

// ?
interface CircleStatic {
    new (radius: number): void
    pi: number
}

const Circle:CircleStatic = class Circle {
    static pi: 3.14
    public radius: number
    public constructor(radius: number) {
        this.radius = radius
    }
}

// enum just same with iota in Go
enum Direction {
    Up,
    Down,
    Left,
    Right,
}

console.log(Direction.Up)
console.log(Direction.Down)
console.log(Direction.Left)
console.log(Direction.Right)
console.log(Direction[0])
console.log(Direction[1])
console.log(Direction[2])
console.log(Direction[3])

enum Status {
    Enable = 1,
    Disable = 0,
}