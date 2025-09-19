package main
import (
"encoding/csv"
"fmt"
"log"
"math"
"os"
"strconv"
"gonum.org/v1/gonum/mat"
)
const (
irisDatasetPath = "exp8iris.csv"
k = 5
)
// irisDataInstance holds features and label for one Iris data point
type irisDataInstance struct {
Features []float64

Label string
}
func main() {
// Load the Iris dataset
irisData, err := loadIrisDataset(irisDatasetPath)
if err != nil {
log.Fatalf("Failed to load dataset: %v", err)
}
// Split the dataset into training and testing sets (70% train, 30% test)
trainData, testData := splitDataset(irisData, 0.7)
// Prepare feature matrices and label slices
trainFeatures, trainLabels := prepareData(trainData)
testFeatures, testLabels := prepareData(testData)
// Create and train KNN classifier
knn := NewKNNClassifier(k)
knn.Fit(trainFeatures, trainLabels)
// Predict on test data
predictedLabels := knn.Predict(testFeatures)
// Calculate accuracy
accuracy := calculateAccuracy(testLabels, predictedLabels)
fmt.Printf("Accuracy: %.2f%%\n", accuracy*100)
}
// loadIrisDataset loads the iris dataset from a CSV file and skips the header
func loadIrisDataset(filepath string) ([]irisDataInstance, error) {
file, err := os.Open(filepath)
if err != nil {
return nil, err
}
defer file.Close()
reader := csv.NewReader(file)
// Skip header row
if _, err := reader.Read(); err != nil {
return nil, err
}

records, err := reader.ReadAll()
if err != nil {
return nil, err
}
irisData := make([]irisDataInstance, len(records))
for i, record := range records {
features := make([]float64, 4)
for j := 0; j < 4; j++ {
features[j], err = strconv.ParseFloat(record[j], 64)
if err != nil {
return nil, err
}
}
irisData[i] = irisDataInstance{
Features: features,
Label: record[4],
}
}
return irisData, nil
}
// splitDataset splits dataset into train and test slices according to ratio
func splitDataset(dataset []irisDataInstance, splitRatio float64) ([]irisDataInstance,
[]irisDataInstance) {
trainSize := int(float64(len(dataset)) * splitRatio)
return dataset[:trainSize], dataset[trainSize:]
}
// prepareData separates features and labels from dataset
func prepareData(dataset []irisDataInstance) (*mat.Dense, []string) {
rows := len(dataset)
cols := len(dataset[0].Features)
features := mat.NewDense(rows, cols, nil)
labels := make([]string, rows)
for i, inst := range dataset {
for j, val := range inst.Features {
features.Set(i, j, val)
}
labels[i] = inst.Label
}
return features, labels

}
// KNNClassifier implements a simple k-nearest neighbor classifier
type KNNClassifier struct {
K int
X *mat.Dense
Y []string
}
// NewKNNClassifier initializes a new KNNClassifier with given k
func NewKNNClassifier(k int) *KNNClassifier {
return &KNNClassifier{K: k}
}
// Fit trains the classifier with training features and labels
func (knn *KNNClassifier) Fit(features *mat.Dense, labels []string) {
knn.X = features
knn.Y = labels
}
// Predict predicts labels for given test features
func (knn *KNNClassifier) Predict(testFeatures *mat.Dense) []string {
testRows, _ := testFeatures.Dims()
predictedLabels := make([]string, testRows)
for i := 0; i < testRows; i++ {
testRow := testFeatures.RowView(i)
distances := make([]float64, len(knn.Y))
for j := 0; j < len(knn.Y); j++ {
trainRow := knn.X.RowView(j)
distances[j] = euclideanDistance(testRow, trainRow)
}
indices := getKNearestIndices(distances, knn.K)
predictedLabels[i] = majorityLabel(indices, knn.Y)
}
return predictedLabels
}
// euclideanDistance calculates the Euclidean distance between two vectors
func euclideanDistance(a, b mat.Vector) float64 {
sum := 0.0
for i := 0; i < a.Len(); i++ {
diff := a.AtVec(i) - b.AtVec(i)
sum += diff * diff

}
return math.Sqrt(sum)
}
// getKNearestIndices returns indices of k smallest distances
func getKNearestIndices(distances []float64, k int) []int {
indices := make([]int, k)
copyDistances := make([]float64, len(distances))
copy(copyDistances, distances)
for i := 0; i < k; i++ {
minIdx := findMinIndex(copyDistances)
indices[i] = minIdx
copyDistances[minIdx] = math.Inf(1)
}
return indices
}
// findMinIndex finds index of smallest value in a slice
func findMinIndex(slice []float64) int {
minIdx := 0
minVal := slice[0]
for i, val := range slice {
if val < minVal {
minVal = val
minIdx = i
}
}
return minIdx
}
// majorityLabel returns the most frequent label among the given indices
func majorityLabel(indices []int, labels []string) string {
counts := make(map[string]int)
maxCount := 0
var majority string
for _, idx := range indices {
counts[labels[idx]]++
if counts[labels[idx]] > maxCount {
maxCount = counts[labels[idx]]
majority = labels[idx]
}
}

return majority
}
// calculateAccuracy computes classification accuracy
func calculateAccuracy(trueLabels, predictedLabels []string) float64 {
correct := 0
total := len(trueLabels)
for i := 0; i < total; i++ {
if trueLabels[i] == predictedLabels[i] {
correct++
}
}
return float64(correct) / float64(total)
}