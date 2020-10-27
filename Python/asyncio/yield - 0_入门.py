# !/usr/bin/env python
# -*- coding: utf-8 -*-
# @Date     : 2020/05/19 10:39:31
# @File     : yield - 入门
# @Link     : https://www.cnblogs.com/BenLam/

import os
import time
import asyncio
from concurrent.futures import ThreadPoolExecutor

# demo8 - 使用
# demo9 - 使用
import aiohttp  # 需要下载
import requests

# 协程
# 从python yield的角度去理解协程比较难理解

def demo():
    y = (i for i in range(3))

    # y.next()
    # y.next()
    # 在python3.x版本中 y.next() 函数已经更名为 y.__next__()，
    # 所以只需要将 y.next() 换成 y.__next__() , 另一种写法就是使用 next(y)

    print(y.__next__())
    print(next(y))

    def gen():
        # 代码来源: https://github.com/dabeaz/python-cookbook/blob/master/src/4/how_to_flatten_a_nested_sequence/example.py
        yield from 'AB'
        yield from range(1, 3)
    print(list(gen()))
    """
    打印:
        ['A', 'B', 1, 2]
    """

def demo1():
    # 生产者与消费者例子
    def consumer():
        r = ''
        while True:
            n = yield r
            if not n:
                return
            # print('[CONSUMER] Consuming %s...' % n)
            print('[-消费者] Consuming %s...' % n)
            time.sleep(1)
            r = '200 OK'

    def produce(c):
        next(c)
        n = 0
        while n < 3:
            n = n + 1
            print('[ 生产者 ] Producing %s...' % n)
            r = c.send(n)  # 通知迭代器返回下一个
            print('[ 生产者 ] Consumer return: %s' % r)
        c.close()

    c = consumer()
    produce(c)
    """
    注意到consumer函数是一个generator（生成器），把一个consumer传入produce后：

    1. 首先调用 c.__next()__ 启动生成器；
    2. 然后，一旦生产了东西，通过 c.send(n) 切换到 consumer 执行；
    3. consumer 通过 yield 拿到消息，处理，又通过 yield 把结果传回；
    4. produce 拿到consumer处理的结果，继续生产下一条消息；
    5. produce 决定不生产了，通过 c.close() 关闭 consumer，整个过程结束。

    整个流程无锁，由一个线程执行，produce和consumer协作完成任务，所以称为“协程”，而非线程的抢占式多任务。
    所以，通过使用yield生成的迭代器，是实现协程的一种方式。

    而在go语言里面，goroutine是协程的另外一种方式。
    所以使用yield实现协程比较难理解。如果将goroutine比喻成已经成品的车，那么yield就相当于提供了车的零件，还需要自己去拼装实现。

    作者：d咚咚呛
    链接：https://www.jianshu.com/p/7d20b5197883
    来源：简书
    """


def demo2():
    # 计算历史数据的平均值
    def avg():
        count, total, avg_num = 0, 0, 0
        while True:
            t = yield
            total += t
            count += 1
            avg_num = round(total / count, 2)
            print(avg_num)
    test = avg()
    # Test 现在是 None
    next(test)
    # list = [10,30,60]
    test.send(10)
    test.send(30)
    test.send(60)

    """
    # 创建协程对象
    test = avg()

    # 启动协程, avg_arg函数开始运行, 进入while循环, 在yield出暂停
    next(test)

    # 向协程发送数据, 直接到t = yield代码的暂停出, yield左边的变量用于接收调用方send过来的数据,
    yield右边暂时没有变量, 表示协程不生成数据给调用方。

    test.send(10)
    test.send(30)
    test.send(60)

    打印:
        10.0
        20.0
        33.33

    作者：lpj24
    链接：https://www.jianshu.com/u/709bf3373f34
    来源：简书
    """


