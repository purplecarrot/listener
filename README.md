# listener

*listener* is a server that simultaneously listens on multiple TCP and UDP ports and logs any incoming connection details to STDOUT. It sends a small response back to the client as an acknowledgement.

This program was written to help in the functional testing of *Consul Service Mesh* running inside Kubernetes by being able to easily visually test connectivity between services within the mesh, and to test UDP ingress to pods in a Kubernetes cluster with *NGINX Ingress Controller*

To start *listener* to listen on TCP ports 2001 and 3001, and UDP ports 4001 and 5001, run the following command:
```
$ listener -t 2001,3001 -u 4001,5001
2022/02/19 10:10:50 IP address is 192.168.1.2
2022/02/19 10:10:50 [pid 1477440] alive with listeners on [{2001 tcp 0} {3001 tcp 0} {4001 udp 0} {5001 udp 0}]
2022/02/19 10:10:50 [2001/tcp] listening on TCP port 2001 for HTTP connections
2022/02/19 10:10:50 [3001/tcp] listening on TCP port 3001 for HTTP connections
2022/02/19 10:10:50 [4001/udp] listening on UDP port 4001
2022/02/19 10:10:50 [5001/udp] listening on UDP port 5001
```

You can verify this using ss:
```
# ss -lnp | grep "001"
udp   UNCONN 0      0            *:4001         *:*    users:(("listener",pid=1477440,fd=7))
udp   UNCONN 0      0            *:5001         *:*    users:(("listener",pid=1477440,fd=3))
tcp   LISTEN 0      4096         *:2001         *:*    users:(("listener",pid=1477440,fd=8))
tcp   LISTEN 0      4096         *:3001         *:*    users:(("listener",pid=1477440,fd=10))
```

## Usage


### TCP/HTTP 
```
# client making basic http connection
$ curl http://192.168.1.2:2001/
<head><title>Listener</title></head><body><p>Listener here, nothing posted</p></body>

# server STDOUT
2022/02/19 11:33:42 [2001/tcp] >192.168.1.2:35186 GET / HTTP/1.1 () 
2022/02/19 11:33:42 [2001/tcp] <192.168.1.2:35186 responding: <head><title>Listener</title></head><body><p>Listener here, nothing posted</p></body>
```

```
# client sending JSON POST (JSON is validated and sent back)
$ curl -X POST -H "Accept: application/json" -H "Content-type: application/json"  http://192.168.1.2:3001/ -d '{"hello":"world"}'
{"received": {"hello": "world"}}

# server STDOUT
2022/02/19 11:36:30 [3001/tcp] >192.168.1.2:45638 POST / HTTP/1.1 (application/json) {"hello":"world"}
2022/02/19 11:36:30 [3001/tcp] <192.168.1.2:45638 responding (application/json): {"received": {"hello": "world"}}
```

### UDP
```
# client interactively sending some text over UDP
$ nc -u 192.168.1.2 4001
hello port 401
pid 1477440 received 15 bytes from you at 127.0.0.1:34466

# server STDOUT
2022/02/19 11:30:07 [4001/udp] >127.0.0.1:34466 sent us 15 bytes "hello port 401"
2022/02/19 11:30:07 [4001/udp] <127.0.0.1:34466 sent back "1549101 received 15 bytes from you at 127.0.0.1:34466"
```

