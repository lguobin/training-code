# !/usr/bin/env python
# -*- coding: utf-8 -*-
# @Date     : 2020/05/18 14:34:28
# @File     : selectors - 1.py
# @Link     : https://www.cnblogs.com/BenLam/

import os
import selectors
import socket

# 官网例子 + 注释

# 生成select实例对象
sel = selectors.DefaultSelector()

def accept(sock, mask):

    # 接收链接
    conn, addr = sock.accept()  # Should be ready
    print("[客户端接入] %s connected! " % str(addr))

    # 链接设置非阻塞模式
    conn.setblocking(False)

    # 注册conn，回调 read函数
    sel.register(conn, selectors.EVENT_READ, read)


def read(conn, mask):
    data = conn.recv(1024)  # Should be ready
    if data:
        print("echoing", repr(data), "to", conn)
        conn.send(data)  # Hope it won't block
    else:
        print("[-客户端断开] %s disconnected! " %str(conn))

        # 取消注册
        sel.unregister(conn)

        # 关闭链接
        conn.close()


sock = socket.socket()
sock.bind(('localhost', 1111))
sock.listen(6)
sock.setblocking(False)

# 注册server事件：
# 参数一：sock 进行监听
# 参数二：selectors.EVENT_READ 执行动作
# 参数三：accept,只要来一个链接就回调这个函数
sel.register(sock, selectors.EVENT_READ, accept)

# 第一次调用server，register accept
# 第二次调用client，register read
print("就绪...")
while True:

    # 调用select：优先使用epoll
    events = sel.select()

    # 只要不阻塞就有调用的数据，返回一个列表
    # 默认阻塞，有活动链接就返回活动的链接列表
    for key, mask in events:

        # callback相当于调accept函数
        callback = key.data
        
        print(key.fileobj)
        print(key.fd)
        print(key.events)
        print(key.data)

        """
        这个对象虽然看似有点复杂，实际很灵活
        sel.select() → return ready
        key.fileobj  = <socket.socket fd=240, family=AddressFamily.AF_INET, type=SocketKind.SOCK_STREAM, proto=0, laddr=()>
        key.fd       = 240
        key.events   = 1
        key.data     = <function accept at 0x0000000001D21EA0>
        """

        # 获取函数内存地址，加入参数
        # key.fileobj = 文件句柄
        callback(key.fileobj, mask)


if __name__ == '__main__':
    pass
