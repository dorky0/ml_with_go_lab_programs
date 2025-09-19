package main
import (
"fmt"
)
func main() {
// Example confusion matrix values
truePositive := 80
falsePositive := 20
trueNegative := 50
falseNegative := 10
// Calculate metrics
accuracy := calculateAccuracy(truePositive, falsePositive, trueNegative,
falseNegative)
precision := calculatePrecision(truePositive, falsePositive)
recall := calculateRecall(truePositive, falseNegative)
auc := calculateAUC(truePositive, falsePositive, trueNegative, falseNegative)
// Print results
fmt.Println("Classification Results:")
fmt.Println("True Positive:", truePositive)
fmt.Println("False Positive:", falsePositive)
fmt.Println("True Negative:", trueNegative)
fmt.Println("False Negative:", falseNegative)
fmt.Printf("Accuracy: %.4f\n", accuracy)
fmt.Printf("Precision: %.4f\n", precision)
fmt.Printf("Recall: %.4f\n", recall)
fmt.Printf("AUC (Area Under Curve): %.4f\n", auc)
}
// Calculate Accuracy
func calculateAccuracy(tp, fp, tn, fn int) float64 {
total := tp + fp + tn + fn
correct := tp + tn
return float64(correct) / float64(total)
}
// Calculate Precision

func calculatePrecision(tp, fp int) float64 {
if tp+fp == 0 {
return 0.0
}
return float64(tp) / float64(tp+fp)
}
// Calculate Recall
func calculateRecall(tp, fn int) float64 {
if tp+fn == 0 {
return 0.0
}
return float64(tp) / float64(tp+fn)
}
// Calculate AUC (Area Under Curve)
// Approximate using trapezoidal rule for a single point (simplified)
// AUC = (TPR + (1 - FPR)) / 2
func calculateAUC(tp, fp, tn, fn int) float64 {
tpr := float64(tp) / float64(tp+fn) // True Positive Rate (Recall)
fpr := float64(fp) / float64(fp+tn) // False Positive Rate
return (tpr + (1 - fpr)) / 2
}