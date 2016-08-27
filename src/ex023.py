N = 28123

def get_sum_of_divisable_numbers(n):
    s = 0
    for j in range(1, n):
        if n % j == 0:
            s += j
    return s

def is_abundant_number(n):
    return get_sum_of_divisable_numbers(n) > n

def  printer(n):
    print(n)
    return n

# 過剰数のlist
abundant_numbers = [printer(n) for n in range(N + 1) if is_abundant_number(n)]

# 過剰数の和として得られる数のlist
s = set()
for n1 in abundant_numbers:
    print('n1 = {}'.format(n1))
    for n2 in abundant_numbers:
        if n1 + n2 > N:
            continue
        s.add(n1 + n2)


print(N*(N+1)//2 - sum(list(s)))
