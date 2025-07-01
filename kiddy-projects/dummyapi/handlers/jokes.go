package handlers

import (
	"encoding/json"
	"math/rand"
	"net/http"

	"dummyapi/models"

	"github.com/julienschmidt/httprouter"
)

// GET ALL JOKES
func GetAllJokes(w http.ResponseWriter , r *http.Request , _ httprouter.Params) {
	w.Header().Set("Content-Type" , "application/json");
	json.NewEncoder(w).Encode(models.JokesDB);
}

// GET RANDOM JOKE
func GetRandomJoke(w http.ResponseWriter , r *http.Request , _ httprouter.Params) {
	randomJoke := models.JokesDB[rand.Intn(len(models.JokesDB))]
	json.NewEncoder(w).Encode(randomJoke);
}