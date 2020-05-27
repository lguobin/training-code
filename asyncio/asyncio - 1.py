# !/usr/bin/env python
# -*- coding: utf-8 -*-
# @Date     : 2020/05/11 17:52:53
# @File     : asyncio - 1.py
# @Author   : BenLam(864551538@qq.com)
# @Link     : https://www.cnblogs.com/BenLam/
# @Version  : 1.0

import asyncio
from time import sleep

async def washing1():
    await asyncio.sleep(3)  # 使用 asyncio.sleep(), 它返回的是一个可等待的对象
    print(" 1 ".center(20, "-"))

async def washing2():
    await asyncio.sleep(4)
    print(" 2 ".center(20, "-"))

# 1. 创建一个事件循环
loop = asyncio.get_event_loop()

# 2. 将异步函数加入事件队列
tasks = [
    washing1(),
    washing2(),
]

# 3. 执行事件队列, 直到最晚的一个事件被处理完毕后结束
loop.run_until_complete(asyncio.wait(tasks))
loop.close()




if __name__ == '__main__':
    pass
