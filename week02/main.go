package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

var db *sql.DB

func initMySQL() (err error) {
	dsn := "root:123qaz@tcp(127.0.0.1:3306)/test"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return errors.Wrap(err, "数据库建立链接失败！")
	}
	err = db.Ping()
	if err != nil {
		return errors.Wrap(err, "数据库建立链接失败！")
	}
	return
}

type User struct {
	id   int
	name string
}

func getUserNameById(id int) (User, error) {
	user := User{id: id}
	err := db.QueryRow("SELECT name FROM user WHERE id = ?", id).Scan(&user.name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			//对sql.ErrNoRows包装后返回
			err = errors.Wrapf(err, "[用户查询]失败  不存在id=%d的用户", id)
		} else {
			err = errors.Wrap(err, "查询数据失败！")
		}
	}

	return user, err
}

func main() {
	err := initMySQL()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	uid := 1
	user, err := getUserNameById(uid)
	if err != nil {
		fmt.Printf("query sql failed, err:%v\n", err)
	} else {
		fmt.Printf("user.id is %d, user.name is %s\n", user.id, user.name)
	}
}
