package main
import (
"fmt"
"log"
"math/rand"
"time"
"gonum.org/v1/plot"
"gonum.org/v1/plot/plotter"
"gonum.org/v1/plot/vg"
)
func main() {
// Seed random number generator
rand.Seed(time.Now().UnixNano())
// Generate example dataset
dataset := generateDataset(100)
// Create histogram
histogram, err := createHistogram(dataset)
if err != nil {
log.Fatal(err)
}
// Save histogram as PNG
savePlot(histogram, "histogram.png")
// Create box plot
boxPlot, err := createBoxPlot(dataset)
if err != nil {
log.Fatal(err)
}
// Save box plot as PNG
savePlot(boxPlot, "boxplot.png")
fmt.Println("Plots generated successfully.")
}
// Function to generate example dataset

func generateDataset(size int) plotter.Values {
values := make(plotter.Values, size)
for i := range values {
values[i] = rand.Float64()
}
return values
}
// Function to create a histogram
func createHistogram(dataset plotter.Values) (*plot.Plot, error) {
p := plot.New() // No error returned here
// Create histogram with 10 bins
hist, err := plotter.NewHist(dataset, 10)
if err != nil {
return nil, err
}
p.Add(hist)
p.Title.Text = "Histogram"
p.X.Label.Text = "Values"
p.Y.Label.Text = "Frequency"
return p, nil
}
// Function to create a box plot
func createBoxPlot(dataset plotter.Values) (*plot.Plot, error) {
p := plot.New() // No error returned here
// Create box plot at x=0
boxPlot, err := plotter.NewBoxPlot(vg.Points(40), 0, dataset)
if err != nil {
return nil, err
}
p.Add(boxPlot)
p.Title.Text = "Box Plot"
p.X.Label.Text = "Dataset"
p.NominalX("") // Hide X axis label for single box plot
return p, nil

}
// Function to save plot as PNG image
func savePlot(p *plot.Plot, filename string) {
// Save plot as PNG image with width=8 inches and height=4 inches
if err := p.Save(8*vg.Inch, 4*vg.Inch, filename); err != nil {
log.Fatal(err)
}
}