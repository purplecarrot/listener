# listen

Program to simultaneously listen on multiple TCP/UDP ports and reply back to anything sent along with IP addresses and lengths of data received and sent.

This was written to help test NGINX Plus Ingress Controller and Consul Service Mesh running inside SDN/Service Mesh inside Kubernetes clusters.

```
# Listen on TCP ports 8000 and 8001 and UDP port 2000
$ listen -t 8000,8001 -u 2000
```