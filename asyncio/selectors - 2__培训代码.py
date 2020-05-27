# !/usr/bin/env python
# -*- coding: utf-8 -*-
# @Date     : 2020/05/18 14:54:15
# @File     : selectors - 2__培训代码.py
# @Link     : https://www.cnblogs.com/BenLam/
# @Version  : 1.0

import os
import time

import socket
from urllib.parse import urlparse
from selectors import DefaultSelector, EVENT_READ, EVENT_WRITE

#1. epoll并不代表一定比select好
# 在并发高的情况下，连接活跃度不是很高， epoll比select
# 并发性不高，同时连接很活跃， select比epoll好

#通过非阻塞io实现http请求
# select + 回调 + 事件循环
#  并发性高
#  使用单线程

selector = DefaultSelector()
# 全局变量 selector 是 poll 或 epoll 由 DefaultSelector() 自己选择


#使用select完成http请求
urls = []
stop = False

class Fetcher:
    def connected(self, key):
        selector.unregister(key.fd)
        self.client.send("GET {} HTTP/1.1\r\nHost:{}\r\nConnection:close\r\n\r\n".format(self.path, self.host).encode("utf-8"))
        selector.register(self.client.fileno(), EVENT_READ, self.readable)

    def readable(self, key):
        d = self.client.recv(1024)
        if d:
            self.data += d
        else:
            selector.unregister(key.fd)
            data = self.data.decode("utf8")
            html_data = data.split("\r\n\r\n")[1]
            # 不想让他打印
            # print(html_data)
            self.client.close()
            urls.remove(self.spider_url)
            if not urls:
                global stop
                stop = True

    def get_url(self, url):
        self.spider_url = url
        url = urlparse(url)
        self.host = url.netloc
        self.path = url.path
        self.data = b""
        if self.path == "":
            self.path = "/"

        # 建立socket连接
        self.client = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        self.client.setblocking(False)

        try:
            self.client.connect((self.host, 80))  # 阻塞不会消耗cpu
        except BlockingIOError as e:
            pass

        #注册
        selector.register(self.client.fileno(), EVENT_WRITE, self.connected)


def loop():
    #事件循环，不停的请求socket的状态并调用对应的回调函数
    #1. select本身是不支持register模式
    #2. socket状态变化以后的回调是由程序员完成的
    while not stop:
        ready = selector.select()
        for key, mask in ready:
            call_back = key.data
            call_back(key)
    #回调+事件循环+select(poll\epoll)

if __name__ == "__main__":
    print(" 方案一 - 同步处理".center(30, "-"))
    fetcher = Fetcher()
    start_time = time.time()

    for url in range(1,20):
        url = "http://shop.projectsedu.com/goods/{}/".format(url)
        urls.append(url)
        fetcher = Fetcher()
        fetcher.get_url(url)
    loop()
    print(f"总耗时 - {round(time.time() - start_time, 2)}")

    
    
    print(" 方案二 - 单进程处理".center(30, "-"))
    import requests
    from urllib.request import urlopen

    def _response(url):
        # print(urlopen(url).status)
        # print(urlopen(url).read().decode("utf-8"))
        return urlopen(url).status
        # return requests.get(url).status_code
    
    start_time = time.time()
    for _ in range(1, 20):
        url = "http://shop.projectsedu.com/goods/{}/".format(_)
        _response(url)
    print(f"总耗时 - {round(time.time() - start_time, 2)}")