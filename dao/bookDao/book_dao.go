package bookDao

import (
	"fmt"
	"github.com/echo-test/dao/coreDao"
	"github.com/echo-test/domain"
	"github.com/uuidcode/coreutil"
)

func FindAll() []domain.Book {
	db := coreDao.ConnectDB()

	sql := `
	select * from book
	
	`

	rows, err := db.Queryx(sql)
	coreutil.CheckErr(err)

	defer rows.Close()
	defer db.Close()

	var bookList []domain.Book

	for rows.Next() {
		var currentBook domain.Book
		err := rows.StructScan(&currentBook)
		coreutil.CheckErr(err)

		bookList = append(bookList, currentBook)
	}

	for i, currentBook := range bookList {
		fmt.Printf("%v %+v\n", i, currentBook)
	}

	return bookList
}
