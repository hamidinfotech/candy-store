package repository

import (
	"candystore/entity"
	"database/sql"
	"log"
	_ "github.com/mattn/go-sqlite3"
)

type CandystoreRepositorySqlite struct {

}

const database string = "data/candystore.db"

func (c CandystoreRepositorySqlite) GetTopCustomers() []entity.CustomerStat {
	var customersStat []entity.CustomerStat

	sql := `select *
			from (
				select name, 
				(select candy from (select candy, sum(eaten) as total from candystore where name = c.name group by name, candy) order by total desc limit 1) favourite_snack,
				SUM(eaten) total_snacks
				from candystore c
				GROUP by name
			) order by total_snacks desc`

	rows := query(sql)

	for rows.Next() {
		customerStat := entity.CustomerStat{}
		rows.Scan(&customerStat.Name, &customerStat.FavouriteSnack, &customerStat.TotalSnacks)
		customersStat = append(customersStat, customerStat)
	}

	return customersStat
}

func getDb() *sql.DB {
	db, err := sql.Open("sqlite3", database)

	// defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	return db
}

func query(query string) *sql.Rows {
	rows, error := getDb().Query(query)

	if error != nil {
		log.Fatal(error)
	}

	return rows
}