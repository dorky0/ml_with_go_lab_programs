package main
import (
"encoding/csv"
"fmt"
"log"
"math/rand"
"os"
"strconv"
"time"
"github.com/sjwhitworth/golearn/base"

"github.com/sjwhitworth/golearn/evaluation"
"github.com/sjwhitworth/golearn/trees"
)
const irisDatasetPath = "iris.csv"
func main() {
// Load data
data, err := loadIrisDataset(irisDatasetPath)
if err != nil {
log.Fatal(err)
}
// Split data
trainData, testData := TrainTestSplit(data, 0.7)
// Train model
tree := trees.NewID3DecisionTree(0.6)
err = tree.Fit(trainData)
if err != nil {
log.Fatal(err)
}
// Predict
predictions, err := tree.Predict(testData)
if err != nil {
log.Fatal(err)
}
// Evaluate
confusionMat, err := evaluation.GetConfusionMatrix(testData, predictions)
if err != nil {
log.Fatal(err)
}
accuracy := evaluation.GetAccuracy(confusionMat)
fmt.Printf("Accuracy: %.2f%%\n", accuracy*100)
}
// ------------------------ Load and Convert CSV ------------------------
func loadIrisDataset(path string) (*base.DenseInstances, error) {
file, err := os.Open(path)
if err != nil {
return nil, err

}
defer file.Close()
reader := csv.NewReader(file)
records, err := reader.ReadAll()
if err != nil {
return nil, err
}
// Attributes
attrs := []base.Attribute{
base.NewFloatAttribute("sepal_length"),
base.NewFloatAttribute("sepal_width"),
base.NewFloatAttribute("petal_length"),
base.NewFloatAttribute("petal_width"),
}
classAttr := base.NewCategoricalAttribute()
// Create dataset
numInstances := len(records) - 1
instances := base.NewDenseInstances()
for _, attr := range attrs {
instances.AddAttribute(attr)
}
instances.AddClassAttribute(classAttr)
instances.Extend(numInstances)
for i, row := range records[1:] {
for j := 0; j < 4; j++ {
val, err := strconv.ParseFloat(row[j], 64)
if err != nil {
return nil, err
}
instances.Set(i, attrs[j], base.PackFloatToBytes(val))
}
instances.Set(i, classAttr, []byte(row[4]))
}
return instances, nil
}
// ------------------------ Train/Test Split ------------------------

func TrainTestSplit(instances *base.DenseInstances, trainRatio float64)
(*base.DenseInstances, *base.DenseInstances) {
rand.Seed(time.Now().UnixNano())
numRows := instances.Size()
indices := rand.Perm(numRows)
trainSize := int(float64(numRows) * trainRatio)
trainIndices := indices[:trainSize]
testIndices := indices[trainSize:]
trainSet := base.NewDenseInstances()
testSet := base.NewDenseInstances()
for _, attr := range instances.AllAttributes() {
trainSet.AddAttribute(attr)
testSet.AddAttribute(attr)
}
classAttr := instances.GetClassAttribute()
trainSet.AddClassAttribute(classAttr)
testSet.AddClassAttribute(classAttr)
trainSet.Extend(len(trainIndices))
testSet.Extend(len(testIndices))
for i, idx := range trainIndices {
for _, attr := range instances.AllAttributes() {
val := instances.Get(idx, attr)
trainSet.Set(i, attr, val)
}
}
for i, idx := range testIndices {
for _, attr := range instances.AllAttributes() {
val := instances.Get(idx, attr)
testSet.Set(i, attr, val)
}
}
return trainSet, testSet
}