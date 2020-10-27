# !/usr/bin/env python
# -*- coding: utf-8 -*-
# @Date     : 2020/05/25 16:28:53
# @File     : 回调机制.py
# @Link     : https://www.cnblogs.com/BenLam/
# @Version  : 1.0

import os
import time
import requests
from concurrent.futures import ThreadPoolExecutor  # 线程池模块

def get(url):
    print(f"GET {url}")
    response = requests.get(url)  # 下载页面
    time.sleep(3)  # 模拟网络延时
    return {"url": url, "content": response.text}  # 页面地址和页面内容

def parse(res):
    # !取到res结果 【回调函数】带参数需要这样
    res = res.result()
    print(f"{res['url']} res is {len(res['content'])} Kb")

if __name__ == "__main__":
    urls = {
        "http://www.jd.com",
        "http://www.baidu.com",
        "http://www.cnblogs.com"
    }
    thread = ThreadPoolExecutor(2)
    for i in urls:
        # 【回调函数】执行完线程后，跟一个函数
        thread.submit(get, i).add_done_callback(parse)