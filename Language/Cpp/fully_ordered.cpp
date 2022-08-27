#include <iostream>
#include <atomic>
#include <thread>
#include <assert.h>

std::atomic<bool> x, y;
std::atomic<int> z;

void write_to_x() {
    x.store(true, std::memory_order_seq_cst);
}

void write_to_y() {
    y.store(true, std::memory_order_seq_cst);
}

void read_x_then_y() {
    while (!x.load(std::memory_order_seq_cst));
    if (y.load(std::memory_order_seq_cst)) {
        ++z;
    }
    std::cout << "z = " << z << std::endl;
}

void read_y_then_x() {
    while (!y.load(std::memory_order_seq_cst));
    if (x.load(std::memory_order_seq_cst)) {
        ++z;
    }
    std::cout << "z = " << z << std::endl;
}

int main()
{
    x = false;
    y = false;
    z = 0;
    
    std::thread a(write_to_x);
    std::thread b(write_to_y);
    std::thread c(read_x_then_y);
    std::thread d(read_y_then_x);

    a.join();
    b.join();
    c.join();
    d.join();

    assert(z.load() != 0);
    
    return 0;
}
