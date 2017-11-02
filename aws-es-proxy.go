package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/namsral/flag"

	awsauth "github.com/smartystreets/go-aws-auth"
)

var (
	endpoint = flag.String("elastic-url", "", "elastic url")
	listen   = flag.String("listen", ":3000", "listen port")
)

func main() {
	u, err := url.Parse(*endpoint)
	if err != nil {
		log.Fatal(err)
	}

	p := &httputil.ReverseProxy{Director: func(r *http.Request) {
		r.Host, r.URL.Scheme, r.URL.Host, r.URL.Path = u.Host, u.Scheme, u.Host, u.Path
		awsauth.Sign(r)
	}}

	log.Fatal(http.ListenAndServe(*listen, p))
}
