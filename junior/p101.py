# -*- coding: utf-8 -*-
#!/usr/bin/env python
# 2018/08/09


import numpy as np
from sympy import Matrix, linsolve, symbols


__author__ = ['chaonan99']


def func_10(n):
  s = 1 - n
  nn = -n
  for _ in range(9):
    nn *= (-n)
    s += nn
  return s

def get_coeff_b_mat(n, f):
  coeff_mat = np.vander(range(1, n+3), n)[:, ::-1]
  b_mat = np.array([f(i) for i in range(1, n+2)])
  return coeff_mat, b_mat


def main():
  p = 10
  x, y, z = symbols("x, y, z")
  res = 0
  coeff_mat, b_mat = get_coeff_b_mat(p, func_10)
  solution = None

  for i in range(p):
    for j in range(i, p+1):
      A = Matrix(coeff_mat[:j+1, :i+1])
      b = Matrix(b_mat[:j+1])
      prev_solution = solution
      solution = linsolve((A, b), [x, y, z])
      if len(solution) == 0:
        res += sum(np.array(next(iter(prev_solution))) * coeff_mat[j, :i+1])
        break

  print(res)


if __name__ == '__main__':
  main()