def demo3():
    import functools
    class coroutine(object):
        def __init__(self, fun):
            self.fun = fun
            functools.update_wrapper(self, fun)

        def __call__(self, *args, **kwargs):
            gen = self.fun(*args, **kwargs)
            next(gen)
            # gen.send(None)
            return gen

    @coroutine
    def avg():
        count, total, avg_num = 0, 0, 0
        while True:
            t = yield
            if t is None:
                break
            total += t
            count +=1
            avg_num = round(total / count, 2)
            print(avg_num)

    test = avg()
    li = [10, 30, 60]
    for _ in li:
        test.send(_)
    """
    1. 首先实现一个类包装器(http://www.jianshu.com/p/30dae1d3e72c), 来包装协程, 该装饰器主要用来启动协程, 不用每次手动next启动协程
    2. yield出多了一个avg_num变量, 这个就是用来生成数据返回给调用者， 所以每次调用者send变量之后, term计算完之后, 协程yield出avg_num给调用方, 然后在yield出继续暂停等待send数据处理
    3. 终止协程和异常处理。主要记住协程中未处理的异常会向上抛出。 如果发送send(None), break协程会向上抛出StopIteration, 协程此刻处于GEN_CLOSED状态不能再接受数据, 可以直接使用close()函数终止协程。

    打印:
        10.0
        20.0
        33.33

    作者：lpj24
    链接：https://www.jianshu.com/u/709bf3373f34
    来源：简书
    """

def demo4():
        def chain(*args):
            for i in args:
                yield from i    # 委托生成器yield的值来自i迭代器的每一个值
        print(list(chain(['A', 'B', 'C'], [1, 2, 3])))
        """
        打印:
            ['A', 'B', 'C', 1, 2, 3]
        """

def demo5():
    # BEGIN YIELD_FROM_AVERAGER
    from collections import namedtuple

    Result = namedtuple('Result', 'count average')

    # the subgenerator
    def averager():  # <1>
        total, count, average = 0.0, 0, None
        while True:
            term = yield  # <2>
            if term is None:  # <3>
                break
            total += term
            count += 1
            average = total/count
        return Result(count, average)  # <4>


    # the delegating generator
    def grouper(results, key):  # <5>
        while True:  # <6>
            results[key] = yield from averager()  # <7>

    # the client code, a.k.a. the caller
    def main(data):  # <8>
        results = {}
        for key, values in data.items():
            group = grouper(results, key)  # <9>
            next(group)  # <10>
            for value in values:
                group.send(value)  # <11>
            group.send(None)  # important! <12>

        # print(results)  # uncomment to debug
        report(results)

    # output report
    def report(results):
        for key, result in sorted(results.items()):
            group, unit = key.split(';')
            print('{:2} {:5} averaging {:.2f}{}'.format(
                result.count, group, result.average, unit))

    data = {
        'girls;kg':
            [40.9, 38.5, 44.3, 42.2, 45.2, 41.7, 44.5, 38.0, 40.6, 44.5],
        'girls;m':
            [1.6, 1.51, 1.4, 1.3, 1.41, 1.39, 1.33, 1.46, 1.45, 1.43],
        'boys;kg':
            [39.0, 40.8, 43.2, 40.8, 43.1, 38.6, 41.4, 40.6, 36.3],
        'boys;m':
            [1.38, 1.5, 1.32, 1.25, 1.37, 1.48, 1.25, 1.49, 1.46],
    }
    main(data)
    """
    代码说明
        1. 创建一个Result的自定义数据结构, 包含count和average变量, 作为返回值
        2. 定义委托生成器grouper, yield from avgerager()启动之后在此处暂停, results[key]的结果from于avgerager()子生成器的返回结果
        3. 定义子生成器avgerager(), 执行方式和第一章的一样, 当循环结束的时候, 返回结果给委托生成器, 赋值给results[key]
        4. 创建调用方main函数
            * group = grouper(results, key) 创建grouper委托生成器对象.
            * next(group) 激活委托生成器
            * 循环data数据中的列表, send(value)把数据发送给子生成器的term
            * 当列表中的数据发送完成之后, send(None), 终止子生成器, 子生成器终止之后, return计算出来的数据返回给委托生成器, 委托生成器中继续循环重复上述4步,
            * 完成列表中的数据之后print(results)
        
        最后yield from替代嵌套for循环产出值
        def chain(*args):
            for i in args:
                yield from i       #委托生成器yield的值来自i迭代器的每一个值
        if __name__ == '__main__':
        print(list(chain(['A', 'B', 'C'], [1, 2, 3])))

    打印:
         9 boys  averaging 40.42kg
         9 boys  averaging 1.39m
        10 girls averaging 42.04kg
        10 girls averaging 1.43m

    作者：lpj24
     Git：https://github.com/fluentpython/example-code/blob/master/16-coroutine/coroaverager3.py
    链接：https://www.jianshu.com/p/d6ff07a2edce
    来源：简书

    下面是书中的一个列子, 可以看看
    https://github.com/dabeaz/python-cookbook/blob/master/src/4/how_to_flatten_a_nested_sequence/example.py
    拆分包含列表的列表
    """

