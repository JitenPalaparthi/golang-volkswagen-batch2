package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var (
	PORT string
)

//public static void main([]string args)

func main() {
	//args := os.Args
	PORT = os.Getenv("PORT") // take from environment if not then take can commandline argument
	if PORT == "" {
		flag.StringVar(&PORT, "port", "8080", "--port=8080")
		flag.Parse()
	}

	log.Println("Server started and running on", PORT)

	router := mux.NewRouter()
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})
	router.HandleFunc("/ping", ping)
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello World!"))
	})

	router.HandleFunc("/user/{id}", createUser)
	if err := http.ListenAndServe(":"+PORT, router); err != nil {
		Fatal(err.Error())
	}
}
func Fatal(elems ...any) {
	fmt.Println(elems...)
	os.Exit(1)
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}

func createUser(w http.ResponseWriter, r *http.Request) {
	//if r.Method != http.MethodPost {
	if r.Method != "POST" {
		http.Error(w, "unimplemented http method", http.StatusNotImplemented)
		return
	}
	status := r.URL.Query().Get("status")

	vars := mux.Vars(r)
	id := vars["id"]

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	user := new(User)

	err = json.Unmarshal(bytes, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user.Id = id
	if status != "" {
		user.Status = status
	} else {
		user.Status = "inactive"
	}

	mbytes, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = WriteToFile("data.txt", string(mbytes))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(string(mbytes)))
}

type User struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Status string `json:"status"`
}

func WriteToFile(fileName string, data string) error {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write([]byte(data))
	if err != nil {
		return err
	}
	return nil
}

/*

curl --location 'http://localhost:8080/user' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id":102,
    "name":"Priya",
    "email":"Priya@outlook.com"
}'*/
