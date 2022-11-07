# Rust Notes

> TODO: 复合类型

## Chapter

- 前置基础
  - 可变 mut
  - 变量绑定、解构
  - 常量 const，变量 let
  - 变量遮蔽（覆盖）
  - 基本类型
    - 数值
      - 有符号整数、无符号整数、浮点数、有理数、复数 
    - 字符 char，单个 Unicode 字符，存储为 4 个字节
    - 布尔
    - 单元，即`()`，类似 Go 中的 struct{}
    - 语句
      - `let x = x + 1;`
      - 语句后面需要加分号
    - 表达式
      - `x + x`
      - 表达式后面不加分号，如果加了分号，则返回一个`()`
      - 表达式如果不返回任何值，会隐式地返回一个`()`
      - 函数也是表达式
    - 函数
      ![](https://pic2.zhimg.com/80/v2-54b3a6d435d2482243edc4be9ab98153_1440w.png)
      - 当用`!`作函数返回类型的时候，表示该函数永不返回(diverge function)，这种语法往往用做会导致程序崩溃的函数
  - 类型推导与标注
  - 所有权
    > 栈中的所有数据都必须占用已知且固定大小的内存空间；而对于大小未知或者可能变化的数据，我们需要将它存储在堆上。
    - 所有权原则
      - Rust 中每一个值都被一个变量所拥有，该变量被称为值的所有者
      - 一个值同时只能被一个变量所拥有，或者说一个值只能拥有一个所有者
      - 当所有者(变量)离开作用域范围时，这个值将被丢弃(drop)
    - 变量绑定背后的数据交互
      - 所有权转移
        - 即上述原则的第三条在 Rust 所有权转移中会自动进行
        - 不同于浅拷贝，Rust 中称为移动(move)
      - 克隆（深拷贝）
        - clone() 
        - Rust 永远也不会自动创建数据的 “深拷贝”，任何自动的复制都不是深拷贝
      - 拷贝（浅拷贝）
        - 浅拷贝只发生在栈上，因此性能很高，在日常编程中，浅拷贝无处不在
        - 任何基本类型的组合是 Copyable 的，不需要分配内存或某种形式资源的类型是 Copyable 的
          - 不可变引用 `&T`
      - 引用与借用
        - 获取变量的引用，称之为借用(borrowing)
        - 同一作用域，特定数据只能有一个可变引用
        - 可变引用与不可变引用不能同时存在
        - 悬垂引用(Dangling References)
          - 即悬垂指针，意思为指针指向某个值后，这个值被释放掉了，而指针仍然存在，其指向的内存可能不存在任何值或已被其它变量重新使用
        - 引用必须总是有效的

## Advices

- Don't upload Cargo.lock to your git repo except your project is a runnable program

## References

- [How to disable warning message?](https://stackoverflow.com/questions/25877285/how-to-disable-unused-code-warnings-in-rust)
