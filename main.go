package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	router := mux.NewRouter()
	api := router.PathPrefix("/dad").Subrouter()
	api.HandleFunc("/", joke).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func joke(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(getJoke())
}

func getJoke() string {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://icanhazdadjoke.com/", nil)
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("an error occurred")
		os.Exit(0)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var data map[string]interface{}
	json.Unmarshal(body, &data)
	return data["joke"].(string)
}