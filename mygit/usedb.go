package main
import(
	"fmt"
	"database/sql"
	"time"
	_ "go-sqlite3"
)
/*
 * 添加文件到仓库主入口
 */
func use_db(){
	fmt.Printf("usedb :\n")
	db,err:= sql.Open("sqlite3","./foo.db")
	checkErr(err)
	sql_table := `
    CREATE TABLE IF NOT EXISTS userinfo(
        uid INTEGER PRIMARY KEY AUTOINCREMENT,
        username VARCHAR(64) NULL,
        departname VARCHAR(64) NULL,
        created DATE NULL
    );
	`
	db.Exec(sql_table)
	stmt, err := db.Prepare("INSERT INTO userinfo(username, departname, created) values(?,?,?)")
    checkErr(err)

    res, err := stmt.Exec("wangshubo", "国务院", "2017-04-21")
    checkErr(err)

    id, err := res.LastInsertId()
    checkErr(err)

    fmt.Println(id)

    // update
    stmt, err = db.Prepare("update userinfo set username=? where uid=?")
    checkErr(err)

    res, err = stmt.Exec("wangshubo_new", id)
    checkErr(err)

    affect, err := res.RowsAffected()
    checkErr(err)

    fmt.Println(affect)

    // query
    rows, err := db.Query("SELECT * FROM userinfo")
    checkErr(err)
    var uid int
    var username string
    var department string
    var created time.Time

    for rows.Next() {
        err = rows.Scan(&uid, &username, &department, &created)
        checkErr(err)
        fmt.Println(uid)
        fmt.Println(username)
        fmt.Println(department)
        fmt.Println(created)
    }

    rows.Close()

    // delete
    stmt, err = db.Prepare("delete from userinfo where uid=?")
    checkErr(err)

    res, err = stmt.Exec(id)
    checkErr(err)

    affect, err = res.RowsAffected()
    checkErr(err)

    fmt.Println(affect)

    db.Close()
	return
}
func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}