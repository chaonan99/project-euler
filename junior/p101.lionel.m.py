def predict(u):
    if len(u)==0: return 0
    return u[-1]+predict([u[i+1]-u[i] for i in range(len(u)-1)])

u = [sum((-n)**k for k in range(11)) for n in range(1,12)]
print(sum(predict(u[:n]) for n in range(1,len(u))))
