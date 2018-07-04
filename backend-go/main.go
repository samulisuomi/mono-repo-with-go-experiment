package main

import (
	"encoding/json"
  "log"
	"net/http"
)

type Hello struct {
	Id int `json:"id"`
	Description	string `json:"description"`
}

func main() {
  http.HandleFunc("/hellos", func(w http.ResponseWriter, r *http.Request) {
		hellos := []Hello{{1, "Hello"}, {2, "from"}, {3, "Go"}, {4, "API!"}}

		hellosJson, err := json.Marshal(hellos)

		if err != nil {
			log.Fatal("Cannot encode to JSON ", err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(hellosJson)
	})

	log.Println("Running...")

	log.Fatal(http.ListenAndServe(":8080", nil))

}
