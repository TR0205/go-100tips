package code

import (
	"database/sql"
	"log"
	"os"
)

var db *sql.DB

// init関数を使うべきではない例
// init関数はテストの際にテストよりも先に実行されてしまう。
// そのためこのロジックのみの単体テストができない
func init() {
	dataSourceName := os.Getenv("MYSQL_DATA_SOURCE_NAME")
	d, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		// return err
		// init関数はエラーを返せないため、呼び出し側でエラー処理を実装することができない
		// この場合、エラーが発生したらアプリケーションは終了してしまい、リトライ処理の実装などが不可
		log.Panic(err)
	}
	err = d.Ping()
	if err != nil {
		log.Panic(err)
	}
	// グローバル変数はどのパッケージからでも値を書き換える事ができてしまうため、使用するべきではない
	db = d
}

// init関数を別関数に切り出すことで、上の3つの問題を解消できる
func createClient(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
