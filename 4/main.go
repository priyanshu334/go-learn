package main

import (
	"fmt"
	"log"
	"net/http"
)
func formHandler(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err !=nil{
		fmt.Fprint(w,"parseForm() err: %v",err)
		return
	}
	fmt.Fprint(w,"Post request successful")
	name := r.FormValue("name")
	address := r.FormValue("Address")
	fmt.Fprintf(w,"Name = %s\n",name)
	fmt.Fprintf(w,"Addresss= $s\n",address)

}
func helloHandler(w http.ResponseWriter,r *http.Request){
	if r.URL.Path != "/hello"{
		http.Error(w,"404 not founc",http.StatusNotFound)
		return
	}

	if r.Method != "GET"{
		http.Error(w,"method is not supproted",http.StatusNotFound)
		return
	}

	fmt.Fprint(w,"Hello")
}

func main(){
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/",fileserver)
	http.HandleFunc("/form",formHandler)
	http.HandleFunc("/hello",helloHandler)


	fmt.Print("starting server at port 8080\n")

	if err:= http.ListenAndServe(":8000",nil); err!=nil{
		log.Fatal(err)
	}

}