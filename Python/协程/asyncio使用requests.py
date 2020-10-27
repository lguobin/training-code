# !/usr/bin/env python
# -*- coding: utf-8 -*-
# @Date     : 2020/05/22 14:21:36
# @File     : asyncio使用requests.py
# @Link     : https://www.cnblogs.com/BenLam/
# @Version  : 1.0

import os
import asyncio
import multiprocessing
import random
import time
import requests
from concurrent.futures import ProcessPoolExecutor
from threading import Event


def worker(item):
    file_name = str(int(time.time()*1000))+"-"+str(item)+".html"
    print(file_name)
    with open(file_name, "wb") as f:
        res = requests.get("http://www.jianshu.com/p/968be2f00119")
        print(res.status_code)
        time.sleep(30)
        f.write(res.content)


class MySpider(object):
    def __init__(self):
        self.queue = asyncio.Queue(maxsize=10)  # 任务队列
        self.pool = ProcessPoolExecutor(max_workers=5)  # 进程池
        self.event = Event()

    @asyncio.coroutine
    def producer(self):
        # 生产控制者
        for i in range(10):
            yield from self.producer_001()
        self.event.set()

    @asyncio.coroutine
    def producer_001(self):
        # 其中一个生产者
        yield from asyncio.sleep(1)
        yield from self.producer_002()

    @asyncio.coroutine
    def producer_002(self):
        # 其中一个生产者
        yield from asyncio.sleep(1)
        yield from self.producer_003()

    @asyncio.coroutine
    def producer_003(self):
        # 最后一个生产者，负责把生产的东西加到队列里
        item = random.randint(1, 100)
        print("put Item", item)
        yield from self.queue.put(item)

    @asyncio.coroutine
    def customer(self, future):
        # 消费者
        while True:
            if self.event.is_set() and self.queue.empty():
                break
            item = yield from self.queue.get()

            self.pool.submit(worker, item)
        future.set_result("customer done")

    @asyncio.coroutine
    def start(self, future):

        asyncio.ensure_future(self.producer())
        asyncio.ensure_future(self.customer(future))

        # asyncio.ensure_future(self.producter())
        # asyncio.ensure_future(self.custom(future))


if __name__ == "__main__":
    multiprocessing.freeze_support()
    loop = asyncio.get_event_loop()

    future = loop.create_future()
    loop.create_task(MySpider().start(future))
    loop.run_until_complete(future)
    print(future.result())
    loop.close()
