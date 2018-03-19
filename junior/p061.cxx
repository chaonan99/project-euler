#include <iostream>
#include <functional>
#include <list>
#include <algorithm>
#include <vector>
#include <tuple>
#include <string>
#include <time.h>

using namespace std;

int process(const vector<list<string>*> &plist)
{
    list<vector<string>*> tlistold;
    vector<string> v{ plist[0]->begin(), plist[0]->end() };
    tlistold.push_back(&v);

    for (int i = 1; i < 6; ++i)
    {
        list<vector<string>*> tlistnew;
        for (int j = 0; j < i+1; ++j)
            tlistnew.push_back(new vector<string>);
        for (auto s1 : *plist[i])
        {
            vector<string>* s2list = tlistold.front();
            for (int ind = 0; ind < s2list->size(); ++ind)
            {
                string s2 = (*s2list)[ind];
                if (s1.compare(2, 2, s2.substr(0,2)) == 0)
                {
                    auto it = tlistnew.begin();
                    (*it)->push_back(s1);
                    for (auto t : tlistold)
                    {
                        ++it;
                        (*it)->push_back((*t)[ind]);
                    }
                }
            }
        }
        tlistold = tlistnew;
        // cout << tlistnew.front()->size() << endl;
        if (tlistnew.front()->size() == 0 && i != 5)
        {
            // Not found
            return 0;
        }
    }

    for (int i = 0; i < tlistold.front()->size(); ++i)
    {
        string s1 = (*tlistold.back())[i];
        string s2 = (*tlistold.front())[i];
        if (s1.compare(2, 2, s2.substr(0,2)) == 0)
        {
            cout << "Find" << endl;
            int sum = 0;
            for (auto t : tlistold)
            {
                cout << (*t)[i] << ", ";
                sum += stoi((*t)[i]);
            }
            cout << endl;
            cout << "sum = " << sum << endl;
            return sum;
        }
    }
    return 0;
}

int main(int argc, char const *argv[])
{
    clock_t t1,t2;
    t1=clock();
    vector<list<string>*> polynumentry;
    vector<function<int(int)>> polynumfunc;
    polynumentry.push_back(new list<string>);
    polynumentry.push_back(new list<string>);
    polynumentry.push_back(new list<string>);
    polynumentry.push_back(new list<string>);
    polynumentry.push_back(new list<string>);
    polynumentry.push_back(new list<string>);
    polynumfunc.push_back([](int n){return n*(n+1)/2; });
    polynumfunc.push_back([](int n){return n*n; });
    polynumfunc.push_back([](int n){return n*(3*n-1)/2; });
    polynumfunc.push_back([](int n){return n*(2*n-1); });
    polynumfunc.push_back([](int n){return n*(5*n-3)/2; });
    polynumfunc.push_back([](int n){return n*(3*n-2); });

    for (int n = 0; n < 141; ++n)
    {
        for (int i = 0; i < 6; ++i)
        {
            string polynum = to_string(polynumfunc[i](n));
            if (polynum.length() == 4)
                polynumentry[i]->push_back(polynum);
        }
    }

    list<string>* octs = polynumentry.back();
    polynumentry.pop_back();
    sort(polynumentry.begin(), polynumentry.end());
    vector<list<string>*> pne;
    pne.reserve(6);
    pne[0] = octs;

    process(polynumentry);
    do {
        int i = 0;
        for (auto it = polynumentry.begin(); it != polynumentry.end(); ++it)
            pne[++i] = *it;
        if (process(pne) > 0)
            break;
    } while (next_permutation(polynumentry.begin(), polynumentry.end()));

    t2=clock();
    float diff ((float)t2-(float)t1);
    cout<<diff/CLOCKS_PER_SEC<<endl;

    return 0;
}