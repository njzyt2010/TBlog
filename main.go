package main

import (
	"TBlog/pkg/config"
	_ "TBlog/pkg/config"
	_ "TBlog/pkg/database"
	"TBlog/pkg/router"
	"net/http"
	"time"
)

func main() {
	r := router.NewRouter()
	port := config.GetServerConf().Port
	server := &http.Server{
		Addr:           ":" + port,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
