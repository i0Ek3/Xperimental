#include <iostream>
#include <map>
#include <vector>
#include <algorithm>
#include <iomanip>
#include <string>

// words count


std::string filter_punctuation(const std::string &s) {
    const char *forbidden {".,:;?_-"};
    const auto idx_start (s.find_first_not_of(forbidden));
    const auto idx_end (s.find_last_not_of(forbidden));
    return s.substr(idx_start, idx_end - idx_start + 1);
}

int main()
{
    std::map<std::string, size_t> words;
    int max_word_len {0};
    std::string s;
    while (std::cin >> s) {
        auto filtered (filter_punctuation(s));
        max_word_len = std::max<int>(max_word_len, filtered.length());
        ++words[filtered];
    }
    std::vector<std::pair<std::string, size_t>> word_counts;
    word_counts.reserve(words.size());
    move(begin(words), end(words), back_inserter(word_counts));

    std::sort(begin(word_counts), end(word_counts), [](const auto &a, const auto &b) { return a.second > b.second; });
    std::cout << "# " << std::setw(max_word_len) << "<WORD>" << " #<COUNT>\n";
    for (const auto & [word, count] : word_counts) {
        std::cout << std::setw(max_word_len + 2) << word << " #" << count << '\n';
    }
    
    return 0;
}