def demo6():
    print("多线程")
    def square_val(val):
        time.sleep(0.1)
        return val*val

    start_time = time.time()
    res = []
    with ThreadPoolExecutor() as executor:
        res = [executor.submit(square_val, i) for i in range(1000)]
    print(sum([f.result() for f in res]))
    print(f"总耗时 - {round(time.time() - start_time, 2)}")

    """
    代码说明
        1. 改代码的主要功能是实现把1-1000列表中每个数平方, 最后求和, square_val()模拟一个耗时请求, 休眠0.1秒
        2. 使用concurrent.futures多线程模块，使用默认线程数, 耗时0.9秒

    打印:
        332833500
        总耗时 - 5.0

    作者：lpj24
    链接：https://www.jianshu.com/p/a06ce6dc4d6b
    来源：简书
    """

def demo7():
    print("单线程协程异步io实现方式")
    @asyncio.coroutine
    def asyncio_square_val(val):
        yield from asyncio.sleep(0.1)
        return val*val
    
    start_time = time.time()
    loop = asyncio.get_event_loop()
    tasks = [asyncio_square_val(i) for i in range(1000)]
    res, _ = loop.run_until_complete(asyncio.wait(tasks))
    loop.close()
    print(sum([i.result() for i in res]))
    print(f"总耗时 - {round(time.time() - start_time, 2)}")

    """
    代码说明
        1. 使用装饰器@asyncio.coroutine装饰async_square_val, 将改协程函数交给asyncio处理,
        2. yield from asyncio.sleep(.1)这样休眠不会阻塞时间循环, 把控制权交给主循环, 继续执行事件循环中的其它任务
        3. 获取事件循环loop = asyncio.get_event_loop()
        4. tasks = [async_square_val(i) for i in test] 创建任务列表
        5. res, _ = loop.run_until_complete(asyncio.wait(tasks))将任务放入事件循环, 最后打印时间

    关于python yield大概介绍这么多, 只是说了大概, 更加细节的知识点需要去阅读流畅的python这本书。另外, 从python3.5开始使用了async和await替换了asyncio协程装饰器和yield from, 只需要替换就好了, 这里不多介绍。
    补充: 模拟非阻塞式休眠的时候我们使用asyncio.sleep()， 当正式使用比如http请求的时候我们要使用aiohttp非阻塞式的模块请求http
        https://github.com/aio-libs 这里提供了一些非阻塞式的第三方模块, 如mysql, redis等等

    打印:
        332833500
        总耗时 - 0.12

    作者：lpj24
    链接：https://www.jianshu.com/p/a06ce6dc4d6b
    来源：简书
    """

