// 逆波兰解析器

#include <iostream>
#include <stack>
#include <iterator>
#include <vector>
#include <cmath>
#include <map>
#include <sstream>
#include <stdexcept>
#include <cassert>

template <typename T>
double evaluate_rpn(T it, T end) {
    std::stack<double> val_stack;

    auto pop_stack ([&]() {
        auto r (val_stack.top()); // 获取栈顶元素的副本
        val_stack.pop(); // 从栈中移除栈顶元素
        return r; // 返回栈顶元素副本
    });

    std::map<std::string, double (*)(double, double)> ops {
        {"+", [](double a, double b) {return a + b; }},
        {"-", [](double a, double b) {return a - b; }},
        {"*", [](double a, double b) {return a * b; }},
        {"/", [](double a, double b) {return a / b; }},
        {"^", [](double a, double b) {return pow(a, b); }},
        {"%", [](double a, double b) {return fmod(a, b); }},

    };

    for (; it != end; it++) {
        std::stringstream ss {*it}; //为了获取每个单词，将流中的类型解析并转换成一个都变了变量。
        if (double val; ss >> val) {
            val_stack.push(val);
        } else {
            const auto r {pop_stack()};
            const auto l {pop_stack()};
            try {
                const auto & op (ops.at(*it));
                const double result {op(l, r)};
                val_stack.push(result);
    
            } catch (const std::out_of_range &) {
                throw std::invalid_argument(*it);
            }
        }
    }
    return val_stack.top();
}


int main()
{
    // way1
    try {
        std::cout << evaluate_rpn(std::istream_iterator<std::string>{std::cin}, {}) << std::endl;
    } catch (const std::invalid_argument &e) {
        std::cout << "Invalid operator: " << e.what() << std::endl;
    }
    
    std::cout << "-------------------------" << std::endl;
    
    // way2
    std::stringstream s {"3 2 1 + * 2 /"};
    std::cout << evaluate_rpn(std::istream_iterator<std::string>{s}, {}) << std::endl; 
    
    std::cout << "-------------------------" << std::endl;
    
    // way3
    std::vector<std::string> v {"3", "2", "1", "+", "*", "2", "/"};
    std::cout << evaluate_rpn(begin(v), end(v)) << std::endl;


    return 0;
}

