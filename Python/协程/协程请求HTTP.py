# !/usr/bin/env python
# -*- coding: utf-8 -*-
# @Date     : 2020/05/21 17:05:27
# @File     : 协程请求HTTP.py
# @Link     : https://www.cnblogs.com/BenLam/

import os
import requests
import asyncio


def demo1():

    class Reques_Error(Exception):
        pass

    def que(option=None, url=None, data=None):
        if option == None and url == None:
            raise Reques_Error("传入参数错误")
        elif option == "get" or option == "GET":
            x = requests.get(url)
        elif option == "post" or option == "POST":
            x = requests.post(url, json=data)
        else:
            raise Reques_Error("请填入请求方式")

        try:        
            body = x.json()
        except Exception:
            body = x.text
        except BaseException:
            body = x.text
        return x.status_code, body

    def request():
        url = "http://www.01happy.com/demo/accept.php"
        data = {"key":"Value","id":"password"}
        x = yield from que("post", url, data)
        y = yield from que("get", url)
        # z = yield from que()

    print(list(request()))
    """
    打印:
        [200, 'get:Array\n(\n)\ncookie:Array\n(\n)\npost:Array\n(\n)\n']
    """

def demo2():
    print("asyncio执行多个requests请求")


if __name__ == '__main__':
    # demo1()
    demo2()
