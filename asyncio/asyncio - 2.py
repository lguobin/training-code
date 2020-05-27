# !/usr/bin/env python
# -*- coding: utf-8 -*-
# @Date     : 2020/05/11 18:12:02
# @File     : asyncio - 2.py
# @Author   : BenLam(864551538@qq.com)
# @Link     : https://www.cnblogs.com/BenLam/
# @Version  : 1.0

import os
import time
import requests
import asyncio


async def test2(i):
    r = await other_test(i)
    print(i,r)

async def other_test(i):
    r = requests.get(i)
    await asyncio.sleep(4)
    print(f"耗时 - {time.time()-start}")
    return r

url = ["https://segmentfault.com/p/1210000013564725",
       "https://www.jianshu.com/",
       "https://www.cnblogs.com/",
       "https://www.jd.com/",
       "https://www.baidu.com/"]

loop = asyncio.get_event_loop()
tasks = [asyncio.ensure_future(test2(i)) for i in url]

start = time.time()
loop.run_until_complete(asyncio.wait(tasks))
print(f"总耗时 - {time.time()-start}")
loop.close()


# a = requests.get("https://www.jd.com/")
# print(a.headers)









if __name__ == '__main__':
    pass
