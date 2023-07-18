package db

import (
	"database/sql"
	"log"
)

type MetricModel struct {
	Metric float64
}

func PersistData(metric float64) {

	db, err := sql.Open("sqlite3", "./data.db")

	if err != nil {
		log.Default().Fatal(err)
	}

	defer db.Close()
	m := []MetricModel{
		{Metric: metric},
	}

	for _, data := range m {
		if _, err = db.Exec("INSERT INTO metrics (metric) VALUES (?)", data.Metric); err != nil {
			log.Default().Fatal(err)
		}
	}
}

func GetData() []MetricModel {
	var res []MetricModel
	db, err := sql.Open("sqlite3", "./data.db")

	if err != nil {
		log.Default().Fatal(err)
	}

	defer db.Close()
	rows, err := db.Query("SELECT DISTINCT metric FROM metrics WHERE metric > 0")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		var data MetricModel
		if err := rows.Scan(&data.Metric); err != nil {
			log.Default().Fatal(err)
		}

		res = append(res, data)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return res
}
