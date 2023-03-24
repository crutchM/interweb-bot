package pg

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"log"
)

type dbLogger struct{}

func (d dbLogger) BeforeQuery(q *pg.QueryEvent) {}

func (d dbLogger) AfterQuery(q *pg.QueryEvent) {
	sql, _ := q.FormattedQuery()
	log.Println(sql)
}

func New(path string, isDebug bool) *pg.DB {
	orm.SetTableNameInflector(func(s string) string {
		return s
	})

	opt, err := pg.ParseURL(path)
	if err != nil {
		panic(err.Error())
	}

	db := pg.Connect(opt)
	if isDebug {
		db.AddQueryHook(dbLogger{})
	}
	return db
}
