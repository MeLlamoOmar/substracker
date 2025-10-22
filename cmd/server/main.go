package main

import (
	"SubsTracker/internal/Renderer"
	"fmt"
	"net/http"
)

type Contact struct {
	Name string
	Email string
}

func NewContact(name, email string) Contact {
	return Contact{
		Name: name,
		Email: email,
	}
}

type Contacts = []Contact

type Data struct {
	Contacts Contacts
}

func NewData() Data {
	return Data{
		Contacts: Contacts{
			NewContact("Omar", "omar@example.com"),
			NewContact("Alice", "alice@example.com"),
		},
	}
}

func (d *Data) hasEmail(email string) bool {
	for _, c := range d.Contacts {
		if c.Email == email {
			return true
		}
	}
	return false
}

type FormData struct {
	Values map[string]string
	Errors map[string]string
}

func NewFormData() FormData {
	return FormData{
		Values: make(map[string]string),
		Errors: make(map[string]string),
	}
}

type Page struct {
	Data Data
	Form FormData
}

func NewPage() Page {
	return Page{
		Data: NewData(),
		Form: NewFormData(),
	}
}

func main() {
	router := http.NewServeMux()
	renderer := Renderer.NewRenderer()
	page := NewPage()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		renderer.Render(w, "base", page)
	})

	router.HandleFunc("GET /home", func(w http.ResponseWriter, r *http.Request) {
		renderer.Render(w, "dashboard", page)
	})

	router.HandleFunc("GET /subscriptions", func(w http.ResponseWriter, r *http.Request) {
		renderer.Render(w, "mySubs", page)
	})

	fmt.Println("Listening on port :8080")
	http.ListenAndServe(":8080", router)
}