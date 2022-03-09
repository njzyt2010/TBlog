package main

import (
	"TBlog/internal/controller"
	"TBlog/pkg/config"
	_ "TBlog/pkg/config"
	_ "TBlog/pkg/database"
	"net/http"
	"time"
)

func main() {
	r := controller.NewRouter()
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
