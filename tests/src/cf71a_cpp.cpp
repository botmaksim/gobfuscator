#include <iostream>
#include <string>

using namespace std;

int main() {
    int n;
    if (!(cin >> n)) return 0;
    while (n--) {
        string s;
        cin >> s;
        if (s.length() > 10) {
            cout << s[0] << s.length() - 2 << s[s.length() - 1] << endl;
        } else {
            cout << s << endl;
        }
    }
    return 0;
}
