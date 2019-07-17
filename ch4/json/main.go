package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Movie represents information about a movie.
type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title: "a", Year: 2000, Color: true, Actors: []string{"x1", "x2", "x3"}},
	{Title: "b", Year: 2001, Color: false, Actors: []string{"y1", "y2", "y3"}},
}

func main() {
	fmt.Println(movies)
	fmt.Printf("%s\n", Marshal(movies))
	data := humanMarshal(movies)
	fmt.Printf("%s\n", data)

	var decodedMovies []Movie
	if err := json.Unmarshal(data, &decodedMovies); err != nil {
		log.Fatalf("ERROR in Unmarshal %s", err)
	}
	fmt.Println(decodedMovies)

	type Title struct{ Title string }
	var decodedTitles []Title
	if err := json.Unmarshal(data, &decodedTitles); err != nil {
		log.Fatalf("ERROR in Unmarshal %s", err)
	}
	fmt.Println(decodedTitles)

	type Actors struct{ Actors []string }
	var decodedActors []Actors
	if err := json.Unmarshal(data, &decodedActors); err != nil {
		log.Fatalf("ERROR in Unmarshal %s", err)
	}
	fmt.Println(decodedActors)
}

// Marshal converts a Movie slice into its json representation
// and returns it as a byte slice.
func Marshal(movies []Movie) []byte {
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON.Marshal FAILED: %s", err)
	}
	return data
}

func humanMarshal(movies []Movie) []byte {
	data, err := json.MarshalIndent(movies, "", "  ")
	if err != nil {
		log.Fatalf("JSON.Marshal FAILED: %s", err)
	}
	return data
}
