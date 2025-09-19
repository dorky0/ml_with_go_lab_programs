package main
import (
"fmt"
"math"
"math/rand"
"time"
)
type Point struct {
X, Y float64
}
type Cluster struct {
Center Point
Points []Point

}
func main() {
// Example data points
data := []Point{
{2, 4}, {2, 6}, {4, 6}, {6, 4}, {6, 6}, {8, 6},
{20, 2}, {22, 4}, {24, 4}, {26, 2}, {26, 4}, {28, 4},
}
// Number of clusters
k := 2
// Run K-Means clustering
clusters := kMeans(data, k)
// Display the resulting clusters
for i, cluster := range clusters {
fmt.Printf("Cluster %d:\n", i+1)
fmt.Printf("Center: (%.2f, %.2f)\n", cluster.Center.X, cluster.Center.Y)
fmt.Println("Points:", cluster.Points)
fmt.Println()
}
}
// kMeans implements the K-Means clustering algorithm
func kMeans(data []Point, k int) []Cluster {
rand.Seed(time.Now().UnixNano())
// Initialize cluster centers using K-Means++ initialization
centers := randomInit(data, k)
clusters := make([]Cluster, k)
for i := range clusters {
clusters[i].Center = centers[i]
}
for {
// Clear points in each cluster before assignment
for i := range clusters {
clusters[i].Points = nil
}
// Assign each point to the nearest cluster center
for _, point := range data {

nearestCluster := findNearestCluster(point, clusters)
nearestCluster.Points = append(nearestCluster.Points, point)
}
changed := false
// Calculate new centers and check for changes
for i := range clusters {
if len(clusters[i].Points) == 0 {
continue
}
newCenter := calculateCenter(clusters[i].Points)
if newCenter != clusters[i].Center {
clusters[i].Center = newCenter
changed = true
}
}
if !changed {
break
}
}
return clusters
}
// randomInit performs K-Means++ initialization
func randomInit(data []Point, k int) []Point {
n := len(data)
centers := make([]Point, k)
centers[0] = data[rand.Intn(n)]
for i := 1; i < k; i++ {
weights := make([]float64, n)
sum := 0.0
for j, point := range data {
minDist := distance(point, centers[0])
for _, center := range centers[:i] {
d := distance(point, center)
if d < minDist {
minDist = d
}
}
weights[j] = minDist * minDist

sum += weights[j]
}
// Normalize weights
for j := range weights {
weights[j] /= sum
}
r := rand.Float64()
for j, w := range weights {
r -= w
if r <= 0 {
centers[i] = data[j]
break
}
}
}
return centers
}
// findNearestCluster returns a pointer to the nearest cluster for the given point
func findNearestCluster(point Point, clusters []Cluster) *Cluster {
nearest := &clusters[0]
minDist := distance(point, nearest.Center)
for i := 1; i < len(clusters); i++ {
d := distance(point, clusters[i].Center)
if d < minDist {
minDist = d
nearest = &clusters[i]
}
}
return nearest
}
// calculateCenter calculates the centroid of the points
func calculateCenter(points []Point) Point {
var sumX, sumY float64
for _, p := range points {
sumX += p.X
sumY += p.Y
}
n := float64(len(points))

return Point{sumX / n, sumY / n}
}
// distance calculates Euclidean distance between two points
func distance(p1, p2 Point) float64 {
dx := p1.X - p2.X
dy := p1.Y - p2.Y
return math.Sqrt(dx*dx + dy*dy)
}