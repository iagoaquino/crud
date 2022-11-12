package controler

import (
	"html/template"
	"log"
	"net/http"

	"server.com/conect"
	"server.com/model"
)

var tmpl = template.Must(template.ParseGlob("view/*"))

func Find(w http.ResponseWriter, r *http.Request) {
	bd := conect.Conect()
	if r.Method == "POST" {
		pesquisa := r.FormValue("busca")
		aluno := model.Aluno{}
		alunos := []model.Aluno{}
		dados, err := bd.Query("SELECT * FROM aluno WHERE nome LIKE '%"+pesquisa+"%' OR curso LIKE '%"+pesquisa+"%' OR idade=? OR matricula=?", pesquisa, pesquisa)
		if err != nil {
			log.Fatal("erro com o sql:", err)
		}
		for dados.Next() {
			dados.Scan(&aluno.Nome, &aluno.Idade, &aluno.Matricula, &aluno.Curso, &aluno.Id)
			alunos = append(alunos, aluno)
		}
		tmpl.ExecuteTemplate(w, "Show", alunos)
		defer bd.Close()
	}
}
func ShowAll(w http.ResponseWriter, r *http.Request) {
	bd := conect.Conect()

	dados, err := bd.Query("SELECT * FROM aluno")

	if err != nil {
		log.Fatal("error com recebimento dos dados", err)
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
func Delete(w http.ResponseWriter, r *http.Request) {
	var id = r.URL.Query().Get("id")
	bd := conect.Conect()
	dados, err := bd.Query("SELECT * FROM aluno WHERE id =" + id)
	if err != nil {
		log.Println("valor não encontrado")
	} else {
		var nome, curso string
		var idade, matricula, idshow int
		aluno := model.Aluno{}
		alunos := []model.Aluno{}
		for dados.Next() {
			dados.Scan(&nome, &idade, &matricula, &curso, &idshow)
			aluno.Nome = nome
			aluno.Curso = curso
			aluno.Matricula = matricula
			aluno.Idade = idade
			aluno.Id = idshow
			alunos = append(alunos, aluno)
		}
		comando, err := bd.Prepare("DELETE FROM aluno WHERE id = ?")
		if err != nil {
			log.Println("erro ao deletar")
		} else {
			comando.Exec(id)
		}
		tmpl.ExecuteTemplate(w, "delete", alunos)
	}
	defer bd.Close()
}
func Edit(w http.ResponseWriter, r *http.Request) {
	bd := conect.Conect()
	aluno := model.Aluno{}
	alunos := []model.Aluno{}
	var idade, matricula, id int
	var nome, curso string
	dados, err := bd.Query("SELECT * FROM aluno WHERE id =?", r.FormValue("id"))
	if err != nil {
		log.Println("não encontrado")
	}
	for dados.Next() {
		dados.Scan(&nome, &idade, &matricula, &curso, &id)
		aluno.Nome = nome
		aluno.Curso = curso
		aluno.Matricula = matricula
		aluno.Idade = idade
		aluno.Id = id
		alunos = append(alunos, aluno)
	}
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		idade := r.FormValue("idade")
		matricula := r.FormValue("matricula")
		curso := r.FormValue("curso")
		idE := r.FormValue("id")
		comando, err := bd.Prepare("UPDATE aluno SET nome=? , idade=? , matricula=? , curso=? WHERE id=?")
		if err != nil {
			log.Println("error com os dados")
		} else {
			comando.Exec(nome, idade, matricula, curso, idE)
		}
	}
	dados, err = bd.Query("SELECT * FROM aluno WHERE id =?", r.FormValue("id"))
	if err != nil {
		log.Println("não encontrado")
	}
	for dados.Next() {
		dados.Scan(&nome, &idade, &matricula, &curso, &id)
		aluno.Nome = nome
		aluno.Curso = curso
		aluno.Matricula = matricula
		aluno.Idade = idade
		aluno.Id = id
		alunos = append(alunos, aluno)
	}
	tmpl.ExecuteTemplate(w, "edit", alunos)
	defer bd.Close()
}
func Insert(w http.ResponseWriter, r *http.Request) {
	bd := conect.Conect()
	if r.Method == "POST" {
		nome := r.FormValue("nomeI")
		idade := r.FormValue("idadeI")
		matricula := r.FormValue("matriculaI")
		curso := r.FormValue("cursoI")
		comando, err := bd.Prepare("INSERT INTO aluno(nome,idade,matricula,curso) values(?,?,?,?)")
		if err != nil {
			log.Println("error com os dados")
		} else {
			comando.Exec(nome, idade, matricula, curso)
		}
		http.Redirect(w, r, "/show", 301)
		defer bd.Close()
	}
}
