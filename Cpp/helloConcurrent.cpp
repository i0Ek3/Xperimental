#include <iostream>
#include <thread> //用来支持多线程

void hello() {
        std::cout << "Hello Concurrent!" << std::endl;
}

int main() // 初始线程
{
    std::thread t(hello); //t的初始函数为hello()
    t.join(); //等待
}

