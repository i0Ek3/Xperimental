//
// clang++ future.cpp -std=c++11
// 


#include <iostream>
#include <future>
#include <thread>


int main()
{
        std::packaged_task<int()> task([](){return 7;}); //lambda --> task 
        std::future<int> result = task.get_future();
        std::thread(std::move(task)).detach();
        std::cout << "Waiting...";
        result.wait();
        std::cout << "Done!" << std::endl << "Result is " << result.get() << '\n';
        
        return 0;
}

