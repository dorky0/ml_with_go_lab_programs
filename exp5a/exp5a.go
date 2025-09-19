package main
import (
"fmt"
"math"
)
func main() {
// Example predicted and actual values
predicted := []float64{1.2, 2.5, 3.7, 4.8, 5.1}
actual := []float64{1.0, 2.7, 3.5, 4.5, 5.2}
// Calculate metrics
mse := calculateMSE(predicted, actual)
mae := calculateMAE(predicted, actual)
rSquared := calculateRSquared(predicted, actual)
// Print results
fmt.Printf("Predicted: %v\n", predicted)
fmt.Printf("Actual: %v\n", actual)
fmt.Printf("Mean Squared Error (MSE): %.4f\n", mse)
fmt.Printf("Mean Absolute Error (MAE): %.4f\n", mae)
fmt.Printf("R2 (R Squared): %.4f\n", rSquared)
}
// Calculate Mean Squared Error (MSE)
func calculateMSE(predicted, actual []float64) float64 {
sum := 0.0
for i := range predicted {
diff := predicted[i] - actual[i]
sum += diff * diff
}
return sum / float64(len(predicted))
}
// Calculate Mean Absolute Error (MAE)
func calculateMAE(predicted, actual []float64) float64 {
sum := 0.0
for i := range predicted {
diff := math.Abs(predicted[i] - actual[i])

sum += diff
}
return sum / float64(len(predicted))
}
// Calculate R Squared (R2)
func calculateRSquared(predicted, actual []float64) float64 {
mean := calculateMean(actual)
ssTotal := 0.0
ssResidual := 0.0
for i := range predicted {
ssTotal += math.Pow(actual[i]-mean, 2)
ssResidual += math.Pow(predicted[i]-actual[i], 2)
}
return 1 - (ssResidual / ssTotal)
}
// Calculate mean of float64 slice
func calculateMean(values []float64) float64 {
sum := 0.0
for _, v := range values {
sum += v
}
return sum / float64(len(values))
}