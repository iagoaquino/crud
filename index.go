package main

import (
	// Pacote Database SQL para realizar Query
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql" // Driver Mysql para Go
	"server.com/controler"
)

func main() {
	log.Println("inicializando servidor na porta:9090")
	http.HandleFunc("/show", controler.ShowAll)
	http.HandleFunc("/find", controler.Find)
	http.HandleFunc("/delete", controler.Delete)
	http.HandleFunc("/edit", controler.Edit)
	http.HandleFunc("/insert", controler.Insert)
	http.HandleFunc("/", controler.Greet)
	http.ListenAndServe(":9090", nil)
}
