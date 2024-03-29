package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	hnd "traductor/handlers"
)

func main() {
	port := 8080

	fmt.Println("holi")

	for path, handler := range hnd.Manejadores {
		http.HandleFunc(path, handler)
	}

	//http.HandleFunc(hnd.PathInicio, hnd.Lista[hnd.PathInicio])

	//http.HandleFunc("/js/", hnd.Js)
	//http.HandleFunc("/envio", hnd.Insert)
	//http.HandleFunc("/lista", hnd.List)

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))
	fmt.Println("Servidor abierto en http://localhost:" + strconv.Itoa(port))
}
