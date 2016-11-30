// Wechat access token
package runner

import (
	"database/sql"
	"time"
	_ "github.com/lib/pq"
	"log"
)

type atdb struct {
	//id          int
	jikan       time.Time
	accesstoken sql.NullString
	expiresin   sql.NullString
	errcode     sql.NullInt64
	errmsg      sql.NullString
}

func InsertAccToken() {
	if time.Since(lastTime()) < -7200*time.Second {

	}
}

func lastTime() time.Time {
	// shadow it
	exn, err := sql.Open("postgres", "postgres:///wechat?sslmode=disable")
	if err != nil {
		log.Println(err)
		return nil
	}
	rows, err := exn.Query(`SELECT
				  jikan,
				  accesstoken,
				  expiresin,
				  errcode,
				  errmsg
				FROM accesstoken
				WHERE id = (SELECT max(id)
					    FROM
					      (SELECT *
					       FROM accesstoken
					       WHERE accesstoken IS NOT NULL) a)`)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer rows.Close()
	var a atdb
	rows.Scan(&a.jikan, &a.accesstoken, &a.expiresin, &a.errcode, &a.errmsg)
	return a.jikan
}
