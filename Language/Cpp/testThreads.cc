#include <iostream>
#include <thread>

int main()
{
    unsigned long const len = std::thread::hardware_concurrency();
    std::cout << "len = " << len << std::endl;
    return 0;
}
