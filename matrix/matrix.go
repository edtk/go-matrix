package matrix

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Errors
var (
	ErrEmptyMatrix = errors.New("empty matrix")
	ErrNotSquare   = errors.New("matrix must be square")
)

// Matrix represents a 2D matrix of integers
type Matrix struct {
	data [][]int
}

// New creates a new Matrix from string records
func New(records [][]string) (*Matrix, error) {
	if len(records) == 0 {
		return nil, ErrEmptyMatrix
	}

	// Validate square matrix and check for empty rows
	rows := len(records)
	cols := len(records[0])

	// Check for empty rows and consistent row lengths
	for i, row := range records {
		if len(row) == 0 {
			return nil, fmt.Errorf("empty row at position %d", i+1)
		}
		if len(row) != cols {
			return nil, fmt.Errorf("inconsistent row length at position %d: expected %d, got %d", i+1, cols, len(row))
		}
	}

	if rows != cols {
		return nil, ErrNotSquare
	}

	// Convert string records to integers
	data := make([][]int, rows)
	for i, row := range records {
		data[i] = make([]int, cols)
		for j, val := range row {
			num, err := strconv.Atoi(strings.TrimSpace(val))
			if err != nil {
				return nil, fmt.Errorf("invalid integer at position [%d,%d]: %s", i+1, j+1, val)
			}
			data[i][j] = num
		}
	}

	return &Matrix{data: data}, nil
}

// Echo returns the matrix as a string in matrix format
func (m *Matrix) Echo() string {
	var result strings.Builder
	for i, row := range m.data {
		if i > 0 {
			result.WriteString("\n")
		}
		result.WriteString(strings.Join(convertToStrings(row), ","))
	}
	return result.String()
}

// Invert returns the matrix with columns and rows inverted
func (m *Matrix) Invert() string {
	rows := len(m.data)
	cols := len(m.data[0])
	inverted := make([][]int, cols)
	for i := range inverted {
		inverted[i] = make([]int, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			inverted[j][i] = m.data[i][j]
		}
	}

	var result strings.Builder
	for i, row := range inverted {
		if i > 0 {
			result.WriteString("\n")
		}
		result.WriteString(strings.Join(convertToStrings(row), ","))
	}
	return result.String()
}

// Flatten returns the matrix as a single line string
func (m *Matrix) Flatten() string {
	var result strings.Builder
	for i, row := range m.data {
		if i > 0 {
			result.WriteString(",")
		}
		result.WriteString(strings.Join(convertToStrings(row), ","))
	}
	return result.String()
}

// Sum returns the sum of all integers in the matrix
func (m *Matrix) Sum() int {
	sum := 0
	for _, row := range m.data {
		for _, val := range row {
			sum += val
		}
	}
	return sum
}

// Multiply returns the product of all integers in the matrix
func (m *Matrix) Multiply() int {
	product := 1
	for _, row := range m.data {
		for _, val := range row {
			product *= val
		}
	}
	return product
}

// convertToStrings converts a slice of integers to a slice of strings
func convertToStrings(nums []int) []string {
	result := make([]string, len(nums))
	for i, num := range nums {
		result[i] = strconv.Itoa(num)
	}
	return result
}
