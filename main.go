package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
)

func main(){
	http.HandleFunc("/",Inicio)

	http.ListenAndServe(":8080",nil)
}

func Inicio(w http.ResponseWriter, r*http.Request){
	//fmt.Fprintf(w, "Sevidor funcionando")
	
	//log.Println("Hellow World")
	d,_:=ioutil.ReadAll(r.Body)
	log.Printf("Data=  %s\n",d)
	fmt.Fprintf(w, "Data %s\n",d)
}