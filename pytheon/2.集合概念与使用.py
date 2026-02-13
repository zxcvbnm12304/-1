class1 = ['tom','nimi']
class2 = ['kifr','gkoi','nimi']

for name1 in class1:
    for name2 in class2:
        if name1 == name2:
            print(name1)

a={1,2,3,4}
b={3,4,5,6}
# 交集
print(a&b)
print(a.intersection(b))

print(b&a)
print(b.intersection(a))
# 并集
print(a|b)
print(a.union(b))

print(b|a)
print(b.union(a))
# 差集
print(a-b)
print(a.difference(b))

print(b-a)
print(b.difference(a))