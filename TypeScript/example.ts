// Code reference: https://mp.weixin.qq.com/s/3fqC8VxazoaiPvNu6eVI6A 

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

// type protection
function getLen(arg: number | string): number {
    // return arg.length // wrong
    if (typeof arg === 'string') {
        return arg.length
    } else {
        return arg.toString().length
    }
}

// type assertion
function getLen2(arg: number | string): number {
    const str = arg as string
    if (str.length) {
        return str.length
    } else {
        const number = arg as number
        return number.toString().length
    }
}

// type literal
type Size = 'L' | 'M' | 'X' | 'XL'

// generic: almost same with Go
function print<T>(arg: T): T {
    console.log(arg)
    return arg
}
print<string>('hi')
print('i0Ek3')

// T = number denotes default type in generic
interface Print<T = number> {
    (arg: T): T
}

//const p_: Print<number> = print
const p_: Print = print

function swap<T, U>(tuple: [T, U]): [U, T] {
    return [tuple[1], tuple[0]]
}

const t_ = swap(['i0Ek3', 18])

interface UserInfo {
    name: string
    age: number
}

function request<T>(url: string): Promise<T> {
    return fetch(url).then(res => res.json())
}

request<UserInfo>('user/info').then(res => {
    console.log(res.age)
})

// generic T also can implement interface
interface Length {
    length: number
}

function printLen<T extends Length>(arg: T): T {
    console.log(arg.length)
    return arg
}

// A generic can constrain a class, but can not
// constrain static members of a class
class Stack<T> {
    private data: T[] = []
    push(item: T) {
        return this.data.push(item)
    }
    pop(): T | undefined {
        return this.data.pop()
    }
}

const st = new Stack<number>()
st.push(18)
//st.push('i0Ek3')

// also generic can constrain interface
interface KV<T, U> {
    key: T
    val: U
}

const k1: KV<number, string> = { key: 18, val: 'i0Ek3' }
const k2: KV<string, number> = { key: 'i0Ek3', val: 18 }

const arr: Array<number> = [1, 2, 3]

// type index

// common use case for get values from userInfo
interface userInfo {
    name: string
    age: string
}

const userInfo = {
    name: 'i0Ek3',
    age: '18',
}

function getVals(userInfo: any, keys: string[]) {
    return keys.map(key => userInfo[key])
}

console.log(getVals(userInfo, ['name', 'age']))
console.log(getVals(userInfo, ['a', 'b']))

interface IPerson {
    name: string
    age: number
}

type keyname = keyof IPerson

let t1: IPerson['name']
let t2: IPerson['age']

// after generic
function getValues<T, K extends keyof T>(userInfo: T, keys: K[]): T[K][] {
    return keys.map(key => userInfo[key])
}

// type mapping: Partial, Readonly, Pick and Record
// use in keyword to traversal array
type PersonInfo = 'name' | 'school' | 'grade'
type Obj = {
    [pi in PersonInfo]: string
}

// Partial<T> maps all attributes of T to optional
type IPartial = Partial<IPerson>
let p1: IPartial = {}

// How?
type Partial<T> = {
    [P in keyof T]?: T[P]
}

// Readonly<T> maps all attributes of T to readonly
type IReadOnly = Readonly<IPerson>
let ro: IReadOnly = {
    name: 'i0Ek3',
    age: 18,
}

// How?
type ReadOnly<T> = {
    readonly [P in keyof T]: T[P]
}

interface OtherPerson {
    name: string
    age: number
    sex: boolean
}

type IPick = Pick<OtherPerson, 'name' | 'sex'>
let pick: IPick = {
    name: 'i0Ek3',
    sex: true,
}

// How?
type Pick<T, K extends keyof T> = {
    [P in K]: T[P]
}

// the type Partial, Pick, Readonly we mentioned above are all homomorphic mapping types
// which means they are only work on obj properties instead of introduced new properties
// but Record is non-homomorphic mapping type, it will introduced new properties

type IRecord = Record<string, IPerson>

let pm: IRecord = {
    p1: {
        name: 'i0Ek3',
        age: 18,
    },
    p2: {
        name: 'John',
        age: 19,
    }
}

// How?
type Record<K extends keyof any, T> = {
    [P in K]: T
}

// type condition
// T extends U ? X : Y

// Exclude<T, U>
type exclude = Exclude<'a' | 'b' | 'c', 'a'>
// How: type Exclude<T, U> = T extends U ? never : T

type test = string | number | never

// Extract<T, U>
type extract = Extract<'a' | 'b', 'a'>
// How: type Extract<T, U> = T extends U ? T : never

// Utils

// Omit<T, U>, from the type T eliminate all the properties in U
type IOmit = Omit<IPerson, 'age'>
// How: 
// 1 type Omit<T, K extends keyof any> = Pick<T, Exclude<keyof T, K>>
// 2 type Omit<T, K extends keyof any> = {
//    [P in Exclude<keyof T, K>]: T[P]
// }

// NonNullable<T>, use to filter null, undefined type in T
type t0 = NonNullable<string | number | null | undefined>
// How: type NonNullable<T> = T extends null | undefined ? never : T

// Parameters use to get the type of function argument
type t1 = Parameters<() => string>
type t2 = Parameters<(arg: string) => void>
// How: type Parameters<T extends (...args: any) => any> = T extends (...args: infer P) => any ? P : never

// ReturnType
type t3 = ReturnType<() => string>
type t4 = ReturnType<(arg: string) => void>
// How: type ReturnType<T extends (...args: any) => any> = T extends (...args: any) => infer R ? R : any

// declare
interface VueOption {
    el: string,
    data: any,
}

declare class Vue {
    options: VueOption
    constructor(options: VueOption)
}

const app = new Vue({
    el: '#app',
    data: {
        message: 'Vue in TS'
    }
})