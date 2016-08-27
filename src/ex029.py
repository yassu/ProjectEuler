arange = range(2, 100 + 1)
brange = range(2, 100 + 1)

numbers = [a**b for a in arange for b in brange]
print(len(set(numbers)))
