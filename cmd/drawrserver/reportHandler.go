package main

import (
	"net/http"

	"github.com/pressly/chi"
)

func reportRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/", statReport)
	r.Get("/db", dbReport)
	return r
}

func statReport(w http.ResponseWriter, r *http.Request) {
	var out string

	out = out + "drawr backend:\n"
	out = out + "connected clients:\n"
	for id, hub := range wsHubs {
		out = out + "> " + id + ":"
		ls := hub.hub.ListConnections()
		for _, s := range ls {
			out = out + s + "\n"
		}
		out = out + "\n"
	}

	w.Write([]byte(out))
}

func dbReport(w http.ResponseWriter, r *http.Request) {
	var out string

	out = out + "drawr backend:\n"
	out = out + "database:\n"

	out = out + dbClient.Stats() + "\n\n"
	out = out + dbClient.String() + "\n\n"

	w.Write([]byte(out))
}
