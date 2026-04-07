package main

import (
	"encoding/json"
	"net/http"
)

type Input struct {
	X float64 `json:"x"`
}

type Output struct {
	Prediction float64 `json:"prediction"`
}

func predict(w http.ResponseWriter, r *http.Request) {
	var input Input
	json.NewDecoder(r.Body).Decode(&input)

	result := 2*input.X + 3

	json.NewEncoder(w).Encode(Output{Prediction: result})
}

func main() {
	http.HandleFunc("/predict", predict)
	http.ListenAndServe(":3000", nil)
}
