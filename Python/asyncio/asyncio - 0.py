# import asyncio

# async def hello():
#     print("Hello world!")
#     r = await asyncio.sleep(1)
#     print("回复: Hello again!")

# # 获取EventLoop:
# loop = asyncio.get_event_loop()
# # 执行coroutine
# loop.run_until_complete(hello())
# loop.close()

# @asyncio.coroutine
# def hello():
#     print("Hello world!")
#     # 异步调用asyncio.sleep(1):
#     r = yield from asyncio.sleep(1)
#     print("回复: Hello again!")

# # 获取EventLoop:
# loop = asyncio.get_event_loop()
# # 执行coroutine
# loop.run_until_complete(hello())
# loop.close()

# import asyncio
# import threading

# async def hello():
#     print("Hello world! (%s)" % threading.currentThread())
#     await asyncio.sleep(1)
#     print("Hello again! (%s)" % threading.currentThread())

# loop = asyncio.get_event_loop()
# tasks = [hello(), hello(), hello()]
# loop.run_until_complete(asyncio.wait(tasks))
# loop.close()

# import asyncio

# async def curl(url):
#     print(f"url 地址列表 - {url}" )
#     connect = asyncio.open_connection(url, 80)
#     reader, writer = await connect
#     header = "GET / HTTP/1.1\r\nHost: %s\r\n\r\n" % url
#     writer.write(header.encode("utf-8"))
#     await writer.drain()
#     while True:
#         line = await reader.readline()
#         if line == b"\r\n":
#             break
#         print("%s header > %s" % (url, line.decode("utf-8").rstrip()))
#     writer.close()

# urls =  ["www.z.cn", "www.jd.com"]
# loop = asyncio.get_event_loop()
# tasks = [curl(_) for _ in urls]
# loop.run_until_complete(asyncio.wait(tasks))
# loop.close()