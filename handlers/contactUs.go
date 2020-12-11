package handlers

import (
	"html/template"
	"net/http"
)

//ContactUs is ...
func ContactUs(w http.ResponseWriter, r *http.Request) {
	tmpFiles := []string{
		"view/contact_view/contact.html",
		"view/contact_view/header.html",
		"view/footer.html",
	}

	tmpl := template.Must(template.ParseFiles(tmpFiles...))

	err := tmpl.ExecuteTemplate(w, "view", nil)
	if err != nil {
		panic(err)
	}
}
