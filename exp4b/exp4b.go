package main
import (
"fmt"
"math"
"sort"
)
func main() {
// Example dataset
dataset := []float64{1.2, 2.5, 3.7, 4.8, 5.1, 6.3, 7.4, 8.6, 9.9}
// Calculate and display statistical measures
fmt.Printf("Dataset: %v\n", dataset)
fmt.Printf("Mean: %.2f\n", calculateMean(dataset))
fmt.Printf("Median: %.2f\n", calculateMedian(dataset))
fmt.Printf("Standard Deviation: %.2f\n", calculateStandardDeviation(dataset))
}
// Function to calculate the mean of a dataset
func calculateMean(dataset []float64) float64 {
sum := 0.0
for _, value := range dataset {
sum += value
}
return sum / float64(len(dataset))
}
// Function to calculate the median of a dataset
func calculateMedian(dataset []float64) float64 {
sorted := make([]float64, len(dataset))
copy(sorted, dataset)
sort.Float64s(sorted)
length := len(sorted)
middle := length / 2
if length%2 == 0 {
return (sorted[middle-1] + sorted[middle]) / 2

}
return sorted[middle]
}
// Function to calculate the standard deviation of a dataset
func calculateStandardDeviation(dataset []float64) float64 {
mean := calculateMean(dataset)
sumSquaredDiff := 0.0
for _, value := range dataset {
diff := value - mean
sumSquaredDiff += diff * diff
}
variance := sumSquaredDiff / float64(len(dataset))
return math.Sqrt(variance)
}