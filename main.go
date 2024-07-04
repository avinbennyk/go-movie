package main

import{
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
}

type movie struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`

}

type Director struct{
	Firstname string `json:"fname"`
	Lastname string `json:"lname"`
}

var movies []Movie

func main(){
	r :=mux.NewRouter()


	movies = append(movies, Movie{ID:"1", Isbn:"438227", Title:"Movie One", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID:"2", Isbn:"45455", Title:"Movie Two", Director: &Director{Firstname: "Steve", Lastname: "Smith"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", getMovies).Methods("POST")
	r.HandleFunc("/movies/{id}", getMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", getMovie).Methods("DELETE")

	fmt.Println("Starting server at port 8000 \n")
	log.Fatal(http.ListenAndServe(":8000",r))

}


