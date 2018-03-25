# 2018-03-18
counter = 0
for i in range(1, 10):
    for j in range(1, 50):
        if len(str(i**j)) < j:
            break
#         print("{}^{}={}".format(i,j,i**j))
        counter += 1
print(counter)