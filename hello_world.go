package main

import "fmt"

type httpserver struct {
	max int
}
func NewHttpServer() *httpserver {
	s := &httpserver{2}
	return s
}
func(hs *httpserver) Run() int{
	fmt.Println(hs.max)
	return hs.max
}
func main() {
	NewHttpServer().Run()
}
