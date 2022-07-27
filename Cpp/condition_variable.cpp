#include <iostream>
#include <condition_variable>
#include <mutex>
#include <thread>
#include <queue>
#include <chrono>

int main()
{
        std::queue<int> produced_nums; // producer
        std::mutex m; // mutex lock
        std::condition_variable cond_var; 
        bool done = false;
        bool notified = false;
        
        // producer thread
        std::thread producer([&]() {
                for (int i = 0; i < 5; ++i) {
                        std::this_thread::sleep_for(std::chrono::seconds(1));
                        std::unique_lock<std::mutex> lock(m);
                        std::cout << "producing..." << i << '\n';
                        produced_nums.push(i);
                        notified = true;
                        cond_var.notify_one();
                }
                done = true;
                notified = true;
                cond_var.notify_one();
        });

        // consumer thread
        std::thread consumer([&]() {
                std::unique_lock<std::mutex> lock(m);
                while (!done) {
                        while (!notified) {
                                cond_var.wait(lock);
                        }
                        while (!produced_nums.empty()) {
                                std::cout << "consuming..." << '\n';
                                produced_nums.pop();
                        }
                        notified = false;
                }
        });

        producer.join();
        consumer.join();

        return 0;
}

