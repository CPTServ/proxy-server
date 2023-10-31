import socket
import json
import time
import util

HOST = "127.0.0.1"
PORT = 15003


def get_hosts() -> str:
    l = {"v6": [], "v4": [], "port": 15003}
    addrs = socket.getaddrinfo(socket.gethostname(), None)
    for addr in addrs:
        key = None
        if addr[0] == socket.AF_INET6:
            key = "v6"
        elif addr[0] == socket.AF_INET:
            key = "v4"
        if key and addr[-1][0] not in l[key]:
            l[key].append(addr[-1][0])
    return json.dumps(l)


hosts = get_hosts()
# print(hosts)

b = b""


# type
b = util.add_string(b, "server")

# key
b = util.add_string(b, "213213213123")

# data
b = util.add_string(b, hosts)

print(b)


s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
s.connect((HOST, PORT))
s.send(b)
while 1:
    msg = s.recv(1024)
    if len(msg) == 1:
        if msg[0] == 200:
            print("ok")
            break
    time.sleep(1)
s.close()
