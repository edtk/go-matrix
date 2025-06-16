package matrix

import (
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name        string
		records     [][]string
		wantErr     bool
		errContains string
	}{
		{
			name: "valid square matrix",
			records: [][]string{
				{"1", "2", "3"},
				{"4", "5", "6"},
				{"7", "8", "9"},
			},
			wantErr:     false,
			errContains: "",
		},
		{
			name: "non-square matrix",
			records: [][]string{
				{"1", "2", "3"},
				{"4", "5", "6"},
			},
			wantErr:     true,
			errContains: "matrix must be square",
		},
		{
			name:        "empty matrix",
			records:     [][]string{},
			wantErr:     true,
			errContains: "empty matrix",
		},
		{
			name: "invalid integer",
			records: [][]string{
				{"1", "2", "3"},
				{"4", "abc", "6"},
				{"7", "8", "9"},
			},
			wantErr:     true,
			errContains: "invalid integer at position [2,2]: abc",
		},
		{
			name: "matrix with empty row",
			records: [][]string{
				{"1", "2", "3"},
				{},
				{"7", "8", "9"},
			},
			wantErr:     true,
			errContains: "empty row at position 2",
		},
		{
			name: "matrix with inconsistent row lengths",
			records: [][]string{
				{"1", "2", "3"},
				{"4", "5"},
				{"7", "8", "9"},
			},
			wantErr:     true,
			errContains: "inconsistent row length at position 2: expected 3, got 2",
		},
		{
			name: "matrix with decimal numbers",
			records: [][]string{
				{"1", "2", "3"},
				{"4", "5.5", "6"},
				{"7", "8", "9"},
			},
			wantErr:     true,
			errContains: "invalid integer at position [2,2]: 5.5",
		},
		{
			name: "matrix with leading/trailing spaces",
			records: [][]string{
				{"1", " 2 ", "3"},
				{"4", "5", "6"},
				{"7", "8", "9"},
			},
			wantErr:     false,
			errContains: "",
		},
		{
			name: "matrix with negative numbers",
			records: [][]string{
				{"-1", "-2", "-3"},
				{"-4", "-5", "-6"},
				{"-7", "-8", "-9"},
			},
			wantErr:     false,
			errContains: "",
		},
		{
			name: "matrix with zero values",
			records: [][]string{
				{"0", "0", "0"},
				{"0", "0", "0"},
				{"0", "0", "0"},
			},
			wantErr:     false,
			errContains: "",
		},
		{
			name: "matrix with mixed signs",
			records: [][]string{
				{"1", "-2", "3"},
				{"-4", "5", "-6"},
				{"7", "-8", "9"},
			},
			wantErr:     false,
			errContains: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := New(tt.records)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && !strings.Contains(err.Error(), tt.errContains) {
				t.Errorf("New() error message = %v, want to contain %v", err, tt.errContains)
			}
		})
	}
}

func TestMatrixOperations(t *testing.T) {
	tests := []struct {
		name      string
		records   [][]string
		operation func(*Matrix) interface{}
		want      interface{}
	}{
		{
			name: "Echo with positive numbers",
			records: [][]string{
				{"1", "2", "3"},
				{"4", "5", "6"},
				{"7", "8", "9"},
			},
			operation: func(m *Matrix) interface{} { return m.Echo() },
			want:      "1,2,3\n4,5,6\n7,8,9",
		},
		{
			name: "Echo with negative numbers",
			records: [][]string{
				{"-1", "-2", "-3"},
				{"-4", "-5", "-6"},
				{"-7", "-8", "-9"},
			},
			operation: func(m *Matrix) interface{} { return m.Echo() },
			want:      "-1,-2,-3\n-4,-5,-6\n-7,-8,-9",
		},
		{
			name: "Invert with positive numbers",
			records: [][]string{
				{"1", "2", "3"},
				{"4", "5", "6"},
				{"7", "8", "9"},
			},
			operation: func(m *Matrix) interface{} { return m.Invert() },
			want:      "1,4,7\n2,5,8\n3,6,9",
		},
		{
			name: "Invert with negative numbers",
			records: [][]string{
				{"-1", "-2", "-3"},
				{"-4", "-5", "-6"},
				{"-7", "-8", "-9"},
			},
			operation: func(m *Matrix) interface{} { return m.Invert() },
			want:      "-1,-4,-7\n-2,-5,-8\n-3,-6,-9",
		},
		{
			name: "Flatten with positive numbers",
			records: [][]string{
				{"1", "2", "3"},
				{"4", "5", "6"},
				{"7", "8", "9"},
			},
			operation: func(m *Matrix) interface{} { return m.Flatten() },
			want:      "1,2,3,4,5,6,7,8,9",
		},
		{
			name: "Flatten with negative numbers",
			records: [][]string{
				{"-1", "-2", "-3"},
				{"-4", "-5", "-6"},
				{"-7", "-8", "-9"},
			},
			operation: func(m *Matrix) interface{} { return m.Flatten() },
			want:      "-1,-2,-3,-4,-5,-6,-7,-8,-9",
		},
		{
			name: "Sum with positive numbers",
			records: [][]string{
				{"1", "2", "3"},
				{"4", "5", "6"},
				{"7", "8", "9"},
			},
			operation: func(m *Matrix) interface{} { return m.Sum() },
			want:      45,
		},
		{
			name: "Sum with negative numbers",
			records: [][]string{
				{"-1", "-2", "-3"},
				{"-4", "-5", "-6"},
				{"-7", "-8", "-9"},
			},
			operation: func(m *Matrix) interface{} { return m.Sum() },
			want:      -45,
		},
		{
			name: "Sum with mixed signs",
			records: [][]string{
				{"1", "-2", "3"},
				{"-4", "5", "-6"},
				{"7", "-8", "9"},
			},
			operation: func(m *Matrix) interface{} { return m.Sum() },
			want:      5,
		},
		{
			name: "Sum with zeros",
			records: [][]string{
				{"0", "0", "0"},
				{"0", "0", "0"},
				{"0", "0", "0"},
			},
			operation: func(m *Matrix) interface{} { return m.Sum() },
			want:      0,
		},
		{
			name: "Multiply with positive numbers",
			records: [][]string{
				{"1", "2", "3"},
				{"4", "5", "6"},
				{"7", "8", "9"},
			},
			operation: func(m *Matrix) interface{} { return m.Multiply() },
			want:      362880,
		},
		{
			name: "Multiply with negative numbers",
			records: [][]string{
				{"-1", "-2", "-3"},
				{"-4", "-5", "-6"},
				{"-7", "-8", "-9"},
			},
			operation: func(m *Matrix) interface{} { return m.Multiply() },
			want:      -362880,
		},
		{
			name: "Multiply with mixed signs",
			records: [][]string{
				{"1", "-2", "3"},
				{"-4", "5", "-6"},
				{"7", "-8", "9"},
			},
			operation: func(m *Matrix) interface{} { return m.Multiply() },
			want:      362880,
		},
		{
			name: "Multiply with zeros",
			records: [][]string{
				{"0", "0", "0"},
				{"0", "0", "0"},
				{"0", "0", "0"},
			},
			operation: func(m *Matrix) interface{} { return m.Multiply() },
			want:      0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := New(tt.records)
			if err != nil {
				t.Fatalf("Failed to create matrix: %v", err)
			}
			got := tt.operation(m)
			if got != tt.want {
				t.Errorf("%s() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
