package main
import (
"database/sql"
"fmt"
"log"
_ "github.com/mattn/go-sqlite3"
)
func main() {
// Replace with your SQLite database file path
db, err := sql.Open("sqlite3", "mydatabase.db")
if err != nil {
log.Fatal(err)

}
defer db.Close()
rows, err := db.Query("SELECT id, name, email FROM users")
if err != nil {
log.Fatal(err)
}
defer rows.Close()
for rows.Next() {
var id int
var name, email string
err := rows.Scan(&id, &name, &email)
if err != nil {
log.Fatal(err)
}
fmt.Printf("ID: %d, Name: %s, Email: %s\n", id, name, email)
}
if err = rows.Err(); err != nil {
log.Fatal(err)
}
}