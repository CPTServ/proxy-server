# simple-proxy-server
Proxy provide address communicate between clent and server

# Usage
## Start up the proxy server
```
go run .
```
## client or server:
### Get host data
```python
def get_hosts() -> str:
    # doesn't have to be this kind of structure, the data can be anything
    l = {
        "6": [],
        "4": [],
        "port": 15003
    }
    addrs = socket.getaddrinfo(socket.gethostname(), None)
    for addr in addrs:
        key = None
        if addr[0] == socket.AF_INET6:
            key = "6"
        elif addr[0] == socket.AF_INET:
            key = "4"
        if key and addr[-1][0] not in l[key]:
            l[key].append(addr[-1][0])
    return json.dumps(l)

hosts = get_hosts()
```
### Convert to bytes
the `key` is what relates two clients connected to the proxy server, both two clients have to give the same key
```python
b = b""

def get_len(content: bytes) -> list[int]:
    l: list[int] = []
    total = len(content)
    while total >= 255:
        l.append(total%255)
        total //=255
    l.append(total)
    if len(l) > 8:
        raise Exception("Length over 8")
    while len(l) < 8:
        l.append(0)
    return l

# type
t = "client".encode()
## or
t = "server".encode()
tl = get_len(t)
b = b + bytes(tl) + t

# key
k = "213213213123".encode()
kl = get_len(k)
b = b + bytes(kl) + k

# data
c = hosts.encode()
cl = get_len(c)
b = b + bytes(cl) + c
```
# Connect to the server
```python
s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
s.connect((HOST, PORT))
s.send(b)
while 1:
    msg = s.recv(1024)
    if msg:
        print(msg)
        break
    time.sleep(1)
s.close()
```
