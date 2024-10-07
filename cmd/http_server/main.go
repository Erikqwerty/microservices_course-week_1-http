package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

const (
	baseURL       = "127.0.0.1:8080"
	createPostfix = "/notes"
	getPostfix    = "/notes/{id}"
)

func main() {
	// создаем роутер chi
	r := chi.NewRouter()

	// метод POST по url /notes, обработчик CreateNote
	r.Post(createPostfix, CreateNote)

	// метод GET по url /notes/%d, обработчик GetNote
	r.Get(getPostfix, GetNote)

	fmt.Printf("Server listen: %v \n", baseURL)

	// Запусе http сервера
	err := http.ListenAndServe(baseURL, r)
	if err != nil {
		log.Fatal(err)
	}
}
