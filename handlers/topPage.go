package handlers

import (
	"html/template"
	"net/http"
)

// TopPage is ...
func TopPage(w http.ResponseWriter, r *http.Request) {
	tmpFiles := []string{
		"view/view.html",
		"view/footer.html",
		"view/header.html",
	}

	tmpl := template.Must(template.ParseFiles(tmpFiles...))

	err := tmpl.ExecuteTemplate(w, "view", nil)
	if err != nil {
		panic(err)
	}
}
