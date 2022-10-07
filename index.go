package main

import (
	"log"
	"net/http"
	"text/template"

	"server.com/conect"
	"server.com/model"
)

var tmpl = template.Must(template.ParseGlob("view/*"))

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
			aluno.Id = id

			alunos = append(alunos, aluno)
		}
	}
	tmpl.ExecuteTemplate(w, "Show", alunos)
	defer bd.Close()
}
func Greet(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "greet", "olá")
}
func main() {
	log.Println("inicializando servidor na porta:9090")
	http.HandleFunc("/show", ShowAll)
	http.HandleFunc("/", Greet)
	http.ListenAndServe(":9090", nil)
}
