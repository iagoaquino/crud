package main

import (
	"log"
	"net/http"

	"server.com/conect"
	"server.com/model"
)

func ShowAll(w http.ResponseWriter, r *http.Request) {
	bd := conect.Conect()

	dados, err := bd.Query("SELECT * FROM aluno")

	if err != nil {
		log.Println("error com recebimento dos dados")
	}
	aluno := model.Aluno{}
	alunos := []model.Aluno{}

	for dados.Next() {
		var nome, curso string
		var idade int
		var matricula int
		var id int

		err = dados.Scan(&nome, &idade, &matricula, &curso, &id)
		if err != nil {
			log.Println("erro nenhum dado obtido")
		} else {
			aluno.Curso = curso
			aluno.Idade = idade
			aluno.Matricula = matricula
			aluno.Nome = nome

			alunos = append(alunos, aluno)
		}
	}
	for n := range alunos {
		log.Println(alunos[n].Nome)
	}
	defer bd.Close()
}
func Greet(w http.ResponseWriter, r *http.Request) {
	log.Println("ola")
}
func main() {
	log.Println("inicializando servidor na porta:9090")
	http.HandleFunc("/show", ShowAll)
	http.HandleFunc("/", Greet)
	http.ListenAndServe(":9090", nil)
}
