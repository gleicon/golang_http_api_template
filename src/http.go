// Copyright 2015 %name% authors.  All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/fiorix/go-redis/redis"
)

type httpServer struct {
	config *configFile
	redis  *redis.Client
	mysql  *sql.DB
}

func (s *httpServer) init(cf *configFile, rc *redis.Client, db *sql.DB) {
	s.config = cf
	s.redis = rc
	s.mysql = db

	// Initialize http handlers.
	s.route()
}

func (s *httpServer) ListenAndServe() {
	if s.config.HTTP.Addr == "" {
		return
	}
	srv := http.Server{
		Addr: s.config.HTTP.Addr,
	}
	log.Println("Starting HTTP server on", s.config.HTTP.Addr)
	log.Fatal(srv.ListenAndServe())
}

func (s *httpServer) ListenAndServeTLS() {
	if s.config.HTTPS.Addr == "" {
		return
	}
	srv := http.Server{
		Addr: s.config.HTTPS.Addr,
	}
	log.Println("Starting HTTPS server on", s.config.HTTPS.Addr)
	log.Fatal(srv.ListenAndServeTLS(
		s.config.HTTPS.CertFile,
		s.config.HTTPS.KeyFile,
	))
}
