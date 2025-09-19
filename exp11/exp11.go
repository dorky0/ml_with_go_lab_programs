package main
import (
"fmt"
"log"
"gonum.org/v1/gonum/stat"
)
func main() {
// Example time series data
data := []float64{
10.0, 12.0, 14.0, 16.0, 18.0, 20.0, 22.0, 24.0, 26.0, 28.0,
}

// Prepare lagged values (inputs) and actual values (outputs)
inputs := make([]float64, len(data)-1)
outputs := make([]float64, len(data)-1)
for i := 1; i < len(data); i++ {
inputs[i-1] = data[i-1] // previous time step
outputs[i-1] = data[i] // current time step
}
// Perform linear regression using least squares
alpha, beta := stat.LinearRegression(inputs, outputs, nil)
// Print the coefficients (Intercept and Slope)
fmt.Printf("Intercept (alpha): %v\n", alpha)
fmt.Printf("Slope (beta): %v\n", beta)
// Predict next value using the last input
lastInput := inputs[len(inputs)-1]
predictedNextValue := alpha + beta*lastInput
fmt.Printf("Predicted next value: %.2f\n", predictedNextValue)
}