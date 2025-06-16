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
	records := [][]string{
		{"1", "2", "3"},
		{"4", "5", "6"},
		{"7", "8", "9"},
	}

	m, err := New(records)
	if err != nil {
		t.Fatalf("Failed to create matrix: %v", err)
	}

	tests := []struct {
		name      string
		operation func() interface{}
		want      interface{}
	}{
		{
			name:      "Echo",
			operation: func() interface{} { return m.Echo() },
			want:      "1,2,3\n4,5,6\n7,8,9",
		},
		{
			name:      "Invert",
			operation: func() interface{} { return m.Invert() },
			want:      "1,4,7\n2,5,8\n3,6,9",
		},
		{
			name:      "Flatten",
			operation: func() interface{} { return m.Flatten() },
			want:      "1,2,3,4,5,6,7,8,9",
		},
		{
			name:      "Sum",
			operation: func() interface{} { return m.Sum() },
			want:      45,
		},
		{
			name:      "Multiply",
			operation: func() interface{} { return m.Multiply() },
			want:      362880,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.operation()
			if got != tt.want {
				t.Errorf("%s() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}

	zeroMatrix, _ := New([][]string{
		{"0", "0", "0"},
		{"0", "0", "0"},
		{"0", "0", "0"},
	})
	t.Run("Sum_ZeroMatrix", func(t *testing.T) {
		if got := zeroMatrix.Sum(); got != 0 {
			t.Errorf("Sum() = %v, want 0", got)
		}
	})
	t.Run("Multiply_ZeroMatrix", func(t *testing.T) {
		if got := zeroMatrix.Multiply(); got != 0 {
			t.Errorf("Multiply() = %v, want 0", got)
		}
	})

	negMatrix, _ := New([][]string{
		{"-1", "-2", "-3"},
		{"-4", "-5", "-6"},
		{"-7", "-8", "-9"},
	})
	t.Run("Sum_NegativeMatrix", func(t *testing.T) {
		if got := negMatrix.Sum(); got != -45 {
			t.Errorf("Sum() = %v, want -45", got)
		}
	})
	t.Run("Multiply_NegativeMatrix", func(t *testing.T) {
		if got := negMatrix.Multiply(); got != -362880 {
			t.Errorf("Multiply() = %v, want -362880", got)
		}
	})

	mixedMatrix, _ := New([][]string{
		{"1", "0", "-1"},
		{"2", "0", "-2"},
		{"3", "0", "-3"},
	})
	t.Run("Sum_MixedMatrix", func(t *testing.T) {
		if got := mixedMatrix.Sum(); got != 0 {
			t.Errorf("Sum() = %v, want 0", got)
		}
	})
	t.Run("Multiply_MixedMatrix", func(t *testing.T) {
		if got := mixedMatrix.Multiply(); got != 0 {
			t.Errorf("Multiply() = %v, want 0", got)
		}
	})
}
