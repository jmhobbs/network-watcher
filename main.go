package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

var (
	port         int
	verbose      bool
	pollInterval time.Duration
)

func init() {
	flag.IntVar(&port, "port", 4444, "port to listen on")
	flag.BoolVar(&verbose, "verbose", false, "extra output")
	flag.DurationVar(&pollInterval, "interval", time.Second*5, "time between polling targets")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: %s [options] hostname:port ... \n", os.Args[0])
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()

	targets := flag.Args()
	if len(targets) == 0 {
		flag.Usage()
		return
	}

	if verbose {
		log.Println("Network targets:")
		for _, target := range targets {
			log.Println("  ", target)
		}
	}

	addr := fmt.Sprintf(":%d", port)

	if verbose {
		log.Println("Starting server on", addr)
	}
	go serve(addr)

	for {
		for _, target := range targets {
			check(target)
		}
		time.Sleep(pollInterval)
	}
}

func check(target string) {
	conn, err := net.DialTimeout("tcp", target, time.Second)
	if err != nil {
		log.Println(target, "DOWN", err)
		return
	}
	defer conn.Close()

	reply := make([]byte, 1024)
	conn.SetReadDeadline(time.Now().Add(time.Second))

	_, err = conn.Read(reply)
	if err != nil {
		if verbose {
			log.Println(target, "DOWN", err)
		} else {
			log.Println(target, "DOWN")
		}
		return
	}

	log.Println(target, "UP")
}

func serve(addr string) {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("error on net.Listen: %s", err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println("error accepting connection: ", err)
			continue
		}
		conn.Write([]byte(time.Now().String()))
		conn.Close()
	}
}
