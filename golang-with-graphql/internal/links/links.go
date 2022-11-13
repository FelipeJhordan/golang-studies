package links

import (
	"log"
	"os/user"

	database "github.com/felipejhordan/hackernews/internal/pkg/db/mysql"
)

type Link struct {
	ID      string
	Title   string
	Address string
	User    *user.User
}

func (link Link) Save() int64 {
	stmt, err := database.DB.Prepare("INSERT INTO Links(Title, Address) Values ( ?, ? ) ")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(link.Title, link.Address)
	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Error", err.Error())
	}

	log.Print("Row created/inserted")
	return id
}

func GetAll() []Link {
	stmt, err := database.DB.Prepare("select id, title, address from Links")
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	rows, err := stmt.Query()

	if err != nil {
		log.Fatal(err)
	}

	var links []Link

	for rows.Next() {
		var link Link
		err := rows.Scan(&link.ID, &link.Title, &link.Address)
		if err != nil {
			log.Fatal(err)
		}

		links = append(links, link)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return links
}
