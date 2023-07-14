package db

import (
	"fmt"
	"log"
)

type MetricModel struct {
	Metric float64
}

func PersistData(metric float64) {
	db := GetDB()
	m := []MetricModel{
		{Metric: metric},
	}

	for _, data := range m {
		_, err = db.Exec("INSERT INTO data (metric) VALUES (?)", data.Metric)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func GetData() string {
	var res string
	db := GetDB()
	rows, err := db.Query("SELECT * FROM metrics")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var metric float64
		if err := rows.Scan(&metric); err != nil {
			log.Default().Fatal(err)
		}
		res = fmt.Sprintf("%0.f", metric)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return res
}
