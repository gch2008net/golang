package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	fmt.Println("this is a data-access example ")

	// Capture connection properties.
	dbaddr := os.Getenv("dbaddr")
	fmt.Println(dbaddr)

	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   os.Getenv("dbaddr"),
		DBName: "recordings",
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	//查询多行
	albums, err := albumsByArtist("John Coltrane")

	if err == nil {

		fmt.Printf("Albums found: %v\n", albums)

		for _, a := range albums {

			fmt.Printf("ID: %v  Title: %v  Artist: %v  Price: %v \n", a.ID, a.Title, a.Artist, a.Price)
		}

	}

	//查询单行
	album, err := albumByID(1)

	if err == nil {

		fmt.Printf("album found: %v\n", album)

	}

	//新增
	var alb Album = Album{
		ID:     5,
		Title:  "小蘑菇头",
		Artist: "李刚",
		Price:  55.2,
	}
	id, err := addAlbum(alb)
	if err == nil {
		fmt.Printf("新增记录的ID：%v \n", id)
	}

	//修改
	alb = Album{
		ID:     5,
		Title:  "77777",
		Artist: "7777",
		Price:  55,
	}
	id, err = updateAlbum(alb)
	if err == nil {
		fmt.Printf("受影响的行数：%v \n", id)
	} else {
		fmt.Printf("err2  %v \n", err)
	}

	//删除
	id, err = delAlbum(8)
	if err == nil {
		fmt.Printf("受影响的行数：%v \n", id)
	} else {
		fmt.Printf("err2  %v \n", err)
	}
}

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

// albumsByArtist queries for albums that have the specified artist name.
func albumsByArtist(name string) ([]Album, error) {
	// An albums slice to hold data from returned rows.
	var albums []Album

	rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	return albums, nil
}

// albumByID queries for the album with the specified ID.
func albumByID(id int64) (Album, error) {
	// An album to hold data from the returned row.
	var alb Album

	row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)
	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("albumsById %d: no such album", id)
		}
		return alb, fmt.Errorf("albumsById %d: %v", id, err)
	}
	return alb, nil
}

// addAlbum adds the specified album to the database,
// returning the album ID of the new entry
func addAlbum(alb Album) (int64, error) {
	result, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", alb.Title, alb.Artist, alb.Price)
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	return id, nil
}

func updateAlbum(alb Album) (int64, error) {
	result, err := db.Exec("UPDATE album set title=?,artist=?,price=? WHERE id=?;", alb.Title, alb.Artist, alb.Price, alb.ID)
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	line, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	return line, nil
}

func delAlbum(id int64) (int64, error) {
	result, err := db.Exec("DELETE from album WHERE id=?;", id)
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	line, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	return line, nil
}