def demo8():
    print("读取当当网的图书")
    # 子生成器
    @asyncio.coroutine
    def get_image(img_url):
        yield from asyncio.sleep(1)
        # resp = yield from aiohttp.request('GET', img_url)
        resp = yield from aiohttp.ClientSession().get(img_url)
        image = yield from resp.read()
        return image

    def save_image(img, img_url):
        pwd = os.path.join(os.path.dirname(os.path.abspath(__file__)) + '.\\img_file', img_url.split('/')[-1])
        with open(pwd, 'wb') as f:
            f.write(img)

    @asyncio.coroutine
    def download_one(img_url):
        image = yield from get_image(img_url)
        save_image(image, img_url)

    images_list = [
        'http://img3m0.ddimg.cn/67/4/24003310-1_b_5.jpg'
        'http://img3m2.ddimg.cn/43/13/23958142-1_b_12.jpg',
        'http://img3m0.ddimg.cn/60/17/24042210-1_b_5.jpg',
        'http://img3m4.ddimg.cn/20/11/23473514-1_b_5.jpg',
        'http://img3m4.ddimg.cn/40/14/22783504-1_b_1.jpg',
        'http://img3m7.ddimg.cn/43/25/23254747-1_b_3.jpg',
        'http://img3m9.ddimg.cn/30/36/23368089-1_b_2.jpg',
        'http://img3m1.ddimg.cn/77/14/23259731-1_b_0.jpg',
        'http://img3m2.ddimg.cn/33/18/23321562-1_b_21.jpg',
        'http://img3m3.ddimg.cn/2/21/22628333-1_b_2.jpg',
        'http://img3m8.ddimg.cn/85/30/23961748-1_b_10.jpg',
        'http://img3m1.ddimg.cn/90/34/22880871-1_b_3.jpg',
        'http://img3m2.ddimg.cn/62/27/23964002-1_b_6.jpg',
        'http://img3m5.ddimg.cn/84/16/24188655-1_b_3.jpg',
        'http://img3m6.ddimg.cn/46/1/24144166-1_b_23081.jpg',
        'http://img3m9.ddimg.cn/79/8/8766529-1_b_0.jpg']

    start_time = time.time()
    loop = asyncio.get_event_loop()
    to_do_tasks = [download_one(img) for img in images_list]

    res, _ = loop.run_until_complete(asyncio.wait(to_do_tasks))
    # print(len(res))
    print(f"总耗时 - {round(time.time() - start_time, 2)}")

    """
    代码解读
        1. 对书本中的例子进行了改动, 创建一个images_list, 里面存储的是当当网一个页面所有书籍的图片地址, 我们要把这些图片下载下来存到本地， 大家可以根据以下代码获取图片地址
        2. 对于main函数的中的代码如果看过前面几张不会陌生, 同样的事件循环, 这里要提一下书中提到的期物, 协程, 任务, 粗略的理解, 协程asyncio.wait()把协程函数转换成任务添加到事件循环, 
            任务完成之后返回future对象, future提供了获取对象信息的api， 大致这样吧(线程的高度抽象接口concurrent.futures也是返回的future对象)。
        3. 创建子生成器, get_image()
            * sleep()一下是因为, 图片太小, 速度太快, 这里模拟一下耗时的请求
            * aiohttp.request()请求会有Unclosed client session的警告, 参见https://github.com/aio-libs/aiohttp/issues/2036, 建议使用
            * with aiohttp.ClientSession上下文的方式自动关闭, 这里不影响结果我们就按书上的来做
            * yield from resp.read() 非阻塞式读取图片二进制数据返回给get_image函数的委派生成器
        4. sav_image保存图片到本地
        5. 创建download_one协程用于下载每一张图片, 然后存储到本地
        6. 终端返回执行结果耗时1.12s左右, 图片成功存储

        没有对比就没有伤害, 大家应该使用concurrent.futures多线程下载对比一下, 这里就不写代码了, 我这边测试多线程要慢一些

    输出:
        文件保存在当前目录 + //img_file(几张图片)
        总耗时 - 1.14

    resp = yield from aiohttp.request('GET', img_url) - 会引发异常
    使用 - resp = yield from aiohttp.ClientSession().get(img_url) - 代替

    作者：lpj24
    链接：https://www.jianshu.com/p/5f41e9c7054c
    来源：简书
    """

