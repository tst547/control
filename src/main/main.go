package main

import (
	"fmt"
	"net/http"
	"log"
	"cm"
)

func theMain(w http.ResponseWriter, req *http.Request) {
	log.Println("Hello,"+req.URL.Path[1:])
}
func fileList(w http.ResponseWriter, req *http.Request) {
	log.Println("Inside HelloServer handler")
	fmt.Fprintf(w, "Hello,"+req.URL.Path[1:])
}
func forWin(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside HelloServer handler")
	fmt.Fprintf(w, "Hello,"+req.URL.Path[1:])
}


func main() {
	go cm.ScanPort(22555)
	http.HandleFunc("/theMain",theMain)
	http.HandleFunc("/fileList",fileList)
	http.HandleFunc("/forWin",forWin)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Println("ListenAndServe: ", err.Error())
		return
	}
}
