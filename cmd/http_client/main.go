package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/brianvoe/gofakeit"
	"github.com/erikqwerty/http/models"
	"github.com/fatih/color"
)

const (
	baseUrl       = "http://127.0.0.1:8080"
	createPostfix = "/notes"
	getPostfix    = "/notes/%d"
)

func CreateNote() (models.Note, error) {
	note := models.NoteInfo{
		Title:    gofakeit.BeerName(),
		Context:  gofakeit.IPv4Address(),
		Author:   gofakeit.Name(),
		IsPublic: gofakeit.Bool(),
	}

	data, err := json.Marshal(note)
	if err != nil {
		return models.Note{}, err
	}

	resp, err := http.Post(baseUrl+createPostfix, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return models.Note{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return models.Note{}, err
	}

	var createdNote models.Note
	if err := json.NewDecoder(resp.Body).Decode(&createdNote); err != nil {
		return models.Note{}, err
	}

	return createdNote, nil
}

func getNote(noteID int64) (models.Note, error) {
	// Выполняем GET-запрос для получения заметки
	url := fmt.Sprintf(baseUrl+getPostfix, noteID)
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		return models.Note{}, fmt.Errorf("failed to get note: %w", err)
	}
	defer resp.Body.Close()

	// Проверка статуса 404 Not Found
	if resp.StatusCode == http.StatusNotFound {
		return models.Note{}, fmt.Errorf("note with ID %d not found", noteID)
	}

	// Проверка статуса отличного от 200 OK
	if resp.StatusCode != http.StatusOK {
		return models.Note{}, fmt.Errorf("failed to get note, status code: %d", resp.StatusCode)
	}

	// Декодируем JSON из ответа в структуру `models.Note`
	var note models.Note
	if err := json.NewDecoder(resp.Body).Decode(&note); err != nil {
		return models.Note{}, fmt.Errorf("failed to decode note data: %w", err)
	}
	return note, nil
}

func main() {
	note, err := CreateNote()
	if err != nil {
		log.Fatal("failed to create note:", err)
	}

	log.Printf(color.RedString("Note created:\n"), color.GreenString("%+v", note))

	note, err = getNote(note.ID)
	if err != nil {
		log.Fatal("failed to get note:", err)
	}

	log.Printf(color.RedString("Note info:\n"), color.GreenString("%+v", note))
}
