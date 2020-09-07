package v1

import (
	"database/sql"
	"fmt"
	"github.com/ClickHouse/clickhouse-go"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

var (
	connect *sql.DB
	err     error
)

func init() {
	connect, err = sql.Open("clickhouse", "tcp://192.168.10.10:9000?debug=true&user=homestead&password=han")
	if err != nil {
		log.Fatal(err)
	}

	if err = connect.Ping(); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			fmt.Printf("[%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		} else {
			fmt.Println(err)
		}
		return
	}
}

func ClickCreateTable(c *gin.Context) {
	_, err = connect.Exec(`
		CREATE TABLE IF NOT EXISTS example (
			country_code FixedString(2),
			os_id        UInt8,
			browser_id   UInt8,
			categories   Array(Int16),
			action_day   Date,
			action_time  DateTime
		) engine=Memory
	`)
}

func ClickInsert(c *gin.Context) {
	var (
		tx, _   = connect.Begin()
		stmt, _ = tx.Prepare("INSERT INTO example (country_code, os_id, browser_id, categories, action_day, action_time) VALUES (?, ?, ?, ?, ?, ?)")
	)
	defer stmt.Close()

	for i := 0; i < 100; i++ {
		if _, err := stmt.Exec(
			"RU",
			10+i,
			100+i,
			clickhouse.Array([]int16{1, 2, 3}),
			time.Now(),
			time.Now(),
		); err != nil {
			log.Fatal(err)
		}
	}

	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}
}

func ClickQuery(c *gin.Context) {

	rows, err := connect.Query("SELECT country_code, os_id, browser_id, categories, action_day, action_time FROM example")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			country               string
			os, browser           uint8
			categories            []int16
			actionDay, actionTime time.Time
		)
		if err := rows.Scan(&country, &os, &browser, &categories, &actionDay, &actionTime); err != nil {
			log.Fatal(err)
		}
		log.Printf("country: %s, os: %d, browser: %d, categories: %v, action_day: %s, action_time: %s", country, os, browser, categories, actionDay, actionTime)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}

func ClickDropTable(c *gin.Context) {
	if _, err := connect.Exec("DROP TABLE example"); err != nil {
		log.Fatal(err)
	}
}
