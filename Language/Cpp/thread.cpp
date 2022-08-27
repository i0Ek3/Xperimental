#include <iostream>
#include <thread>

void foo() {
        std::cout << "Hello World." << std::endl;
}

int main()
{
        std::thread t(foo); // create a thread
        t.join(); // join to a thread

        return 0;
}

