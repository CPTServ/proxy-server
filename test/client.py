import socket
import time
import util

HOST = "127.0.0.1"
PORT = 15003


b = b""


# type
b = util.add_string(b, "client")

# key
b = util.add_string(b, "213213213123")

print(b)


s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

s.connect((HOST, PORT))
s.send(b)
while 1:
    msg = s.recv(1024)
    if msg:
        print(msg)
    time.sleep(1)
s.close()
