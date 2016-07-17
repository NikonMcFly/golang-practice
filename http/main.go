package main


import (
	"net/http"
	"io"
	"fmt"
)

type myHandler struct{

}

func serve(resp http.ResponseWriter, r *http.Request){
	io.WriteString(resp, "hello")
	// if r.URL.Path == "/cat"{
	// 	fmt.Println("hello")
	// }
	if r.Method == "Post"{
		fmt.Println("post")
	}
}

func Create(resp http.ResponseWriter, r *http.Request) {
	http.ServeFile(resp, r, "http/page.html")
}

func here(resp http.ResponseWriter, r *http.Request){
	key := "q"
	val := r.FormValue(key)
	fmt.Println("value:", val)
	resp.Header().Set("Content-Type", "text/html")
} 

func main(){
	// var h myHandler
	// mux := http.NewServeMux()

	http.HandleFunc("/h", here)
	http.HandleFunc("/c", Create)
	http.HandleFunc("/", serve)
	http.ListenAndServe(":8080", nil)
}