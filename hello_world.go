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
	fmt.Println("test")
	fmt.Println("123")

	fmt.Println("456")

	fmt.Println("789")
	fmt.Println("abc")

}
