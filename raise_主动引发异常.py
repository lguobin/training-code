# !/usr/bin/env python
# -*- coding: utf-8 -*-
# @Date     : 2020/05/20 10:21:59
# @File     : raise_主动引发异常.py
# @Link     : https://www.cnblogs.com/BenLam/
# @Version  : 1.0

import os
import random

class MyError(Exception):
    pass

def add(x, y):
    print(x, y)
    if x < 0 or y < 0:
        raise MyError('自定义的异常')
    r = x + y
    return r

# print(add(-123, random.randint(1, 100000)))

def test():
    raise MyError("主动引发异常")
    print("1")
    print("2")

test()








if __name__ == '__main__':
    pass
