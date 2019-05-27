import ptvsd

# 5678 is the default attach port in the VS Code debug configurations
print("Waiting for debugger attach")
ptvsd.enable_attach(address=('localhost', 5678), redirect_output=True)
ptvsd.wait_for_attach()
breakpoint()


import sys
import urwid
import asyncio
from pubmarine import PubPen

aioloop = asyncio.get_event_loop()
buffer = []
pubpen = PubPen(aioloop)

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
        pubpen.publish("line", line)




aioloop.create_task(update_buffer_from_stdin(aioloop))
palette = [('I say', 'default,bold', 'default', 'bold'),]
ask = urwid.Edit(('I say', u"What is your name?\n"))
reply = urwid.Text(u"")
button = urwid.Button(u'Exit')
div = urwid.Divider()
pile = urwid.Pile([ask, div, reply, div, button])
top = urwid.Filler(pile, valign='top')

def new_line_added_to_buffer(line):
    reply.set_text(('I say', line))

pubpen.subscribe("line", new_line_added_to_buffer)

def on_ask_change(edit, new_edit_text):
    reply.set_text(('I say', u"Nice to meet you, %s" % new_edit_text))

def on_exit_clicked(button):
    raise urwid.ExitMainLoop()

urwid.connect_signal(ask, 'change', on_ask_change)
urwid.connect_signal(button, 'click', on_exit_clicked)
urwid.MainLoop(top, palette).run()