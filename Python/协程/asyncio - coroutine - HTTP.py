# !/usr/bin/env python
# -*- coding: utf-8 -*-
# @Date     : 2020/05/23 10:04:28
# @File     : asyncio - coroutine - HTTP.py
# @Link     : https://www.cnblogs.com/BenLam/
# @Version  : 1.0

import os
import asyncio
import functools

async def test_coroutine(msg="这是协程启动"):
    print(msg)
    return "返回值"

print("---- 分割线 ----")

async def main():
    print("主协程")
    res_A = await result_A()
    print("1号执行")
    res_B = await result_B()
    print("2号执行")
    return res_A, res_B

async def result_A():
    print("协程 1 号")
    return 1

async def result_B():
    print("协程 2 号")
    return 2

print("---- 分割线 ----")

def call_back(args, *, kwargs="defalut"):
    print(f"回调函数获取传参: {args} - {kwargs}")

async def main_b(loop):
    print("注册回调函数", loop)
    loop.call_soon(call_back, "call_soon")
    wrappen = functools.partial(call_back, kwargs="not defalut")
    loop.call_soon(wrappen, "B")

    wrappen = functools.partial(call_back, kwargs="参数·2")
    loop.call_soon(wrappen, "参数1")
    
    loop.call_later(.1, call_back, "call_later")
    loop.call_at(.1, call_back, "call_at")
    await asyncio.sleep(1)

print("---- 分割线 ----")

def foo(future, result):
    print(f"future 状态: {future}")
    print(f"result 状态: {result}")
    future.set_result(result)
    print(f"future 状态: {future}")

def main_c():
    loop = asyncio.get_event_loop()
    try:
        all_done = asyncio.Future()
        loop.call_soon(foo, all_done, "Future is done!")
        result = loop.run_until_complete(all_done)
        print(result)
    finally:
        loop.close()
    print(all_done.result())



async def main_d(loop):
    print("将协程包装成任务")
    task = loop.create_task(test_coroutine())
    print("通过 cancel() 方法取消任务")
    task.cancel()
    try:
        await task
    except asyncio.CancelledError:
        print("取消任务报出异常")
    else:
        print(f"取消结果 - {task.result()}")


if __name__ == '__main__':
    # loop = asyncio.get_event_loop()
    # try:
    #     # a = loop.run_until_complete(test_coroutine())
    #     # a = loop.run_until_complete(main())
    #     a = loop.run_until_complete(main_b(loop))
    # finally:
    #     loop.close()

    # main_c()
    loop = asyncio.get_event_loop()
    try:
        loop.run_until_complete(main_d(loop))
    finally:
        loop.close()