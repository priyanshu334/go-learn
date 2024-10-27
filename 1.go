package main

import(
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter,r *http.Request){
	fmt.Fprint(w,"Hello World")
}
func main(){
	http.HandleFunc("/",handler)

	fmt.Println("starting server")
	err :=http.ListenAndServe(":8000",nil)
	if err !=nil{
		fmt.Println("Error staring server",err)
	}

}