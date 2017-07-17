
package main

import(
	"net/http"
	"fmt"
	"strconv"
	"os"
)

//Init global variables
var num = 0
var id

//Handles requests from root ("/")
func handleRoot(w http.ResponseWriter, r *http.Request) {
	num++
	fmt.Fprintf(w , "This is a Go webserver! \nThis webpage has been loaded "+strconv.Itoa(num)+" times.")
}

//Handles requests for favicon.ico to take these requests away from root
func handleIcon(w http.ResponseWriter, r *http.Request) {
	//Code for favicon goes here!
}

//Main function
func main(){
	//Sets port number and gets dockerID
	port := 7001
	id := getDockerID()
	
	//Registers handlers for GET and POST requests
	http.HandleFunc("/favicon.ico", handleIcon)
	http.HandleFunc("/", handleRoot)
	
	//Prints start message
	fmt.Println("Server Started on :"+strconv.Itoa(port))
	
	//Initializes listening thread
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}

