package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/ian-kent/go-log/log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

func getSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost")

	if err != nil {
		panic(err)
	}

	return session
}

func responseMovie(w http.ResponseWriter, status int, results Movie) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(results)
}

func responseMovies(w http.ResponseWriter, status int, results []Movie) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(results)
}

var collection = getSession().DB("curso_go").C("movies")

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo desde mi servidor web con GO")
}

func MovieList(w http.ResponseWriter, r *http.Request) {
	var results []Movie
	err := collection.Find(nil).Sort("-_id").All(&results)

	if err != nil {
		log.Fatal(err)
	}

	responseMovies(w, 200, results)
}

func MovieShow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie_id := params["id"]

	if !bson.IsObjectIdHex(movie_id) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(movie_id)

	var result Movie
	err := collection.FindId(oid).One(&result) /*Por mas que individualicemos con con la id One es necesario para especificar
	que queremos un solo objeto json y no un array de un solo elemento*/

	if err != nil {
		w.WriteHeader(404)
		return
	}

	responseMovie(w, 200, result)

}

func MovieUpdate(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie_id := params["id"]

	if !bson.IsObjectIdHex(movie_id) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(movie_id)

	decoder := json.NewDecoder(r.Body)

	var movie_data Movie
	err := decoder.Decode(&movie_data)

	if err != nil {
		panic(err)
		w.WriteHeader(500)
		return
	}

	defer r.Body.Close()

	document := bson.M{"_id": oid}            //Recupero el documento bson (binary json) correspondiente a la id pasada
	change := bson.M{"$set": movie_data}      //Creo un documento bson con los datos de movie_data
	err = collection.Update(document, change) //Piso el documento original con el nuevo en base sin cambiar el id

	if err != nil {
		w.WriteHeader(404)
		return
	}

	responseMovie(w, 200, movie_data)

}

type ResponseMessage struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func MovieDelete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie_id := params["id"]

	if !bson.IsObjectIdHex(movie_id) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(movie_id)

	err := collection.RemoveId(oid)

	if err != nil {
		w.WriteHeader(404)
		return
	}

	result := ResponseMessage{200, "La película con ID= " + movie_id + " ha sido borrada exitosamente"}
	w.Header().Set("ContentType", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(result)

}

func MovieAdd(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var movie_data Movie
	err := decoder.Decode(&movie_data)

	if err != nil {
		panic(err)
	}

	defer r.Body.Close()

	err = collection.Insert(movie_data)

	if err != nil {
		w.WriteHeader(500)
		return
	}

	responseMovie(w, 200, movie_data)
}
