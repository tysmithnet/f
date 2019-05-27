import sys
import urwid
import asyncio

buffer = []

async def stream_as_generator(loop, stream):
    reader = asyncio.StreamReader(loop=loop)
    reader_protocol = asyncio.StreamReaderProtocol(reader)
    await loop.connect_read_pipe(lambda: reader_protocol, stream)

    while True:
        line = await reader.readline()
        if not line:  # EOF.
            break
        yield line

async def update_buffer_from_stdin(loop):
    global buffer
    async for line in stream_as_generator(loop, sys.stdin):
        buffer.append(line)

aioloop = asyncio.get_event_loop()
aioloop.create_task(update_buffer_from_stdin(aioloop))
txt = urwid.Text(u"Hello World")
fill = urwid.Filler(txt, 'top')
evl = urwid.AsyncioEventLoop(loop=aioloop)
urwidloop = urwid.MainLoop(fill, event_loop=evl)
urwidloop.run()