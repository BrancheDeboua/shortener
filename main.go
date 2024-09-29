package main

import "github.com/BrancheDeboua/url-shortener/internal/app"

func main() {
	port := ":8089"
	app.Serve(port)
}
