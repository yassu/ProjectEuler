def fact(n):
    if n == 0:
        return 1
    else:
        return n * fact(n-1)

f = fact(100)
print(sum([int(n) for n in str(f)]))
