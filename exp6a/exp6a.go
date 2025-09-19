package main
import (
"fmt"
)
func main() {
// Example training dataset
x := []float64{1, 2, 3, 4, 5}
y := []float64{2, 4, 6, 8, 10}
// Train the linear regression model with learning rate and iterations
theta0, theta1 := trainLinearRegression(x, y, 0.01, 1000)
// Predict using the trained model
predicted := predictLinearRegression(x, theta0, theta1)
// Display the trained model and predictions
fmt.Printf("Trained Model: y = %.2f + %.2fx\n", theta0, theta1)
fmt.Println("Predicted Values:", predicted)
}
// Function to train a linear regression model using gradient descent

func trainLinearRegression(x, y []float64, learningRate float64, numIterations int) (float64,
float64) {
m := len(x)
theta0 := 0.0 // Intercept
theta1 := 0.0 // Slope
for i := 0; i < numIterations; i++ {
gradient0 := 0.0
gradient1 := 0.0
for j := 0; j < m; j++ {
predicted := theta0 + theta1*x[j]
error := predicted - y[j]
gradient0 += error
gradient1 += error * x[j]
}
gradient0 /= float64(m)
gradient1 /= float64(m)
theta0 -= learningRate * gradient0
theta1 -= learningRate * gradient1
}
return theta0, theta1
}
// Function to predict values using a trained linear regression model
func predictLinearRegression(x []float64, theta0, theta1 float64) []float64 {
predicted := make([]float64, len(x))
for i := 0; i < len(x); i++ {
predicted[i] = theta0 + theta1*x[i]
}
return predicted
}