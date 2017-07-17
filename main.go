
package main

import(
	"net/http"
	"fmt"
	"strconv"
	"os"
)

var num = 0
var id

func handleRoot(w http.ResponseWriter, r *http.Request) {
	num++
	fmt.Fprintf(w , "This is a Go webserver! \nThis webpage has been loaded "+strconv.Itoa(num)+" times.")
}

func handleIcon(w http.ResponseWriter, r *http.Request) {
	
}


func main(){
	port := 7001
	id := getDockerID()
	
	http.HandleFunc("/favicon.ico", handleIcon)
	http.HandleFunc("/", handleRoot)
	
	fmt.Println("Server Started on :"+strconv.Itoa(port))
	
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}

