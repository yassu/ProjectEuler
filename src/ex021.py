def d(n):
    s = 0
    for j in range(1, n):
        if n % j == 0:
            s += j
    return s

def is_amicable_number(n):
    n1 = n
    n2 = d(n)
    return n1 != n2 and n1 == d(n2)

s = 0
for n in range(1, 10**4):
    print(n)
    if(is_amicable_number(n)):
        s += n

print(s)
# => 31626
