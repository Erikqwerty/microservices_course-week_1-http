package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/erikqwerty/http/models"
	"github.com/go-chi/chi"
)

// карта с мьютексом для безопасного паралельного доступа к данным
var notes = &models.SyncMap{
	Elems: make(map[int64]*models.Note),
	M:     &sync.RWMutex{}, // мьютекс блокирующий запись пока кто-то пишет, но дает читать данные всем
}

func CreateNote(w http.ResponseWriter, r *http.Request) {

	// декодируем полученный json в info
	info := &models.NoteInfo{}
	if err := json.NewDecoder(r.Body).Decode(info); err != nil {
		http.Error(w, "Failed to decode note data", http.StatusBadRequest)
		return
	}

	// псевдорандом
	rand.Seed(time.Now().UnixNano())
	now := time.Now()

	note := &models.Note{
		ID:        rand.Int63(),
		Info:      *info,
		CreatedAt: now,
		UpdatedAt: now,
	}

	// устанавливаем в ответ заголовки тип данных json, код ответа 201
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	// кодируем данные в json
	if err := json.NewEncoder(w).Encode(note); err != nil {
		http.Error(w, "Failed to encode note data", http.StatusInternalServerError)
		return
	}

	// Мьютекс блокируется для записи, заметка добавляется в карту, а затем мьютекс разблокируется.
	notes.M.Lock()
	defer notes.M.Unlock()
	notes.Elems[note.ID] = note
}

func GetNote(w http.ResponseWriter, r *http.Request) {
	// Извлекаем параметр id и выводим для отладки
	noteID := chi.URLParam(r, "id")

	// Преобразуем строку в int64 и выводим для отладки
	id, err := parseNoteID(noteID)
	if err != nil {
		http.Error(w, "Invalid note ID", http.StatusBadRequest)
		return
	}

	// Блокируем чтение карты и проверяем наличие заметки по ID
	notes.M.RLock()
	defer notes.M.RUnlock()
	note, ok := notes.Elems[id]
	if !ok {
		fmt.Println("Note not found for ID:", id)
		http.Error(w, "Note not found", http.StatusNotFound)
		return
	}

	// Устанавливаем заголовки и кодируем заметку в JSON для ответа
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(note); err != nil {
		http.Error(w, "Failed to encode note data", http.StatusInternalServerError)
	}
}

func parseNoteID(idstr string) (int64, error) {
	id, err := strconv.ParseInt(idstr, 10, 64)
	if err != nil {
		return 0, err
	}
	return id, err
}