def demo9():
    # 子生成器
    @asyncio.coroutine
    def get_image(img_url):
        yield from asyncio.sleep(1)
        # resp = yield from aiohttp.request('GET', img_url)
        resp = yield from aiohttp.ClientSession().get(img_url)
        image = yield from resp.read()
        return image


    def save_image(img, img_url):
        time.sleep(0.5)
        # with open(os.path.join('./img_file', img_url.split('/')[-1]), 'wb') as f:
        pwd = os.path.join(os.path.dirname(os.path.abspath(__file__)) + '.\\img_file', img_url.split('/')[-1])
        with open(pwd, 'wb') as f:
            f.write(img)


    @asyncio.coroutine
    def download_one(img_url):
        image = yield from get_image(img_url)
        save_image(image, img_url)


    def thread_download_one(img_url):
        time.sleep(1)
        resp = requests.get(img_url)
        image = resp.text
        save_image(image, img_url)

    images_list = [
        'http://img3m0.ddimg.cn/67/4/24003310-1_b_5.jpg'
        'http://img3m2.ddimg.cn/43/13/23958142-1_b_12.jpg',
        'http://img3m0.ddimg.cn/60/17/24042210-1_b_5.jpg',
        'http://img3m4.ddimg.cn/20/11/23473514-1_b_5.jpg',
        'http://img3m4.ddimg.cn/40/14/22783504-1_b_1.jpg',
        'http://img3m7.ddimg.cn/43/25/23254747-1_b_3.jpg',
        'http://img3m9.ddimg.cn/30/36/23368089-1_b_2.jpg',
        'http://img3m1.ddimg.cn/77/14/23259731-1_b_0.jpg',
        'http://img3m2.ddimg.cn/33/18/23321562-1_b_21.jpg',
        'http://img3m3.ddimg.cn/2/21/22628333-1_b_2.jpg',
        'http://img3m8.ddimg.cn/85/30/23961748-1_b_10.jpg',
        'http://img3m1.ddimg.cn/90/34/22880871-1_b_3.jpg',
        'http://img3m2.ddimg.cn/62/27/23964002-1_b_6.jpg',
        'http://img3m5.ddimg.cn/84/16/24188655-1_b_3.jpg',
        'http://img3m6.ddimg.cn/46/1/24144166-1_b_23081.jpg',
        'http://img3m9.ddimg.cn/79/8/8766529-1_b_0.jpg']
    
    start_time = time.time()
    loop = asyncio.get_event_loop()
    to_do_tasks = [download_one(img) for img in images_list]

    res, _ = loop.run_until_complete(asyncio.wait(to_do_tasks))
    # print(len(res))
    print(f"asyncio cost 总耗时 - {round(time.time() - start_time, 2)}")

    # ======================多线程版本===============================
    start_time = time.time()
    with ThreadPoolExecutor() as executor:
        res = [executor.submit(thread_download_one, i) for i in images_list]
    # print(len(res))
    print(f"Thread cost 总耗时 - {round(time.time() - start_time, 2)}")

    """
    代码解读
        1. 增加了多线程的下载函数thread_download_one, 和asyncio的方式一样在http请求的时候阻塞1s
        2. 承接我们上一章的问题, 上一章的问题主要就是在save_image()函数, save_image操作硬盘保存文件, 控制权交还给主循环, 此刻有很多子生成器都返回了数据等待主线程的处理, 
            会导致主线程阻塞, 我们模拟耗时操作硬盘(休眠0.5s), 最终耗时8.63s, 而多线程耗时6.63s左右, asyncio比多线程效率更低了, 线程池多个线程并发的写硬盘, 而此刻asyncio需要主线程处理完一个任务的写硬盘操作之后才能处理下一个任务, 所以效率会很低。
    
    备注: 硬盘的操作会阻塞主线程 - save_image()，需要优化

    输出:
        文件保存在当前目录 + //img_file(几张图片)
        asyncio cost 总耗时 - 8.73
        Thread cost 总耗时 - 7.07

    resp = yield from aiohttp.request('GET', img_url) - 会引发异常
    使用 - resp = yield from aiohttp.ClientSession().get(img_url) - 代替

    作者：lpj24
    链接：https://www.jianshu.com/p/2d6a509fcc1f
    来源：简书
    """

