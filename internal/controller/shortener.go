package controller

import (
	"context"
	"fmt"
	"html/template"
	"net/http"

	"github.com/jackc/pgx/v4"
	"golang.org/x/exp/rand"
)

type Connector interface {
	Connect() (*pgx.Conn, error)
}

type Shortener struct {
	connector Connector
}

func NewShortener(c Connector) *Shortener {
	return &Shortener{
		connector: c,
	}
}

func (s *Shortener) HandleIndex(w http.ResponseWriter, r *http.Request) {
	templFile := "templates/index.html"
	t, err := template.ParseFiles(templFile)
	if err != nil {
		panic(err)
	}
	t.Execute(w, nil)
}

func (s *Shortener) HandleShorten(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handling shorten")
	r.ParseForm()
	url := r.Form.Get("url")

	if url == "" {
		http.Error(w, "<p style=\"color: red;\">URL is required</p>", http.StatusBadRequest)
		return
	}

	if _, err := http.Get(url); err != nil {
		http.Error(w, "<p style=\"color: red;\">Invalid URL</p>", http.StatusBadRequest)
		return
	}

	conn, err := s.connector.Connect()
	if err != nil {
		panic(err)
	}
	defer conn.Close(context.Background())

	var existingURL string
	var existingId string
	_ = conn.QueryRow(context.Background(), "SELECT id, url FROM urls WHERE url = $1", url).Scan(&existingId, &existingURL)
	if existingURL != "" {
		fmt.Fprint(w, "<p>http://localhost:8089/"+existingId+"</p>")
		return
	}

	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	var id string
	idNotFound := true
	for idNotFound {
		var existingId string
		b := make([]rune, 3)
		for i := range b {
			b[i] = letters[rand.Intn(len(letters))]
		}
		id = string(b)
		err := conn.QueryRow(context.Background(), "SELECT id FROM urls WHERE id = $1", id).Scan(&existingId)
		if err != nil {
			idNotFound = false
		}
	}

	_, err = conn.Exec(context.Background(), "INSERT INTO urls VALUES ($1, $2, $3)", id, url, id)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "<p style=\"color: red;\">Failed to insert URL</p>", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, "<p>http://localhost:8089/"+id+"</p>")
}

func (s *Shortener) HandleRedirect(w http.ResponseWriter, r *http.Request) {
	conn, err := s.connector.Connect()
	if err != nil {
		panic(err)
	}
	defer conn.Close(context.Background())

	id := r.PathValue("id")

	var url string
	err = conn.QueryRow(context.Background(), "SELECT url FROM urls WHERE id = $1", id).Scan(&url)
	if err != nil {
		http.Redirect(w, r, "/404", http.StatusSeeOther)
		return
	}
	http.Redirect(w, r, url, http.StatusPermanentRedirect)
}

func (s *Shortener) HandleError(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/404.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, nil)
}
