
# csvsimultaneo

Package for concurrent and traceless CSV writing in Go

## Overview

This package provides functions for writing data to CSV files in a concurrent and traceless manner, making it ideal for handling sensitive data or scenarios where performance is critical.

## Key Features

Concurrent Writing: Employs channels and goroutines to write CSV data asynchronously, optimizing performance.
Traceless Handling: Prevents data tracing by directly writing to files without intermediate buffers, enhancing data security.
Error Handling: Includes robust error handling to ensure data integrity and prevent unexpected failures.
Simple Usage: Offers straightforward functions with clear parameters for easy integration into Go projects.
## Usage

Import the package:

Go
import (
    "github.com/your-username/csvsimultaneo"
)
Prepare data for writing:

Go
data := [][]string{
    {"header1", "header2"},
    {"value1", "value2"},
    // ... more data
}
filename := "output.csv"
Write data concurrently:

Go
csvsimultaneo.WriteCsv(csvsimultaneo.CsvLine{Text: data, File: filename})

## License

This package is licensed under the MIT License. See the LICENSE file for details.

## Author

Emilio Gerdez emiliojgerdezd@gmail.com
