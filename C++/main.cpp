#include <iostream>
#include <fstream>
#include "Filewc.hpp"

using namespace Filewc;
int main(int argc, char *argv[]) {
    if (argc < 2){ return 0; }
    std::ifstream file(argv[1]);
    if (!file){ return 0; }

    countLines(file);

    file.close();
    return 0;
}
