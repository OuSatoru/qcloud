// Wechat access token
package runner

import (
	"database/sql"
	"time"
	_ "github.com/lib/pq"
	"log"
	"github.com/OuSatoru/qcloud/wechat"
	"fmt"
)

type atdb struct {
	//id          int
	jikan       time.Time
	accesstoken sql.NullString
	expiresin   sql.NullInt64
	errcode     sql.NullInt64
	errmsg      sql.NullString
}

type DbLogin struct {
	DbUser string
	DbPwd string
}

func (db DbLogin) RunningGetAccToken(wat wechat.AccessToken) {
	for {
		fmt.Println("Here Here")
		insertAccToken(db, wat)
		_, lastExpire := lastTimeExpire(db)
		// 1 hour suggested, now a little time before 7200 secs
		time.Sleep(time.Duration(lastExpire*1000-233) * time.Millisecond)

	}
}

func insertAccToken(db DbLogin, wat wechat.AccessToken) {
	r, err := wat.FetchAtResp()
	if err != nil {
		log.Println(err)
		return
	}
	if r.AccessToken != "" {
		jikan := time.Now()
		exn, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@localhost/wechat?sslmode=disable", db.DbUser, db.DbPwd))
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
		exn, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@localhost/wechat?sslmode=disable", db.DbUser, db.DbPwd))
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

func lastTimeExpire(db DbLogin) (time.Time, int64) {
	exn, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@localhost/wechat?sslmode=disable", db.DbUser, db.DbPwd))
	if err != nil {
		log.Println(err)
		return time.Now(), 7200
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
		return time.Now(), 7200
	}
	defer rows.Close()
	var a atdb
	for rows.Next() {
		rows.Scan(&a.jikan, &a.accesstoken, &a.expiresin, &a.errcode, &a.errmsg)
	}
	return a.jikan, a.expiresin.Int64
}
