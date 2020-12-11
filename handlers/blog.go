package handlers

import (
	"html/template"
	"net/http"
)

// Blog is ...
func Blog(w http.ResponseWriter, r *http.Request) {
	tmpFiles := []string{
		"view/blog_view/blog.html",
		"view/blog_view/header.html",
		"view/footer.html",
	}

	tmpl := template.Must(template.ParseFiles(tmpFiles...))

	err := tmpl.ExecuteTemplate(w, "view", nil)
	if err != nil {
		panic(err)
	}
}
