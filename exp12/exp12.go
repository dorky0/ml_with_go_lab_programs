package main
import (
"fmt"
"log"
"gorgonia.org/gorgonia"
"gorgonia.org/tensor"
)
func main() {
// Create a new computation graph
g := gorgonia.NewGraph()
// Define input data and expected output
inputData := []float32{0.2, 0.4, 0.6} // Input features
expectedOutput := []float32{0.8} // Expected output (target)
// Create input tensor (matrix)
input := gorgonia.NewMatrix(g, tensor.Float32, gorgonia.WithShape(1,
len(inputData)),

gorgonia.WithName("input"), gorgonia.WithShape(len(inputData)))
// Create weights and biases for the neural network
weights := gorgonia.NewMatrix(g, tensor.Float32,
gorgonia.WithShape(len(inputData), 1),
gorgonia.WithName("weights"))
biases := gorgonia.NewMatrix(g, tensor.Float32, gorgonia.WithShape(1),

gorgonia.WithName("biases"))
// Perform forward pass (input * weights + biases)
output := gorgonia.Must(gorgonia.Add(gorgonia.Must(gorgonia.Mul(input,
weights)), biases))
// Define loss function (mean squared error)
loss := gorgonia.Must(gorgonia.Square(gorgonia.Must(gorgonia.Sub(output,
gorgonia.NewMatrix(g,

tensor.Float32, gorgonia.WithShape(1, 1),

gorgonia.WithValue(expectedOutput))))))
// Compute gradients for weights and biases
grads, err := gorgonia.Gradient(loss, weights, biases)
if err != nil {
log.Fatal(err)
}
// Create VM (Virtual Machine) to run the computation graph
vm := gorgonia.NewTapeMachine(g, gorgonia.BindDualValues(weights, biases))
// Initialize weights and biases with values
gorgonia.Read(weights, tensor.New(tensor.WithShape(len(inputData), 1),
tensor.WithShape(1, 1), tensor.WithValue([]float32{0.1})))
gorgonia.Read(biases, tensor.New(tensor.WithShape(1), tensor.WithShape(1),
tensor.WithValue([]float32{0.0})))
// Perform gradient descent to optimize weights and biases
learningRate := float32(0.1)
for i := 0; i < 100; i++ {
// Run the graph
if err := vm.RunAll(); err != nil {
log.Fatal(err)
}
// Update weights and biases using gradient descent
gorgonia.Apply(weights, func(x float32) float32 {
return x - learningRate*grads[weights].Data().([]float32)[0]
})
gorgonia.Apply(biases, func(x float32) float32 {
return x - learningRate*grads[biases].Data().([]float32)[0]
})
// Reset VM state

vm.Reset()
// Print loss at each iteration to track progress
fmt.Printf("Iteration %d - Loss: %.4f\n", i+1, loss.Value().Data().([]float32)[0])
}
// Print final weights and biases after training
fmt.Println("Final Weights:", weights.Value())
fmt.Println("Final Biases:", biases.Value())
}