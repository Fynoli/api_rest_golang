package main

/*Fijate que si los nombres de los campos no empiezan con mayusculas
el encoder de JSON se chotea y te manda los campos vacios, y lo mismo el encoder
Una cagada*/
type Movie struct {
	Name     string `json:"name"`
	Year     int    `json:"year"`
	Director string `json:"director"`
}

type Movies []Movie
