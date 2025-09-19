package main
import (
"encoding/csv"
"fmt"
"log"
"os"
"strconv"
)
type Iris struct {
SepalLength float64
SepalWidth float64
PetalLength float64
PetalWidth float64
Species string
}
func main() {
// Open the CSV file
file, err := os.Open("iris.csv")
if err != nil {
log.Fatal(err)
}
defer file.Close()
// Create a new CSV reader
reader := csv.NewReader(file)
// Read the header row
header, err := reader.Read()
if err != nil {
log.Fatal(err)
}
// Validate the header columns
expectedColumns := []string{"sepal_length", "sepal_width", "petal_length", "petal_width",
"species"}
for i, col := range header {
if col != expectedColumns[i] {
log.Fatalf("Unexpected column at index %d: %s", i, col)
}

}
// Initialize a slice to store the Iris data
irisData := make([]Iris, 0)
// Read the remaining rows
for {
row, err := reader.Read()
if err != nil {
break // EOF or other error
}
// Parse the values from the row
sepalLength, err := strconv.ParseFloat(row[0], 64)
if err != nil {
log.Printf("Error parsing SepalLength in row %d: %v", len(irisData)+1, err)
continue
}
sepalWidth, err := strconv.ParseFloat(row[1], 64)
if err != nil {
log.Printf("Error parsing SepalWidth in row %d: %v", len(irisData)+1, err)
continue
}
petalLength, err := strconv.ParseFloat(row[2], 64)
if err != nil {
log.Printf("Error parsing PetalLength in row %d: %v", len(irisData)+1, err)
continue
}
petalWidth, err := strconv.ParseFloat(row[3], 64)
if err != nil {
log.Printf("Error parsing PetalWidth in row %d: %v", len(irisData)+1, err)
continue
}
iris := Iris{
SepalLength: sepalLength,
SepalWidth: sepalWidth,
PetalLength: petalLength,
PetalWidth: petalWidth,
Species: row[4],
}

irisData = append(irisData, iris)
}
// Example: Print the number of records and the average sepal length
numRecords := len(irisData)
totalSepalLength := 0.0
for _, iris := range irisData {
totalSepalLength += iris.SepalLength
}
averageSepalLength := totalSepalLength / float64(numRecords)
// Print the results
fmt.Println("Number of records:", numRecords)
fmt.Println("Average sepal length:", averageSepalLength)
}