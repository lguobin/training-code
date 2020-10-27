# !/usr/bin/env python
# -*- coding: utf-8 -*-
# @Date     : 2020/05/22 17:41:06
# @File     : asyncio - 3.py
# @Link     : https://www.cnblogs.com/BenLam/
# @Version  : 1.0

import os
import time
import asyncio


async def taskIO_1():
    print('开始运行IO任务1...')
    await asyncio.sleep(3)  # 假设该任务耗时2s
    print('IO任务1已完成，耗时3s')
    return taskIO_1.__name__


async def taskIO_2():
    print('开始运行IO任务2...')
    await asyncio.sleep(3)  # 假设该任务耗时3s
    print('IO任务2已完成，耗时3s')
    return taskIO_2.__name__


"""
async def main():  # 调用方
    # 此处并发运行传入的aws(awaitable objects)，同时通过await返回一个包含(done, pending)的元组，
    # done表示已完成的任务列表，pending表示未完成的任务列表。
    # 注：
    # ①只有当给wait()传入timeout参数时才有可能产生pending列表。
    # ②通过wait()返回的结果集是按照事件循环中的任务完成顺序排列的，所以其往往和原始任务顺序不同。
    # gather()通过await直接返回一个结果集列表，我们可以清晰的从执行结果看出来，
    # 虽然任务2是先完成的，但最后返回的结果集的顺序是按照初始传入的任务顺序排的
    tasks = [taskIO_1(), taskIO_2()]  # 把所有任务添加到task中
    done, pending = await asyncio.wait(tasks)  # 子生成器
    for r in done:  # done和pending都是一个任务，所以返回结果需要逐个调用result()
        print('协程无序返回值：'+r.result())
"""

"""
async def main(): # 调用方
    # 只关心协程并发运行后的结果集合，可以使用gather()，
    # 它不仅通过await返回仅一个结果集，
    # 而且这个结果集的结果顺序是传入任务的原始顺序。
    resualts = await asyncio.gather(taskIO_1(), taskIO_2()) # 子生成器
    print(resualts)
"""


async def main(): # 调用方
    tasks = [taskIO_1(), taskIO_2()]  # 把所有任务添加到task中
    for completed_task in asyncio.as_completed(tasks):
        resualt = await completed_task # 子生成器
        print('协程无序返回值：'+resualt)

if __name__ == '__main__':

    start = time.time()
    loop = asyncio.get_event_loop()  # 创建一个事件循环对象loop
    try:
        loop.run_until_complete(main())  # 完成事件循环，直到最后一个任务结束
    finally:
        loop.close()  # 结束事件循环
    print(f'所有IO任务总耗时{round(time.time()-start, 3)}秒')
