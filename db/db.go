package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type MyDB struct {
	Id           uint   `json:"id"`
	LongLink     string `json:"longLink"`
	ShortLink    string `json:"shortLink"`
	ClickCounter uint   `json:"clickCounter"`
}

// добавление записи в БД
func InsertToDB(m MyDB) {
	// подключение БД
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/URL_Shortener")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// fmt.Println("База данных подключена")

	insert, err := db.Query(fmt.Sprintf("INSERT INTO `Links` (`long_link`, `short_link`, `click_counter`) VALUES('%s', '%s', '%d')", m.LongLink, m.ShortLink, m.ClickCounter))
	if err != nil {
		panic(err)
	}
	defer insert.Close()
}

// выборка данных из БД
// проверка на дубликаты по полю Long_Link
// выборка данных при редиректе
func LookForDB(column, loadedLink string) string {
	// подключение БД
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/URL_Shortener")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	res := db.QueryRow(fmt.Sprintf("SELECT * FROM `Links` WHERE `%s` = '%s'", column, loadedLink))
	var row MyDB
	err = res.Scan(&row.Id, &row.LongLink, &row.ShortLink, &row.ClickCounter)
	if err != nil {
		log.Print(err)
	}
	switch column {
	case "short_link":
		row.ClickCounter++
		db.Exec(fmt.Sprintf("UPDATE `Links` SET `click_counter` = '%d' WHERE `id` = '%d'", row.ClickCounter, row.Id))
		return (row.LongLink)
	case "long_link":
		return (row.ShortLink)
	}
	return ""
}
