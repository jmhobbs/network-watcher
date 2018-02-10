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

# Example

In this example run, three instances exist: `isolated`, `unisolated` and `unisolated-too`.  The connection between `isolated` and `unisolated-too` is blocked, everything else is permitted.

    $ kubectl logs unisolated-66cf5b9fc8-hn7t7
    2018/02/10 17:05:41 isolated:4444 UP
    2018/02/10 17:05:41 unisolated-too:4444 UP
    2018/02/10 17:05:46 isolated:4444 UP
    2018/02/10 17:05:46 unisolated-too:4444 UP
    2018/02/10 17:05:51 isolated:4444 UP
    2018/02/10 17:05:51 unisolated-too:4444 UP
    $ kubectl logs unisolated-too-5468649df4-zl6vp
    2018/02/10 17:05:42 isolated:4444 DOWN dial tcp 10.63.42.6:4444: i/o timeout
    2018/02/10 17:05:42 unisolated:4444 UP
    2018/02/10 17:05:48 isolated:4444 DOWN dial tcp 10.63.42.6:4444: i/o timeout
    2018/02/10 17:05:48 unisolated:4444 UP
    2018/02/10 17:05:54 isolated:4444 DOWN dial tcp 10.63.42.6:4444: i/o timeout
    2018/02/10 17:05:54 unisolated:4444 UP
    2018/02/10 17:06:00 isolated:4444 DOWN dial tcp 10.63.42.6:4444: i/o timeout
    2018/02/10 17:06:00 unisolated:4444 UP

That's all there is to it.
