# Network Watcher

This is an application that acts as both client and server to monitor network connections.

It's designed to respond to other instances of itself, so you would run this in multiple machines and watch the traffic between them.

I originally created to watch the effect of [Kubernetes network policies](https://kubernetes.io/docs/concepts/services-networking/network-policies/) as I modified them.

# Installation

To install, just run `go get github.com/jmhobbs/network-watcher`

Alternatively, I have a Docker image you can use at https://hub.docker.com/r/jmhobbs/network-watcher/

# Usage

    usage: ./network-watcher [options] hostname:port ...
      -interval duration
          time between polling targets (default 5s)
      -port int
          port to listen on (default 4444)
      -verbose
          extra output
