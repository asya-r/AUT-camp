package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// Camp type
type Camp struct {
	Id            int
	Title         string
	Specification string
	StartLat      float64
	StartLng      float64
	StartDate     string
	EndDate       string
	MembersNum    int
}

func DatabaseConnection() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}
	return db
}

var Db = DatabaseConnection()

// GetCamps returns all camps
func GetCamps() []Camp {
	query := fmt.Sprintf("SELECT id, title, specification, start_lat, start_lng, start_date, end_date, members_num FROM camp")
	rows, err := Db.Query(query)
	camps := []Camp{}
	if err != nil {
		panic(err)
	} else {
		for rows.Next() {
			c := Camp{}
			err := rows.Scan(&c.Id, &c.Title, &c.Specification, &c.StartLat, &c.StartLng, &c.StartDate, &c.EndDate, &c.MembersNum)
			if err != nil {
				log.Println(err)
				continue
			}
			camps = append(camps, c)
		}
	}
	return camps

}
