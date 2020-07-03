package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3307)/apitest")
	if err != nil {
		log.Println(err)
	}
	//在这里进行一些数据库操作
	defer db.Close()
	log.Println("succ")
	err = db.Ping()
	if err != nil {
		log.Println(err)
	}
	var name string
	var age string
	rows, err := db.Query("select id,name from article where status = ? ", 0)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&name, &age)
		if err != nil {
			fmt.Println(err)
		}
	}
	err = rows.Err()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("name:", name, "age:", age)

	stmt, err := db.Prepare("insert into log(id,status)values(?,?)")
	if err != nil {
		log.Println(err)
	}
	rs, err := stmt.Exec("go-test", 0)
	if err != nil {
		log.Println(err)
	}
	//我们可以获得插入的id
	id, err := rs.LastInsertId()
	//可以获得影响行数
	affect, err := rs.RowsAffected()
	log.Println(id, affect, err)
}
