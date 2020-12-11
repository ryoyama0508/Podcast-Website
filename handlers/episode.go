package handlers

import (
	"html/template"
	"net/http"

	"github.com/mmcdole/gofeed"
)

func episodeGet() []*gofeed.Item {
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

	return items
}

// RecentEpisodes is for serving 10 recent episodes at html
func RecentEpisodes(w http.ResponseWriter, r *http.Request) {
	items := episodeGet()

	if len(items) > 10 {
		items = items[:10]
	}
	makeTemplate(w, r, items)
}

// AllEpisodes is for serving all episodes at html
func AllEpisodes(w http.ResponseWriter, r *http.Request) {
	items := episodeGet()
	makeTemplate(w, r, items)
}

func makeTemplate(w http.ResponseWriter, r *http.Request, items []*gofeed.Item) {
	tmpFiles := []string{
		"view/episode_view/episode.html",
		"view/episode_view/header.html",
		"view/footer.html",
	}

	tmpl := template.Must(template.ParseFiles(tmpFiles...))

	err := tmpl.ExecuteTemplate(w, "view", items)
	if err != nil {
		panic(err)
	}
}

//create generateHTML func after figure it out how to create header.html template
//which has different color depending on which url called
