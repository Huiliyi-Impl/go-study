package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  int      `json:"price"`
	Actors []string `json:"actors"`
}

func main() {
	defer fmt.Println("defer")
	movie := Movie{
		Title:  "test",
		Year:   2019,
		Price:  100,
		Actors: []string{"a", "b"},
	}
	//结构体 ---》json
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("json marshal error", err)
		return
	}
	fmt.Printf("jsonStr = %s\n", jsonStr)

	myMovie := Movie{}
	newErr := json.Unmarshal(jsonStr, &myMovie)
	if newErr != nil {
		fmt.Println("json unmarshal error", newErr)
		return
	}
	fmt.Printf("my_movie = %v\n", myMovie)
}
