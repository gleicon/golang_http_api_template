// Copyright 2015 %name% authors.  All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

func (s *httpServer) route() {
	// Static file server.
	http.Handle("/static/", handlers.CombinedLoggingHandler(os.Stdout, http.FileServer(http.Dir(s.config.DocumentRoot))))

	// Other handlers.
	http.Handle("/", handlers.CombinedLoggingHandler(os.Stdout, http.HandlerFunc(s.indexHandler)))
	http.Handle("/test", handlers.CombinedLoggingHandler(os.Stdout, http.HandlerFunc(s.testHandler)))
	http.Handle("/api/v1/echo", handlers.CombinedLoggingHandler(os.Stdout, http.HandlerFunc(s.echoHandler)))
}

func (s *httpServer) indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, world\r\n")
}

func (s *httpServer) testHandler(w http.ResponseWriter, r *http.Request) {
	// Set the "hello" key in redis first: redis-cli set hello world
	// Then call this handler: curl localhost:8080/test

	// The redis connection is fault-tolerant. Try killing redis and
	// calling /test again. Then run redis and call /test again.

	if v, err := s.redis.Get("hello"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		fmt.Fprintf(w, "hello %s\r\n", v)
	}
}

func (s *httpServer) echoHandler(w http.ResponseWriter, r *http.Request) {
	appname := getURIParameter("/api/v1/echo/", r)
	if appname == "" {
		fmt.Fprintf(w, "no name after /api/v1/echo/")
		http.NotFound(w, r)
	} else {
		fmt.Fprintf(w, appname)
	}
}
