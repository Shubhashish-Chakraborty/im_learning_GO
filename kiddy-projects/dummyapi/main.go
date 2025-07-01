package main

import (
	"dummyapi/handlers"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	PORT := "3000";
	r := httprouter.New();

	// // use this if returning json on any end point 
	// type Response struct {
	// 	Message string `json:"message"`
	// }

	// Routes Handling:

	r.GET("/" , func (w http.ResponseWriter , r * http.Request , _ httprouter.Params) {
		// w.Header().Set("Content-Type" , "application/json");
		// json.NewEncoder(w).Encode(Response{Message: "Server is Up"})

		// so either pass json or simple plain text:

		// w.Header().Set("Content-Type" , "text/plain");
		// fmt.Fprintln(w , "HELLO THE SERVER IS UP");

		// or a web page (html) :

		w.Header().Set("Content-Type" , "text/html");
		fmt.Fprintf(w , "Server is Up!")
		
	})

	r.GET("/jokes" , handlers.GetAllJokes);
	r.GET("/jokes/random" , handlers.GetRandomJoke);

	r.POST("/jokes/add" , handlers.AddJoke);
	
	// Serve static files (like Express's express.static())
	r.ServeFiles("/public/*filepath", http.Dir("public"));

	log.Printf("Server running at http://localhost:%v" , PORT);
	log.Fatal(http.ListenAndServe(":" + PORT , r));
	
}