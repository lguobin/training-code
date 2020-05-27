# !/usr/bin/env python
# -*- coding: utf-8 -*-
# @Date     : 2020/05/19 10:09:42
# @File     : yield - 2_用法.py
# @Link     : https://www.cnblogs.com/BenLam/
# @Version  : 1.0

import os

a = [1, 2, 3, 4, 5, 6]
b = {
    "A": "-----",
    "B": "....."
}

def g1(iterable):
    print("返回传参对象")
    yield iterable

def g2(iterable):
    yield from iterable
    print("循环传参对象 - 结束")

for value in g1(range(10)):
    print(value)
for value in g2(range(10)):
    print(value)

# ----------------
# from itertools import chain
# for _ in chain(a, b, range(4)):
#     print(_)

def my_chain(*args, **kwargs):
    for my_iterable in args:
        yield from my_iterable
        for value in my_iterable:
            yield value

for _ in my_chain(a, b, range(4)):
    print(_)






if __name__ == '__main__':
    pass
