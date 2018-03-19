#include <vector>
#include <iostream>
using namespace std;

int main(int argc, char const *argv[])
{
    vector<vector<string>> sarr;
    sarr.push_back(vector<string>());
    sarr.push_back(vector<string>());
    sarr.push_back(vector<string>());
    for (int i = 0; i < 3; ++i)
    {
        sarr[i].push_back("a");
    }
    for (int i = 0; i < 3; ++i)
    {
        sarr[i].push_back("b");
    }
    for (int i = 0; i < 3; ++i)
    {
        for (int j = 0; j < 2; ++j)
            cout << sarr[i][j] << ", ";
        cout << endl;
    }
    return 0;
}