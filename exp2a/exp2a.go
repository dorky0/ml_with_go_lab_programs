package main
import (
"encoding/json"
"fmt"
"log"
)
type Person struct {
Name string `json:"name"`
Age int `json:"age"`
Email string `json:"email"`
}
func main() {
// Example JSON data
jsonData := `{
"name": "John Doe",
"age": 30,
"email": "johndoe@example.com"
}`
// Parse the JSON data into a Person struct

var person Person
err := json.Unmarshal([]byte(jsonData), &person)
if err != nil {
log.Fatal(err)
}
// Access the parsed data
fmt.Println("Name:", person.Name)
fmt.Println("Age:", person.Age)
fmt.Println("Email:", person.Email)
}