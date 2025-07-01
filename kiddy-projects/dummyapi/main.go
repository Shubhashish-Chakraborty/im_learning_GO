package main

import (
	"log"
	"dummyapi/handlers"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	PORT := "3000";
	r := httprouter.New();

	// Routes Handling:
	
	r.GET("/jokes" , handlers.GetAllJokes);
	r.GET("/jokes/random" , handlers.GetRandomJoke);
	
	// Serve static files (like Express's express.static())
	r.ServeFiles("/public/*filepath", http.Dir("public"));

	log.Printf("Server running at http://localhost:%v" , PORT);
	log.Fatal(http.ListenAndServe(":" + PORT , r));
	
}