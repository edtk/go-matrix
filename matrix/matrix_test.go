package matrix

import (
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name    string
		records [][]string
		wantErr bool
	}{
		{
			name: "valid square matrix",
			records: [][]string{
				{"1", "2", "3"},
				{"4", "5", "6"},
				{"7", "8", "9"},
			},
			wantErr: false,
		},
		{
			name: "non-square matrix",
			records: [][]string{
				{"1", "2", "3"},
				{"4", "5", "6"},
			},
			wantErr: true,
		},
		{
			name:    "empty matrix",
			records: [][]string{},
			wantErr: true,
		},
		{
			name: "invalid integer",
			records: [][]string{
				{"1", "2", "3"},
				{"4", "abc", "6"},
				{"7", "8", "9"},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := New(tt.records)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
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
}
