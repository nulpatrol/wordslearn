package main

import (
    "sort"
    "fmt"
    "strings"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

type WordForm struct {
    Id int
    WordId int
    Form string
}

type TodoPageData struct {
    PageTitle string
    WordForms []WordForm
}

type preparedQuery struct {
    sql string
    bindings []string
}

func dbConn() (db *sql.DB) {
    dbDriver := "mysql"
    dbUser := "enigma"
    dbPass := "secret"
    dbName := "words-app"
    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(127.0.0.1:3308)/"+dbName)
    if err != nil {
        fmt.Println("Problem")
    }
    return db
}

func getQueryForWordsForms(m map[string]int) []preparedQuery {
	keys := make([]string, 0, len(m))
    for k := range m {
        keys = append(keys, k)
    }
    sort.Strings(keys)

    var queries []preparedQuery

	offset := 0
	for {
		page := paginate(keys, offset, 10)
		if (len(page) == 0) {
			break
		}

		query := preparedQuery{
            sql: "select * from words_forms where form in (?" + strings.Repeat(",?", len(page)-1) + ")",
            bindings: page,
        }
		queries = append(queries, query)

		offset += 10
	}

	return queries
}

func paginate(x []string, skip int, size int) []string {
    if skip > len(x) {
        skip = len(x)
    }

    end := skip + size
    if end > len(x) {
        end = len(x)
    }

    return x[skip:end]
}

func fetch(db *sql.DB) []WordForm {
    rows, err := db.Query("SELECT * FROM words_forms")
    if err != nil {
        fmt.Println("Problem")
    }

    var wordForms []WordForm
    for rows.Next() {
        var id, wordId int
        var form string
        err = rows.Scan(&id, &wordId, &form)
        if err != nil {
            panic(err.Error())
        }
        wordForm := WordForm{id, wordId, form}
        wordForms = append(wordForms, wordForm)
    }
    return wordForms
}