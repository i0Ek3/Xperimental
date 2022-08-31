#include <iostream>
#include <thread>
#include <mutex>

std::mutex mtx;

void block_area() {
        std::unique_lock<std::mutex> lock(mtx);
        // critical zone
        lock.unlock();
        // other code 
        lock.lock();
}

int main()
{
        std::thread thd1(block_area);
        thd1.join();

        return 0;
}

