package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB // Database handle
// db is a global variable to simplifies this example 

type Album struct {
  ID int64
  Title string
  Artist string
  Price float32
}

func main() {
  // Capture connection properties
  cfg := mysql.Config{
    User: os.Getenv("DBUSER"),
    Passwd: os.Getenv("DBPASS"),
    Net: "tcp",
    Addr: "127.0.0.1:3306",
    DBName: "gorecordings",
    AllowNativePasswords: true,
  }

  // Get a database handle
  var err error
  db, err = sql.Open("mysql", cfg.FormatDSN()) // initialize the db variable, passing the return value of FormatDSN
  if err != nil {
    log.Fatal(err)
  }

  pingErr := db.Ping() // Confirm that connecting to the database works when it needs to 
  if pingErr != nil { // Check for an error from Ping in case that connection failed
    log.Fatal(pingErr)
  }

  fmt.Println("Connected!") 

  // Search by artist
  albums, err := albumsByArtist("John Coltrane")
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("Albums found: %v\n", albums)

  // Search by Id
  alb, err := albumByID(2)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("Album found: %v\n",alb)

  // Add new album 
  albID, err := addAlbum(Album{
    Title: "The Modern Sound of Betty Carter",
    Artist: "Betty Carter",
    Price: 49.99,
  })
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("ID of added album: %v\n", albID)
}

// albumsByArtist queries for albums that have the specified artist name
func albumsByArtist(name string) ([]Album, error) {
  // An albums slice to hold data from returned rows.
  var albums []Album

  rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
  if err != nil {
    return nil, fmt.Errorf("albumsByArtist %q: %v",name, err)
  } 

  defer rows.Close()
  // Loop through rows, using Scan to assign column data to struct fields
  for rows.Next() {
    var alb Album

    if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
      return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
    }
    albums = append(albums, alb)
  }

  if err := rows.Err(); err != nil { // Always check for an error from sql.Rows after looping over query results. If the query failed, this is how your code finds out
    return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
  }

  return albums, nil
}

// albumByID queries for the album with the specified ID
func albumByID(id int64) (Album, error) {
  // An album to hold data from the returned row
  var alb Album

  row := db.QueryRow("select * from album where id = ?", id)
  if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
    if err == sql.ErrNoRows {
      return alb, fmt.Errorf("albumById %d: no such album", id)
    }
    return alb, fmt.Errorf("albumById %d: %v", id, err)
  }

  return alb, nil
}

// addAlbum adds the specified album to the database,
// returning the album ID of the new entry
func addAlbum(alb Album) (int64, error){
  result, err := db.Exec("insert into album (title,artist, price) values(?,?,?)", alb.Title, alb.Artist, alb.Price) 
  if err != nil {
    return 0, fmt.Errorf("addAlbum: %v", err)
  }

  id, err := result.LastInsertId()
  if err != nil {
    return 0, fmt.Errorf("addAlbum: %v", err)
  }

  return id, nil
}
