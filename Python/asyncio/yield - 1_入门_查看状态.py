# !/usr/bin/env python
# -*- coding: utf-8 -*-
# @Date     : 2020/05/20 10:32:35
# @File     : yield - 0_入门_查看状态.py
# @Link     : https://www.cnblogs.com/BenLam/
# @Version  : 1.0

import os
import inspect
import requests

def demo1():
    def test():
        yield 1
        return "OK"

    t = test()
    print(inspect.getgeneratorstate(t))
    next(t)
    print(inspect.getgeneratorstate(t))
    try:
        next(t)
    except StopIteration:
        print(inspect.getgeneratorstate(t))

    def test2():
        # 如果你想实现一种新的迭代模式，使用一个生成器函数来定义它。
        # 下面是一个生产某个范围内浮点数的生成器：
        def frange(start, stop, increment):
            print(" 自定义迭代器 ".center(30, "-"))
            x = start
            while x < stop:
                # yield x
                yield round(x, 3)
                x += increment

        print([n for n in frange(0, 4, 0.3)])
        print(list(frange(0, 1, 0.125)))
    test2()

    class Countdown:
        # a = reversed(range(10))
        # print(next(a))
        # print(next(a))
        # print(list(a))
        print(" 反向迭代 ".center(30, "-"))
        def __init__(self, start):
            self.start = start

        # Forward iterator
        def __iter__(self):
            n = self.start
            while n > 0:
                yield n
                n -= 1

        # Reverse iterator
        def __reversed__(self):
            n = 1
            while n <= self.start:
                yield n
                n += 1

    print([rr for rr in reversed(Countdown(30))])
    print([rr for rr in Countdown(30)])


    """
    打印:
        GEN_CREATED
        GEN_SUSPENDED
        GEN_CLOSED
        ----------- 自定义迭代器 -----------
        [0, 0.3, 0.6, 0.9, 1.2, 1.5, 1.8, 2.1, 2.4, 2.7, 3.0, 3.3, 3.6, 3.9]
        ----------- 自定义迭代器 -----------
        [0, 0.125, 0.25, 0.375, 0.5, 0.625, 0.75, 0.875]
        ------------ 反向迭代 ------------
        [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30]
        [30, 29, 28, 27, 26, 25, 24, 23, 22, 21, 20, 19, 18, 17, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1]
    """


if __name__ == '__main__':
    demo1()