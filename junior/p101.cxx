// 2018-03-24
#include <iostream>
#include <time.h>

using namespace std;

int func10(int n)
{
  int it = 1;
  int sum = 1;
  for (int i = 0; i < 10; i++)
  {
    it *= (-n);
    sum += it;
  }
  return sum;
}

void work()
{
  for (int i = 0; i < 20; i++)
  {
    cout << func10(i) << endl;
  }
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