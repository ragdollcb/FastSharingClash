package main

import (
	"FastSharingClash/pkg"
	"FastSharingClash/router"
	"fmt"
	"net/http"
)

func main() {
	Addr := fmt.Sprint(":" + pkg.GetConfig("default.port"))
	s := &http.Server{
		Addr:           Addr,
		Handler:        router.InitRouter(),
		MaxHeaderBytes: 1 << 20,
	}

	_ = s.ListenAndServe()
}
