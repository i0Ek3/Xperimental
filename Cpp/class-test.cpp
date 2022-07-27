#include <iostream>

class A {
public:
    void print() {
        std::cout << "this is A!" << std::endl;
    }
};

class B : public A {};
class C : public A {};
class D : public B, public C {};

void test1() {
    // 当然，下面的二义性问题可以将基类A声明为virtual来解决

    D d;
    //A* pa = (A*)&d; // 产生二义性
    A* pa1 = (A*)(B*)&d; // 需要指明具体转换的类型
    A* pa2 = (A*)(C*)&d;
    d.B::print(); // 需要指明调用哪个类的print函数
    d.C::print();
    //d.print(); // 这将产生二义性
}

void test2() {
    D d;
    A* pa = (A*)&d;
    d.print();
}

int main()
{
    //test1();
    test2();

    return 0;
}
