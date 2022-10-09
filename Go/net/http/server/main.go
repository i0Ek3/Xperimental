package main

import "net/http"

func hander(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		w.Write([]byte("Get: GET"))
	case "POST":
		w.Write([]byte("Post: POST"))
	case "PUT":
		w.Write([]byte("Put: PUT"))
	case "DELETE":
		w.Write([]byte("Delete: DELETE"))
	}

}

func main() {
	http.HandleFunc("/hi", hander)
	http.ListenAndServe(":8080", nil)
}
