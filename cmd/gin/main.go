package main

import (
	"fmt"
	"go-api/internal/gin"
	"go-api/pkg"
	"net/http"
	"time"
)

func main() {
	config := pkg.Must(gin.GetConfig())

	db := pkg.Must(gin.BuildDb(config))

	router := pkg.Must(gin.Build(config, pkg.Must(gin.BuildLogger(config)), db))

	server := http.Server{
		Addr:         fmt.Sprintf("%s:%d", config.Http.Host, config.Http.Port),
		Handler:      router,
		ReadTimeout:  time.Duration(config.Http.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(config.Http.WriteTimeout) * time.Second,
		IdleTimeout:  time.Duration(config.Http.IdleTimeout) * time.Second,
	}

	server.ListenAndServe()
}
