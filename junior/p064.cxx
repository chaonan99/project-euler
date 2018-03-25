// 2018-03-24
#include <iostream>
#include <cmath>
#include <list>
#include <algorithm>
#include <iterator>
#include <time.h>

using namespace std;

int get_period(int n)
{
  // Take n = 23 for example
  float sqrt_n = sqrt(n);  // sqrt(23)
  int sqrt_n_int = (int) sqrt_n;  // 4
  if (sqrt_n_int*sqrt_n_int - n != 0)
  {
    int a, b, c;
    a = n - sqrt_n_int*sqrt_n_int;  // 7
    b = (sqrt_n + sqrt_n_int) / a;  // 1
    c = a*b - sqrt_n_int;           // 3
    list<int> hash_list;
    int hash_key = a*n*n+b*n+c;
    list<int>::iterator res;
    do {
      hash_list.push_back(hash_key);
      b = (int) (a / (sqrt_n-c));
      a = (n-c*c)/a;
      c = a*b-c;
      hash_key = a*n*n+b*n+c;
      res = find(hash_list.begin(), hash_list.end(), hash_key);
    } while (res == end(hash_list));
    return hash_list.size() - distance(hash_list.begin(), res);
  }
  return 0;
}

void work()
{
  int counter = 0;
  for (int i = 2; i <= 10000; ++i)
    if (get_period(i) % 2 == 1)
      counter++;
  cout << "Number of odd periods: " << counter << endl;
}

int main(int argc, char const *argv[])
{
  clock_t t1,t2;
  t1=clock();
  work();
  t2=clock();
  float diff ((float)t2-(float)t1);
  cout << "Running time (s): " << diff/CLOCKS_PER_SEC << endl;
  return 0;
}