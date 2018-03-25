# 2018-03-23
from bisect import bisect_left

cubes = []
cube_tuple = []
for i in range(300, 9000):
    on = int(''.join(sorted(str(i**3),reverse=True)))
    ind = bisect_right(cubes, on)
    cubes.insert(ind, on)
    cube_tuple.insert(ind, (on, i))
    if ind - 4 >= 0 and cubes[ind-4] == on:
        print(cube_tuple[ind-4][1]**3)
        break