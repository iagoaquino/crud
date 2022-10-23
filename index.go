package main

import (
	"database/sql" // Pacote Database SQL para realizar Query
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql" // Driver Mysql para Go
)

type Aluno struct {
	Nome      string
	Idade     int
	Matricula int
	Curso     string
	Id        int
}

func Conect() (bd *sql.DB) {
	driver := "mysql"
	name := "alunos"

	bd, err := sql.Open(driver, "root:teste123@tcp(172.17.0.2:3306)/"+name)

	if err != nil {
		log.Print("error com a conexão do banco")
	} else {
		log.Print("conectou ao banco")
	}
	return bd
}

var tmpl = template.Must(template.ParseGlob("view/*"))

func ShowAll(w http.ResponseWriter, r *http.Request) {
	bd := Conect()

	dados, err := bd.Query("SELECT * FROM aluno")

	if err != nil {
		log.Println("error com recebimento dos dados")
	}
	aluno := Aluno{}
	alunos := []Aluno{}
	log.Println("cheguei aqui")
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
	bd := Conect()
	dados, err := bd.Query("SELECT * FROM aluno WHERE id =" + id)
	if err != nil {
		log.Println("valor não encontrado")
	} else {
		var nome, curso string
		var idade, matricula, idshow int
		aluno := Aluno{}
		alunos := []Aluno{}
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
			log.Println("deletado com sucesso")
		}
		tmpl.ExecuteTemplate(w, "delete", alunos)
	}
	defer bd.Close()
}
func Edit(w http.ResponseWriter, r *http.Request) {
	bd := Conect()
	aluno := Aluno{}
	alunos := []Aluno{}
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
	bd := Conect()
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
func main() {
	log.Println("inicializando servidor na porta:9090")
	http.HandleFunc("/show", ShowAll)
	http.HandleFunc("/delete", Delete)
	http.HandleFunc("/edit", Edit)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/", Greet)
	http.ListenAndServe(":9090", nil)
}
