#include <iostream>
#include <queue>
#include <list>
#include <string>


// Todo.cpp
// implement by priority_queue



int main()
{
    using item = std::pair<int, std::string>;
    std::priority_queue<item> q;

    std::initializer_list<item> il {
        {1, "paper"},
        {0, "music"},
        {2, "coding"},
        {3, "graduate"},
    };

    for (const auto &p : il) {
        q.push(p);
    }
    
    while (!q.empty()) {
        std::cout << q.top().first << ": " << q.top().second << std::endl;
        q.pop();
    }

    return 0;
}

