package main

import(
"encoding/json"
"log"
"net/http"
"github.com/gorilla/mux"
)

type Person struct {
    ID        string   `json:"id,omitempty"`
    Firstname string   `json:"firstname,omitempty"`
    Lastname  string   `json:"lastname,omitempty"`

}

var people []Person

func GetPersonEndpoint(w http.ResponseWriter, req *http.Request){
	params := mux.Vars(req)
	for _,item := range people{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for index, item := range people {
        if item.ID == params["id"] {
            people = append(people[:index], people[index+1]...)
            break
    }
    json.NewEncoder(w).Encode(people)
}
}

func main() {
    router := mux.NewRouter()
    people = append(people, Person{ID: "1", Firstname: "Nic", Lastname: "Raboy"})
    people = append(people, Person{ID: "2", Firstname: "Maria", Lastname: "Raboy"})
    router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
    
    log.Fatal(http.ListenAndServe(":3000", router))
}