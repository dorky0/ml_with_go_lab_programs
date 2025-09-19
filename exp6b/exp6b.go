package main

import (
    "fmt"
)

func main() {
    // Example training dataset
    x1 := []float64{1, 2, 3, 4, 5}
    x2 := []float64{0.5, 1, 1.5, 2, 2.5}
    y := []float64{3, 5, 7, 9, 11}

    // Train the multiple linear regression model
    theta0, theta1, theta2 := trainMultipleLinearRegression(x1, x2, y, 0.01, 1000)

    // Predict using the trained model
    predicted := predictMultipleLinearRegression(x1, x2, theta0, theta1, theta2)

    // Display the trained model and predictions
    fmt.Printf("Trained Model: y = %.2f + %.2fx1 + %.2fx2\n", theta0, theta1, theta2)
    fmt.Println("Predicted Values:", predicted)
}

// Function to train a multiple linear regression model using gradient descent
func trainMultipleLinearRegression(
    x1, x2, y []float64,
    learningRate float64,
    numIterations int,
) (float64, float64, float64) {
    m := len(x1)
    theta0 := 0.0
    theta1 := 0.0
    theta2 := 0.0

    for i := 0; i < numIterations; i++ {
        gradient0 := 0.0
        gradient1 := 0.0
        gradient2 := 0.0

        for j := 0; j < m; j++ {
            predicted := theta0 + theta1*x1[j] + theta2*x2[j]
            error := predicted - y[j]
            gradient0 += error
            gradient1 += error * x1[j]
            gradient2 += error * x2[j]
        }

        gradient0 /= float64(m)
        gradient1 /= float64(m)
        gradient2 /= float64(m)

        theta0 -= learningRate * gradient0
        theta1 -= learningRate * gradient1
        theta2 -= learningRate * gradient2
    }

    return theta0, theta1, theta2
}

// Function to predict values using a trained multiple linear regression model
func predictMultipleLinearRegression(
    x1, x2 []float64,
    theta0, theta1, theta2 float64,
) []float64 {
    predicted := make([]float64, len(x1))
    for i := 0; i < len(x1); i++ {
        predicted[i] = theta0 + theta1*x1[i] + theta2*x2[i]
    }
    return predicted
}