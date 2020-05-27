# !/usr/bin/env python
# -*- coding: utf-8 -*-
# @Date     : 2020/05/11 15:55:47
# @File     : asyncio-1.py
# @Author   : BenLam(864551538@qq.com)
# @Link     : https://www.cnblogs.com/BenLam/
# @Version  : 1.0


import asyncio
import itertools
import sys


@asyncio.coroutine # 打算交给asyncio 处理的协程要使用 @asyncio.coroutine 装饰
def spin(msg):
    write, flush = sys.stdout.write, sys.stdout.flush
    for char in itertools.cycle('|/-\\'):  # itertools.cycle 函数从指定的序列中反复不断地生成元素
        status = char + ' ' + msg
        write(status)
        flush()
        write('\x08' * len(status))  # 使用退格符把光标移回行首
        try:
            yield from asyncio.sleep(0.1)  # 使用 yield from asyncio.sleep(0.1) 代替 time.sleep(.1), 这样的休眠不会阻塞事件循环
        except asyncio.CancelledError:  # 如果 spin 函数苏醒后抛出 asyncio.CancelledError 异常，其原因是发出了取消请求
            break

    write(' ' * len(status) + '\x08' * len(status))  # 使用空格清除状态消息，把光标移回开头


@asyncio.coroutine
def slow_function():  # 5 现在此函数是协程，使用休眠假装进行I/O 操作时，使用 yield from 继续执行事件循环
    # 假装等待I/O一段时间
    yield from asyncio.sleep(3)  # 此表达式把控制权交给主循环，在休眠结束后回复这个协程
    return 42


@asyncio.coroutine
def supervisor():  #这个函数也是协程，因此可以使用 yield from 驱动 slow_function
    spinner = asyncio.async(spin('thinking!'))  # asyncio.async() 函数排定协程的运行时间，使用一个 Task 对象包装spin 协程，并立即返回
    print('spinner object:', spinner)  # Task 对象，输出类似 spinner object: <Task pending coro=<spin() running at spinner_asyncio.py:6>>
    # 驱动slow_function() 函数，结束后，获取返回值。同事事件循环继续运行，
    # 因为slow_function 函数最后使用yield from asyncio.sleep(3) 表达式把控制权交给主循环
    result = yield from slow_function()
    # Task 对象可以取消；取消后会在协程当前暂停的yield处抛出 asyncio.CancelledError 异常
    # 协程可以捕获这个异常，也可以延迟取消，甚至拒绝取消
    spinner.cancel()

    return result

def main():
    loop = asyncio.get_event_loop()  # 获取事件循环引用
    # 驱动supervisor 协程，让它运行完毕；这个协程的返回值是这次调用的返回值
    result = loop.run_until_complete(supervisor())
    loop.close()
    print('Answer', result)


if __name__ == '__main__':
    main()