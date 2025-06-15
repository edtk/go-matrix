package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strconv"

	"github.com/edtk/go-matrix/matrix"
)

// handleMatrixOperation is a helper function to handle matrix operations
func handleMatrixOperation(w http.ResponseWriter, r *http.Request, operation func(*matrix.Matrix) interface{}) {
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, fmt.Sprintf("error reading file: %s", err.Error()), http.StatusBadRequest)
		return
	}
	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		http.Error(w, fmt.Sprintf("error parsing CSV: %s", err.Error()), http.StatusBadRequest)
		return
	}

	m, err := matrix.New(records)
	if err != nil {
		http.Error(w, fmt.Sprintf("error processing matrix: %s", err.Error()), http.StatusBadRequest)
		return
	}

	result := operation(m)
	fmt.Fprint(w, result)
}

func main() {
	// Echo endpoint
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		handleMatrixOperation(w, r, func(m *matrix.Matrix) interface{} {
			return m.Echo()
		})
	})

	// Invert endpoint
	http.HandleFunc("/invert", func(w http.ResponseWriter, r *http.Request) {
		handleMatrixOperation(w, r, func(m *matrix.Matrix) interface{} {
			return m.Invert()
		})
	})

	// Flatten endpoint
	http.HandleFunc("/flatten", func(w http.ResponseWriter, r *http.Request) {
		handleMatrixOperation(w, r, func(m *matrix.Matrix) interface{} {
			return m.Flatten()
		})
	})

	// Sum endpoint
	http.HandleFunc("/sum", func(w http.ResponseWriter, r *http.Request) {
		handleMatrixOperation(w, r, func(m *matrix.Matrix) interface{} {
			return strconv.Itoa(m.Sum())
		})
	})

	// Multiply endpoint
	http.HandleFunc("/multiply", func(w http.ResponseWriter, r *http.Request) {
		handleMatrixOperation(w, r, func(m *matrix.Matrix) interface{} {
			return strconv.Itoa(m.Multiply())
		})
	})

	fmt.Println("Server starting on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
