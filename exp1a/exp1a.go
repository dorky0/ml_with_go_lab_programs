package main
import (
"encoding/csv"
"fmt"
"io"
"log"
"os"
"strconv"
)

func main() {
// Open the CSV file
file, err := os.Open("data.csv")
if err != nil {
log.Fatal("Failed to open file:", err)
}
defer file.Close()
// Create CSV reader
reader := csv.NewReader(file)
// Read the header row
header, err := reader.Read()
if err != nil {
log.Fatal("Failed to read header:", err)
}
// Target column to find max value in
columnName := "Score"
// Find the index of the target column
columnIndex := -1
for i, col := range header {
if col == columnName {
columnIndex = i
break
}
}
if columnIndex == -1 {
log.Fatalf("Column %s not found in header", columnName)
}
// Initialize maxValue as negative infinity to handle all values correctly
maxValue := -1e308 // a very small number
// Read and process each row
for {
row, err := reader.Read()
if err == io.EOF {
break
}
if err != nil {
log.Fatal("Failed to read row:", err)

}
// Convert value to float64
value, err := strconv.ParseFloat(row[columnIndex], 64)
if err != nil {
log.Printf("Skipping invalid value: %s\n", row[columnIndex])
continue
}
// Update max value if current value is greater
if value > maxValue {
maxValue = value
}
}
// Output the maximum value found
fmt.Printf("Maximum value in column %s is %.2f\n", columnName, maxValue)
}

//create excel name "data" and save it in same folder(go-test) as csv-msdol in options

