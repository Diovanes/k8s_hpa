package main

import (
	"fmt"
	"math"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloServer)
	http.ListenAndServe(":8000", nil)
}

func executaCalculo() string {
	var x float64 = 0.0001
	for i := 1; i < 100000000; i++ {
		x += math.Sqrt(x)
	}
	return "Code.education Rocks!"
}

func helloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", executaCalculo())
}
