// Wechat access token
package runner

import (
	"database/sql"
	"time"
	_ "github.com/lib/pq"
	"log"
	"github.com/OuSatoru/qcloud/wechat"
)

type atdb struct {
	//id          int
	jikan       time.Time
	accesstoken sql.NullString
	expiresin   sql.NullString
	errcode     sql.NullInt64
	errmsg      sql.NullString
}

func InsertAccToken(wat wechat.AccessToken) {
	r, err := wat.FetchAtResp()
	if err != nil {
		log.Println(err)
		return
	}
	if r.AccessToken != "" {
		jikan := time.Now()
		//shadow
		exn, err := sql.Open("postgres", "postgres:///wechat?sslmode=disable")
		if err != nil {
			log.Println(err)
			return
		}
		stmt, err := exn.Prepare(`INSERT INTO accesstoken (jikan, accesstoken, expiresin) VALUES ($1, $2, $3)`)
		if err != nil {
			log.Println(err)
			return
		}
		stmt.Exec(jikan, r.AccessToken, r.ExpiresIn)
	} else if r.ErrCode != 0 {
		jikan := time.Now()
		exn, err := sql.Open("postgres", "postgres:///wechat?sslmode=disable")
		if err != nil {
			log.Println(err)
			return
		}
		stmt, err := exn.Prepare(`INSERT INTO accesstoken (jikan, errcode, errmsg) VALUES ($1, $2, $3)`)
		if err != nil {
			log.Println(err)
			return
		}
		stmt.Exec(jikan, r.ErrCode, r.ErrMsg)
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
