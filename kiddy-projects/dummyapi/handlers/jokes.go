package handlers

import (
	"encoding/json"
	"math/rand"
	"net/http"

	"dummyapi/models"

	"github.com/julienschmidt/httprouter"
)

// GET ALL JOKES
func GetAllJokes(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.JokesDB)
}

// GET RANDOM JOKE
func GetRandomJoke(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	randomJoke := models.JokesDB[rand.Intn(len(models.JokesDB))]
	json.NewEncoder(w).Encode(randomJoke)
}

// MAKE A POST REQUEST AND ADD A JOKE!
func AddJoke(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var newJoke models.Joke;
	// handle errors:
	err := json.NewDecoder(r.Body).Decode(&newJoke)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// but if everythings okay, pass the execution:

	newJoke.JokeNumber = len(models.JokesDB) + 1;

	models.JokesDB = append(models.JokesDB, newJoke)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newJoke)
}
