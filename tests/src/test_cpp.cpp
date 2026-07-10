#include <iostream>

int main() {
    const char* secretKey = "SuperSecretKey123";
    int myVariable = 42;

    std::cout << "My logic runs: " << secretKey << std::endl;

    if (myVariable == 42) {
        std::cout << "Logic is intact!" << std::endl;
    }

    return 0;
}
