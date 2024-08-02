package main

import (
	"fmt"
	"net/http"
)

// ex func
// func handlerIndex(w http.ResponseWriter, r *http.Request){
// 	message := "welcome"
// 	w.Write([]byte(message))
// }
// func handlerHello(w http.ResponseWriter, r *http.Request){
// 	message := "hello world"
// 	w.Write([]byte(message))
// }

func main(){
	// closure 
	// handlerEx := func(w http.ResponseWriter, r *http.Request){
	// 	messages := "hellow example"
	// 	w.Write([]byte(messages))
	// }
	
	// http.HandleFunc("/ex", handlerEx)
	// http.HandleFunc("/", handlerIndex)
	// http.HandleFunc("/index", handlerIndex)
	// http.HandleFunc("/hello", handlerHello)
	
	// // anonym
	// http.HandleFunc("/data",func(w http.ResponseWriter, r *http.Request){
	// 	w.Write([]byte("hello again"))
	// })
	
	// address := "localhost:9000"
	// fmt.Printf("server started at %s\n", address)
	// server := new(http.Server)
	// server.Addr = address
	// err := server.ListenAndServe()
	
	// if err != nil {
	// 	fmt.Printf(err.Error())
	// }	
	http.Handle("/static/",
		http.StripPrefix("/static", 
			http.FileServer(http.Dir("assets"))))
	
	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000",nil)
	
}