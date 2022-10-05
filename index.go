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
		log.Println("error")
	}
	aluno := model.Aluno{}
	alunos := []model.Aluno{}

	for dados.Next() {
		var nome, curso string
		var idade int
		var matricula int

		err = dados.Scan(&nome, &idade, &matricula, &curso)
		if err != nil {
			log.Println("erro")
		}
		aluno.Curso = curso
		aluno.Idade = idade
		aluno.Matricula = matricula
		aluno.Nome = nome

		alunos = append(alunos, aluno)
	}

}
func main() {
	log.Println("inicializando servidor na porta:9090")
	http.HandleFunc("/show", ShowAll)
}
