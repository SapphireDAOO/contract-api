package main

import "net/http"

func Route() *http.ServeMux {
	mux := http.NewServeMux()

	return mux
}
