package main

import (
	"fmt"
	"math"
)

// sigmoid computes the sigmoid function value for input z
func sigmoid(z float64) float64 {
	return 1 / (1 + math.Exp(-z))
}

// trainLogisticRegression trains a logistic regression model using gradient descent
func trainLogisticRegression(x [][]float64, y []int, learningRate float64, numIterations int) (float64, float64, float64) {
	m := len(x)
	theta0, theta1, theta2 := 0.0, 0.0, 0.0

	for iter := 0; iter < numIterations; iter++ {
		var gradient0, gradient1, gradient2 float64

		for i := 0; i < m; i++ {
			z := theta0 + theta1*x[i][0] + theta2*x[i][1]
			pred := sigmoid(z)
			err := pred - float64(y[i])

			gradient0 += err
			gradient1 += err * x[i][0]
			gradient2 += err * x[i][1]
		}

		// Average the gradients
		gradient0 /= float64(m)
		gradient1 /= float64(m)
		gradient2 /= float64(m)

		// Update parameters
		theta0 -= learningRate * gradient0
		theta1 -= learningRate * gradient1
		theta2 -= learningRate * gradient2

		// Optional: print loss every 100 iterations
		/*
			if iter%100 == 0 {
				loss := 0.0
				for i := 0; i < m; i++ {
					z := theta0 + theta1*x[i][0] + theta2*x[i][1]
					h := sigmoid(z)
					loss += -float64(y[i])*math.Log(h) - (1-float64(y[i]))*math.Log(1-h)
				}
				loss /= float64(m)
				fmt.Printf("Iteration %d - Loss: %.4f\n", iter, loss)
			}
		*/
	}

	return theta0, theta1, theta2
}

// predictLogisticRegression predicts probabilities for input data using the trained model
func predictLogisticRegression(x [][]float64, theta0, theta1, theta2 float64) []float64 {
	m := len(x)
	predicted := make([]float64, m)

	for i := 0; i < m; i++ {
		z := theta0 + theta1*x[i][0] + theta2*x[i][1]
		predicted[i] = sigmoid(z)
	}

	return predicted
}

func main() {
	// Training dataset (2 features per sample)
	x := [][]float64{
		{1, 1}, {1, 2}, {2, 1}, {2, 2},
		{3, 3}, {4, 4}, {4, 5}, {5, 4},
	}

	// Corresponding binary labels
	y := []int{0, 0, 0, 0, 1, 1, 1, 1}

	// Train the model
	theta0, theta1, theta2 := trainLogisticRegression(x, y, 0.01, 1000)

	// Predict probabilities on training data
	probabilities := predictLogisticRegression(x, theta0, theta1, theta2)

	// Output the learned model and predictions
	fmt.Printf("Trained Model: y = 1 / (1 + exp(- (%.4f + %.4f*x1 + %.4f*x2)))\n\n", theta0, theta1, theta2)
	fmt.Println("Predictions on training data:")

	for i, prob := range probabilities {
		label := 0
		if prob >= 0.5 {
			label = 1
		}
		fmt.Printf("Sample %d: x = [%.1f, %.1f], Predicted Probability = %.4f, Predicted Label = %d, True Label = %d\n",
			i+1, x[i][0], x[i][1], prob, label, y[i])
	}
}