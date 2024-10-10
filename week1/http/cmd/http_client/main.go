package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/brianvoe/gofakeit"
	"github.com/fatih/color"
	"github.com/pkg/errors"
)

const (
	baseURL       = "http://localhost:8081"
	createPostFix = "/notes"
	getPostFix    = "/notes/%d"
)

// NoteInfo содержит информацию о заметке.
type NoteInfo struct {
	Title    string `json:"title"`
	Context  string `json:"context"`
	Author   string `json:"author"`
	IsPublic bool   `json:"is_public"`
}

// Note сущность заметки
type Note struct {
	ID        int64     `json:"id"`
	Info      NoteInfo  `json:"info"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func createNote() (Note, error) {
	note := NoteInfo{
		Title:    gofakeit.BeerName(),
		Context:  gofakeit.IPv4Address(),
		Author:   gofakeit.Name(),
		IsPublic: gofakeit.Bool(),
	}

	data, err := json.Marshal(note)
	if err != nil {
		return Note{}, err
	}

	resp, err := http.Post(baseURL+createPostFix, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return Note{}, err
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}()

	if resp.StatusCode != http.StatusCreated {
		return Note{}, err
	}

	var createdNote Note
	if err := json.NewDecoder(resp.Body).Decode(&createdNote); err != nil {
		return Note{}, err
	}

	return createdNote, nil
}

func getNote(id int64) (Note, error) {
	fmt.Println(fmt.Sprintf(baseURL+getPostFix, id))
	resp, err := http.Get(fmt.Sprintf(baseURL+getPostFix, id))
	if err != nil {
		log.Fatal("Failed to get note", err)
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}()

	if resp.StatusCode == http.StatusNotFound {
		return Note{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return Note{}, errors.Errorf("failed to get note: %d", resp.StatusCode)
	}

	var note Note
	if err := json.NewDecoder(resp.Body).Decode(&note); err != nil {
		return Note{}, err
	}

	return note, nil
}

func main() {
	note, err := createNote()
	if err != nil {
		log.Fatal("Failed to create client:", err)
	}
	log.Printf(color.RedString("Note created:\n"), color.GreenString("$+v", note))

	note, err = getNote(note.ID)
	if err != nil {
		log.Fatal("Failed to get note:", err)
	}
	log.Printf(color.RedString("Note info got:\n"), color.GreenString("$+v", note))
}
