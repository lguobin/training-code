# !/usr/bin/env python
# -*- coding: utf-8 -*-
# @Date     : 2020/05/21 10:29:56
# @File     : asyncio - 1_入门.py
# @Link     : https://www.cnblogs.com/BenLam/
# @Version  : 1.0

import os
import asyncio
import itertools
import sys
import time
import datetime

def demo1():
    @asyncio.coroutine
    def spin():
        # for i in itertools.cycle('|/-\\'):
        #     write, flush = sys.stdout.write, sys.stdout.flush
        #     write(i)
        #     flush()
        #     write('\x08'*len(i))
        write, flush = sys.stdout.write, sys.stdout.flush
        for x in itertools.cycle("|/-\\"):  # itertools.cycle 函数从指定的序列中反复不断地生成元素
            status = x + ' loading'
            write(status)
            flush()
            write('\x08' * len(status))  # 使用退格符把光标移回行首
            time.sleep(0.2)  # 每 0.2 秒刷新一次
            try:
                yield from asyncio.sleep(1)
            except asyncio.CancelledError:
                break

    @asyncio.coroutine
    def slow_f():
        yield from asyncio.sleep(8)
        return 3

    @asyncio.coroutine
    def sup():
        spiner = asyncio.async(spin())
        print("spiner:",spiner)
        r = yield from slow_f()
        spiner.cancel()
        return r

    def main():
        loop = asyncio.get_event_loop()
        r = loop.run_until_complete(sup())
        loop.close()
        print("r:",r)

    main()
    # 输出结果：
    # spiner: <Task pending coro=<spin() running at c:/Users/DELL/Desktop/ssj/search/descrip.py:7>>
    # r: 3    # 运行期间会有动画指针

def demo2():
    async def display_date(loop):
        end_time = loop.time()
        for _ in range(10):
            print(datetime.datetime.now())
            await asyncio.sleep(1)

    loop = asyncio.get_event_loop()
    loop.run_until_complete(display_date(loop))
    print("stop")
    loop.close()

    # 2020-05-21 11:00:06.163600
    # 2020-05-21 11:00:07.163600
    # 2020-05-21 11:00:08.163600
    # 2020-05-21 11:00:09.163600
    # 2020-05-21 11:00:10.163600
    # stop

def demo3():
    def hello_world(loop):
        for _ in range(3):
            print('Hello World')
        loop.stop()

    loop = asyncio.get_event_loop()
    loop.call_soon(hello_world, loop)
    loop.run_forever()
    loop.close()


def demo4():
    async def test(sleep_time):
        print("test")
        await asyncio.sleep(sleep_time)
        print("Hello World - {}".format(sleep_time))

    task = [test(a) for a in range(3)]
    loop = asyncio.get_event_loop()
    try:
        loop.run_until_complete(asyncio.wait(task))
    except KeyboardInterrupt as e:
        all_tasks = asyncio.Task.all_tasks()
        for task in all_tasks:
            print("cancel task")
            print(task.cancel())
        loop.stop()
        loop.run_forever()
    finally:
        loop.close()


def demo5():
    print("----- lgb -------")
    def test(sleep_time):
        print("Hello World - {}".format(sleep_time))
        # print("Hello World")

    def stop(loop):
        loop.stop()

    loop = asyncio.get_event_loop()


    # 立即执行
    loop.call_soon(test, 22)
    loop.call_soon(test, 33)
    loop.call_soon(test, 11)

    # 延迟执行
    loop.call_later(3, test, 221)
    loop.call_later(2, test, 122)
    loop.call_later(1, test, 131)

    now = loop.time()
    loop.call_at(now+4, test, 221)
    loop.call_at(now+5, test, 122)
    loop.call_at(now+7, test, 131)


    # break out
    # loop.call_soon(stop, loop)
    loop.run_forever()



if __name__ == '__main__':
    # demo1()
    # demo2()
    # demo3()
    # demo4()
    demo5()