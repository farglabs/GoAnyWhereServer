package status

import (
	"html/template"
	"net/http"
	"log"
)

func S_404(w http.ResponseWriter)  {
	type Todo struct {
		Title string
		Done  bool
	}

	type TodoPageData struct {
		PageTitle string
		Todos     []Todo
	}
	
	tmpl := template.Must(template.ParseFiles("template/status/layout.html"))
	data := TodoPageData{
		PageTitle: "My TODO list",
		Todos: []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		},
	}
	err :=tmpl.Execute(w, data)
	log.Print(err)
}