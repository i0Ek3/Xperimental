// 
// ref.cpp
// @ianpasm(kno30826@gmail.com)
// 2018-07-28 16:23:55
// 

#include <iostream>
using std::cout;
using std::endl;

int main()
{
    int a = 100;
    int & refa = a;

    cout << "a = " << a << " refa = " << refa << endl;

    refa++;
    cout << "a = " << a << " refa = " << refa << endl;

    cout << "a's addr = " << &a << " refa's addr = " << &refa << endl;

    return 0;
}
