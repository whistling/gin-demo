package core

import (
	"fmt"
	"gin/src/initialize"
	"net/http"
	"time"
)

func RunServer() {

	Router := initialize.Routers()

	s := &http.Server{
		Addr:           ":8080",
		Handler:        Router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Println("server run success on 8080")

	err := s.ListenAndServe()

	if err != nil {
		fmt.Println("error : ", err)
	}
}
