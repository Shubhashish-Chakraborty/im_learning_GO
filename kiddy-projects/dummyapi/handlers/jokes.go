package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"dummyapi/models"
)

// GET ALL JOKES
func GetAllJokes(w http.ResponseWriter , r *http.Request , _ httprouter.Params) {
	w.Header().Set("Content-Type" , "application/json");
	json.NewEncoder(w).Encode(models.JokesDB);
}