def demo10():
    # 子生成器
    @asyncio.coroutine
    def get_image(img_url):
        yield from asyncio.sleep(1)
        # resp = yield from aiohttp.request('GET', img_url)
        resp = yield from aiohttp.ClientSession().get(img_url)
        image = yield from resp.read()
        return image


    def save_image(img, img_url):
        time.sleep(0.5)
        
        # with open(os.path.join('./img_file', img_url.split('/')[-1]), 'wb') as f:
        pwd = os.path.join(os.path.dirname(os.path.abspath(__file__)) + '.\\img_file', img_url.split('/')[-1])
        with open(pwd, 'wb') as f:
            f.write(img)


    @asyncio.coroutine
    def download_one(img_url):
        image = yield from get_image(img_url)
        loop = asyncio.get_event_loop()
        loop.run_in_executor(None, save_image, image, img_url)


    def thread_download_one(img_url):
        time.sleep(1)
        resp = requests.get(img_url)
        image = resp.text
        save_image(image, img_url)

    images_list = [
        'http://img3m0.ddimg.cn/67/4/24003310-1_b_5.jpg'
        'http://img3m2.ddimg.cn/43/13/23958142-1_b_12.jpg',
        'http://img3m0.ddimg.cn/60/17/24042210-1_b_5.jpg',
        'http://img3m4.ddimg.cn/20/11/23473514-1_b_5.jpg',
        'http://img3m4.ddimg.cn/40/14/22783504-1_b_1.jpg',
        'http://img3m7.ddimg.cn/43/25/23254747-1_b_3.jpg',
        'http://img3m9.ddimg.cn/30/36/23368089-1_b_2.jpg',
        'http://img3m1.ddimg.cn/77/14/23259731-1_b_0.jpg',
        'http://img3m2.ddimg.cn/33/18/23321562-1_b_21.jpg',
        'http://img3m3.ddimg.cn/2/21/22628333-1_b_2.jpg',
        'http://img3m8.ddimg.cn/85/30/23961748-1_b_10.jpg',
        'http://img3m1.ddimg.cn/90/34/22880871-1_b_3.jpg',
        'http://img3m2.ddimg.cn/62/27/23964002-1_b_6.jpg',
        'http://img3m5.ddimg.cn/84/16/24188655-1_b_3.jpg',
        'http://img3m6.ddimg.cn/46/1/24144166-1_b_23081.jpg',
        'http://img3m9.ddimg.cn/79/8/8766529-1_b_0.jpg']
    
    start_time = time.time()
    loop = asyncio.get_event_loop()
    to_do_tasks = [download_one(img) for img in images_list]

    res, _ = loop.run_until_complete(asyncio.wait(to_do_tasks))
    # print(len(res))
    print(f"修改后 asyncio cost 总耗时 - {round(time.time() - start_time, 2)}")

    # ======================多线程版本===============================
    start_time = time.time()
    with ThreadPoolExecutor() as executor:
        res = [executor.submit(thread_download_one, i) for i in images_list]
    # print(len(res))
    print(f"修改后 Thread cost 总耗时 - {round(time.time() - start_time, 2)}")

    """
    只需要改一下download_one函数
    
    代码解读
        1. download_one函数中创建的loop循环对象和main函数中的loop对象是同一个, 可以看看源码或者id()一下
        2. 主函数中不要loop.close(), run_in_executor函数每次都会调用self._check_closed()检测循环是否关闭
        3. 书本中还介绍了yield from semaphore来限制并发请求数量, 由于asyncio不向多线程那样阻塞, 加入循环事件任务被快速驱动, 并发访问人家的网页, 
            所以使用semaphore来及限制并发的数量, 让你的程序温柔对待他人的网站。这一块可以结合书中的代码学习, 这里不展开
    
    输出:
        文件保存在当前目录 + //img_file(几张图片)
        修改后 asyncio cost 总耗时 - 1.26
        修改后 Thread cost 总耗时 - 4.06

    作者：lpj24
    链接：https://www.jianshu.com/p/2d6a509fcc1f
    来源：简书
    """


if __name__ == '__main__':
    demo()
    demo1()
    demo2()
    demo3()
    demo4()
    demo5()
    demo6()
    demo7()

    # print("↓↓↓ - 下面代码需要分开执行 - ↓↓↓")
    # demo8()
    # demo9()
    # demo10()


