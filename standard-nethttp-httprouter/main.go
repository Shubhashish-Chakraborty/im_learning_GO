package main

import (
	"fmt";
	"net/http";

	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New();
	PORT := "6969";
	
	r.GET("/" , func(rw http.ResponseWriter , r *http.Request , p httprouter.Params) {
		fmt.Fprintln(rw , "Server is Up, My first Go lang Backend!");
	})

	fmt.Printf("Server is Running at http://localhost:%v \n" , PORT);
	http.ListenAndServe(":" + PORT , r);
}