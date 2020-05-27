# !/usr/bin/env python
# -*- coding: utf-8 -*-
# @Date     : 2020/05/23 13:47:40
# @File     : hashlib加密 - sign.py
# @Link     : https://www.cnblogs.com/BenLam/

import os
import time
import hashlib
import requests


def get_item_sign_content(item):
    if item is None:
        return ''
    if isinstance(item, (str, int)):
        return str(item)
    if isinstance(item, list):
        v = ''
        for it in item:
            v += "".join(str(it.get(i)) for i in sorted(it.keys()))
        return v
    return ''

def get_items_sign_content(items={}):
    result = ''
    if items is None:
        return ''
    for key in sorted(items.keys()):
        result += get_item_sign_content(items.get(key))
    return result

# 根据签名内容content生成 sign
def sign_generate(content, key):
    code = "{}{}".format(content, key)
    for i in range(7):
        code = hashlib.sha256(code.encode('utf-8')).hexdigest()
    return code


def request(option, url, headers={}, user={"name":"admin","password":"123456"}):
    if option == "get" or option == "GET":
        try:
            response = requests.get(url, headers=headers).json()
            if response.get("message"):
                print(response.get("message"))
            elif response.get("orders") == []:
                print("成功启动本地点餐...")
                return True
            else:
                print(f"订单列表...\n{response}")
                return True
        except Exception:
            return False

    elif option == "post" or option == "POST":
        response = requests.post(url, json=user, headers=headers).json()
        try:
            response = response.get("access_token")
            response = "JWT " + response
            headers = {"Authorization": response}
            return headers
        except Exception:
            return False
    else:
        print("请求方式不对......[例子: GET\get| 或 |POST\post ]")
    return None

def test_signal():
    #生成content
    items = {"remote":0, "size":20, "start":int(time.time())}
    del items["start"]
    content = get_items_sign_content(items)
    sign = sign_generate(content, 123456)
    # ec23ad6ad4c495a5c06bcbfc3d8f6ae8b20e045d8cf9bbbd5d48ca17648bb47d
    # print(f"测试环境\n\thttp://test-food.vm-shopping.com/v1/bridge/orders?remote=0&start=&size=20&sign={sign}")
    
    url = "http://test-food.vm-shopping.com/v1/users/login"
    token = request("post", url)

    # url = "http://test-food.vm-shopping.com/v1/orders/"
    url = f"http://test-food.vm-shopping.com/v1/bridge/orders?remote=0&start=&size=20&sign={sign}"
    request("get", url, token)

    return token


if __name__ == '__main__':
    test_signal()