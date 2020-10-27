# !/usr/bin/env python
# -*- coding: utf-8 -*-
# @Date     : 2020/05/12 09:40:30
# @File     : yield - 用法.py
# @Author   : BenLam(864551538@qq.com)
# @Link     : https://www.cnblogs.com/BenLam/
# @Version  : 1.0

import os


def demo1():
    # 原文代码地址: https://www.jianshu.com/p/2dfaacdd0a90
    def consumer():       # 有yield的函数就是生成器，没的跑
        r = "启动......"
        print(r)   # 😄发送None时，函数从头开始执行的，到 yield r 停止，此后的send(xxx)都是从 n = yield 开始。记住，n = yield 是启动点， yield r 暂停点，并返回yield r结果给produce函数
        while True:
            n = yield r        
    # 注意，yield r 是代码终止点，n = yield是启动点，一个正常的循环♻️过程是从 n = yield开始执行，到下面，执行到r ='200k'后，再回到 yield r处暂停，此时暂停的yield r 应该是经过新的循环，这里没有 for in 函数，但是，r最新的200k就是它的新循环，所以此时yield r为200k时生成器程序也就是consumer停止，但是，新的yield r 要回给send，send发送消息，也会要求得到消息的结果
            if not n:
                return
            print('- [CONSUMER] Consuming %s...' % n)
            # r = '200 OK'
            r = "运行次数" + str(n)

    def produce(c):
        c.send(None)    # yield 特性，未启动前必须传 send(None)或 next(function()) / function.__next__
        n = 0     # 没有上面的c.send，系统报错can't send non-None value to a just-started generator
        while n < 5:   #  
            n = n + 1
            print('[PRODUCER] Producing %s...' % n)
            r = c.send(n)
            print('[PRODUCER] Consumer return: %s' % r)
        c.close()
    c = consumer()
    produce(c)


def demo2():
    def _yield():
        r = "启动"
        while True:
            x = yield r
            if not x:
                return
            print(f"打印 - {x}")
            r = "运行次数" + str(x)

    def test(c):
        c.send(None)
        n = 0
        for _ in range(5):
            n += 1
            c.send(n)
        c.close()

    _ = _yield()
    test(_)


def deco(func):
    # yield 装饰器, 解决第一步不需要传 None 或 next
    def wrapper():
        res = func()
        next(res)
        return res
    return wrapper

@deco
def foo():
    food_list = []
    while True:
        food = yield food_list  #返回添加food的列表
        food_list.append(food)
        print("列表现有元素:",food_list)
g = foo()
print(g.send('苹果'))
print(g.send('香蕉'))
print(g.send('菠萝'))

if __name__ == '__main__':
    # demo1()
    # demo2()

    pass