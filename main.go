package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/kaloterrace_site/handlers"
)

func main() {
	fmt.Println("The Server runs with http://localhost:8080/")

	files := http.FileServer(http.Dir("public/"))
	http.Handle("/public/", http.StripPrefix("/public/", files))
	http.Handle("/public/pic", http.StripPrefix("/public/", files))

	http.HandleFunc("/", handlers.TopPage)

	/* http.HandleFunc("/episode", handlers.RecentEpisodes)
	http.HandleFunc("/episodeAll", handlers.AllEpisodes) */

	http.HandleFunc("/blog", handlers.Blog)

	http.HandleFunc("/contactUs", handlers.ContactUs)

	/* http.ListenAndServe(":8080", nil) */
	server := http.Server{
		Addr: ":" + os.Getenv("PORT"),
	}
	server.ListenAndServe()
}
