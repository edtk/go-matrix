# Matrix Operations Web Service

A Go web service that performs various operations on square matrices provided via CSV files.

## Features

- Echo: Returns the matrix as a string in matrix format
- Invert: Returns the matrix with columns and rows inverted
- Flatten: Returns the matrix as a single line string
- Sum: Returns the sum of all integers in the matrix
- Multiply: Returns the product of all integers in the matrix

## Requirements

- Go 1.24.4 or higher
- Git

## Installation

1. Clone the repository:
```bash
git clone https://github.com/edtk/go-matrix.git
cd go-matrix
```

## Running the Service

1. Start the server:
```bash
go run .
```

2. The server will start on port 8080.

## Testing the Endpoints

You can test the endpoints using curl. Here are examples for each operation:

```bash
# Echo - Returns the matrix as is
curl -F 'file=@matrix.csv' "localhost:8080/echo"
# Expected output:
# 1,2,3
# 4,5,6
# 7,8,9

# Invert - Returns the matrix with columns and rows inverted
curl -F 'file=@matrix.csv' "localhost:8080/invert"
# Expected output:
# 1,4,7
# 2,5,8
# 3,6,9

# Flatten - Returns the matrix as a single line
curl -F 'file=@matrix.csv' "localhost:8080/flatten"
# Expected output:
# 1,2,3,4,5,6,7,8,9

# Sum - Returns the sum of all integers
curl -F 'file=@matrix.csv' "localhost:8080/sum"
# Expected output:
# 45

# Multiply - Returns the product of all integers
curl -F 'file=@matrix.csv' "localhost:8080/multiply"
# Expected output:
# 362880
```

## Input Format

The input should be a CSV file containing a square matrix of integers. For example:

```
1,2,3
4,5,6
7,8,9
```

Requirements for the input file:
- Must be a square matrix (equal number of rows and columns)
- Must contain only integers
- No header row
- Values separated by commas
- Each row on a new line

## Error Handling

The service handles various error cases:
- Invalid CSV format
- Non-square matrices
- Non-integer values
- Empty matrices
- Missing or invalid file uploads

## Running Tests

To run the tests:

```bash
go test ./matrix
```
