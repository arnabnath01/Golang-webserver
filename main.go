package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer:= http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form",formhandler)
	http.HandleFunc("/hello",hellohandler)

	fmt.Println("Server is listening port 3000...")
	
	if err:=http.ListenAndServe(("localhost:3000"),nil); err!=nil{
		log.Fatal(err)
	}


}

// hellohandler is a function to handle /hello request
func hellohandler(w http.ResponseWriter, r *http.Request){

	if r.URL.Path!="/hello"{
		http.Error(w,"404 not found",http.StatusNotFound)
	}

	if r.Method!="GET"{
		http.Error(w,"Only get method allowed",http.StatusNotFound)
	}
	fmt.Fprintf(w, "hello");
}

// formhandler is a function to handle /form request
func formhandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path!="/form"{
		http.Error(w,"404 not found",http.StatusNotFound)
	}

	if r.Method!="POST"{
		http.Error(w,"bhai sirf post method allowed hai, code change kr yaar ",http.StatusNotFound)
	}

	fmt.Fprintf(w,"Post request successful \n")

	name:=r.FormValue("name")
	// this shows log in the server
	fmt.Println(name,"name entered in the frontend")
	
	address:=r.FormValue("address")
	// this shows log in the server
	  fmt.Println(address,"address entered in the frontend")

	fmt.Fprintf(w,"Your name is : %s\n",name)
	fmt.Fprintf(w,"Your address is :%s\n",address)

}
