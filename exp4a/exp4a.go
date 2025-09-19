package main
import (
"fmt"
)
func main() {
// Creating a matrix
matrix := [][]int{
{1, 2, 3},
{4, 5, 6},
{7, 8, 9},
}
// Creating a vector
vector := []int{1, 2, 3}
// Displaying the matrix
fmt.Println("Matrix:")
displayMatrix(matrix)
// Displaying the vector
fmt.Println("Vector:")
displayVector(vector)
// Accessing matrix elements
element := matrix[1][2]
fmt.Printf("Element at row 1, column 2: %d\n", element)
// Accessing vector elements
element = vector[0]
fmt.Printf("First element of the vector: %d\n", element)
// Modifying matrix elements
matrix[0][1] = 10
// Modifying vector elements
vector[2] = 20
// Displaying the modified matrix and vector
fmt.Println("Modified matrix:")
displayMatrix(matrix)

fmt.Println("Modified vector:")
displayVector(vector)
}
// Helper function to display the matrix
func displayMatrix(matrix [][]int) {
for _, row := range matrix {
for _, element := range row {
fmt.Printf("%d ", element)
}
fmt.Println()
}
fmt.Println()
}
// Helper function to display the vector
func displayVector(vector []int) {
for _, element := range vector {
fmt.Printf("%d ", element)
}
fmt.Println()
}