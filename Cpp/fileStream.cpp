// ref: https://wenku.baidu.com/view/0cb75850f01dc281e53af075.html



#include <iostream>
#include <fstream> // 文件输入输出函数的头文件
#include <string>

void w() {
    std::string str = "String test";
    std::ofstream saveFile("test.txt"); // ofstream实际上是一个类，输出文件流，saveFile是一个句柄，为了方便文件的输入输出
    //std::ofstream saveFile;
    //saveFile.open("test.txt", ios::app); // 当然了，你也可以用这种方式来打开一个文件，并指定好打开的模式。如果你需要指定多个打开模式，使用|即可。如，ios::in | ios::ate.
    saveFile << "Just a test file!\n"; // 将内容写入文件
    saveFile << str;
    saveFile.close(); // 关闭文件流，关闭之后便不能再继续访问，除非你再次打开文件流
}

void r() {
    // iftream为输入文件流，同上面的ofstream
    std::ifstream openFile("text.txt"); // 如果文件不在当前目录，你需要文件的全部路径名
    char ch;
    while (!openFile.eof()) { // 遍历整个文件，检测是个到了末尾
        openFile.get(ch); // get()函数从相应的流文件中读出一个字符，并将其返回给变量
        std::cout << ch;
    }
    openFile.close();
}

int main() 
{
    w();
    r();

    return 0;
}

