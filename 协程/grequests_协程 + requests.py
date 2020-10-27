# !/usr/bin/env python
# -*- coding: utf-8 -*-
# @Date     : 2020/05/22 14:52:35
# @File     : grequests_协程 + requests.py
# @Link     : https://www.cnblogs.com/BenLam/
# @Version  : 1.0

import os
import time
import requests
import grequests

# urls = [
#     'https://docs.python.org/2.7/library/index.html',
#     'https://docs.python.org/2.7/library/dl.html',
#     'http://www.iciba.com/partial',
#     'http://2489843.blog.51cto.com/2479843/1407808',
#     'http://blog.csdn.net/woshiaotian/article/details/61027814',
#     'https://docs.python.org/2.7/library/unix.html',
#     'http://2489843.blog.51cto.com/2479843/1386820',
#     'http://www.bazhuayu.com/tutorial/extract_loop_url.aspx?t=0',
# ]

urls = []
for url in range(1, 60):
	url = "http://shop.projectsedu.com/goods/{}/".format(url)
	urls.append(url)

def consum_time(func):
    # 查看代码耗时装饰器
    def wrapper(*args, **kw):
        start_time = time.time()
        res = func(*args,**kw)
        print(f"{func.__name__}: 耗时 - {round(time.time() - start_time, 2)}")
        return res
    return wrapper

@consum_time
def demo1():
    for url in urls:
        res = requests.get(url)
        # print(res.status_code)

@consum_time
def demo2():
    tasks = [grequests.get(u) for u in urls]
    res = grequests.map(tasks, size=3)
    # print(res)

@consum_time
def demo3():
    tasks = [grequests.get(u) for u in urls]
    res = grequests.map(tasks, size=6)
    # print(res)

@consum_time
def demo4():
    tasks = [grequests.get(u) for u in urls]
    res = grequests.map(tasks, size=None)
    # print(res)

def test():
    # grequests的底层库，是requests，因此它也支持事件钩子
    print("事件钩子")
    def print_url(r, *args, **kwargs):
            print(r.url, r.status_code)

    url = "http://test-api-mall.vm-shopping.com/v2/operators/login"
    # 1.
    res = requests.get(url, hooks={"response":print_url})
    print(res.text)

    # 2.
    tasks = []
    req = grequests.get(url, callback=print_url)
    tasks.append(req)
    res = grequests.map(tasks)


def _requests():
    start_time = time.time()
    response = [requests.get(url) for url in urls]
    print(f"requests 总耗时 - {round(time.time() - start_time, 2)}")
    return response

def _grequests():
    start_time = time.time()
    response = grequests.map([grequests.get(u) for u in urls], size=20)
    print(f"grequests 总耗时 - {round(time.time() - start_time, 2)}")
    return response

def main():
    import threading
    r = threading.Thread(target=_requests)
    g = threading.Thread(target=_grequests)
    
    r.start()
    g.start()

if __name__ == '__main__':
    # demo1()
    # demo2()
    # demo3()
    # demo4()
    # test()
    main()
