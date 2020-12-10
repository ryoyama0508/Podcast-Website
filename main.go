package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/mmcdole/gofeed"
	_ "github.com/mmcdole/gofeed"
)

func index(w http.ResponseWriter, r *http.Request) {
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

func rssFeed(w http.ResponseWriter, r *http.Request) {
	url := "https://anchor.fm/s/31a64358/podcast/rss"

	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(url)
	if err != nil {
		panic(err)
	}

	items := feed.Items

	for _, item := range items {
		if item.Description[:3] == "<p>" {
			item.Description = item.Description[3 : len(item.Description)-5]
		}
	}

	if len(items) > 10 {
		items = items[:10]
	}

	tmpFiles := []string{
		"view/episode_view/episode.html",
		"view/episode_view/header.html",
		"view/footer.html",
	}

	tmpl := template.Must(template.ParseFiles(tmpFiles...))

	err = tmpl.ExecuteTemplate(w, "view", items)
	if err != nil {
		panic(err)
	}
}

func rssFeed2(w http.ResponseWriter, r *http.Request) {
	url := "https://anchor.fm/s/31a64358/podcast/rss"

	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(url)
	if err != nil {
		panic(err)
	}

	items := feed.Items

	for _, item := range items {
		if item.Description[:3] == "<p>" {
			item.Description = item.Description[3 : len(item.Description)-5]
		}
	}

	tmpFiles := []string{
		"view/episode_view/episode.html",
		"view/episode_view/header.html",
		"view/footer.html",
	}

	tmpl := template.Must(template.ParseFiles(tmpFiles...))

	err = tmpl.ExecuteTemplate(w, "view", items)
	if err != nil {
		panic(err)
	}
}

func blog(w http.ResponseWriter, r *http.Request) {
	tmpFiles := []string{
		"view/blog_view/blog.html",
		"view/blog_view/header.html",
		"view/blog_view/footer.html",
	}

	tmpl := template.Must(template.ParseFiles(tmpFiles...))

	err := tmpl.ExecuteTemplate(w, "view", nil)
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("The Server runs with http://localhost:8080/")

	files := http.FileServer(http.Dir("public/"))
	http.Handle("/public/", http.StripPrefix("/public/", files))
	http.Handle("/public/pic", http.StripPrefix("/public/", files))

	http.HandleFunc("/", index)

	http.HandleFunc("/episode", rssFeed)
	http.HandleFunc("/episodeAll", rssFeed2)
	//don't know how to put number permanent link on every episode.

	http.HandleFunc("/blog", blog)

	http.ListenAndServe(":8080", nil)
}
