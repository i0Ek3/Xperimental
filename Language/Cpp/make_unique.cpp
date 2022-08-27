// std::make_unique not support in C++11, we can implement it as follows:


#include <memory>


template <typename T, typename ...Args>
std::unique_ptr<T> make_unique( Args&& ...args ) {
        return std::unique_ptr<T>( new T( std::forward<Args>(args) ... ) );
}
