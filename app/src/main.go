package main

import (
    "encoding/json"
    "github.com/gorilla/mux"
    "log"
    "net/http"
)

// "Person type" (tipo um objeto)
type Materia struct {
    Content_id       string   `json:"content_id,omitempty"`
    Email 	         string   `json:"email,omitempty"`
    Comment          string   `json:"comment,omitempty"`
}

var people []Materia

// GetPeople mostra todos os contatos da variável people
func GetPeople(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(people)
}

// GetPerson mostra apenas um contato
func GetMateria(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for _, item := range people {
        if item.Content_id == params["content_id"] {
            json.NewEncoder(w).Encode(item)
        }
    }
    //json.NewEncoder(w).Encode(&Materia{})
}

// CreatePerson cria um novo contato
func CreateMateria(w http.ResponseWriter, r *http.Request) {
    var materia Materia
    err := json.NewDecoder(r.Body).Decode(&materia)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    people = append(people, materia)
    json.NewEncoder(w).Encode(people)
}


// função principal para executar a api
func main() {
    router := mux.NewRouter()
    // people = append(people, Materia{Content_id: "1", Email: "John", Comment: "Doe"})
    // people = append(people, Materia{Content_id: "1", Email: "blabla", Comment: "Dooooe"})
    // people = append(people, Materia{Content_id: "2", Email: "Koko", Comment: "Does"})
	// people = append(people, Materia{Content_id: "2", Email: "asdfasdf", Comment: "Does"})
    router.HandleFunc("/api/comment/list", GetPeople).Methods("GET")
    router.HandleFunc("/api/comment/list/{content_id}", GetMateria).Methods("GET")
    router.HandleFunc("/api/comment/new/{content_id}", CreateMateria).Methods("POST")
	router.HandleFunc("/api/comment/new/", CreateMateria).Methods("POST")
    log.Fatal(http.ListenAndServe(":8000", router))
}
