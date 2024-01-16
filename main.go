package main

import (
	"go-web-native/config"
	// contactcontroller "go-web-native/controller/contactController"
	contactcontroller "go-web-native/controller/contactController"
	"go-web-native/controller/homecontroller"
	"log"
	"net/http"
)

func main() {
	config.ConnnectDB()

	// 1. Home Page
	http.HandleFunc("/", homecontroller.Home)

	http.HandleFunc("/tambah", contactcontroller.Add)
	http.HandleFunc("/edit", contactcontroller.Edit)
	http.HandleFunc("/delete", contactcontroller.Hapus)

	log.Println("Server Running on Port 3030")
	http.ListenAndServe(":3030", nil)
}
