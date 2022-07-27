// 
// clang++ regex.cpp -std=c++11
// pass these three parameters std::string/std::smatch/std::regex also can do so.
//

#include <iostream>
#include <string>
#include <regex>

int main()
{
        std::string strings[] = {"012.txt", "abns.txt", "asdfhilsdhf", "sa34hj34.txt", "AKSH.txt", "adsl.tct"};
        std::regex txt_regex("[a-zA-Z]+\\.txt");
        
        for (const auto &strings: strings) {
                std::cout << strings << ": " << std::regex_match(strings, txt_regex) << std::endl;
        }

        return 0;
}

