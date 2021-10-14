package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/ian-kent/go-log/log"
	"net/http"
	"strconv"
)

var movies = Movies{
	Movie{"Sin limites", 2013, "Desconocido"},
	Movie{"Batman Begins", 2008, "Desconocido"},
	Movie{"Rapido y furioso", 2005, "Desconocido"},
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo desde mi servidor web con GO")
}

func MovieList(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(movies)
}

func MovieShow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie_id, _ := strconv.Atoi(params["id"])

	json.NewEncoder(w).Encode(movies[movie_id-1])

}

func MovieAdd(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var movie_data Movie
	err := decoder.Decode(&movie_data)

	if err != nil {
		panic(err)
	}

	defer r.Body.Close()

	log.Println(movie_data)
	movies = append(movies, movie_data)
}
