package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type handle struct {
	param string
}

func (this *handle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	targets, _ := q[this.param]
	target := targets[0]

	remote, err := url.Parse(target)
	if err != nil {
		log.Fatalln(err)
	}

	q.Del(this.param)

	proxy := httputil.NewSingleHostReverseProxy(remote)
	r.URL.Host = remote.Host
	r.URL.Scheme = remote.Scheme
	r.Host = remote.Host
	r.URL.RawQuery = q.Encode()

	log.Println(r.RemoteAddr + " " + r.Method + " " + r.URL.String() + " " + r.UserAgent())
	proxy.ServeHTTP(w, r)
}
