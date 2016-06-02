// Simple proxy.
// Based on http://blog.charmes.net/2015/07/reverse-proxy-in-go.html

package main

import (
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

// NewMultipleHostReverseProxy creates a reverse proxy that will randomly
// select a host from the passed `targets`
func newPeterHostReverseProxy(target *url.URL) *httputil.ReverseProxy {
	director := func(req *http.Request) {
		fmt.Printf("\nDIRECTOR !!!\n")

		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.URL.Path = target.Path
		fmt.Printf("DIRECTOR====> Scheme='%s' Host='%s' Path='%s'\n",
			req.URL.Scheme, req.URL.Host, req.URL.Path)
	}

	return &httputil.ReverseProxy{
		Director: director,
		Transport: &http.Transport{
			Proxy: func(req *http.Request) (*url.URL, error) {
				fmt.Printf("PROXY====>%#v\n", req)
				return http.ProxyFromEnvironment(req)
			},
			Dial: func(network, addr string) (net.Conn, error) {
				fmt.Printf("DIAL====>addr='%s'\n", addr)
				conn, err := (&net.Dialer{
					Timeout:   30 * time.Second,
					KeepAlive: 30 * time.Second,
				}).Dial(network, addr)
				if err != nil {
					println("Error during DIAL:", err.Error())
				}
				return conn, err
			},
			TLSHandshakeTimeout: 10 * time.Second,
		},
	}
}

const (
	fromPort = 631
	toPort   = 9631
	// fromPort = 9090
	//toPort   = 9091
)

func main() {
	fmt.Printf("Redirecting %d=>%d\n", fromPort, toPort)
	// proxy := httputil.NewSingleHostReverseProxy(&url.URL{
	// 	Scheme: "http",
	// 	Host:   "localhost:9091",
	// })
	// http://localhost:9631/
	target := url.URL{
		Scheme: "http",
		Host:   fmt.Sprintf("localhost:%d", toPort),
	}
	fmt.Printf("target=%#v\n", target)
	proxy := newPeterHostReverseProxy(&target)
	// proxy := httputil.NewSingleHostReverseProxy(&target)

	fromAddr := fmt.Sprintf(":%d", fromPort)
	fmt.Printf("Listening on %s\n", fromAddr)
	http.ListenAndServe(fromAddr, proxy)
}
