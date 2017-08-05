//
// Created by Arthur  on 6.08.17.
//

#include "Filewc.hpp"

int Filewc::countLines(std::ifstream &file) {
    if (file.bad()){
        return 0;
    }
    unsigned int countLines = 0;
    while(!file.eof()) {
        if(file.get() == '\n') countLines++;
    }
    countLines++;
    std::cout << "Count of lines: " << countLines << std::endl;
    file.clear();
    file.seekg(0);
    return 0;
}
